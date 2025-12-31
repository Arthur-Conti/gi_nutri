package resultsrepository

import (
	"context"
	"fmt"

	"github.com/Arthur-Conti/gi_nutri/internal/infra/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ResultsRepository struct {
	db         *configs.MongoDB
	collection *mongo.Collection
}

func NewResultsRepository(db *configs.MongoDB) *ResultsRepository {
	return &ResultsRepository{
		db:         db,
		collection: db.GetCollection("results"),
	}
}

func (rr *ResultsRepository) Create(results ResultsModel) (string, error) {
	insertResult, err := rr.collection.InsertOne(context.TODO(), results)
	if err != nil {
		fmt.Println("error creating results:", err)
		return primitive.ObjectID{}.String(), err
	}
	id := insertResult.InsertedID.(primitive.ObjectID)
	return id.Hex(), nil
}

func (rr *ResultsRepository) List() ([]ResultsModel, error) {
	list, err := rr.collection.Find(context.TODO(), nil)
	if err != nil {
		fmt.Println("error listing results:", err)
		return nil, err
	}
	var modelList []ResultsModel
	err = list.All(context.TODO(), &modelList)
	if err != nil {
		fmt.Println("error deconding results list:", err)
		return nil, err
	}
	return modelList, nil
}

func (rr *ResultsRepository) GetByID(id primitive.ObjectID) (ResultsModel, error) {
	var results ResultsModel
	filter := bson.M{"_id": id}
	err := rr.collection.FindOne(context.TODO(), filter).Decode(&results)
	if err != nil {
		fmt.Println("Error decoding results:", err)
		return results, err
	}
	return results, err
}
