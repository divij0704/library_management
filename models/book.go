// models/book.go
package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title  string             `json:"title"`
	Author string             `json:"author"`
	ISBN   string             `json:"isbn"`
}
