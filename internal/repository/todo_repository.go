package repository

import (
	"context"

	"github.com/lluxy8/todo-app-go/internal/model"
)

type TodoRepository interface {
	GetAll(ctx context.Context) ([]model.Todo, error)
	GetById(id string, ctx context.Context) (model.Todo, error)
	Create(todo model.Todo, ctx context.Context) error
	Delete(id string, ctx context.Context) error
	Update(id string, todo model.Todo, ctx context.Context) error
}
