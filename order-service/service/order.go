package service

import (
	"fmt"
	"github.com/GOAT-prod/goatcontext"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"order-service/cluster/cart"
	"order-service/cluster/warehouse"
	"order-service/database"
	"order-service/domain"
	"order-service/repository"
	"time"
)

type Order interface {
	Order(ctx goatcontext.Context, cartItemsIds []int) error
	GetUserOrders(ctx goatcontext.Context) ([]domain.Order, error)
}

type OrderServiceImpl struct {
	cartService       *cart.Client
	warehouseService  *warehouse.Client
	orderRepository   repository.Order
	financeRepository repository.Finance
}

func NewOrderService(cartService *cart.Client, warehouseService *warehouse.Client, orderRepository repository.Order, financeRepository repository.Finance) Order {
	return &OrderServiceImpl{
		cartService:       cartService,
		warehouseService:  warehouseService,
		orderRepository:   orderRepository,
		financeRepository: financeRepository,
	}
}

func (s *OrderServiceImpl) Order(ctx goatcontext.Context, cartItemsIds []int) error {
	itemsForOrder, err := s.cartService.GetCartItems(ctx, cartItemsIds)
	if err != nil {
		return fmt.Errorf("не удалось получить айдемы из корзины для заказа: %w", err)
	}

	orderId, err := uuid.NewV7()
	if err != nil {
		return fmt.Errorf("не удалось создать id заказа: %w", err)
	}

	order := database.Order{
		Id:         orderId,
		Status:     string(domain.Pending),
		CreateDate: time.Now(),
		UserId:     ctx.Authorize().UserId,
	}

	if err = s.orderRepository.CreateOrder(ctx, order); err != nil {
		return fmt.Errorf("не удалось создать заказ: %w", err)
	}

	totalOrderPrice := decimal.Zero
	for _, item := range itemsForOrder {
		//TODO: проверка на количество товара на складе для реквестов
		totalOrderPrice = totalOrderPrice.Add(item.Price.Mul(decimal.NewFromInt(int64(item.Count))))

		orderItemId, uErr := uuid.NewV7()
		if uErr != nil {
			return fmt.Errorf("не удалось создать id айтема заказа: %w", err)
		}

		orderItem := database.OrderItem{
			Id:            orderItemId,
			OrderId:       orderId,
			ProductItemId: item.ProductItemId,
			Quantity:      item.Count,
		}

		if err = s.orderRepository.CreateOrderItem(ctx, orderItem); err != nil {
			return fmt.Errorf("не удалось создать айтем заказа: %w", err)
		}
	}

	operationId, err := uuid.NewV7()
	if err != nil {
		return fmt.Errorf("не удалось создать id фин операции: %w", err)
	}

	operation := database.Operation{
		Id:          operationId,
		Date:        time.Now(),
		Description: fmt.Sprintf("Покупка кроссовок. Пользователь %s", ctx.Authorize().Username),
		OrderId:     orderId,
	}

	if err = s.financeRepository.CreateOrderOperation(ctx, operation); err != nil {
		return fmt.Errorf("не удалось создать оперцию: %w", err)
	}

	operationDetailId, err := uuid.NewV7()
	if err != nil {
		return fmt.Errorf("не удалось создать id детали операци: %w", err)
	}

	operationDetail := database.OperationDetail{
		Id:          operationDetailId,
		OperationId: operationId,
		Type:        database.OrderOperation,
		Price:       totalOrderPrice,
	}

	if err = s.financeRepository.CreateOperationDetail(ctx, operationDetail); err != nil {
		return fmt.Errorf("не удалось создать деталь фин операции: %w", err)
	}

	return nil
}

func (s *OrderServiceImpl) GetUserOrders(ctx goatcontext.Context) ([]domain.Order, error) {
	dbOrders, err := s.orderRepository.GetUserOrders(ctx, ctx.Authorize().UserId)
	if err != nil {
		return nil, err
	}

	userOrders := make([]domain.Order, 0, len(dbOrders))
	for _, dbOrder := range dbOrders {
		order := domain.Order{
			Id:         dbOrder.Id,
			CreateDate: dbOrder.CreateDate,
			Status:     domain.OrderStatus(dbOrder.Status),
		}

		operation, oErr := s.financeRepository.GetOrderOperation(ctx, dbOrder.Id)
		if oErr != nil {
			return nil, oErr
		}

		operationDetails, odErr := s.financeRepository.GetOperationDetails(ctx, operation.Id)
		if odErr != nil {
			return nil, odErr
		}

		orderOperationDetail, _ := lo.Find(operationDetails, func(item database.OperationDetail) bool { return item.Type == database.OrderOperation })
		order.Total = orderOperationDetail.Price

		orderItems, oiErr := s.orderRepository.GetOrderItems(ctx, dbOrder.Id)
		if oiErr != nil {
			return nil, oiErr
		}

		productItemsIds := lo.Map(orderItems, func(item database.OrderItem, _ int) int { return item.ProductItemId })

		productItemsInfo, pErr := s.warehouseService.GetProductItemsInfo(ctx, productItemsIds)
		if pErr != nil {
			return nil, pErr
		}

		totalWeightFunc := func(infos []warehouse.ProductItemInfo) decimal.Decimal {
			totalWeight := decimal.Zero
			for _, info := range infos {
				totalWeight = totalWeight.Add(info.Weight)
			}

			return totalWeight
		}

		order.DeliveryWeight = totalWeightFunc(productItemsInfo)

		userOrders = append(userOrders, order)
	}

	return userOrders, nil
}
