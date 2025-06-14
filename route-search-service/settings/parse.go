package settings

import "github.com/GOAT-prod/goatsettings"

func Get() (cfg Config, err error) {
	return cfg, goatsettings.ReadConfig(&cfg)
}
