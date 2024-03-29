package taskmodel

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskDTO struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Priority    int                `json:"priority" bson:"priority"`
	Status      int                `json:"status" bson:"status"`
	DueDate     time.Time          `json:"due_date,omitempty" bson:"due_date,omitempty"`
	Timestamp   time.Time          `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
	CreatedAt   time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
