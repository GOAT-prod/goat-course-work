package domain

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type OrderRequest struct {
	CartItemIds []int `json:"cartItemIds"`
}

type Order struct {
	Id             uuid.UUID       `json:"id"`
	CreateDate     time.Time       `json:"createDate"`
	DeliveryWeight decimal.Decimal `json:"deliveryWeight"`
	Total          decimal.Decimal `json:"total"`
	Status         OrderStatus     `json:"status"`
}

type ReportItem struct {
	Date        time.Time       `json:"date" db:"date"`
	ProductName string          `json:"productName" db:"product_name"`
	FactoryId   int             `json:"factoryId" db:"factory_id"`
	UserId      int             `json:"userId" db:"user_id"`
	Color       string          `json:"color" db:"color"`
	Size        int             `json:"size" db:"size"`
	Count       int             `json:"count" db:"count"`
	Price       decimal.Decimal `json:"price" db:"price"`
}
