package handlerstest

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"library_management/handlers"
)

var db *mongo.Database // Declare db as a global variable

// Replace the following with your actual MongoDB connection string
const mongoDBConnectionString = "mongodb://localhost:27017"

// Replace the following with your actual MongoDB setup logic
func initDB() *mongo.Client {
	// Replace with your MongoDB connection string
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoDBConnectionString))
	if err != nil {
		log.Fatal(err)
	}

	// Context with timeout to handle connection establishment
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	if err := client.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}

// Replace the following with your actual MongoDB setup logic
func setupMongoDB() {
	client := initDB()
	db = client.Database("library_management_test")
}

func TestBookHandlers(t *testing.T) {
	// Set up the MongoDB client
	setupMongoDB()

	// Test CreateBook
	t.Run("CreateBook", func(t *testing.T) {
		// Create a request to pass to the handler
		req, err := http.NewRequest("POST", "/api/books/", bytes.NewBuffer([]byte(`{"title": "Test Book", "author": "Test Author"}`)))
		assert.NoError(t, err)

		// Create a ResponseRecorder to record the response
		rr := httptest.NewRecorder()
		router := setupRoutes()

		// Serve the HTTP request to the ResponseRecorder
		router.ServeHTTP(rr, req)

		fmt.Println(rr.Body.String())

		// Check the status code
		assert.Equal(t, http.StatusCreated, rr.Code)

		// Additional assertions based on your business logic
		// For example, you can check if the book was actually created in the real database
	})

	// Add similar tests for other handlers

	// Test Route Setup
	t.Run("RouteSetup", func(t *testing.T) {
		router := setupRoutes()

		// Test if the router is set up correctly
		assert.NotNil(t, router)
	})
}

// Function to set up routes for testing
func setupRoutes() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/books")
	{
		api.POST("/", handlers.CreateBook)
		api.GET("/", handlers.GetBooks)
		api.GET("/:id", handlers.GetBookByID)
		api.PUT("/:id", handlers.UpdateBook)
		api.DELETE("/:id", handlers.DeleteBook)
	}

	return router
}
