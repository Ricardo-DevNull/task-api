package services

import (
	"github.com/Ricardo-DevNull/task-api/internal/middleware"
	"github.com/Ricardo-DevNull/task-api/internal/models/dtos"
	"github.com/Ricardo-DevNull/task-api/internal/models/entities"
	"github.com/Ricardo-DevNull/task-api/internal/repository"
)

type UserService interface {
	LoginUser(authUser *dtos.UserLogin) (string, error)
	CreateUser(newUser *dtos.UserRequest) (*entities.User, error)
	GetUserByID(id uint) (*entities.User, error)
	UpdateUser(id uint, newUser *dtos.UpdateUserRequest) (*entities.User, error)
	DeleteUserByID(id uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) LoginUser(authUser *dtos.UserLogin) (string, error) {

	user, err := s.repo.FindByEmail(authUser.Email)
	if err != nil {
		return "", err
	}

	err = middleware.ComparePassword(authUser.Password, user.Password)
	if err != nil {
		return "", err
	}

	token, err := middleware.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) CreateUser(newUser *dtos.UserRequest) (*entities.User, error) {

	user := &entities.User{
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

func (s *userService) GetUserByID(id uint) (*entities.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) UpdateUser(id uint, newUser *dtos.UpdateUserRequest) (*entities.User, error) {
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
