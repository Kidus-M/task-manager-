package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"task_manager/data"
	"task_manager/models"
)

// TaskController handles task-related HTTP requests
type TaskController struct {
	service *data.TaskService
}

// NewTaskController creates a new TaskController instance
func NewTaskController() *TaskController {
	return &TaskController{
		service: data.NewTaskService(),
	}
}

// GetTasks handles GET /tasks
func (tc *TaskController) GetTasks(c *gin.Context) {
	tasks := tc.service.GetAllTasks()
	c.JSON(http.StatusOK, tasks)
}

// GetTaskByID handles GET /tasks/:id
func (tc *TaskController) GetTaskByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, found := tc.service.GetTaskByID(id)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// CreateTask handles POST /tasks
func (tc *TaskController) CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	createdTask := tc.service.CreateTask(newTask)
	c.JSON(http.StatusCreated, createdTask)
}

// UpdateTask handles PUT /tasks/:id
func (tc *TaskController) UpdateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	task, ok := tc.service.UpdateTask(id, updatedTask)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// DeleteTask handles DELETE /tasks/:id
func (tc *TaskController) DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	if ok := tc.service.DeleteTask(id); !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
