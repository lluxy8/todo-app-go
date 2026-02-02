package handler

import (
	"encoding/hex"
	"errors"
	"net/http"

	"github.com/lluxy8/todo-app-go/internal/service"

	"github.com/gin-gonic/gin"
)

func handleBindJsonError(ctx *gin.Context, dto any) bool {
	if err := ctx.ShouldBindJSON(dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return true
	}
	return false
}

func requireQueryParam[T any](
	ctx *gin.Context,
	name string,
	parse func(string) (T, error),
) (T, bool) {

	value := ctx.Query(name)
	if value == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": name + " query param is required",
		})
		var zero T
		return zero, true
	}

	parsed, err := parse(value)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": name + " is invalid",
		})
		var zero T
		return zero, true
	}

	return parsed, false
}

func requirePathParam[T any](
	ctx *gin.Context,
	name string,
	parse func(string) (T, error),
) (T, bool) {

	value := ctx.Param(name)
	if value == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": name + " path param is required",
		})
		var zero T
		return zero, true
	}

	parsed, err := parse(value)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": name + " is invalid",
		})
		var zero T
		return zero, true
	}

	return parsed, false
}

func handleServiceError(ctx *gin.Context, err error) bool {
	if err == nil {
		return false
	}

	switch {
	case errors.Is(err, service.ErrTodoDoesNotExist):
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "todo not found",
		})
	case errors.Is(err, service.ErrInternal):
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
	default:
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
	}

	return true
}

func parseHexString(s string) (string, error) {
	if _, err := hex.DecodeString(s); err != nil {
		return "", err
	}
	return s, nil
}
