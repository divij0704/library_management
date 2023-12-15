// handlers/book_handlers.go
package handlers

import (
	"context"
	"net/http"

	"library_management/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateBook handles the creation of a new book
func CreateBook(c *gin.Context) {
	var book models.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	db := getDB(c)
	if db == nil {
		return
	}

	_, err = db.Collection("books").InsertOne(context.Background(), book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert book"})
		return
	}

	c.JSON(http.StatusCreated, book)
}

// GetBooks retrieves a list of all books
func GetBooks(c *gin.Context) {
	var books []models.Book
	db := getDB(c)

	cursor, err := db.Collection("books").Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var book models.Book
		if err := cursor.Decode(&book); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode book"})
			return
		}
		books = append(books, book)
	}

	c.JSON(http.StatusOK, books)
}



