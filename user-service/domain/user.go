package domain

type User struct {
	Id       int        `json:"id"`
	Username string     `json:"username"`
	Password string     `json:"password"`
	Status   UserStatus `json:"status"`
	Role     UserRole   `json:"role"`
	ClientId int        `json:"clientId"`
}

type UpdatePasswordRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
