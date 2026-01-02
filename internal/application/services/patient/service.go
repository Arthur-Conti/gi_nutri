package patientservice

import (
	"context"
	"fmt"
	"time"

	portsrepositories "github.com/Arthur-Conti/gi_nutri/internal/application/ports/repositories"
	"github.com/Arthur-Conti/gi_nutri/internal/domain/dtos"
	"github.com/Arthur-Conti/gi_nutri/internal/domain/entities/patient"
	resultsrepository "github.com/Arthur-Conti/gi_nutri/internal/infra/repositories/results"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PatientService struct {
	patientRepo portsrepositories.PatientRepository
	resultsRepo portsrepositories.ResultsRepository
}

func NewPatientService(
	patientRepo portsrepositories.PatientRepository,
	resultsRepo portsrepositories.ResultsRepository,
) *PatientService {
	return &PatientService{
		patientRepo: patientRepo,
		resultsRepo: resultsRepo,
	}
}

func (ps *PatientService) Create(ctx context.Context, patientDto dtos.PatientDTO) (string, error) {
	entity := patient.NewPatientFromDTO(patientDto)

	patientModel := entity.PatientToModel()
	patientModel.CreatedAt = time.Now()
	patientModel.UpdatedAt = time.Now()
	patientModel.Deleted = false
	patientModel.Finished = false

	patientID, err := ps.patientRepo.Create(ctx, patientModel)
	if err != nil {
		return "", fmt.Errorf("failed to create patient: %w", err)
	}

	objectID, err := primitive.ObjectIDFromHex(patientID)
	if err != nil {
		return "", fmt.Errorf("invalid patient id: %w", err)
	}

	resultsModel := entity.ResultsToModel()
	resultsModel.PatientID = objectID
	resultsModel.RecordedAt = time.Now()
	resultsModel.Measures = resultsrepository.Measures(entity.GetResults().Measures)
	
	_, err = ps.resultsRepo.Create(ctx, resultsModel)
	if err != nil {
		return "", fmt.Errorf("failed to create initial results: %w", err)
	}

	return patientID, nil
}

func (ps *PatientService) List(ctx context.Context) ([]dtos.PatientDTO, error) {
	modelList, err := ps.patientRepo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list patients: %w", err)
	}

	var dtoList []dtos.PatientDTO
	for _, model := range modelList {
		dtoList = append(dtoList, dtos.PatientFromModel(model))
	}

	return dtoList, nil
}
