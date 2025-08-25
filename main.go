package main

import (
	apiHandlers "github.com/hippopop/full-stack-todo-w-go/src/api"
	database "github.com/hippopop/full-stack-todo-w-go/src/database"
	apiRouter "github.com/hippopop/full-stack-todo-w-go/src/services"
	appConfig "github.com/hippopop/full-stack-todo-w-go/src/utils/config"
)

func main() {
	// Initiate App Configurations and initiate Database connection!
	config := appConfig.GetAppConfig()
	_ = database.GetCurrentDatabase()

	//* Initiate API Router!
	router := apiRouter.GetAPIRouter()

	//* Attach API Routes!
	apiHandlers.RegisterAuthenticationRoutes(router)

	// Finally run the server!
	router.Run(config.ServerConf.Port.Value)
}
