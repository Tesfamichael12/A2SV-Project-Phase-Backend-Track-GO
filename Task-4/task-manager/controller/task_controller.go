package controller

import (
	"net/http"
	"strconv"
	"strings"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

/*  This controller is clean and uses centralized helper/middleware functions to validate and clean up data,
    thereby following the DRY (Don't Repeat Yourself) principle and ensuring maintainable, consistent request handling.*/

/* HTTP Request Handlers/logic */

var validate = validator.New()

// GetAllTasks returns all tasks
func GetAllTasks(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, data.GetAllTasks())
}

// GetTaskByID returns a task by its ID
func GetTaskByID(ctx *gin.Context) {
	id, validID := validateID(ctx)
	if !validID {return}

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

	if !validateAndCleanupTask(ctx, &newTask) {
		return
	}

	data.AddTask(newTask)
	ctx.IndentedJSON(http.StatusCreated, gin.H{"message": "Successfully Added"})
}

// UpdateTask updates an existing task
func UpdateTask(ctx *gin.Context) {
	id, validID := validateID(ctx)
	if !validID {return}

	var updatedTask models.Task
	if err := ctx.BindJSON(&updatedTask); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Bad Request : " + err.Error()})
		return
	}

	if !validateAndCleanupTask(ctx, &updatedTask) {
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
	id, validID := validateID(ctx)
	if !validID {return}
	
	if found_and_deleted := data.DeleteTask(id); found_and_deleted {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
	} else {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task was not found"})
	}
}

// validateAndCleanupTask is a middleware that trims whitespace and validates the struct
func validateAndCleanupTask(ctx *gin.Context, task *models.Task) bool {

	// Trim trailing white spaces
	task.Title = strings.TrimSpace(task.Title)
	task.Description = strings.TrimSpace(task.Description)
	task.Status = strings.TrimSpace(task.Status)

	// validate the struct
	if err := validate.Struct(*task); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Bad request: " + err.Error()})
		return false
	}

	return true
}

// validateID is a middleware that checks if the task ID parameter is a valid integer.
func validateID(ctx *gin.Context) (int, bool) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error" : "Invalid ID! Please, use an integer value."})
		return 0, false
	}

	return id, true
}
