package config

import (
	customLog "github.com/hippopop/full-stack-todo-w-go/src/utils/logging"
)

type DatabaseConfig struct {
	Port         EnvData
	User         EnvData
	Pass         EnvData
	Host         EnvData
	DatabaseName EnvData
}

var requiredKeys = []EnvKey{
	{"SQLPORT", true},
	{"SQLUSER", true},
	{"SQLPASSWORD", false},
	{"SQLHOST", false},
	{"DATABASENAME", true},
}

func (super *DatabaseConfig) CheckValidity(data map[string]string) (valid bool, err error) {
	log := customLog.LogOptions{Msg: "DatabaseConfig->CheckValidity"}
	return BasicEnvValidationCheck(log, data, requiredKeys...)
}

func (super *DatabaseConfig) Initialize(data map[string]string) (err error) {
	if valid, error := super.CheckValidity(data); !valid {
		err = error
		return
	}

	super.Port = EnvData{Key: requiredKeys[0], Value: data[requiredKeys[0].Key]}
	super.User = EnvData{Key: requiredKeys[1], Value: data[requiredKeys[1].Key]}
	super.Pass = EnvData{Key: requiredKeys[2], Value: data[requiredKeys[2].Key]}
	super.Host = EnvData{Key: requiredKeys[3], Value: data[requiredKeys[3].Key]}
	super.DatabaseName = EnvData{Key: requiredKeys[4], Value: data[requiredKeys[4].Key]}
	return
}
