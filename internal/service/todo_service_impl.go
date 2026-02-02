package service

import (
	"context"
	"errors"
	"fmt"
	"log"

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
	todo, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, mapError(err)
	}

	return todo, nil
}

func (s *todoService) GetById(id string, ctx context.Context) (model.Todo, error) {
	todo, err := s.repo.GetById(id, ctx)
	if err != nil {
		return model.Todo{}, mapError(err)
	}

	return todo, nil
}

func (s *todoService) Create(todo model.Todo, ctx context.Context) error {
	err := s.repo.Create(todo, ctx)
	if err != nil {
		return mapError(err)
	}

	return nil
}

func (s *todoService) Delete(id string, ctx context.Context) error {
	err := s.repo.Delete(id, ctx)
	if err != nil {
		return mapError(err)
	}

	return nil
}

func (s *todoService) Update(id string, todo model.Todo, ctx context.Context) error {
	err := s.repo.Update(id, todo,  ctx)
	if err != nil {
		return mapError(err)
	}

	return nil
}

func mapError(err error) error {
	switch {
	case errors.Is(err, repository.ErrNotFound):
		return ErrTodoDoesNotExist
	default:
		log.Printf("unexpected repo error: %+v", err)
		return fmt.Errorf("%w: %v", ErrInternal, err)
	}
}
