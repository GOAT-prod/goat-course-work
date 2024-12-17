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
