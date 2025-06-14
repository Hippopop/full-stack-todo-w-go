package main

import (
	database "github.com/hippopop/full-stack-todo-w-go/src/database"
)

func main() {
	currentDatabase := database.GetCurrentDatabase()
	currentDatabase.DisconnectFromDatabase()
}
