package service

import "errors"

var (
	ErrTodoDoesNotExist = errors.New("todo does not exist")
	ErrInternal = errors.New("internal server error")
)
