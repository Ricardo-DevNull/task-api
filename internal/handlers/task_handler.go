package handlers

import (
	"net/http"

	"github.com/Ricardo-DevNull/task-api/internal/models/dtos"
	"github.com/Ricardo-DevNull/task-api/internal/models/enums"
	"github.com/Ricardo-DevNull/task-api/internal/services"
	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	TaskService services.TaskService
}

func NewTaskService(taskService services.TaskService) *TaskHandler {
	return &TaskHandler{TaskService: taskService}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var newTask dtos.TaskRequest

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := h.TaskService.CreateTask(&newTask)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	response := &dtos.TaskResponse{
		BaseModel:   task.BaseModel,
		Title:       task.Title,
		Description: task.Description,
		Status:      enums.StatusAvailable,
	}

	c.JSON(http.StatusCreated, response)
}
