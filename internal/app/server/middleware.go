package server

import (
	"context"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Middleware middleware
type Middleware struct {
	// logger log.Logger
}

func (s *Server) setupMiddlewares(ctx context.Context) {
	// middleware setup
	var middleware Middleware
	s.routers.Use(middleware.Context(ctx))
	s.routers.Use(middleware.CORS())
	s.routers.Use(middleware.RequestID())
	// s.routers.Use(middleware.Authorization())
}

// CORS cors middleware
func (m *Middleware) CORS() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	return cors.New(config)
}

// Context injector context
func (m *Middleware) Context(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// RequestID inject requestID(uuid)
func (m *Middleware) RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {

		uuid, err := uuid.NewUUID()
		if err != nil {

			c.Abort()
		}

		requestID := uuid.String()
		ctx := context.WithValue(c.Request.Context(), ContextRequestID, requestID)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

// Authorization inject Authorization(jwtToken)
// func (m *Middleware) Authorization() gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 		reqToken := c.GetHeader("Authorization")
// 		if len(reqToken) >= 1 {
// 			if strings.Index(reqToken, "Bearer") == -1 {
// 				return
// 			}

// 			splitToken := strings.Split(reqToken, "Bearer")

// 			if len(splitToken) != 2 {
// 				c.AbortWithStatus(http.StatusBadRequest)
// 				return
// 			}

// 			token := strings.TrimSpace(splitToken[1])
// 			ctx := context.WithValue(c.Request.Context(), ContextAuthorization, token)
// 			c.Request = c.Request.WithContext(ctx)
// 		}

// 		c.Next()
// 	}
// }
