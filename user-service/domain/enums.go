package domain

type UserRole string

const (
	Admin   UserRole = "admin"
	Shop    UserRole = "shop"
	Factory UserRole = "factory"
)

type UserStatus int

const (
	Undefined UserStatus = iota
	Active    UserStatus = iota
	Deleted   UserStatus = iota
)
