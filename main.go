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
	// Setting up the MongoDB connection
	client := initDB()
	defer client.Disconnect(context.Background())

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Set("mongoClient", client)
		c.Next()
	})

	//API routes
	api := router.Group("/api/books")
	{
		api.POST("/", handlers.CreateBook)
		api.GET("/", handlers.GetBooks)
		api.GET("/:id", handlers.GetBookByID)
		api.PUT("/:id", handlers.UpdateBook)
		api.DELETE("/:id", handlers.DeleteBook)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}

func initDB() *mongo.Client {
	uri := "mongodb://localhost:27017/library_management"

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}
