package todoapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errormsg "github.com/simonteh28/go-todo-app/api/error"
	"github.com/simonteh28/go-todo-app/cmd/webserver"
)

// Get specific todo
func GetById(srv webserver.Services) gin.HandlerFunc {
	return gin.HandlerFunc(func (c *gin.Context){
		id := c.Param("id")

		todo, err := srv.GetById(id)
		if err != nil {
			cerr := errormsg.NewErrorMessage(http.StatusBadRequest, "Could not fetch todo with id: " + id, err.Error())
			c.Error(cerr)
			return
		}		
		
		c.IndentedJSON(http.StatusOK, todo)
	})
}