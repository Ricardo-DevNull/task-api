package models

import (
	"time"

	"github.com/Ricardo-DevNull/task-api/internal/models/enums"
)

type Task struct {
	BaseModel
	Title         string           `gorm:"varchar(255)"`
	Description   string           `gorm:"text"`
	Status        enums.TaskStatus `gorm:"varchar(20)"`
	CompletedDate time.Time        `gorm:"timestamp"`
}
