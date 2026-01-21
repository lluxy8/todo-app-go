package service

import (
	"github.com/lluxy8/todo-app-go/internal/model"
)

type TodoService interface {
	GetAll() ([]model.Todo, error)
}
