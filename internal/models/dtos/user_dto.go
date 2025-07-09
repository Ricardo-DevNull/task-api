package dtos

import (
	"github.com/Ricardo-DevNull/task-api/internal/models/shared"
)

type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=4,containsany=!@#$%^&*()+"`
}

type UserRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=4,containsany=!@#$%^&*()+"`
	Role     string `json:"role" binding:"required,oneof=admin user"`
}

type UserResponse struct {
	shared.BaseModel
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Role  string `json:"role,omitempty"`
}

type UpdateUserRequest struct {
	Name string `json:"name" binding:"min=2,max=100"`
	Role string `json:"role" binding:"oneof=admin user"`
}
