package main

import (
	"fmt"
	"log"
	"os"

	"github.com/simonteh28/go-todo-app/cmd/webserver"
)

func main(){
	
	log.Println("Go Todo App start")

	wb, err := webserver.New()

	if err != nil {
		fmt.Printf("Invalid configuration: %s\n", err)
		os.Exit(1)
	}

	wb.Start(BuildRoutes)
}

// To do
// 1. Need to standardize error messages. Use log fatal
// 2. Deployment
