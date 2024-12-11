package settings

type Config struct {
	Port      int      `json:"port"`
	AppName   string   `json:"app_name"`
	Databases Database `json:"databases"`
	Cluster   Cluster  `json:"cluster"`
}

type Database struct {
	Postgres string `json:"postgres"`
	Kafka    Kafka  `json:"kafka"`
}

type Kafka struct {
	Address string `json:"address"`
	Topic   string `json:"producerTopic"`
}

type Cluster struct {
	ClientService string `json:"client_service"`
}
