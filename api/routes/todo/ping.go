package todoapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/simonteh28/go-todo-app/cmd/webserver"
)

// Ping to ensure routes are available
func Ping(srv webserver.Services) gin.HandlerFunc {
	return gin.HandlerFunc(func (c *gin.Context) {
		c.JSON(http.StatusOK, "Pong")
	})
}