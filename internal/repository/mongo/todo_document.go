package mongo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type todoDocument struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	DueDate     time.Time          `bson:"dueDate"`
	DateCreated time.Time          `bson:"dateCreated"`
	IsCompleted bool               `bson:"isCompleted"`
}
