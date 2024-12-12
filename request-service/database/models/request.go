package models

import "time"

type Request struct {
	Id         int       `db:"id"`
	Status     int       `db:"status"`
	Type       int       `db:"type"`
	UpdateDate time.Time `db:"update_date"`
	Summary    string    `db:"summary"`
	Items      []RequestItem
}

type RequestItem struct {
	Id            int `db:"id"`
	RequestId     int `db:"request_id"`
	ProductId     int `db:"product_id"`
	ProductItemId int `db:"product_item_id"`
}
