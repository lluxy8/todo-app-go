package handler

import (
	"net/http"
	"time"

	"github.com/lluxy8/todo-app-go/internal/handler/dto"
	"github.com/lluxy8/todo-app-go/internal/model"
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

	if handleServiceError(ctx, err) {
		return
	}

	ctx.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) GetById(ctx *gin.Context) {
	id, hasError := requirePathParam(ctx, "id", parseHexString)
	if hasError {
		return
	}

	todo, err := h.todoService.GetById(id, ctx.Request.Context())
	if handleServiceError(ctx, err) {
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) Create(ctx *gin.Context) {
	var req dto.CreateTodoRequest

	if handleBindJsonError(ctx, &req) {
		return
	}

	todo := model.Todo{
		Title:       req.Title,
		Description: req.Description,
		DueDate:     req.DueDate,
		DateCreated: time.Now().UTC(),
		IsCompleted: false,
	}

	err := h.todoService.Create(todo, ctx.Request.Context())
	if handleServiceError(ctx, err) {
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *TodoHandler) Update(ctx *gin.Context) {
	id, hasError := requireQueryParam(ctx, "id", parseHexString)
	if hasError {
		return
	}

	var req dto.UpdateTodoRequest

	if handleBindJsonError(ctx, &req) {
		return
	}

	todo := model.Todo{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
		DueDate:     req.DueDate,
		IsCompleted: req.IsCompleted,
	}

	err := h.todoService.Update(id, todo, ctx)
	if handleServiceError(ctx, err) {
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *TodoHandler) Delete(ctx *gin.Context) {
	id, hasError := requirePathParam(ctx, "id", parseHexString)
	if hasError {
		return
	}

	err := h.todoService.Delete(id, ctx.Request.Context())
	if handleServiceError(ctx, err) {
		return
	}

	ctx.Status(http.StatusOK)
}
