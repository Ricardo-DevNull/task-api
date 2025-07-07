package dtos

import (
	"github.com/Ricardo-DevNull/task-api/internal/models"
	"github.com/Ricardo-DevNull/task-api/internal/models/enums"
)

type TaskRequest struct {
	models.Default
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TaskResponse struct {
	models.Default
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Status      enums.TaskStatus `json:"status"`
}
