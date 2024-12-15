package authservice

// LoginData представляет данные для входа пользователя.
type LoginData struct {
	Username string `json:"username"` // Имя пользователя для входа в систему.
	Password string `json:"password"` // Пароль пользователя для входа в систему.
}

// RegisterData представляет данные для регистрации нового пользователя.
type RegisterData struct {
	UserName      string `json:"username"` // Имя пользователя, которое будет использоваться при регистрации.
	Password      string `json:"password"` // Пароль, который пользователь создаст при регистрации.
	Role          string `json:"role"`     // Роль пользователя, например, администратор или обычный пользователь.
	ClientName    string `json:"name"`     // Имя клиента, связанного с пользователем.
	ClientAddress string `json:"address"`  // Адрес клиента, связанного с пользователем.
	ClientINN     string `json:"inn"`      // ИНН клиента, связанного с пользователем.
}

// Tokens представляет пару токенов для аутентификации пользователя.
type Tokens struct {
	AccessToken  string `json:"access_token"`  // Токен доступа для авторизации запросов.
	RefreshToken string `json:"refresh_token"` // Токен для обновления access_token, если он истек.
}
