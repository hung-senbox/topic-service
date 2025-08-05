package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Topic struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title" binding:"required"`
	Icon      string             `json:"icon" bson:"icon" binding:"required"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
