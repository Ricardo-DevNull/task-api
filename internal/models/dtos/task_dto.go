package dtos

import (
	"github.com/Ricardo-DevNull/task-api/internal/models"
	"github.com/Ricardo-DevNull/task-api/internal/models/enums"
)

type TaskRequest struct {
	models.BaseModel
	Title       string `json:"title" binding:"required,min=3,max=100"`
	Description string `json:"description" binding:"omitempty,max=300"`
}

type TaskResponse struct {
	models.BaseModel
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Status      enums.TaskStatus `json:"status"`
}
