package repository

import (
	"github.com/Ricardo-DevNull/task-api/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByID(id uint) (*models.User, error)
	Update(id uint, user *models.User) error
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByID(id uint) (*models.User, error) {
	var user models.User

	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Update(id uint, user *models.User) error {
	return r.db.Where("id = ?", id).Updates(&user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Where("id = ?", id).Delete(&models.User{}).Error
}
