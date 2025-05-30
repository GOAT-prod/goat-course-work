package warehouse

import "github.com/shopspring/decimal"

type ProductItemInfo struct {
	Id        int             `json:"id"`
	ProductId int             `json:"productId"`
	FactoryId int             `json:"factoryId"`
	Name      string          `json:"name"`
	Price     decimal.Decimal `json:"price"`
	Color     string          `json:"color"`
	Size      int             `json:"size"`
	Weight    decimal.Decimal `json:"weight"`
	Count     int             `json:"count"`
}
