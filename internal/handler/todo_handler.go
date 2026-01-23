package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/lluxy8/todo-app-go/internal/handler/dto"
	"github.com/lluxy8/todo-app-go/internal/model"
	"github.com/lluxy8/todo-app-go/internal/repository"
	"github.com/lluxy8/todo-app-go/internal/service"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	todoService service.TodoService
}

func NewTodoHandler(ts service.TodoService) *TodoHandler {
	return &TodoHandler{todoService: ts}
}

func (h *TodoHandler) GetAll(ctx *gin.Context) {
	todos, err := h.todoService.GetAll(ctx.Request.Context())

	if handleRepoError(ctx, err) {
		return
	}

	ctx.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) Create(ctx *gin.Context) {
	var req dto.CreateTodoRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := model.Todo{
		Title:       req.Title,
		Description: req.Description,
		DueDate:     req.DueDate,
		DateCreated: time.Now(),
		IsCompleted: false,
	}

	err := h.todoService.Create(todo, ctx.Request.Context())
	if handleRepoError(ctx, err) {
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *TodoHandler) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id query param is required",
		})
		return
	}

	todo, err := h.todoService.GetById(id, ctx.Request.Context())
	if handleRepoError(ctx, err) {
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

func handleRepoError(ctx *gin.Context, err error) bool {
	if err == nil {
		return false
	}

	if errors.Is(err, repository.ErrNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "no items found",
		})
		return true
	}

	ctx.JSON(http.StatusInternalServerError, gin.H{
		"error": "internal server error",
	})
	return true
}
