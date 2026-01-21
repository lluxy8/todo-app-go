package repository

import (
	"time"
	"todo-app/internal/model"
)

type MemoryTodoRepo struct{}

func NewMemoryTodoRepo() TodoRepository {
	return &MemoryTodoRepo{}
}

func (r *MemoryTodoRepo) GetAll() ([]model.Todo, error) {
	return inMemoryTodos, nil
}

var inMemoryTodos []model.Todo = []model.Todo{
	{
		Id:          1,
		Title:       "MyTodo1",
		Description: "This is MyTodo1.",
		DueDate:     time.Date(2027, time.April, 12, 17, 30, 12, 53, time.UTC),
	},
	{
		Id:          2,
		Title:       "MyTodo2",
		Description: "This is MyTodo2.",
		DueDate:     time.Date(2026, time.August, 21, 17, 30, 12, 53, time.UTC),
	},
	{
		Id:          3,
		Title:       "MyTodo3",
		Description: "This is MyTodo3.",
		DueDate:     time.Date(2026, time.September, 21, 17, 30, 12, 53, time.UTC),
	},
}
