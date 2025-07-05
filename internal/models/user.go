package models

type User struct {
	// TODO: Usar struct default do gorm
	// ID
	// Created_At
	// Updated_At
	// Deleted_At
	Name     string
	Email    string
	Password string
	Role     string
}
