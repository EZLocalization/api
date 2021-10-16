package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

// Configruation server configuration structure
type Configruation struct {
	Port int
}

// Server app for restful serving
type Server struct {
	Port int

	httpServer *http.Server
	routers    *gin.Engine
	logger     log.Logger

	runningMutex sync.RWMutex
}

// New create http app instance
func New(ctx context.Context, cfg *Configruation) (*Server, error) {

	// set gin values
	gin.SetMode("release")

	// create server
	server := &Server{
		Port:    cfg.Port,
		routers: gin.New(),
	}

	// load datasource
	// datasource, err := datasource.New(ctx, cfg.DataSourceConfig)
	// if err != nil {
	// 	logger.Error(err)
	// 	return nil, err
	// }

	// setup middleware
	server.setupMiddlewares(ctx)

	// setup routers
	if err := server.setupRouters(ctx); err != nil {
		return nil, err
	}

	// logger.Debug("create http server instance")
	return server, nil
}

// Start server start
func (s *Server) Start(ctx context.Context) {
	s.runningMutex.Lock()
	defer s.runningMutex.Unlock()

	s.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%v", s.Port),
		Handler: s.routers,
	}

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {

		}
	}()

	fmt.Printf("listening http://0.0.0.0:%v", s.Port)
}

// Stop server stop
func (s *Server) Stop(ctx context.Context) {
	if s.httpServer != nil {
		s.runningMutex.Lock()

		s.httpServer.RegisterOnShutdown(func() {
			s.runningMutex.Unlock()
			// s.logger.Infof("stop server")
		})
		s.httpServer.Shutdown(ctx)
	}
}
