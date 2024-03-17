package app

func (s *Server) RegisterRoutes() {
	// health check
	s.Engine.GET("/health", s.handleHealthCheck)

	// v1
	v1g := s.Engine.Group("/v1")
	v1g.GET("/hello", s.handleHelloWorld)

	// v1/protected
	protect := v1g.Group("/protected").Use(s.middlewareApiKeyAuth)
	protect.GET("/hello", s.handleHelloWorld)
}
