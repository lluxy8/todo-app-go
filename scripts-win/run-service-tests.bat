@echo off
echo Running service tests...
go test ../internal/service/todo_service_test.go
pause