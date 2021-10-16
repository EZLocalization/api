package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// User is for the User api
type User struct {
	// DataSource *datasource.DataSource
	Name string `form:"name" json:"name"  binding:"required"`
}

// GetUser return user info
func (h *User) GetUser(c *gin.Context) {
	name := c.Param("name")
	id := c.Query("id")
	c.JSON(http.StatusOK, fmt.Sprintf("getUsers: id:%s name: %s", id, name))
}

// GetUser return user info
func (h *User) AddUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("getUsers: name: %s", user.Name))
}
