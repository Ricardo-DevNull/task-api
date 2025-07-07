package models

type User struct {
	Default
	Name     string
	Email    string
	Password string
	Role     string
}
