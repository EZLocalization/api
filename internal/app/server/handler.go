package server

import (
	"net/http"

	"github.com/ez-api/internal/app/server/api"

	"github.com/gin-gonic/gin"
)

// Handle handle
type Handle struct {
	// DataSource *datasource.DataSource
	User api.User
}

// OnHealthCheck health check handle
func (h *Handle) OnHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, "liveness")
}
