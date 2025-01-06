package settings

type Settings struct {
	Port            int    `json:"port"`
	AppName         string `json:"app_name"`
	SmtpCredentials Smtp   `json:"smtp"`
}

type Smtp struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	From     string `json:"from"`
	To       string `json:"to"`
	Password string `json:"password"`
}
