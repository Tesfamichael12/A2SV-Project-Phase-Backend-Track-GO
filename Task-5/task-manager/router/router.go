package router

import (
    "task_manager/controller"

    "github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
    router := gin.Default()

    taskRoutes := router.Group("/tasks")
    {
        taskRoutes.GET("", controller.GetAllTasks)
        taskRoutes.POST("", controller.AddTask)
        taskRoutes.GET("/:id", controller.GetTaskByID)
        taskRoutes.PUT("/:id", controller.UpdateTask)
        taskRoutes.DELETE("/:id", controller.DeleteTask)
    }

    return router
}