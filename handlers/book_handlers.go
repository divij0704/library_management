// handlers/book_handlers.go
package handlers

import (
	"context"
	"log"
	"net/http"
	"os"

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
		log.Println("Error inserting book:", err)
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

// GetBookByID retrieves a book by ID
func GetBookByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	var book models.Book
	db := getDB(c)
	err = db.Collection("books").FindOne(context.Background(), bson.M{"_id": id}).Decode(&book)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// UpdateBook updates a book by ID
func UpdateBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	var updatedBook models.Book
	err = c.ShouldBindJSON(&updatedBook)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	db := getDB(c)
	if db == nil {
		return
	}

	result, err := db.Collection("books").UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.M{"$set": updatedBook},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}

	if result.ModifiedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, updatedBook)
}

// DeleteBook deletes a book by ID
func DeleteBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	db := getDB(c)
	result, err := db.Collection("books").DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.Status(http.StatusNoContent)
}

var db *mongo.Database

func getDB(c *gin.Context) *mongo.Database {
	if os.Getenv("GIN_MODE") == "test" {
		return db
	}

	client, exists := c.Get("mongoClient")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not available"})
		return nil
	}

	return client.(*mongo.Client).Database("library_management")
}
