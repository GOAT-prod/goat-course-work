package models

import "github.com/shopspring/decimal"

type ProductItem struct {
	Id         int             `db:"id" json:"id"`
	ProductId  int             `db:"product_id" json:"productId"`
	StockCount int             `db:"stock_count" json:"stockCount"`
	Size       int             `db:"size" json:"size"`
	Weight     decimal.Decimal `db:"weight" json:"weight"`
	Color      string          `db:"color" json:"color"`
}