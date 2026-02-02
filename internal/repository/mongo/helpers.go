package mongo

import (
	"context"
	"time"

	"github.com/lluxy8/todo-app-go/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func withTimeout(
	parent context.Context,
	timeout time.Duration,
) (context.Context, context.CancelFunc) {

	if parent == nil {
		parent = context.Background()
	}

	return context.WithTimeout(parent, timeout)
}

func toDomain(doc todoDocument) model.Todo {
	return model.Todo{
		ID:          doc.ID.Hex(),
		Title:       doc.Title,
		Description: doc.Description,
		DueDate:     doc.DueDate,
		DateCreated: doc.DateCreated,
		IsCompleted: doc.IsCompleted,
	}
}

func toDocument(todo model.Todo) todoDocument {
	id, _ := primitive.ObjectIDFromHex(todo.ID)

	return todoDocument{
		ID:          id,
		Title:       todo.Title,
		Description: todo.Description,
		DueDate:     todo.DueDate,
		DateCreated: todo.DateCreated,
		IsCompleted: todo.IsCompleted,
	}
}
