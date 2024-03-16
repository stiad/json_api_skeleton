package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stiad/json_api_skeleton/src/app/handlers"
	"github.com/stiad/json_api_skeleton/src/app/middleware"
)

func Register(router *gin.Engine) {
	// v1/
	v1g := router.Group("/v1")
	v1g.GET("/hello", handlers.HandleHelloWorld)

	// v1/protected
	protect := v1g.Group("/protected").Use(middleware.ApiKeyAuth())
	protect.GET("/hello", handlers.HandleHelloWorld)
}
