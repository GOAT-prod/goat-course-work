package domain

type ProductStatus string

const (
	Unknown        ProductStatus = "unknown"
	WaitingApprove ProductStatus = "waiting_approve"
	Approved       ProductStatus = "approved"
	Editing        ProductStatus = "editing"
	Deleted        ProductStatus = "deleted"
)
