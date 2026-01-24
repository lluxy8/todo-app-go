package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/lluxy8/todo-app-go/internal/handler"
	"github.com/lluxy8/todo-app-go/internal/handler/dto"
	"github.com/lluxy8/todo-app-go/internal/model"
	"github.com/lluxy8/todo-app-go/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetTodos(t *testing.T) {
	// arrange
	r, _ := setupRouter()

	req, _ := http.NewRequest(http.MethodGet, "/todos", nil)
	w := httptest.NewRecorder()

	// act
	r.ServeHTTP(w, req)

	// assert
	assert.Equal(t, http.StatusOK, w.Code)

	var actual []model.Todo

	err := json.Unmarshal(w.Body.Bytes(), &actual)
	require.NoError(t, err)

	assert.Equal(t, fakeData(), actual)
}

func TestGetTodosByID_OK(t *testing.T) {
	// arrange
	r, _ := setupRouter()

	req, _ := http.NewRequest(
		http.MethodGet,
		"/todos/69718bdf78dd80d4f16a1792",
		nil,
	)
	w := httptest.NewRecorder()

	// act
	r.ServeHTTP(w, req)

	// assert
	assert.Equal(t, http.StatusOK, w.Code)

	var actuel model.Todo
	err := json.Unmarshal(w.Body.Bytes(), &actuel)
	require.NoError(t, err)

	assert.Equal(t, fakeData()[0], actuel)
}

func TestGetTodosByID_NotFound(t *testing.T) {
	// arrange
	r, _ := setupRouter()

	req, _ := http.NewRequest(
		http.MethodGet, 
		"/todos/unkown", 
		nil,
	)
	w := httptest.NewRecorder()

	//act
	r.ServeHTTP(w, req)

	// assert
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCreateTodos_Created(t *testing.T) {
	// arrange
	r, service := setupRouter()

	createRequest := dto.CreateTodoRequest{
		Title:       "My Todo 2",
		Description: "This is my todo 2!",
		DueDate:     time.Date(2027, time.April, 12, 17, 30, 12, 53, time.UTC),
	}

	body, err := json.Marshal(createRequest)
	require.NoError(t, err)

	req, _ := http.NewRequest(http.MethodPost, "/todos", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	// act
	r.ServeHTTP(w, req)

	// assert
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Len(t, service.todos, 2)
}

func TestCreateTodo_BadRequest(t *testing.T) {
	r, _ := setupRouter()

	req, _ := http.NewRequest(
		http.MethodPost,
		"/todos",
		bytes.NewBufferString(`{invalid-json}`),
	)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

type fakeTodoService struct {
	todos []model.Todo
}

func (s *fakeTodoService) GetAll(ctx context.Context) ([]model.Todo, error) {
	return s.todos, nil
}

func (s *fakeTodoService) GetById(id string, ctx context.Context) (model.Todo, error) {
	for _, t := range s.todos {
		if t.ID == id {
			return t, nil
		}
	}

	return model.Todo{}, repository.ErrNotFound
}

func (s *fakeTodoService) Create(todo model.Todo, ctx context.Context) error {
	s.todos = append(s.todos, todo)
	return nil
}

func setupRouter() (*gin.Engine, *fakeTodoService) {
	gin.SetMode(gin.TestMode)

	r := gin.New()

	todoService := &fakeTodoService{
		todos: fakeData(),
	}
	todoHandler := handler.NewTodoHandler(todoService)

	r.GET("/todos", todoHandler.GetAll)
	r.GET("/todos/:id", todoHandler.GetById)
	r.POST("/todos", todoHandler.Create)

	return r, todoService
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
