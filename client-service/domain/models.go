package domain

type Client struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	INN     string `json:"inn"`
	Address string `json:"address"`
}
