package database

import (
	"log"

	"github.com/Ricardo-DevNull/task-api/internal/models"
)

func Migrate() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Task{},
	)

	if err != nil {
		log.Fatal("Error executing migration")
	}
}
