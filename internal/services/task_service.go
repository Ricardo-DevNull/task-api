package services

import (
	"github.com/Ricardo-DevNull/task-api/internal/models/dtos"
	"github.com/Ricardo-DevNull/task-api/internal/models/entities"
	"github.com/Ricardo-DevNull/task-api/internal/models/enums"
	"github.com/Ricardo-DevNull/task-api/internal/repository"
)

type TaskService interface {
	CreateTask(newTask *dtos.TaskRequest) (*entities.Task, error)
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(r repository.TaskRepository) TaskService {
	return &taskService{repo: r}
}

func (s *taskService) CreateTask(newTask *dtos.TaskRequest) (*entities.Task, error) {

	task := &entities.Task{
		Title:       newTask.Title,
		Description: newTask.Description,
		Status:      enums.StatusAvailable,
	}

	if err := s.repo.Create(task); err != nil {
		return nil, err
	}

	return task, nil
}
