package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) handleHealthCheck(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func (s *Server) handleHelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "Hello World"})
}
