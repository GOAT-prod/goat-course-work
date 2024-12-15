package userservice

// User представляет информацию о пользователе.
type User struct {
	Id       int    `json:"id"`       // Уникальный идентификатор пользователя.
	Username string `json:"username"` // Имя пользователя для входа в систему.
	Password string `json:"password"` // Пароль пользователя для входа в систему.
	Status   int    `json:"status"`   // Статус пользователя (например, активен или заблокирован).
	Role     string `json:"role"`     // Роль пользователя (например, администратор или обычный пользователь).
	ClientId int    `json:"clientId"` // Идентификатор клиента, с которым связан пользователь.
}
