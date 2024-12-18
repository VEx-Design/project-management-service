package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongodb struct {
	client *mongo.Client
}

func ConnectToMongo() *mongodb {
	// Mongodb connection string
	dbPort := os.Getenv("MONGO_DB_DEPLOY_PORT")
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://localhost:%s", dbPort))

	// Getting username and password from .env
	username := os.Getenv("MONGO_DB_USERNAME")
	password := os.Getenv("MONGO_DB_PASSWORD")

	clientOptions.SetAuth(options.Credential{
		Username: username,
		Password: password,
	})

	// Connect to mongo
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
		return nil
	}

	log.Println("Connected to MongoDB")

	return &mongodb{client: client}
}

func (db *mongodb) GetClient() *mongo.Client {
	return db.client
}

func (db *mongodb) Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := db.client.Disconnect(ctx)
	if err != nil {
		log.Fatalf("Failed to disconnect to MongoDB: %v", err)
	}
	log.Println("Disconnected from MongoDB")
}
