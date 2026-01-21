package service

import (
	"todo-app/internal/model"
)

type TodoService interface {
	GetAll() ([]model.Todo, error)
}