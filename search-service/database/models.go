package database

type Filter struct {
	Name          string   `bson:"name"`
	AllowedValues []string `bson:"allowedValues"`
}
