package domain

import "github.com/shopspring/decimal"

type Cart struct {
	Id    int        `json:"id"`
	Total int        `json:"total"`
	Items []CartItem `json:"items"`
}

type CartItem struct {
	Id            int             `json:"id"`
	ProductItemId int             `json:"productId"`
	Name          string          `json:"name"`
	Price         decimal.Decimal `json:"price"`
	Color         string          `json:"color"`
	Size          int             `json:"size"`
	Count         int             `json:"count"`
	IsSelected    bool            `json:"isSelected"`
}
