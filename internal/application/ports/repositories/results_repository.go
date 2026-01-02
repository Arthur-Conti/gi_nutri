package portsrepositories

import (
	"context"

	resultsrepository "github.com/Arthur-Conti/gi_nutri/internal/infra/repositories/results"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ResultsRepository define a interface para operações de repositório de resultados
type ResultsRepository interface {
	Create(ctx context.Context, results resultsrepository.ResultsModel) (string, error)
	List(ctx context.Context, patientID string) ([]resultsrepository.ResultsModel, error)
	GetByID(ctx context.Context, id primitive.ObjectID) (resultsrepository.ResultsModel, error)
	GetLastByPatientID(ctx context.Context, patientID string) (resultsrepository.ResultsModel, error)
}

