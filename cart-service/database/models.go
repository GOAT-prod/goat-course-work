package database

import "time"

type Cart struct {
	Id         int       `db:"id"`
	CreateDate time.Time `db:"create_date"`
	UserId     int       `db:"user_id"`
	Items      []CartItem
}

type CartItem struct {
	Id            int  `db:"id"`
	ProductItemId int  `db:"product_item_id"`
	Quantity      int  `db:"quantity"`
	CartId        int  `db:"cart_id"`
	IsSelected    bool `db:"is_selected"`
}
