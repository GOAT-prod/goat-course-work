package database

import "time"

type Request struct {
	Id         int           `db:"id" json:"id"`
	Status     string        `db:"status" json:"status"`
	Type       string        `db:"type" json:"type"`
	UpdateDate time.Time     `db:"update_date" json:"update_date"`
	Summary    string        `db:"summary" json:"summary"`
	Items      []RequestItem `db:"-" json:"items"`
}

type RequestItem struct {
	Id               int `db:"id" json:"id"`
	RequestId        int `db:"request_id" json:"request_id"`
	ProductId        int `db:"product_id" json:"product_id"`
	ProductItemId    int `db:"product_item_id" json:"product_item_id"`
	ProductItemCount int `db:"product_item_count" json:"product_item_count"`
}
