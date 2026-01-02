package configs

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	db *mongo.Client
}

func NewMongoDB() *MongoDB {
	return &MongoDB{}
}

func (md *MongoDB) Connect() error {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		uri = "mongodb://root:example@localhost:27017/?authSource=admin"
	}

	fmt.Printf("Connecting to MongoDB with URI: %s\n", uri)
	clientOptions := options.Client().ApplyURI(uri)

	// Add retry logic with exponential backoff
	maxRetries := 10
	retryDelay := 2 * time.Second

	var client *mongo.Client
	var err error

	for i := 0; i < maxRetries; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		client, err = mongo.Connect(ctx, clientOptions)
		cancel()

		if err != nil {
			fmt.Printf("Attempt %d/%d: error connecting to mongo: %v\n", i+1, maxRetries, err)
			if i < maxRetries-1 {
				time.Sleep(retryDelay)
				retryDelay *= 2 // Exponential backoff
			}
			continue
		}

		// Try to ping
		ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		err = client.Ping(ctx, nil)
		cancel()

		if err != nil {
			fmt.Printf("Attempt %d/%d: error pinging mongo: %v\n", i+1, maxRetries, err)
			client.Disconnect(context.Background())
			if i < maxRetries-1 {
				time.Sleep(retryDelay)
				retryDelay *= 2
			}
			continue
		}

		// Success!
		md.db = client
		fmt.Println("Successfully connected to MongoDB")
		return nil
	}

	fmt.Println("uri:", uri)
	fmt.Println("error: failed to connect to mongo after", maxRetries, "attempts")
	return err
}

func (md *MongoDB) GetDB() *mongo.Client {
	return md.db
}

func (md *MongoDB) GetCollection(collectionName string) *mongo.Collection {
	return md.db.Database("project").Collection(collectionName)
}
