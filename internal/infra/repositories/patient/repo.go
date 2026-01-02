package patientrepository

import (
	"context"
	"fmt"

	"github.com/Arthur-Conti/gi_nutri/internal/infra/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PatientRepository struct {
	db         *configs.MongoDB
	collection *mongo.Collection
	logger     *configs.Logger
}

func NewRepository(db *configs.MongoDB, logger *configs.Logger) *PatientRepository {
	return &PatientRepository{
		db:         db,
		collection: db.GetCollection("patient"),
		logger:     logger,
	}
}

func (pr *PatientRepository) Create(ctx context.Context, patient PatientModel) (string, error) {
	insertResult, err := pr.collection.InsertOne(ctx, patient)
	if err != nil {
		pr.logger.Error("error creating patient: %v", err)
		return primitive.ObjectID{}.String(), fmt.Errorf("failed to create patient: %w", err)
	}
	id := insertResult.InsertedID.(primitive.ObjectID)
	return id.Hex(), nil
}

func (pr *PatientRepository) List(ctx context.Context) ([]PatientModel, error) {
	filter := bson.M{
		"deleted": bson.M{"$ne": "false"},
	}
	list, err := pr.collection.Find(ctx, filter)
	if err != nil {
		pr.logger.Error("error listing patients: %v", err)
		return nil, fmt.Errorf("failed to list patients: %w", err)
	}
	defer list.Close(ctx)

	var modelList []PatientModel
	err = list.All(ctx, &modelList)
	if err != nil {
		pr.logger.Error("error decoding patients list: %v", err)
		return nil, fmt.Errorf("failed to decode patients list: %w", err)
	}
	return modelList, nil
}

func (pr *PatientRepository) GetByID(ctx context.Context, id string) (PatientModel, error) {
	var patient PatientModel

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		pr.logger.Warn("invalid patient id format: %s", id)
		return patient, fmt.Errorf("invalid id format: %w", err)
	}

	filter := bson.M{"_id": objectID}
	err = pr.collection.FindOne(ctx, filter).Decode(&patient)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			pr.logger.Warn("patient not found: %s", id)
			return patient, fmt.Errorf("patient not found: %w", err)
		}
		pr.logger.Error("error decoding patient: %v", err)
		return patient, fmt.Errorf("failed to get patient: %w", err)
	}
	return patient, nil
}

func (pr *PatientRepository) Update(ctx context.Context, id primitive.ObjectID, patient *PatientModel) (string, error) {
	filter := bson.M{"_id": id}
	result, err := pr.collection.ReplaceOne(ctx, filter, patient)
	if err != nil {
		pr.logger.Error("error updating patient: %v", err)
		return "", fmt.Errorf("failed to update patient: %w", err)
	}
	if result.ModifiedCount < 1 {
		pr.logger.Warn("no patient updated with id: %s", id.Hex())
		return "", fmt.Errorf("no patient updated")
	}
	return id.Hex(), nil
}

func (pr *PatientRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"deleted": true,
		},
	}
	result, err := pr.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		pr.logger.Error("error deleting patient: %v", err)
		return fmt.Errorf("failed to delete patient: %w", err)
	}
	if result.ModifiedCount < 1 {
		pr.logger.Warn("no patient deleted with id: %s", id.Hex())
		return fmt.Errorf("no patient deleted")
	}
	return nil
}
