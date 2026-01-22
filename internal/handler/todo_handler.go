package handler

import (
	"net/http"

	"github.com/lluxy8/todo-app-go/internal/service"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	todoService service.TodoService
}

func NewTodoHandler(ts service.TodoService) *TodoHandler {
	return &TodoHandler{todoService: ts}
}

func (h *TodoHandler) GetAll(c *gin.Context) {
	todos, err := h.todoService.GetAll(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, todos)
}
