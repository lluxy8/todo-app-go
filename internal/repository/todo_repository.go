package repository

import (
	"github.com/lluxy8/todo-app-go/internal/model"
)

type TodoRepository interface {
	GetAll() ([]model.Todo, error)
	Create(todo model.Todo)
}
