package router

import (
	"todo-app/internal/handler"
	"todo-app/internal/repository"
	"todo-app/internal/service"

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
