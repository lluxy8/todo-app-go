package service

import (
	"context"

	"github.com/lluxy8/todo-app-go/internal/model"
)

type TodoService interface {
	GetAll(ctx context.Context) ([]model.Todo, error)
	GetById(id string, ctx context.Context) (model.Todo, error)
	Create(todo model.Todo, ctx context.Context) error
}
