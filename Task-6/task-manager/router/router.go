package router

import (
	"task_manager/controllers"
	"task_manager/middleware"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	// Public routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Protected task routes
	task := r.Group("/tasks")
	task.Use(middleware.AuthMiddleware())
	{
		task.GET("", controllers.GetAllTasks)
		task.GET(":id", controllers.GetTaskByID)
		task.POST("", controllers.AddTask)
		task.PUT(":id", controllers.UpdateTask)
		task.DELETE(":id", controllers.DeleteTask)
	}

	return r
}