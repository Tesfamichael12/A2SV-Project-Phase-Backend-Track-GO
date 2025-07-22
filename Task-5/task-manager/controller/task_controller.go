package controller

import (
	"net/http"
	"strings"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*  This controller is clean and uses centralized helper/middleware functions to validate and clean up data,
    thereby following the DRY (Don't Repeat Yourself) principle and ensuring maintainable, consistent request handling.*/

/* HTTP Request Handlers/logic */

var validate = validator.New()

// bindAndValidateTask binds the JSON, cleans up the data, and validates the task.
func bindAndValidateTask(c *gin.Context) (models.Task, error) {
	var task models.Task

	// bind the incoming JSON to the task struct
	if err := c.ShouldBindJSON(&task); err != nil {
		return models.Task{}, err
	}

	task.Title = strings.TrimSpace(task.Title)
	task.Description = strings.TrimSpace(task.Description)
	task.Status = strings.TrimSpace(task.Status)

	if err := validate.Struct(task); err != nil {
		return models.Task{}, err
	}

	return task, nil
}

// GetAllTasks returns all tasks
func GetAllTasks(c *gin.Context) {
	tasks, err := data.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// GetTaskByID returns a single task by its ID
func GetTaskByID(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	task, err := data.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// AddTask adds a new task
func AddTask(c *gin.Context) {
	task, err := bindAndValidateTask(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	insertedID, err := data.AddTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add task"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": insertedID})
}

// UpdateTask updates an existing task
func UpdateTask(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	task, err := bindAndValidateTask(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	modifiedCount, err := data.UpdateTask(id, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	if modifiedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found or not modified"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

// DeleteTask deletes a task
func DeleteTask(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	deletedCount, err := data.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	if deletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
