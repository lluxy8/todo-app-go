package repository

import (
	"todo-app/internal/model"
)

type TodoRepository interface {
	GetAll() ([]model.Todo, error)
}
