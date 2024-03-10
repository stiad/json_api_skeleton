package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleHelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "Hello World"})
}

func HandleProtectedHelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "Hello World"})
}
