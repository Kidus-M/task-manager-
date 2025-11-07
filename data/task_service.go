package data

import (
	"sync"
	"task_manager/models"
)

// TaskService provides in-memory task management functionality
type TaskService struct {
	mu    sync.Mutex
	tasks map[int]models.Task
	nextID int
}

// NewTaskService initializes and returns a new TaskService
func NewTaskService() *TaskService {
	return &TaskService{
		tasks: make(map[int]models.Task),
		nextID: 1,
	}
}

// GetAllTasks returns all stored tasks
func (s *TaskService) GetAllTasks() []models.Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	tasks := make([]models.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

// GetTaskByID returns a task by ID
func (s *TaskService) GetTaskByID(id int) (models.Task, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, exists := s.tasks[id]
	return task, exists
}

// CreateTask creates a new task
func (s *TaskService) CreateTask(task models.Task) models.Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	task.ID = s.nextID
	s.nextID++
	s.tasks[task.ID] = task
	return task
}

// UpdateTask updates a task if it exists
func (s *TaskService) UpdateTask(id int, updated models.Task) (models.Task, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, exists := s.tasks[id]
	if !exists {
		return models.Task{}, false
	}

	// Update fields
	if updated.Title != "" {
		task.Title = updated.Title
	}
	if updated.Description != "" {
		task.Description = updated.Description
	}
	if updated.DueDate != "" {
		task.DueDate = updated.DueDate
	}
	if updated.Status != "" {
		task.Status = updated.Status
	}

	s.tasks[id] = task
	return task, true
}

// DeleteTask removes a task by ID
func (s *TaskService) DeleteTask(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[id]; !exists {
		return false
	}
	delete(s.tasks, id)
	return true
}
