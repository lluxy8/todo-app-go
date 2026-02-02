package service

import (
	"errors"
	"fmt"
	"log"

	"github.com/lluxy8/todo-app-go/internal/repository"
)

func mapError(err error) error {
	switch {
	case errors.Is(err, repository.ErrNotFound):
		return ErrTodoDoesNotExist
	default:
		log.Printf("unexpected repo error: %+v", err)
		return fmt.Errorf("%w: %v", ErrInternal, err)
	}
}
