package domain

const EarthRadiusKm = 6371.0

type ServiceLocation struct {
	ClientId int    `json:"client_id"`
	Address  string `json:"address"`
}

type Location struct {
	ID  string  `json:"id"`
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Route struct {
	Path          []Location `json:"path"`
	TotalDistance float64    `json:"total_distance"`
}
