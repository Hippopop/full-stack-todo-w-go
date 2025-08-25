package config

import (
	"sync"

	customLog "github.com/hippopop/full-stack-todo-w-go/src/utils/logging"
)

type ApplicationConfig struct {
	DatabaseConf DatabaseConfig
	ServerConf   ServerConfig
}

var once sync.Once
var config *ApplicationConfig

func GetAppConfig() *ApplicationConfig {
	once.Do(
		func() {
			config = &ApplicationConfig{}
			LoadEnv(
				customLog.LogOptions{LogOptionsType: customLog.LogFatal, Msg: "ApplicationConfig->GetAppConfig"},
				&config.DatabaseConf, &config.ServerConf,
			)
		},
	)
	return config
}
