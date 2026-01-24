package mongo_test

import (
	"context"
	"testing"
	"time"

	"github.com/lluxy8/todo-app-go/internal/model"
	"github.com/lluxy8/todo-app-go/internal/repository"
	"github.com/lluxy8/todo-app-go/internal/repository/mongo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestTodoRepository_CreateAndGetAll_OK(t *testing.T) {
	// Arrange
	_, collection := setupTestMongo(t)

	repo := mongo.NewTodoRepo(collection, 5*time.Second)

	todo := model.Todo{
		ID:          "69718bdf78dd80d4f16a1792",
		Title:       "Test Todo",
		Description: "This is my todo.",
		DueDate:     time.Date(2027, time.April, 12, 17, 30, 12, 53, time.UTC),
	}

	// Act
	err := repo.Create(todo, context.Background())
	require.NoError(t, err)

	todos, err := repo.GetAll(context.Background())
	require.NoError(t, err)

	// Assert
	require.Len(t, todos, 1)
	assert.Equal(t, "Test Todo", todos[0].Title)
}

func TestTodoRepository_GetById_NotFound(t *testing.T) {
	// arrange
	_, collection := setupTestMongo(t)
	repo := mongo.NewTodoRepo(collection, 5*time.Second)

	// act
	_, err := repo.GetById(
		primitive.NewObjectID().Hex(),
		context.Background(),
	)

	// assert
	assert.Error(t, err, repository.ErrNotFound)
}

func setupTestMongo(t *testing.T) (*mongoDriver.Client, *mongoDriver.Collection) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	t.Cleanup(cancel)

	client, err := mongoDriver.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27018"))
	require.NoError(t, err)

	db := client.Database("todo_test")
	colllection := db.Collection("todos")

	err = colllection.Drop(ctx)
	require.NoError(t, err)

	t.Cleanup(func() {
		_ = client.Disconnect(context.Background())
	})

	return client, colllection
}
