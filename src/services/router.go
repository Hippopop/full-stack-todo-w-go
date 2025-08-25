package services

import (
	server "github.com/gin-gonic/gin"
)

func GetAPIRouter() (router *server.Engine) {
	router = server.Default()
	router.Static("/public", "")

	router.GET("/ping", func(context *server.Context) {
		context.String(200, "pong")
	})

	return
}
