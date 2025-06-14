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
	CartService      string `json:"cart_service_url"`
	OrderService     string `json:"order_service_url"`
	SearchService    string `json:"search_service_url"`
	RequestService   string `json:"request_service_url"`
	ReportService    string `json:"report_service_url"`
	RouteService     string `json:"route_service_url"`
}
