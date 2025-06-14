package goatsettings

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	_appName    = "storage-service"
	_configPath = ".config"

	_envKey = "ENV"

	_prod  ENV = "prod"
	_local ENV = "local"
	_dev   ENV = "dev"
)

type ENV string

func AppName() string {
	return _appName
}

func ReadConfig(settings any) error {
	path := fmt.Sprintf("%s/%s.json", _configPath, GetEnv())
	configBytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(configBytes, &settings)
}

func GetEnv() ENV {
	env := os.Getenv(_envKey)
	if env == "" {
		return _local
	}

	return ENV(env)
}

func LocalEnv() ENV {
	return _local
}
