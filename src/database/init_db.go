package database

import (
	"fmt"
	"sync"

	appConfig "github.com/hippopop/full-stack-todo-w-go/src/utils/config"
	customLog "github.com/hippopop/full-stack-todo-w-go/src/utils/logging"
	mysqlDriver "gorm.io/driver/mysql"
	orm "gorm.io/gorm"
)

type CurrentDatabase struct {
	Database *orm.DB
}

func (super CurrentDatabase) DisconnectFromDatabase() {
	disconnectDB(super.Database)
	customLog.BaseLog(customLog.LogOptions{
		LogOptionsType: customLog.LogInfo,
		Tag:            "init_db.go->CurrentDatabase->DisconnectFromDatabase",
		Msg:            "Disconnected from database successfully!",
	})
}

var (
	once            sync.Once
	currentDatabase *CurrentDatabase
)

func GetCurrentDatabase() *CurrentDatabase {
	once.Do(func() {
		databaseConfig := appConfig.GetAppConfig().DatabaseConf

		currentDatabase = &CurrentDatabase{
			Database: connectDB(databaseConfig),
		}

		customLog.BaseLog(customLog.LogOptions{
			LogOptionsType: customLog.LogInfo,
			Tag:            "init_db.go->ConnectDB",
			Msg:            fmt.Sprintf("Database connection established successfully (%s)!", currentDatabase.Database.Name()),
		})
	})

	return currentDatabase
}

func connectDB(databaseConfig appConfig.DatabaseConfig) *orm.DB {
	databaseConnectionRoute := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		databaseConfig.User.Value,
		databaseConfig.Pass.Value,
		databaseConfig.Host.Value,
		databaseConfig.Port.Value,
		databaseConfig.DatabaseName.Value,
	)

	database, error := orm.Open(mysqlDriver.Open(databaseConnectionRoute), &orm.Config{})
	if error != nil {
		customLog.IfError(
			customLog.LogOptions{LogOptionsType: customLog.LogFatal},
			fmt.Errorf("(init_db.go) Failed to <Open> database: %w", error),
		)
	}

	return database
}

func disconnectDB(database *orm.DB) {
	currentDatabase, error := database.DB()
	if error != nil {
		customLog.IfError(
			customLog.LogOptions{LogOptionsType: customLog.LogFatal},
			fmt.Errorf("(init_db.go) Failed to <Find(DB)> database: %w", error),
		)
	}

	error = currentDatabase.Close()
	if error != nil {
		customLog.IfError(
			customLog.LogOptions{LogOptionsType: customLog.LogFatal},
			fmt.Errorf("(init_db.go) Failed to <Close> database: %w", error),
		)
	}
}
