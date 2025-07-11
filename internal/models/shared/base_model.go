package shared

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint            `json:"id,omitempty" gorm:"primaryKey"`
	CreatedAt *time.Time      `json:"createdAt,omitempty"`
	UpdatedAt *time.Time      `json:"updatedAt,omitempty"`
	DeletedAt *gorm.DeletedAt `json:"-" gorm:"index"`
}
