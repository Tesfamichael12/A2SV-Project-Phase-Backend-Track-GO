package controller

import (
	"strconv"
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

// GetAllTasks returns all tasks
func GetAllTasks(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, data.GetAllTasks())
}

// GetTaskByID returns a task by its ID
func GetTaskByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error" : "Invalid ID! Please, use an integer value."})
	}

	if task, found := data.GetTaskByID(id); found {
		ctx.IndentedJSON(http.StatusOK, task)
	} else {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "No task found with that id."})
	}
}

// AddTask adds a new task
func AddTask(ctx *gin.Context) {
	var newTask models.Task

	if err := ctx.BindJSON(&newTask); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Bad Request: " + err.Error()})
		return
	}

	data.AddTask(newTask)
	ctx.IndentedJSON(http.StatusCreated, gin.H{"message": "Successfully Added"})
}

// UpdateTask updates an existing task
func UpdateTask(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error" : "Invalid ID! Please, use an integer value."})
	}

	var updatedTask models.Task

	if err := ctx.BindJSON(&updatedTask); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Bad Request : " + err.Error()})
		return
	}
	
	if found_and_updated := data.UpdateTask(id, updatedTask); found_and_updated {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
	} else {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task was not found"})
	}
}

// DeleteTask deletes a task by its ID
func DeleteTask(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error" : "Invalid ID! Please, use an integer value."})
	}
	
	if found_and_deleted := data.DeleteTask(id); found_and_deleted {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
	} else {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task was not found"})
	}
}
