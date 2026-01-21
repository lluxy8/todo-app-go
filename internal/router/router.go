package router

import (
	"github.com/lluxy8/todo-app-go/internal/handler"
	"github.com/lluxy8/todo-app-go/internal/repository"
	"github.com/lluxy8/todo-app-go/internal/service"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	todoRepo := repository.NewMemoryTodoRepo()
	todoService := service.NewTodoService(todoRepo)
	todoHandler := handler.NewTodoHandler(todoService)

	r.GET("/health", handler.Health)
	r.GET("/todos", todoHandler.GetAll)

	return r
}
