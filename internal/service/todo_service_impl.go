package service

import (
	"todo-app/internal/model"
	"todo-app/internal/repository"
)

type todoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) TodoService {
	return &todoService{repo: repo}
}

func (s *todoService) GetAll() ([]model.Todo, error) {
	return s.repo.GetAll()
}
