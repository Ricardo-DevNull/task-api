package services

import (
	"github.com/Ricardo-DevNull/task-api/internal/models"
	"github.com/Ricardo-DevNull/task-api/internal/repository"
)

type UserService interface {
	CreateUser(user *models.User) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) CreateUser(user *models.User) error {
	return s.repo.Create(user)
}
