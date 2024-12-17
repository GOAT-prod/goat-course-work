package domain

type OrderStatus string

const (
	Unknown    OrderStatus = "unknown"
	Pending    OrderStatus = "pending"
	Delivering OrderStatus = "delivering"
	Delivered  OrderStatus = "delivered"
	Cancelled  OrderStatus = "cancelled"
)
