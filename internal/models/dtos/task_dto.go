package dtos

import (
	"github.com/Ricardo-DevNull/task-api/internal/models/enums"
	"github.com/Ricardo-DevNull/task-api/internal/models/shared"
)

type TaskRequest struct {
	shared.BaseModel
	Title       string `json:"title" binding:"required,min=3,max=100"`
	Description string `json:"description" binding:"omitempty,max=300"`
}

type TaskResponse struct {
	shared.BaseModel
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Status      enums.TaskStatus `json:"status"`
}
