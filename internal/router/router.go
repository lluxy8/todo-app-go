package router

import (
	"github.com/lluxy8/todo-app-go/internal/config"
	"github.com/lluxy8/todo-app-go/internal/handler"
	"github.com/lluxy8/todo-app-go/internal/repository/mongo"
	"github.com/lluxy8/todo-app-go/internal/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
)

type RouterDeps struct {
	TodoCollection *mongoDriver.Collection
	Cfg            *config.Config
}

func New(deps RouterDeps) *gin.Engine {
	r := gin.Default()

	todoRepo := mongo.NewTodoRepo(deps.TodoCollection, 5)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := handler.NewTodoHandler(todoService)

	r.GET("/health", handler.Health)

	r.GET("/todos", todoHandler.GetAll)
	r.GET("/todos/:id", todoHandler.GetById)
	r.POST("/todos", todoHandler.Create)
	r.PATCH("/todos/:id", todoHandler.Update)
	r.DELETE("/todos/:id", todoHandler.Delete)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
