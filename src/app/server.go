package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/stiad/json_api_skeleton/src/app/routes"
	"log"
	"net/http"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"POST", "DELETE", "GET", "OPTIONS", "PUT", "PATCH"},
		AllowHeaders:    []string{"Authorization", "Content-Type"},
	}))
	router.GET("/health", func(c *gin.Context) { c.Status(http.StatusNoContent) })

	s := Server{router}
	routes.Register(s.router)

	return &s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) Serve(addr string) {
	log.Fatal(s.router.Run(addr))
}
