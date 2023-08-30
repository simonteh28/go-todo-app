package todoapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errormsg "github.com/simonteh28/go-todo-app/api/error"
	"github.com/simonteh28/go-todo-app/cmd/webserver"
)

// Delete specific todo
func Delete(srv webserver.Services) gin.HandlerFunc {
	return gin.HandlerFunc(func (c *gin.Context){
		id := c.Param("id")

		err := srv.Delete(id)
		if err != nil {
			cerr := errormsg.NewErrorMessage(http.StatusBadRequest, "Unable to delete todo", err.Error())
			c.Error(cerr)
			return
		}

		c.JSON(http.StatusAccepted, "Deleted successfully")
	})
}