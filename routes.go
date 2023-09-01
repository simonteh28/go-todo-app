package main

import (
	"github.com/gin-gonic/gin"

	todoapi "github.com/simonteh28/go-todo-app/api/routes/todo"
	"github.com/simonteh28/go-todo-app/cmd/webserver"
)

// Build routes
func BuildRoutes(srv webserver.Services, r *gin.Engine) {
	r.GET("/ping", todoapi.Ping(srv))

	// Todo API
	r.GET("/todo/:id", todoapi.GetById(srv))
	r.POST("/todo", todoapi.Post(srv))
	r.GET("/todo",  todoapi.Get(srv))
	r.PATCH("/todo/:id", todoapi.Patch(srv))
	r.DELETE("/todo/:id", todoapi.Delete(srv))
}