package services

import (
	"github.com/Ricardo-DevNull/task-api/internal/models"
	"github.com/Ricardo-DevNull/task-api/internal/models/dtos"
	"github.com/Ricardo-DevNull/task-api/internal/repository"
)

type UserService interface {
	CreateUser(newUser *dtos.UserRequest) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(id uint, newUser *dtos.UpdateUserRequest) (*models.User, error)
	DeleteUserByID(id uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) CreateUser(newUser *dtos.UserRequest) (*models.User, error) {

	user := &models.User{
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: newUser.Password,
		Role:     newUser.Role,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) UpdateUser(id uint, newUser *dtos.UpdateUserRequest) (*models.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	user.Name = newUser.Name
	user.Role = newUser.Role

	if err := s.repo.Update(id, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) DeleteUserByID(id uint) error {
	return s.repo.Delete(id)
}
