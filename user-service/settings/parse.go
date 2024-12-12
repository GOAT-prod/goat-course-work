package settings

import (
	"github.com/GOAT-prod/goatsettings"
)

func Parse() (cfg Config, err error) {
	return cfg, goatsettings.ReadConfig(&cfg)
}
