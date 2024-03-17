package app

import (
	"database/sql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Server struct {
	Engine *gin.Engine
	DB     *sql.DB
}

func NewServer() *Server {
	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"POST", "DELETE", "GET", "OPTIONS", "PUT", "PATCH"},
		AllowHeaders:    []string{"Authorization", "Content-Type"},
	}))
	s := &Server{engine, &sql.DB{}}
	s.RegisterRoutes()

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) { s.Engine.ServeHTTP(w, r) }
func (s *Server) Serve(addr string)                                { log.Fatal(s.Engine.Run(addr)) }
