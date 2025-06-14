package settings

type Config struct {
	Port       int    `json:"port"`
	AppName    string `json:"app_name"`
	MapsApiUrl string `json:"maps_api_url"`
	MapsApiKey string `json:"maps_api_key"`
}
