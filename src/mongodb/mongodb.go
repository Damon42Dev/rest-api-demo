package mongodb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectMongoDb take mongodb url and related to connections
func ConnectMongoDb(url string) (*mongo.Client, error) {

	log.Println("Connecting to MongoDB...", url)
	clientOptions := options.Client().ApplyURI(url)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")

	return client, nil
}
