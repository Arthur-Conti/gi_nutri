package patientservice

import (
	"context"
	"fmt"
	"time"

	portsrepositories "github.com/Arthur-Conti/gi_nutri/internal/application/ports/repositories"
	"github.com/Arthur-Conti/gi_nutri/internal/domain/dtos"
	"github.com/Arthur-Conti/gi_nutri/internal/domain/entities/patient"
	"github.com/Arthur-Conti/gi_nutri/internal/infra/mappers"
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

	patientModel := mappers.PatientEntityToModel(entity)
	patientModel.CreatedAt = time.Now()
	patientModel.UpdatedAt = time.Now()
	patientModel.Deleted = false
	patientModel.Finished = false

	patientID, err := ps.patientRepo.Create(ctx, patientModel)
	if err != nil {
		return "", fmt.Errorf("failed to create patient: %w", err)
	}

	resultsModel := mappers.ResultsEntityToModel(entity.GetResults(), patientID)
	resultsModel.RecordedAt = time.Now()
	
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
