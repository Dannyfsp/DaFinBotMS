package database

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectDB(uri string) {
	if uri == "" {
		log.Fatal("invalid uri")
	}

	clientOptions := options.Client().ApplyURI(uri)

	// Create a context with a timeout to avoid hanging connections
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the MongoDB server to verify connection
	if err = Client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")
}

func GetCollection(dbName, collectionName string) (*mongo.Collection, error) {
	if dbName == "" {
		return nil, errors.New("invalid database name: cannot be empty")
	}
	if collectionName == "" {
		return nil, errors.New("invalid collection name: cannot be empty")
	}

	collection := Client.Database(dbName).Collection(collectionName)
	return collection, nil
}

// DisconnectDB closes the MongoDB connection.
func DisconnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := Client.Disconnect(ctx); err != nil {
		log.Fatal("Failed to disconnect from MongoDB:", err)
	}

	log.Println("Disconnected from MongoDB!")
}
