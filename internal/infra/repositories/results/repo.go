package resultsrepository

import (
	"context"
	"fmt"

	"github.com/Arthur-Conti/gi_nutri/internal/infra/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ResultsRepository struct {
	db         *configs.MongoDB
	collection *mongo.Collection
	logger     *configs.Logger
}

func NewResultsRepository(db *configs.MongoDB, logger *configs.Logger) *ResultsRepository {
	return &ResultsRepository{
		db:         db,
		collection: db.GetCollection("results"),
		logger:     logger,
	}
}

func (rr *ResultsRepository) Create(ctx context.Context, results ResultsModel) (string, error) {
	insertResult, err := rr.collection.InsertOne(ctx, results)
	if err != nil {
		rr.logger.Error("error creating results: %v", err)
		return primitive.ObjectID{}.String(), fmt.Errorf("failed to create results: %w", err)
	}
	id := insertResult.InsertedID.(primitive.ObjectID)
	return id.Hex(), nil
}

func (rr *ResultsRepository) List(ctx context.Context, patientID string) ([]ResultsModel, error) {
	objectID, err := primitive.ObjectIDFromHex(patientID)
	if err != nil {
		rr.logger.Warn("invalid patient id format: %s", patientID)
		return nil, fmt.Errorf("invalid id format: %w", err)
	}

	filter := bson.M{"patient_id": objectID}

	list, err := rr.collection.Find(ctx, filter)
	if err != nil {
		rr.logger.Error("error listing results: %v", err)
		return nil, fmt.Errorf("failed to list results: %w", err)
	}
	defer list.Close(ctx)

	var modelList []ResultsModel
	err = list.All(ctx, &modelList)
	if err != nil {
		rr.logger.Error("error decoding results list: %v", err)
		return nil, fmt.Errorf("failed to decode results list: %w", err)
	}
	return modelList, nil
}

func (rr *ResultsRepository) GetByID(ctx context.Context, id primitive.ObjectID) (ResultsModel, error) {
	var results ResultsModel
	filter := bson.M{"_id": id}
	err := rr.collection.FindOne(ctx, filter).Decode(&results)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			rr.logger.Warn("results not found: %s", id.Hex())
			return results, fmt.Errorf("results not found: %w", err)
		}
		rr.logger.Error("error decoding results: %v", err)
		return results, fmt.Errorf("failed to get results: %w", err)
	}
	return results, nil
}

func (rr *ResultsRepository) GetLastByPatientID(ctx context.Context, patientID string) (ResultsModel, error) {
	var result ResultsModel

	objectID, err := primitive.ObjectIDFromHex(patientID)
	if err != nil {
		rr.logger.Warn("invalid patient id format: %s", patientID)
		return result, fmt.Errorf("invalid id format: %w", err)
	}

	opts := options.FindOne().SetSort(bson.D{{Key: "recorded_at", Value: -1}})
	filter := bson.M{"patient_id": objectID}

	err = rr.collection.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			rr.logger.Warn("results not found for patient: %s", patientID)
			return result, fmt.Errorf("results not found: %w", err)
		}
		rr.logger.Error("error decoding result: %v", err)
		return result, fmt.Errorf("failed to get last result: %w", err)
	}
	return result, nil
}
