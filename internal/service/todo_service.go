package service

import (
	"context"

	"github.com/lluxy8/todo-app-go/internal/model"
)

type TodoService interface {
	GetAll(context context.Context) ([]model.Todo, error)
}
