package report

import "time"

type Report struct {
	Date       time.Time    `json:"date"`
	TotalPrice float64      `json:"totalPrice"`
	Items      []ReportItem `json:"items"`
}

type ReportItem struct {
	ProductName string  `json:"productName"`
	Color       string  `json:"color"`
	Size        int     `json:"size"`
	TotalCount  int     `json:"totalCount"`
	TotalPrice  float64 `json:"totalPrice"`
}
