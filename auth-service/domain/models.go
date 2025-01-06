package domain

const (
	UndefinedRole int = iota
	UserRole
	AdminRole
)

type UserStatus int

const (
	Undefined UserStatus = iota
	Active    UserStatus = iota
	Deleted   UserStatus = iota
)

type User struct {
	Id       int        `json:"id"`
	UserName string     `json:"username"`
	Password string     `json:"password"`
	Status   UserStatus `json:"status"`
	Role     string     `json:"role"`
	ClientId int        `json:"clientId"`
}

type CheckUserResponse struct {
	Exist bool `json:"exist"`
}

type RegisterData struct {
	UserName      string `json:"username"`
	Password      string `json:"password"`
	Role          string `json:"role"`
	ClientName    string `json:"name"`
	ClientAddress string `json:"address"`
	ClientINN     string `json:"inn"`
}

type ClientData struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	INN     string `json:"inn"`
	Address string `json:"address"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenBase struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
	UserRole string `json:"user_role"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdatePasswordRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *RefreshTokenBase) IsEquals(compareValue RefreshTokenBase) bool {
	return r.UserId == compareValue.UserId && r.UserName == compareValue.UserName
}
