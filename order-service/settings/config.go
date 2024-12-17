package settings

type Config struct {
	Port      int      `json:"port"`
	AppName   string   `json:"app_name"`
	Databases Database `json:"databases"`
	Cluster   Cluster  `json:"cluster"`
}

type Database struct {
	Postgres string `json:"postgres"`
}

type Cluster struct {
	WarehouseService string `json:"warehouse_service_url"`
	CartService      string `json:"cart_service_url"`
}
