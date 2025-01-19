package service

import (
	"fmt"
	"github.com/GOAT-prod/goatcontext"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"log"
	"order-service/cluster/cart"
	"order-service/cluster/warehouse"
	"order-service/database"
	"order-service/domain"
	"order-service/kafka"
	"order-service/repository"
	"time"
)

type Order interface {
	Order(ctx goatcontext.Context, cartItemsIds []int) error
	GetUserOrders(ctx goatcontext.Context) ([]domain.Order, error)
	GetLatestOrders(ctx goatcontext.Context) ([]domain.ReportItem, error)
}

type OrderServiceImpl struct {
	cartService       *cart.Client
	warehouseService  *warehouse.Client
	orderRepository   repository.Order
	financeRepository repository.Finance
	kafkaProducer     *kafka.Producer
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

	warehouseProductItems, err := s.warehouseService.GetProductItemsInfo(ctx, lo.Map(itemsForOrder, func(item cart.Item, _ int) int {
		return item.ProductItemId
	}))
	if err != nil {
		return fmt.Errorf("не удалось получить покупаемые товары со склада: %w", err)
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

	supplyItems := make([]kafka.RequestItem, 0, len(itemsForOrder))

	totalOrderPrice := decimal.Zero
	for _, item := range itemsForOrder {
		warehouseProductItem, ok := lo.Find(warehouseProductItems, func(productItem warehouse.ProductItemInfo) bool {
			return productItem.Id == item.ProductItemId
		})

		if !ok {
			return fmt.Errorf("вариант продукта %d из корзины отсутствует на складе", item.ProductItemId)
		}

		if warehouseProductItem.Count < item.Count {
			supplyItems = append(supplyItems, kafka.RequestItem{
				ProductId:        warehouseProductItem.ProductId,
				ProductItemId:    item.ProductItemId,
				ProductItemCount: item.Count - warehouseProductItem.Count,
			})
		}

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

	if len(supplyItems) > 0 {
		go s.produceSupplyMessage(supplyItems)
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

func (s *OrderServiceImpl) GetLatestOrders(ctx goatcontext.Context) ([]domain.ReportItem, error) {
	startTime := time.Now().Add(-15 * time.Minute)
	endTime := time.Now()

	latestOrders, err := s.orderRepository.GetLatestOrders(ctx, startTime, endTime)
	if err != nil {
		return nil, err
	}

	productItemIds := lo.Map(latestOrders, func(item database.LatestOrder, _ int) int {
		return item.ProductItemId
	})

	productItemInfos, err := s.warehouseService.GetProductItemsInfo(ctx, productItemIds)
	if err != nil {
		return nil, err
	}

	result := make([]domain.ReportItem, 0, len(productItemInfos))

	for _, order := range latestOrders {
		productItem, ok := lo.Find(productItemInfos, func(item warehouse.ProductItemInfo) bool {
			return item.Id == order.ProductItemId
		})

		if !ok {
			continue
		}

		result = append(result, domain.ReportItem{
			Date:        order.Date,
			ProductName: productItem.Name,
			FactoryId:   productItem.FactoryId,
			UserId:      order.UserId,
			Color:       productItem.Color,
			Size:        productItem.Size,
			Count:       order.Quantity,
			Price:       order.Price,
		})
	}

	return result, nil
}

func (s *OrderServiceImpl) produceSupplyMessage(supplyItems []kafka.RequestItem) {
	supplyItemsByProductId := associateRequestItemByProducts(supplyItems)
	for _, supplyItem := range supplyItemsByProductId {
		request := kafka.Request{
			Status:      "pending",
			Type:        "supply",
			UpdatedDate: time.Now(),
			Summary:     "необходима поставка продукта на склад",
			Items:       supplyItem,
		}

		if err := s.kafkaProducer.Produce(request); err != nil {
			log.Println("не удалось записать сообщение в кафку")
		}
	}
}

func associateRequestItemByProducts(requestItems []kafka.RequestItem) map[int][]kafka.RequestItem {
	result := make(map[int][]kafka.RequestItem)

	for _, requestItem := range requestItems {
		if _, ok := result[requestItem.ProductId]; !ok {
			result[requestItem.ProductId] = []kafka.RequestItem{requestItem}
			continue
		}

		result[requestItem.ProductId] = append(result[requestItem.ProductId], requestItem)
	}

	return result
}
