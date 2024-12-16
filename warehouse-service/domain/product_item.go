package domain

import "github.com/shopspring/decimal"

type ProductItem struct {
	Id         int             `json:"id"`
	StockCount int             `json:"stockCount"`
	Size       int             `json:"size"`
	Weight     decimal.Decimal `json:"weight"`
	Color      string          `json:"color"`
}

type ProductItemInfo struct {
	Id    int             `json:"id"`
	Name  string          `json:"name"`
	Price decimal.Decimal `json:"price"`
	Color string          `json:"color"`
	Size  int             `json:"size"`
}
