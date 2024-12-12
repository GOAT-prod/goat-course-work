package queries

import _ "embed"

var (
	//go:embed sql/get_users.sql
	GetUsers string

	//go:embed sql/get_user_by_id.sql
	GetUserById string

	//go:embed sql/get_user_by_username.sql
	GetUserByUsername string

	//go:embed sql/add_user.sql
	AddUser string

	//go:embed sql/update_user.sql
	UpdateUser string

	//go:embed sql/delete_user.sql
	DeleteUser string

	//go:embed sql/get_role_id_by_name.sql
	GetRoleIdByName string

	//go:embed sql/get_role_by_id.sql
	GetRoleById string
)
