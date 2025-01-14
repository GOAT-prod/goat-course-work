package settings

type Config struct {
	Port      int      `json:"port"`
	AppName   string   `json:"app_name"`
	Databases Database `json:"databases"`
	Cluster   Cluster  `json:"cluster"`
}

type Database struct {
	Mongo Mongo `json:"mongo"`
	Redis Redis `json:"redis"`
}

type Mongo struct {
	Connection string `json:"connection"`
	Database   string `json:"database"`
	Collection string `json:"collection"`
}

type Redis struct {
	Address  string `json:"address"`
	Password string `json:"password"`
	Database int    `json:"database"`
}

type Cluster struct {
	WarehouseService string `json:"warehouse_service_url"`
}
