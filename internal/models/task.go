package models

type Task struct {
	// TODO: Usar struct default do gorm
	// ID
	// Created_At
	// Updated_At
	// Deleted_At
	Title         string
	Description   string
	Status        string
	CompletedDate string
}
