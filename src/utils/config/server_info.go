package config

import (
	customLog "github.com/hippopop/full-stack-todo-w-go/src/utils/logging"
)

type ServerConfig struct {
	Port         EnvData
	AccessToken  EnvData
	RefreshToken EnvData
}

var requiredServerKeys = []EnvKey{
	{"PORT", true},
	{"ACCESS_TOKEN", true},
	{"REFRESH_TOKEN", false},
}

func validatePortValue(data map[string]string) (valid bool, err error) {

	return
}

func (super *ServerConfig) CheckValidity(data map[string]string) (valid bool, err error) {
	log := customLog.LogOptions{Msg: "ServerConfig->CheckValidity"}
	return BasicEnvValidationCheck(
		log,
		data,
		[]CustomKeyValidation{
			{Key: requiredServerKeys[0], Validate: validatePortValue},
		},
		requiredServerKeys...,
	)
}

func (super *ServerConfig) Initialize(data map[string]string) (err error) {
	if valid, error := super.CheckValidity(data); !valid {
		err = error
		return
	}

	super.Port = EnvData{Key: requiredDatabaseKeys[0], Value: data[requiredDatabaseKeys[0].Key]}
	super.AccessToken = EnvData{Key: requiredDatabaseKeys[1], Value: data[requiredDatabaseKeys[1].Key]}
	super.RefreshToken = EnvData{Key: requiredDatabaseKeys[2], Value: data[requiredDatabaseKeys[2].Key]}
	return
}
