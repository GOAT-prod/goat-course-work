package settings

import "github.com/GOAT-prod/goatsettings"

func Parse() (cfg Settings, err error) {
	return cfg, goatsettings.ReadConfig(&cfg)
}
