package authentication

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/hippopop/full-stack-todo-w-go/src/utils/logging"
)

func asStringJson(super *gin.Context) string {
	var value string
	json, error := io.ReadAll(super.Request.Body)
	if error == nil {
		value = string(json)
	}
	return value
}

func asMapJson(super *gin.Context) map[string]any {
	value := make(map[string]any)
	binary, error := io.ReadAll(super.Request.Body)
	if error == nil {
		json.Unmarshal(binary, &value)
	}
	return value
}

func HandleLogin(context *gin.Context) {
	// Extract login credentials from the request
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// print the request body for debugging purposes
	logging.BaseLog(logging.LogOptions{
		Msg: fmt.Sprintf("Request Body String: %s", asStringJson(context)),
	})
	logging.BaseLog(logging.LogOptions{
		Msg: fmt.Sprintf("Request Body Map: %s", asMapJson(context)),
	})

	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	context.JSON(200, gin.H{"message": "Login successful", "user": credentials.Username})

}
