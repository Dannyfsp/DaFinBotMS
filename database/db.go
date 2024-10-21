package database

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectDB(uri string) (context.Context, error) {
	if uri == "" {
		return nil, errors.New("invalid uri")
	}

	clientOptions := options.Client().ApplyURI(uri)

	// Create a context with a timeout to avoid hanging connections
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the MongoDB server to verify connection
	if err = Client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB!")

	return ctx, nil
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

func CreateCollections() error {
	collections := []string{"users", "admin", "user_profiles", "wallet"}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	existingCollections, err := Client.Database("finbotms").ListCollectionNames(ctx, bson.D{})
	if err != nil {
		return err
	}

	existingCollectionsMap := make(map[string]bool)
	for _, name := range existingCollections {
		existingCollectionsMap[name] = true
	}

	for _, collection := range collections {
		if _, exists := existingCollectionsMap[collection]; !exists {
			if err := Client.Database("finbotms").CreateCollection(ctx, collection); err != nil {
				return err
			}
		}
	}
	return nil
}
