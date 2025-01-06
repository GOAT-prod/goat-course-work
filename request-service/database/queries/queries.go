package queries

import _ "embed"

var (
	//go:embed sql/get_request_by_id.sql
	GetRequestById string

	//go:embed sql/get_pending_requests.sql
	GetPendingRequests string

	//go:embed sql/get_request_items_by_request_id.sql
	GetRequestItemsByRequestId string

	//go:embed sql/add_request.sql
	AddRequest string

	//go:embed sql/add_request_item.sql
	AddRequestItem string

	//go:embed sql/update_request_status.sql
	UpdateRequestStatus string
)
