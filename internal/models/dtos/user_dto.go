package dtos

import "github.com/Ricardo-DevNull/task-api/internal/models"

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserResponse struct {
	models.BaseModel
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Role  string `json:"role,omitempty"`
}

type UpdateUserRequest struct {
	Name string `json:"name"`
	Role string `json:"role"`
}
