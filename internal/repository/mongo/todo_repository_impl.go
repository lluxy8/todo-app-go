package mongo

import (
	"context"
	"time"

	"github.com/lluxy8/todo-app-go/internal/model"
	"github.com/lluxy8/todo-app-go/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoMongoRepository struct {
	collection *mongo.Collection
	timeout    time.Duration
}

func NewTodoRepo(collection *mongo.Collection, timeout time.Duration) repository.TodoRepository {
	return &TodoMongoRepository{collection: collection, timeout: timeout}
}

func (r *TodoMongoRepository) GetAll(ctx context.Context) ([]model.Todo, error) {
	ctx, cancel := withTimeout(ctx, r.timeout)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var docs []todoDocument
	if err := cursor.All(ctx, &docs); err != nil {
		return nil, err
	}

	todos := make([]model.Todo, 0, len(docs))

	for _, doc := range docs {
		todos = append(todos, toDomain(doc))
	}

	return todos, nil
}

func (r *TodoMongoRepository) Create(todo model.Todo, ctx context.Context) error {
	ctx, cancel := withTimeout(ctx, r.timeout)
	defer cancel()

	doc := toDocument(todo)
	_, err := r.collection.InsertOne(ctx, doc)
	if err != nil {
		return err
	}

	return nil
}

func (r *TodoMongoRepository) GetById(id string, ctx context.Context) (model.Todo, error) {
	ctx, cancel := withTimeout(ctx, r.timeout)
	defer cancel()

	var doc todoDocument

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Todo{}, err
	}

	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&doc)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Todo{}, repository.ErrNotFound
		}

		return model.Todo{}, err
	}

	return toDomain(doc), nil
}

func (r *TodoMongoRepository) Delete(id string, ctx context.Context) error {
	ctx, cancel := withTimeout(ctx, r.timeout)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	result, err := r.collection.DeleteOne(
		ctx,
		bson.M{"_id": objectID},
	)

	if err != nil {
		return nil
	}

	if result.DeletedCount == 0 {
		return repository.ErrNotFound
	}

	return nil
}

func (r *TodoMongoRepository) Update(id string, todo model.Todo, ctx context.Context) error {
	ctx, cancel := withTimeout(ctx, r.timeout)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"title":       todo.Title,
			"description": todo.Description,
			"dueDate":     todo.DueDate,
		},
	}

	result, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": objectID},
		update,
	)

	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return repository.ErrNotFound
	}

	return nil
}

func withTimeout(
	parent context.Context,
	timeout time.Duration,
) (context.Context, context.CancelFunc) {

	if parent == nil {
		parent = context.Background()
	}

	return context.WithTimeout(parent, timeout)
}
