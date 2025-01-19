package domain

import (
	"github.com/shopspring/decimal"
	"time"
)

type Report struct {
	Date  time.Time       `json:"date"`
	Total decimal.Decimal `json:"total"`
	Items []ReportItem    `json:"items"`
}

type ReportItem struct {
	ProductName string          `json:"productName"`
	Color       string          `json:"color"`
	Size        int             `json:"size"`
	Count       int             `json:"count"`
	Price       decimal.Decimal `json:"price"`
}
