package configs

import (
	"context"
	"fmt"
	"os"

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
		uri = "mongodb://root:example@localhost:27017"
	}
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("error connecting to mongo:", err)
		return err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("error pinging mongo:", err)
		return err
	}
	md.db = client
	return nil
}

func (md *MongoDB) GetDB() *mongo.Client {
	return md.db
}

func (md *MongoDB) GetCollection(collectionName string) *mongo.Collection {
	return md.db.Database("project").Collection(collectionName)
}
