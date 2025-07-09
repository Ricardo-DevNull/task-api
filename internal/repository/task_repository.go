package repository

import (
	"github.com/Ricardo-DevNull/task-api/internal/models/entities"
	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task *entities.Task) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(task *entities.Task) error {
	return r.db.Create(task).Error
}
