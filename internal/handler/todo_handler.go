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

// GetAll retrieves all todos
// @Summary Get all todos
// @Description Retrieve a list of all todos
// @Tags todos
// @Accept json
// @Produce json
// @Success 200 {array} model.Todo
// @Failure 500 {object} map[string]string
// @Router /todos [get]
func (h *TodoHandler) GetAll(ctx *gin.Context) {
	todos, err := h.todoService.GetAll(ctx.Request.Context())

	if handleServiceError(ctx, err) {
		return
	}

	ctx.JSON(http.StatusOK, todos)
}

// GetById retrieves a todo by ID
// @Summary Get a todo by ID
// @Description Retrieve a single todo by its ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {object} model.Todo
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /todos/{id} [get]
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

// Create creates a new todo
// @Summary Create a new todo
// @Description Create a new todo item
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body dto.CreateTodoRequest true "Todo data"
// @Success 201
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /todos [post]
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

// Update updates an existing todo
// @Summary Update a todo
// @Description Update an existing todo item by ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Param todo body dto.UpdateTodoRequest true "Updated todo data"
// @Success 200
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /todos/{id} [patch]
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

// Delete deletes a todo by ID
// @Summary Delete a todo
// @Description Delete a todo item by its ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /todos/{id} [delete]
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
