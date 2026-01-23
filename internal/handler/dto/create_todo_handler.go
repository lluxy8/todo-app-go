package dto

import "time"

type CreateTodoRequest struct {
	Title       string    `json:"title" binding:"required,min=1"`
	Description string    `json:"description" binding:"required,min=1,max=500"`
	DueDate     time.Time `json:"dueDate" binding:"required"`
}
