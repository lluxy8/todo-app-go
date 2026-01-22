package router

import (
	"github.com/lluxy8/todo-app-go/internal/config"
	"github.com/lluxy8/todo-app-go/internal/handler"
	"github.com/lluxy8/todo-app-go/internal/repository/mongo"
	"github.com/lluxy8/todo-app-go/internal/service"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
)

type RouterDeps struct {
	TodoCollection *mongoDriver.Collection
	Cfg *config.Config
}

func New(deps RouterDeps) *gin.Engine {
	r := gin.Default()

	todoRepo := mongo.NewTodoRepo(deps.TodoCollection, 5)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := handler.NewTodoHandler(todoService)

	r.GET("/health", handler.Health)
	r.GET("/todos", todoHandler.GetAll)

	return r
}
