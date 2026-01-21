package model

import "time"

type Todo struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"dueDate"`
	DateCreated time.Time `json:"dateCreted"`
	IsCompleted bool      `json:"isCompleted"`
}
