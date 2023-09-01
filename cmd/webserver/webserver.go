package webserver

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/simonteh28/go-todo-app/api/middleware"
	"github.com/simonteh28/go-todo-app/config"
	todoservice "github.com/simonteh28/go-todo-app/internal/services/todo"
)

// webserver manages internal state
type WebServer struct {
	todoservice.TodoService

	router *gin.Engine // API router
	cfg *config.Config // Config file
}

func New() (*WebServer, error){
	var err error

	// Initialize controller
	ws := &WebServer{}

	// Initialize environment variables from .env file
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading environment variables: %w", err)
	}

	// Initialize config
	ws.cfg, err = config.Get()
	if err != nil {
		return nil, fmt.Errorf("error loading configuration: %w", err)
	}

	// To implement error handling
	ws.TodoService, err = todoservice.New(ws.cfg)
	if err != nil {
		return nil, fmt.Errorf("error initializing service: %w", err)
	}

	return ws, nil
}

func (ws *WebServer) Start(registerRoutes func(s Services, r *gin.Engine)){
	ws.router = gin.Default()

	// Cors config
	config := cors.DefaultConfig();
	config.AllowOrigins = []string{ "http://localhost:4200" }

	// Register error handler
	ws.router.Use(
		middleware.ErrorHandler(),
		cors.New(config),
	)

	registerRoutes(ws, ws.router)

	ws.router.Run()
}