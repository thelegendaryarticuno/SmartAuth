package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UserCollection *mongo.Collection

// ConnectMongo initializes the MongoDB connection and sets up the UserCollection
func ConnectMongo() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, relying on system environment variables")
	}

	// Get MongoDB URI from environment variables
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI not found in environment variables")
	}

	// Create a new MongoDB client
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Failed to create MongoDB client:", err)
	}

	// Set a timeout context for connecting to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Get the database name from environment variables
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Println("DB_NAME not set in environment variables. Using default 'test'")
		dbName = "test"
	}

	// Initialize the UserCollection
	UserCollection = client.Database(dbName).Collection("users")

	log.Println("Successfully connected to MongoDB")
}
