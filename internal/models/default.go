package models

import "time"

type Default struct {
	ID        uint
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
