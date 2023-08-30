package todoapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errormsg "github.com/simonteh28/go-todo-app/api/error"
	"github.com/simonteh28/go-todo-app/cmd/webserver"
)

// Get all todo
func Get(srv webserver.Services) gin.HandlerFunc {
	return gin.HandlerFunc(func (c *gin.Context){
		
		todos, err := srv.Get()
		if err != nil {
			cerr := errormsg.NewErrorMessage(http.StatusBadRequest, "Could not fetch all todos", err.Error())
			c.Error(cerr)
			return
		}		
		
		c.IndentedJSON(http.StatusOK, todos)
	})
}