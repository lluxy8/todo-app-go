package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/lluxy8/todo-app-go/internal/model"
	"github.com/lluxy8/todo-app-go/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTodoService_GetAll_OK(t *testing.T) {
	// arrange
	repo := fakeTodoRepo{
		getAllFn: func(ctx context.Context) ([]model.Todo, error) {
			return fakeData(), nil
		},
	}

	service := service.NewTodoService(&repo)

	//act
	result, err := service.GetAll(context.Background())

	// assert
	require.NoError(t, err)
	assert.Equal(t, fakeData(), result)
}

func TestTodoService_GetAll_Error(t *testing.T) {
	// arrange
	expectedErr := errors.New("db error")

	repo := &fakeTodoRepo{
		getAllFn: func(ctx context.Context) ([]model.Todo, error) {
			return nil, expectedErr
		},
	}

	todoService := service.NewTodoService(repo)

	// act
	result, err := todoService.GetAll(context.Background())

	// assert
	assert.Nil(t, result)
	assert.ErrorIs(t, err, service.ErrInternal)
}

func TestTodoService_GetById_OK(t *testing.T) {
	// arrange
	expected := fakeData()[0]

	repo := &fakeTodoRepo{
		getByIdFn: func(id string, ctx context.Context) (model.Todo, error) {
			assert.Equal(t, expected.ID, id)
			return expected, nil
		},
	}

	service := service.NewTodoService(repo)

	// act
	result, err := service.GetById(expected.ID, context.Background())

	// assert
	require.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestTodoService_Create_OK(t *testing.T) {
	// arrange
	called := false

	repo := &fakeTodoRepo{
		createFn: func(todo model.Todo, ctx context.Context) error {
			called = true
			assert.Equal(t, "My Todo", todo.Title)
			return nil
		},
	}

	service := service.NewTodoService(repo)

	// act
	err := service.Create(model.Todo{Title: "My Todo"}, context.Background())

	// assert
	require.NoError(t, err)
	assert.True(t, called)
}

type fakeTodoRepo struct {
	getAllFn  func(ctx context.Context) ([]model.Todo, error)
	getByIdFn func(id string, ctx context.Context) (model.Todo, error)
	createFn  func(todo model.Todo, ctx context.Context) error
}

func (f *fakeTodoRepo) GetAll(ctx context.Context) ([]model.Todo, error) {
	return f.getAllFn(ctx)
}

func (f *fakeTodoRepo) GetById(id string, ctx context.Context) (model.Todo, error) {
	return f.getByIdFn(id, ctx)
}

func (f *fakeTodoRepo) Create(todo model.Todo, ctx context.Context) error {
	return f.createFn(todo, ctx)
}

func fakeData() []model.Todo {
	return []model.Todo{
		{
			ID:          "69718bdf78dd80d4f16a1792",
			Title:       "My Todo",
			Description: "This is my todo.",
			DueDate:     time.Date(2027, time.April, 12, 17, 30, 12, 53, time.UTC),
		},
	}
}
