package queries

import _ "embed"

var (
	//go:embed sql/get_ordels.sql
	GetOrder string

	//go:embed sql/add_order.sql
	CreateOrder string

	//go:embed sql/get_order_item.sql
	GetOrderItems string

	//go:embed sql/add_order_item.sql
	CreateOrderItem string

	//go:embed sql/get_operation.sql
	GetOperation string

	//go:embed sql/add_operation.sql
	CreateOperation string

	//go:embed sql/get_operation_details.sql
	GetOperationDetails string

	//go:embed sql/add_operation_detail.sql
	CreateOperationDetail string

	//go:embed sql/get_latest_orders.sql
	GetLatestOrders string
)
