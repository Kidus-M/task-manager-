package router

import (
	"github.com/gin-gonic/gin"
	"task_manager/controllers"
)

// SetupRouter initializes and returns a Gin router with all routes configured
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Create a task controller instance
	taskController := controllers.NewTaskController()

	// Define routes for task management
	r.GET("/tasks", taskController.GetTasks)
	r.GET("/tasks/:id", taskController.GetTaskByID)
	r.POST("/tasks", taskController.CreateTask)
	r.PUT("/tasks/:id", taskController.UpdateTask)
	r.DELETE("/tasks/:id", taskController.DeleteTask)

	return r
}
