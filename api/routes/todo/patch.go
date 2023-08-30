package todoapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	tododtos "github.com/simonteh28/go-todo-app/api/dtos/todo"
	errormsg "github.com/simonteh28/go-todo-app/api/error"
	"github.com/simonteh28/go-todo-app/cmd/webserver"
)

// Update specific todo
func Patch(srv webserver.Services) gin.HandlerFunc {
	return gin.HandlerFunc(func (c *gin.Context){
		var updateTodo tododtos.UpdateTodo
		id := c.Param("id")

		if err := c.ShouldBindJSON(&updateTodo); err != nil {
			cerr := errormsg.NewErrorMessage(http.StatusBadRequest, "Could not bind JSON", err.Error())
			c.Error(cerr)
			return
		}

		todo, err := srv.Patch(id, updateTodo)
		if err != nil {
			cerr := errormsg.NewErrorMessage(http.StatusBadRequest, "Could not patch todo", err.Error())
			c.Error(cerr)
			return
		}
	
		c.IndentedJSON(http.StatusOK, todo)
	})
}