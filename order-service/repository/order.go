package repository

import (
	"github.com/GOAT-prod/goatcontext"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"order-service/database"
	"order-service/database/queries"
)

type Order interface {
	GetUserOrders(ctx goatcontext.Context, userId int) (orders []database.Order, err error)
	CreateOrder(ctx goatcontext.Context, order database.Order) error
	GetOrderItems(ctx goatcontext.Context, orderId uuid.UUID) (items []database.OrderItem, err error)
	CreateOrderItem(ctx goatcontext.Context, orderItem database.OrderItem) error
}

type OrderRepositoryImpl struct {
	postgres *sqlx.DB
}

func NewOrderRepository(postgres *sqlx.DB) Order {
	return &OrderRepositoryImpl{
		postgres: postgres,
	}
}

func (r *OrderRepositoryImpl) GetUserOrders(ctx goatcontext.Context, userId int) (orders []database.Order, err error) {
	return orders, r.postgres.SelectContext(ctx, &orders, queries.GetOrder, userId)
}

func (r *OrderRepositoryImpl) CreateOrder(ctx goatcontext.Context, order database.Order) error {
	_, err := r.postgres.NamedExecContext(ctx, queries.CreateOrder, order)
	return err
}

func (r *OrderRepositoryImpl) GetOrderItems(ctx goatcontext.Context, orderId uuid.UUID) (items []database.OrderItem, err error) {
	return items, r.postgres.SelectContext(ctx, &items, queries.GetOrderItems, orderId)
}

func (r *OrderRepositoryImpl) CreateOrderItem(ctx goatcontext.Context, orderItem database.OrderItem) error {
	_, err := r.postgres.NamedExecContext(ctx, queries.CreateOrderItem, orderItem)
	return err
}
