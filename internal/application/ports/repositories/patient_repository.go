package portsrepositories

import (
	"context"

	patientrepository "github.com/Arthur-Conti/gi_nutri/internal/infra/repositories/patient"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PatientRepository define a interface para operações de repositório de pacientes
type PatientRepository interface {
	Create(ctx context.Context, patient patientrepository.PatientModel) (string, error)
	List(ctx context.Context) ([]patientrepository.PatientModel, error)
	GetByID(ctx context.Context, id string) (patientrepository.PatientModel, error)
	Update(ctx context.Context, id primitive.ObjectID, patient *patientrepository.PatientModel) (string, error)
	Delete(ctx context.Context, id primitive.ObjectID) error
}

