package models

type Task struct {
	Default
	Title         string `json:"title"`
	Description   string `json:"description"`
	Status        string `json:"status"`
	CompletedDate string `json:"completedDate"`
}
