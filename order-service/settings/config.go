package settings

type Config struct {
	Port      int      `json:"port"`
	AppName   string   `json:"app_name"`
	Databases Database `json:"databases"`
	Cluster   Cluster  `json:"cluster"`
	Kafka     Kafka    `json:"kafka"`
}

type Database struct {
	Postgres string `json:"postgres"`
}

type Cluster struct {
	WarehouseService string `json:"warehouse_service_url"`
	CartService      string `json:"cart_service_url"`
}

type Kafka struct {
	Address string `json:"address"`
	Topic   string `json:"topic"`
}
