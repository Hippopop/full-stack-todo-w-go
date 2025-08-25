package config

import (
	"fmt"
	"reflect"

	customLog "github.com/hippopop/full-stack-todo-w-go/src/utils/logging"
	envHandler "github.com/joho/godotenv"
)

type EnvKey struct {
	Key      string
	NonEmpty bool
}

func (super EnvKey) IsValid() bool {
	return super.Key != ""
}

type EnvData struct {
	Key   EnvKey
	Value string
}

type ValidateEnv func(data map[string]string) (valid bool, err error)
type CustomKeyValidation struct {
	Key      EnvKey
	Validate ValidateEnv
}

// Implement this only on <pointer> types!
type ParsableEnvData interface {
	Initialize(data map[string]string) error
	CheckValidity(data map[string]string) (valid bool, err error)
}

func BasicEnvValidationCheck(log customLog.LogOptions, data map[string]string, customValidations []CustomKeyValidation, keys ...EnvKey) (valid bool, err error) {
	valid = true
	for _, key := range keys {
		if !key.IsValid() {
			valid = false
			err = fmt.Errorf("env_config.go: (%s) Invalid/Empty <EnvKey> provided: <%v>", log.Msg, key)
			return
		}

		value, exists := data[key.Key]
		if !exists {
			valid = false
			err = fmt.Errorf("env_config.go: (%s) Missing required environment variable: <%s>", log.Msg, key.Key)
			return
		}

		if key.NonEmpty && value == "" {
			valid = false
			err = fmt.Errorf("env_config.go: (%s) Environment variable cannot be empty: <%s>", log.Msg, key.Key)
			return
		}
	}
	return
}

// LoadEnv loads environment variables from a .env file.
// If loading fails, the returned error includes the filename and a custom message.
func LoadEnv(log customLog.LogOptions, parsable ...ParsableEnvData) (err error) {
	err = envHandler.Load()
	if customLog.IfError(log.UpdateShow(true), err) {
		err = fmt.Errorf("env_config.go: Failed to <Load> file (.env): %w", err)
		return
	}

	envData, err := envHandler.Read()
	if customLog.IfError(log.UpdateShow(true), err) {
		err = fmt.Errorf("env_config.go: Failed to <Read> file (.env): %w", err)
		return
	}

	for _, dataStruct := range parsable {
		// Ensure <pointer> type!
		if reflect.TypeOf(dataStruct).Kind() != reflect.Ptr {
			// If not a pointer, Stop the system with the following msg!
			customLog.IfError(
				log.UpdateShow(true).UpdateType(customLog.LogFatal),
				fmt.Errorf("env_config.go: Expected <pointer> type for <ParsableEnvData>, instead got %T", dataStruct),
			)
			return
		}
		// Purse and Initiate!
		pointerName := reflect.TypeOf(dataStruct).Elem().Name()
		isValid, validationError := dataStruct.CheckValidity(envData)
		isError := customLog.IfError(log.UpdateShow(true), validationError)
		if isError || !isValid {
			err = fmt.Errorf("env_config.go: Data validation failed for <%s> (.env): %w", pointerName, validationError)
			return
		} else {
			// Initiate the Data!
			initError := dataStruct.Initialize(envData)
			if initError != nil {
				customLog.IfError(
					log.UpdateShow(true).UpdateType(customLog.LogFatal),
					fmt.Errorf("env_config.go: Initialization Error of <%s>: %w", pointerName, validationError),
				)
			}
		}
	}

	return
}
