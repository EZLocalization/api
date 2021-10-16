package server

import (
	"context"
)

// setupRouters
func (s *Server) setupRouters(ctx context.Context) error {
	// routing path setup
	handle := Handle{
		// DataSource: datasource,
	}
	public := s.routers.Group("/")
	public.Any("/liveness", handle.OnHealthCheck)
	public.GET("/user/:name", handle.User.GetUser)
	public.POST("/user", handle.User.AddUser)
	return nil
}
