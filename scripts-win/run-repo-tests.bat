@echo off
echo Running repository tests...
go test ../internal/repository/mongo/todo_repository_test.go
pause