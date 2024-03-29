package taskmodel

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskDTO struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Priority    int                `json:"priority" bson:"priority"`
	Status      int                `json:"status" bson:"status"`
	DueDate     string             `json:"dueDate" bson:"due_date"`
	CreatedAt   time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
