package service

import (
	"errors"
	"github.com/GOAT-prod/goatcontext"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"request-service/cluster/notifier"
	"request-service/cluster/warehouse"
	"request-service/database"
	"request-service/domain"
	"testing"
	"time"
)

// Мок для репозитория запросов
type MockRequestRepository struct {
	mock.Mock
}

func (m *MockRequestRepository) GetPendingRequests(ctx goatcontext.Context) ([]database.Request, error) {
	args := m.Called(ctx)
	return args.Get(0).([]database.Request), args.Error(1)
}

func (m *MockRequestRepository) UpdateRequestStatus(ctx goatcontext.Context, id int, status string) error {
	args := m.Called(ctx, id, status)
	return args.Error(0)
}

func (m *MockRequestRepository) GetRequestById(ctx goatcontext.Context, id int) (database.Request, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(database.Request), args.Error(1)
}

// Мок для клиента склада
type MockWarehouseClient struct {
	mock.Mock
}

func (m *MockWarehouseClient) GetDetailedProduct(ctx goatcontext.Context, productId string) (domain.Product, error) {
	args := m.Called(ctx, productId)
	return args.Get(0).(domain.Product), args.Error(1)
}

// Мок для клиента уведомлений
type MockNotifierClient struct {
	mock.Mock
}

func (m *MockNotifierClient) SendMail(ctx goatcontext.Context, mail notifier.Mail) error {
	args := m.Called(ctx, mail)
	return args.Error(0)
}

func TestImpl_UpdateStatus(t *testing.T) {
	// Подготовка общих данных для тестов
	ctx := goatcontext.NewContext()
	requestId := 123
	productId := "product-1"
	
	testProduct := domain.Product{
		Id:   productId,
		Name: "Test Product",
		Items: []domain.ProductItem{
			{
				Id:    productId,
				Color: "Red",
				Size:  "M",
			},
		},
	}
	
	testRequest := database.Request{
		Id:         requestId,
		Status:     "PENDING",
		Type:       "SUPPLY",
		UpdateDate: time.Now(),
		Items: []database.RequestItem{
			{
				ProductId:         productId,
				ProductItemCount:  5,
			},
		},
	}

	t.Run("успешное обновление статуса (не ApprovedStatus)", func(t *testing.T) {
		// Arrange
		mockRepo := new(MockRequestRepository)
		mockWarehouse := new(MockWarehouseClient)
		mockNotifier := new(MockNotifierClient)
		
		service := NewRequestService(mockRepo, &warehouse.Client{Client: mockWarehouse}, &notifier.Client{Client: mockNotifier})
		
		mockRepo.On("UpdateRequestStatus", ctx, requestId, "REJECTED").Return(nil)
		
		// Act
		err := service.UpdateStatus(ctx, requestId, domain.Status("REJECTED"))
		
		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
		mockWarehouse.AssertNotCalled(t, "GetDetailedProduct")
		mockNotifier.AssertNotCalled(t, "SendMail")
	})

	t.Run("успешное обновление статуса до ApprovedStatus", func(t *testing.T) {
		// Arrange
		mockRepo := new(MockRequestRepository)
		mockWarehouse := new(MockWarehouseClient)
		mockNotifier := new(MockNotifierClient)
		
		service := NewRequestService(mockRepo, &warehouse.Client{Client: mockWarehouse}, &notifier.Client{Client: mockNotifier})
		
		mockRepo.On("UpdateRequestStatus", ctx, requestId, "APPROVED").Return(nil)
		mockRepo.On("GetRequestById", ctx, requestId).Return(testRequest, nil)
		mockWarehouse.On("GetDetailedProduct", ctx, productId).Return(testProduct, nil)
		mockNotifier.On("SendMail", ctx, mock.Anything).Return(nil)
		
		// Act
		err := service.UpdateStatus(ctx, requestId, domain.ApprovedStatus)
		
		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
		
		// Даем время для выполнения горутины
		time.Sleep(100 * time.Millisecond)
		mockWarehouse.AssertExpectations(t)
		mockNotifier.AssertExpectations(t)
	})

	t.Run("ошибка при обновлении статуса", func(t *testing.T) {
		// Arrange
		mockRepo := new(MockRequestRepository)
		mockWarehouse := new(MockWarehouseClient)
		mockNotifier := new(MockNotifierClient)
		
		service := NewRequestService(mockRepo, &warehouse.Client{Client: mockWarehouse}, &notifier.Client{Client: mockNotifier})
		
		expectedErr := errors.New("database error")
		mockRepo.On("UpdateRequestStatus", ctx, requestId, "APPROVED").Return(expectedErr)
		
		// Act
		err := service.UpdateStatus(ctx, requestId, domain.ApprovedStatus)
		
		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		mockRepo.AssertExpectations(t)
		mockWarehouse.AssertNotCalled(t, "GetDetailedProduct")
		mockNotifier.AssertNotCalled(t, "SendMail")
	})

	t.Run("ошибка при отправке сообщения после обновления до ApprovedStatus", func(t *testing.T) {
		// Arrange
		mockRepo := new(MockRequestRepository)
		mockWarehouse := new(MockWarehouseClient)
		mockNotifier := new(MockNotifierClient)
		
		service := NewRequestService(mockRepo, &warehouse.Client{Client: mockWarehouse}, &notifier.Client{Client: mockNotifier})
		
		mockRepo.On("UpdateRequestStatus", ctx, requestId, "APPROVED").Return(nil)
		mockRepo.On("GetRequestById", ctx, requestId).Return(testRequest, nil)
		mockWarehouse.On("GetDetailedProduct", ctx, productId).Return(testProduct, nil)
		mockNotifier.On("SendMail", ctx, mock.Anything).Return(errors.New("notification error"))
		
		// Act
		err := service.UpdateStatus(ctx, requestId, domain.ApprovedStatus)
		
		// Assert
		assert.NoError(t, err) // Основной метод должен завершиться успешно
		mockRepo.AssertExpectations(t)
		
		// Даем время для выполнения горутины
		time.Sleep(100 * time.Millisecond)
		mockWarehouse.AssertExpectations(t)
		mockNotifier.AssertExpectations(t)
	})
}