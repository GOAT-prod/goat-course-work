package settings

type Config struct {
	Port    int     `json:"port"`
	AppName string  `json:"app_name"`
	Cluster Cluster `json:"cluster"`
}

type Cluster struct {
	AuthService      string `json:"auth_service_url"`
	UserService      string `json:"user_service_url"`
	ClientService    string `json:"client_service_url"`
	WareHouseService string `json:"warehouse_service_url"`
}
