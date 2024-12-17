package order

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type Status string

const (
	Unknown    Status = "unknown"
	Pending    Status = "pending"
	Delivering Status = "delivering"
	Delivered  Status = "delivered"
	Cancelled  Status = "cancelled"
)

type Request struct {
	CartItemIds []int `json:"cartItemIds"` // Идентификаторы товаров в корзине, которые будут включены в заказ
}

type Order struct {
	Id             uuid.UUID       `json:"id"`             // Уникальный идентификатор заказа
	CreateDate     time.Time       `json:"createDate"`     // Дата и время создания заказа
	DeliveryWeight decimal.Decimal `json:"deliveryWeight"` // Общий вес доставки для всех товаров в заказе
	Total          decimal.Decimal `json:"total"`          // Общая сумма заказа
	Status         Status          `json:"status"`         // Текущий статус заказа (например, "в процессе", "доставлен", "отменен")
}
