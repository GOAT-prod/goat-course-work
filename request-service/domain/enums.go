package domain

type Type string

const (
	UndefinedType Type = "undefined"
	ApproveType   Type = "approve"
	SupplyType    Type = "supply"
)

type Status string

const (
	UndefinedStatus Status = "undefined"
	PendingStatus   Status = "pending"
	ApprovedStatus  Status = "approved"
	RejectedStatus  Status = "rejected"
)

var (
	StatusToDomain map[string]Status = map[string]Status{
		"undefined": UndefinedStatus,
		"pending":   PendingStatus,
		"approved":  ApprovedStatus,
		"rejected":  RejectedStatus,
	}
)
