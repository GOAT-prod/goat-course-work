package clientservice

// ClientInfo представляет информацию о клиенте.
type ClientInfo struct {
	Id      int    `json:"id"`      // Уникальный идентификатор клиента.
	Name    string `json:"name"`    // Имя клиента.
	INN     string `json:"inn"`     // ИНН клиента.
	Address string `json:"address"` // Адрес клиента.
}
