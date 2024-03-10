package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stiad/json_api_skeleton/src/app/handlers/v1"
	"github.com/stiad/json_api_skeleton/src/app/middleware"
)

func RegisterV1(router *gin.Engine) {
	// v1
	v1g := router.Group("/v1")
	v1g.GET("/hello", v1.HandleHelloWorld)
	v1g.GET("/protected/hello", middleware.ApiKeyAuth(), v1.HandleProtectedHelloWorld)
}
