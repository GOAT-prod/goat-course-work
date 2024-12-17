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
	Id             uuid.UUID
	CreateDate     time.Time
	DeliveryWeight decimal.Decimal
	Total          decimal.Decimal
	Status         OrderStatus
}
