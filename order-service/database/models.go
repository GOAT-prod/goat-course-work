package database

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type OperationDetailType int

const (
	Undefined OperationDetailType = iota
	OrderOperation
	DeliveryOperation
)

type Order struct {
	Id           uuid.UUID `db:"id"`
	Status       string    `db:"status"`
	CreateDate   time.Time `db:"create_date"`
	DeliveryDate time.Time `db:"delivery_date"`
	UserId       int       `db:"user_id"`
}

type OrderItem struct {
	Id            uuid.UUID `db:"id"`
	OrderId       string    `db:"order_id"`
	ProductItemId int       `db:"product_item_id"`
	Quantity      int       `db:"quantity"`
}

type Operation struct {
	Id          uuid.UUID `db:"id"`
	Date        time.Time `db:"date"`
	Description string    `db:"description"`
	OrderId     uuid.UUID `db:"order_id"`
}

type OperationDetail struct {
	Id          uuid.UUID           `db:"id"`
	OperationId uuid.UUID           `db:"operation_id"`
	Type        OperationDetailType `db:"type"`
	Price       decimal.Decimal     `db:"price"`
}