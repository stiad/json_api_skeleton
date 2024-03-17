package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Server struct {
	engine *gin.Engine
	// add storage interface here if needed
}

func NewServer() *Server {
	engine := gin.Default()
	engine.Use(cors.Default())
	s := &Server{engine}
	s.RegisterRoutes()

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) { s.engine.ServeHTTP(w, r) }
func (s *Server) Serve(addr string)                                { log.Fatal(s.engine.Run(addr)) }
