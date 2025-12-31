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
}

func NewRepository(db *configs.MongoDB) *PatientRepository {
	return &PatientRepository{
		db:         db,
		collection: db.GetCollection("patient"),
	}
}

func (pr *PatientRepository) Create(patient PatientModel) (string, error) {
	insertResult, err := pr.collection.InsertOne(context.TODO(), patient)
	if err != nil {
		fmt.Println("error creating patient:", err)
		return primitive.ObjectID{}.String(), err
	}
	id := insertResult.InsertedID.(primitive.ObjectID)
	return id.Hex(), nil
}

func (pr *PatientRepository) List() ([]PatientModel, error) {
	filter := bson.M{
		"deleted": bson.M{"$ne": "false"},
	}
	list, err := pr.collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("error listing patients:", err)
		return nil, err
	}
	var modelList []PatientModel
	err = list.All(context.TODO(), &modelList)
	if err != nil {
		fmt.Println("error deconding patients list:", err)
		return nil, err
	}
	return modelList, nil
}

func (pr *PatientRepository) GetByID(id primitive.ObjectID) (PatientModel, error) {
	var patient PatientModel
	filter := bson.M{"_id": id}
	err := pr.collection.FindOne(context.TODO(), filter).Decode(&patient)
	if err != nil {
		fmt.Println("Error decoding patient:", err)
		return patient, err
	}
	return patient, err
}

func (pr *PatientRepository) Update(id primitive.ObjectID, patient *PatientModel) (string, error) {
	filter := bson.M{"_id": id}
	result, err := pr.collection.ReplaceOne(context.TODO(), filter, patient)
	if err != nil {
		fmt.Println("error updating patient:", err)
		return "", err
	}
	if result.ModifiedCount < 1 {
		return "", fmt.Errorf("no rows updateds")
	}
	return id.Hex(), nil
}

func (pr *PatientRepository) Delete(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"deleted": true,
		},
	}
	result, err := pr.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println("error deleting patient:", err)
		return err
	}
	if result.ModifiedCount < 1 {
		return fmt.Errorf("no rows updated")
	}
	return nil
}
