package settings

type Config struct {
	Port      int      `json:"port"`
	AppName   string   `json:"app_name"`
	Databases Database `json:"databases"`
	Kafka     Kafka    `json:"kafka"`
	Cluster   Cluster  `json:"cluster"`
}

type Database struct {
	Postgres string `json:"postgres"`
}

type Kafka struct {
	Address             string `json:"address"`
	ProductApproveTopic string `json:"productApproveTopic"`
	SupplyProductsTopic string `json:"supplyProductsTopic"`
}

type Cluster struct {
	WarehouseService string `json:"warehouse_service_url"`
}
