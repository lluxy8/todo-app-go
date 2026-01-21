package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"todo-app/internal/handler"
	"todo-app/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetTodos(t *testing.T) {

	// arrange
	r := setupRouter()

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

type fakeTodoService struct{}

func (h *fakeTodoService) GetAll() ([]model.Todo, error) {
	return fakeData(), nil
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	r := gin.New()

	todoService := &fakeTodoService{}
	todoHandler := handler.NewTodoHandler(todoService)

	r.GET("/todos", todoHandler.GetAll)

	return r
}

func fakeData() []model.Todo {
	return []model.Todo{
		{
			Id:          1,
			Title:       "MyTodo1",
			Description: "This is MyTodo1.",
			DueDate:     time.Date(2027, time.April, 12, 17, 30, 12, 53, time.UTC),
		},
	}
}
