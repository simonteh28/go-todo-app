package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	errormsg "github.com/simonteh28/go-todo-app/api/error"
)

func ErrorHandler()  gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			log.Println(err.Error())
			switch e := err.Err.(type) {
				case errormsg.CustomError: 
					c.AbortWithStatusJSON(e.Code, e)
				default:
					c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"message": "Service Unavailable"})
			}
		
		}
	}
}