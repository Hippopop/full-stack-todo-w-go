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

var requiredDatabaseKeys = []EnvKey{
	{"SQLPORT", true},
	{"SQLUSER", true},
	{"SQLPASSWORD", false},
	{"SQLHOST", false},
	{"DATABASENAME", true},
}

func (super *DatabaseConfig) CheckValidity(data map[string]string) (valid bool, err error) {
	log := customLog.LogOptions{Msg: "DatabaseConfig->CheckValidity"}
	return BasicEnvValidationCheck(log, data, []CustomKeyValidation{}, requiredDatabaseKeys...)
}

func (super *DatabaseConfig) Initialize(data map[string]string) (err error) {
	if valid, error := super.CheckValidity(data); !valid {
		err = error
		return
	}

	super.Port = EnvData{Key: requiredDatabaseKeys[0], Value: data[requiredDatabaseKeys[0].Key]}
	super.User = EnvData{Key: requiredDatabaseKeys[1], Value: data[requiredDatabaseKeys[1].Key]}
	super.Pass = EnvData{Key: requiredDatabaseKeys[2], Value: data[requiredDatabaseKeys[2].Key]}
	super.Host = EnvData{Key: requiredDatabaseKeys[3], Value: data[requiredDatabaseKeys[3].Key]}
	super.DatabaseName = EnvData{Key: requiredDatabaseKeys[4], Value: data[requiredDatabaseKeys[4].Key]}
	return
}
