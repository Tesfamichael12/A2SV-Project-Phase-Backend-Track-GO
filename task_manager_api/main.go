package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Task struct {
 ID          string    `json:"id"`
 Title       string    `json:"title"`
 Description string    `json:"description"`
 DueDate     time.Time `json:"due_date"`
 Status      string    `json:"status"`
}

var tasks = []Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
    {ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
    {ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

func main() {
	router := gin.Default()

	router.GET("/tasks", getAllTasks)
	router.GET("/tasks/:id", getTaskByID)
	router.POST("/create-task", addTask)
	router.PUT("/update-task/:id", updateTask)
	router.DELETE("/delete-task/:id", deleteTask)
	

	router.Run("localhost:8080") // default port is :8080
	fmt.Println("Server listening at : http://localhost:8080")
}

// getAllTasks a handler to list all the available tasks
func getAllTasks(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, tasks)
}

// getTaskByID a handler to get a task by id
func getTaskByID(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, task := range tasks {
		if task.ID == id {
			ctx.IndentedJSON(http.StatusOK, task)
			return
		}
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message" : "No task found with that id."})
}

// addTask a handler to add new tasks 
func addTask(ctx *gin.Context) {
	var newTask Task

	if err := ctx.BindJSON(&newTask); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Bad Request: " + err.Error()})
	}

	tasks = append(tasks, newTask)
	ctx.IndentedJSON(http.StatusOK, gin.H{"message" : "Successfully Added"})
}

// updateTask a handler to update a specific task
func updateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedTask Task

	if err := ctx.BindJSON(&updatedTask); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Bad Request : " + err.Error()})
		return
	}

	for i, task := range tasks	 {
		if task.ID == id {

			if updatedTask.Title != "" {
				tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				tasks[i].Description = updatedTask.Description
			}
			if !updatedTask.DueDate.IsZero() {
				tasks[i].DueDate = updatedTask.DueDate
			}
			if updatedTask.Status != "" {
				tasks[i].Status = updatedTask.Status
			}
			ctx.IndentedJSON(http.StatusOK, gin.H{"message" : "Task updated successfully"})
			return
		}
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{"error" : "Task was not found"})
}

// deleteTask a handler to delete a specific task by id
func deleteTask(ctx * gin.Context) {

	id := ctx.Param("id")

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
			return 
		}
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Task was not found"})
	
}