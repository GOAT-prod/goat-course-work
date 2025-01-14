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
	NotifierService string `json:"notifier_service_url"`
}
