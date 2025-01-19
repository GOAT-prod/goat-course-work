package queries

import _ "embed"

var (
	//go:embed sql/get_clients.sql
	GetClients string

	//go:embed sql/get_client_by_id.sql
	GetClientById string

	//go:embed sql/add_client.sql
	AddClient string

	//go:embed sql/update_client.sql
	UpdateClient string

	//go:embed sql/delete_client.sql
	DeleteClient string

	//go:embed sql/get_clients_by_ids.sql
	GetClientsByIds string
)
