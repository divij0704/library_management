// main.go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"library_management/handlers"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set up the MongoDB connection
	client := initDB()
	defer client.Disconnect(context.Background())

	// Initialize the Gin router
	router := gin.Default()

	// Inject the MongoDB client into the context
	router.Use(func(c *gin.Context) {
		c.Set("mongoClient", client)
		c.Next()
	})

	// Define API routes
	api := router.Group("/api/books")
	{
		api.POST("/", handlers.CreateBook)
		api.GET("/", handlers.GetBooks)
		api.GET("/:id", handlers.GetBookByID)
		api.PUT("/:id", handlers.UpdateBook)
		api.DELETE("/:id", handlers.DeleteBook)
	}

	// Run the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}

func initDB() *mongo.Client {
	// Replace with your MongoDB connection string
	uri := "mongodb://localhost:27017/library_management"

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
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
