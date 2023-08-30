package webserver

import (
	todoservice "github.com/simonteh28/go-todo-app/internal/services/todo"
)

// Exposes all todo requests
type Services interface {
	todoservice.TodoService
}
