package cart

import "github.com/shopspring/decimal"

type Item struct {
	Id            int             `json:"id"`
	ProductItemId int             `json:"productId"`
	Name          string          `json:"name"`
	Price         decimal.Decimal `json:"price"`
	Color         string          `json:"color"`
	Size          int             `json:"size"`
	Count         int             `json:"count"`
	IsSelected    bool            `json:"isSelected"`
}
