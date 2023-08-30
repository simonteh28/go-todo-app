package todoapi

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	tododtos "github.com/simonteh28/go-todo-app/api/dtos/todo"
	errormsg "github.com/simonteh28/go-todo-app/api/error"
	"github.com/simonteh28/go-todo-app/cmd/webserver"
)

// Create todo
func Post(srv webserver.Services) gin.HandlerFunc {
	return gin.HandlerFunc(func (c *gin.Context){
		var todo tododtos.Todo

    	// Call BindJSON to bind the received JSON to
		if err := c.ShouldBindJSON(&todo); err != nil {
        	cerr := errormsg.NewErrorMessage(http.StatusBadRequest, "Could not bind to JSON", err.Error())
			c.Error(cerr)
			return
   		}

		id, err := srv.Post(todo)
		if err != nil {
			cerr := errormsg.NewErrorMessage(http.StatusBadRequest, "Could not save data", err.Error())
			c.Error(cerr)
			return
		}

		c.IndentedJSON(http.StatusCreated, "Todo created with Id: " + strconv.FormatInt(id, 10))
	})
}