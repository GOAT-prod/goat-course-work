package kafka

import "time"

type Request struct {
	Id          int       `json:"id"`
	Status      string    `json:"status"`
	Type        string    `json:"type"`
	UpdatedDate time.Time `json:"updated_date"`
	Summary     string    `json:"summary"`

	Items []RequestItem `json:"items"`
}

type RequestItem struct {
	Id               int `json:"id"`
	RequestId        int `json:"request_id"`
	ProductId        int `json:"product_id"`
	ProductItemId    int `json:"product_item_id"`
	ProductItemCount int `json:"product_item_count"`
}
