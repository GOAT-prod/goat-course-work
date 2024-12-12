package database

type Client struct {
	Id      int    `db:"id"`
	Name    string `db:"name"`
	INN     string `db:"inn"`
	Address string `db:"address"`
}
