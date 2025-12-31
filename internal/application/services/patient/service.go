package patientservice

import (
	"time"

	"github.com/Arthur-Conti/gi_nutri/internal/domain/dtos"
	"github.com/Arthur-Conti/gi_nutri/internal/domain/entities/patient"
	patientrepository "github.com/Arthur-Conti/gi_nutri/internal/infra/repositories/patient"
	resultsrepository "github.com/Arthur-Conti/gi_nutri/internal/infra/repositories/results"
)

type PatientService struct {
	patientRepo    *patientrepository.PatientRepository
	resultsRepo *resultsrepository.ResultsRepository
}

func NewPatientService(
	patientRepo *patientrepository.PatientRepository, 
	resultsRepo *resultsrepository.ResultsRepository,
) *PatientService {
	return &PatientService{
		patientRepo:    patientRepo,
		resultsRepo: resultsRepo,
	}
}

func (ps *PatientService) Create(patientDto dtos.PatientDTO) (string, error) {
	entity := patient.NewPatientFromDTO(patientDto)

	patientModel := entity.PatientToModel()
	patientModel.CreatedAt = time.Now()
	patientModel.UpdatedAt = time.Now()
	patientModel.Deleted = false
	patientModel.Finished = false

	patientID, err := ps.patientRepo.Create(patientModel)
	if err != nil {
		return "", err
	}

	resultsModel := entity.ResultsToModel()
	resultsModel.RecordedAt = time.Now()
	_, err = ps.resultsRepo.Create(resultsModel)
	if err != nil {
		return "", err
	}

	return patientID, nil
}
