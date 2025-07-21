package main

import (
	"fmt"
	"task_manager/router"
)


func main() {
	
	router := router.Router()

	router.Run("localhost:8080") // default port is :8080
	fmt.Println("Server listening at : http://localhost:8080")
}
