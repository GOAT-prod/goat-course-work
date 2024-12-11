package domain

import "github.com/shopspring/decimal"

type ProductItem struct {
	Id         int             `json:"id"`
	StockCount int             `json:"stockCount"`
	Size       int             `json:"size"`
	Weight     decimal.Decimal `json:"weight"`
	Color      string          `json:"color"`
}
