package database

import (
	"log"

	"github.com/Ricardo-DevNull/task-api/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.Task{},
	)

	if err != nil {
		log.Fatal("Error executing migration")
	}
}
