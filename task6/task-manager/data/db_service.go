package data

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var (
	userCollection *mongo.Collection
	taskCollection *mongo.Collection
)

func InitDB() {
    // Load the .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Get the MongoDB URI from environment variables
    mongoURI := os.Getenv("MONGODB_URI")
    if mongoURI == "" {
        log.Fatal("MONGODB_URI not set in environment")
    }

    clientOptions := options.Client().ApplyURI(mongoURI)

    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    taskCollection = client.Database("task_db").Collection("tasks")
	userCollection = client.Database("task_db").Collection("user")
    log.Println("Connected to MongoDB and initialized TaskCollection")
}