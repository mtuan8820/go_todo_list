package models

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title" binding:"required, max=100"`
	Description string `json:"description" binding:"min=0,max=1000"`
	Completed   bool   `json:"completed"`
}
