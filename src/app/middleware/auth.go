package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func ApiKeyAuth() gin.HandlerFunc {
	if viper.GetString("API_KEY") == "" {
		log.Println("WARNING: API_KEY is not configured, API key authentication is disabled.")
	}
	return func(c *gin.Context) {
		if viper.GetString("API_KEY") == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "api authentication is not configured"})
			return
		}
		if c.Request.Header.Get("x-api-key") == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "api authentication failed, authentication header (x-api-key) is missing"})
			return
		}
		if c.Request.Header.Get("x-api-key") != viper.GetString("API_KEY") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "api authentication failed"})
			return
		}
		c.Next()
	}
}
