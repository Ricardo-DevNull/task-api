package repository

import (
	"github.com/Ricardo-DevNull/task-api/internal/models/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entities.User) error
	FindByID(id uint) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
	Update(id uint, user *entities.User) error
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *entities.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByID(id uint) (*entities.User, error) {
	var user entities.User

	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*entities.User, error) {
	var user entities.User

	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Update(id uint, user *entities.User) error {
	return r.db.Where("id = ?", id).Updates(&user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Where("id = ?", id).Delete(&entities.User{}).Error
}
