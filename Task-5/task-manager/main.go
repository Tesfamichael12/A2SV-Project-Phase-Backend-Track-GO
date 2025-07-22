package main

import (
	"fmt"
	"log"
	"task_manager/data"
	"task_manager/router"
)

func main() {
	if err := data.InitDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	router := router.Router()

	fmt.Println("Server listening at : http://localhost:8080")
	router.Run("localhost:8080") // default port is :8080
}
