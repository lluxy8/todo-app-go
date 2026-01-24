@echo off
echo Running handler tests...
go test ../internal/handler/todo_handler_test.go
pause