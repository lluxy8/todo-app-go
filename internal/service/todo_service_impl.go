package service

import (
	"context"

	"github.com/lluxy8/todo-app-go/internal/model"
	"github.com/lluxy8/todo-app-go/internal/repository"
)

type todoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) TodoService {
	return &todoService{repo: repo}
}

func (s *todoService) GetAll(ctx context.Context) ([]model.Todo, error) {
	return s.repo.GetAll(ctx)
}

func (s *todoService) GetById(id string, ctx context.Context) (model.Todo, error) {
	return s.repo.GetById(id, ctx)
}

func (s *todoService) Create(todo model.Todo, ctx context.Context) error {
	return s.repo.Create(todo, ctx)
}
