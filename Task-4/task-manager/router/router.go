package router

import (
	"task_manager/controller"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
    router := gin.Default()

    router.GET("/tasks", controller.GetAllTasks)
    router.GET("/tasks/:id", controller.GetTaskByID)
    router.POST("/tasks", controller.AddTask)
    router.PUT("/tasks/:id", controller.UpdateTask)
    router.DELETE("/tasks/:id", controller.DeleteTask)

    return router
}