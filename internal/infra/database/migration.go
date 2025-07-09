package database

import (
	"log"

	"github.com/Ricardo-DevNull/task-api/internal/models/entities"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&entities.User{},
		&entities.Task{},
	)

	if err != nil {
		log.Fatal("Error executing migration")
	}
}
