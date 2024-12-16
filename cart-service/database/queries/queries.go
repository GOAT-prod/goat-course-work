package queries

import _ "embed"

var (
	//go:embed sql/get_cart.sql
	GetCart string

	//go:embed sql/create_cart.sql
	CreateCart string

	//go:embed sql/get_cart_items.sql
	GetCartItems string

	//go:embed sql/add_cart_item.sql
	AddCartItem string

	//go:embed sql/update_cart_item.sql
	UpdateCartItem string

	//go:embed sql/delete_cart_items.sql
	DeleteCartItems string

	//go:embed sql/clear_cart_items.sql
	ClearCartItems string
)
