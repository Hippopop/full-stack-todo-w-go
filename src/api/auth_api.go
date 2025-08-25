package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hippopop/full-stack-todo-w-go/src/api/authentication"
)

func RegisterAuthenticationRoutes(router *gin.Engine) {
	authRouter := router.Group("/auth")

	authRouter.POST("/login", authentication.HandleLogin)
}
