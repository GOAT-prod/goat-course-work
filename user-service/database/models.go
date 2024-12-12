package database

type User struct {
	Id       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Status   int    `db:"status"`
	RoleId   int    `db:"role_id"`
	ClientId int    `db:"client_id"`
}

type Role struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}
