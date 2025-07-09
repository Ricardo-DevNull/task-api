package entities

import "github.com/Ricardo-DevNull/task-api/internal/models/shared"

type User struct {
	shared.BaseModel
	Name     string
	Email    string
	Password string
	Role     string
}
