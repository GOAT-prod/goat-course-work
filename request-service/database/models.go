package database

import "time"

type Request struct {
	Id         int       `db:"id"`
	Status     string    `db:"status"`
	Type       string    `db:"type"`
	UpdateDate time.Time `db:"update_date"`
	Summary    string    `db:"summary"`

	Items []RequestItem
}

type RequestItem struct {
	Id               int `db:"id"`
	RequestId        int `db:"request_id"`
	ProductId        int `db:"product_id"`
	ProductItemId    int `db:"product_item_id"`
	ProductItemCount int `db:"product_item_count"`
}
