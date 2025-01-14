package domain

type Filter struct {
	Name          string   `json:"name"`
	AllowedValues []string `json:"allowedValues"`
}
