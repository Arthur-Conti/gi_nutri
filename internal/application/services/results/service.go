package resultsservice

import (
	"context"
	"fmt"
	"strconv"
	"time"

	portsrepositories "github.com/Arthur-Conti/gi_nutri/internal/application/ports/repositories"
	"github.com/Arthur-Conti/gi_nutri/internal/domain/dtos"
	"github.com/Arthur-Conti/gi_nutri/internal/domain/entities/patient"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ResultsService struct {
	patientRepo portsrepositories.PatientRepository
	resultsRepo portsrepositories.ResultsRepository
}

func NewResultsService(
	patientRepo portsrepositories.PatientRepository,
	resultsRepo portsrepositories.ResultsRepository,
) *ResultsService {
	return &ResultsService{
		patientRepo: patientRepo,
		resultsRepo: resultsRepo,
	}
}

func (rs *ResultsService) List(ctx context.Context, patientID string) ([]dtos.ResultsDTO, error) {
	resultsModel, err := rs.resultsRepo.List(ctx, patientID)
	if err != nil {
		return nil, fmt.Errorf("failed to list results: %w", err)
	}

	var resultsDto []dtos.ResultsDTO
	for _, model := range resultsModel {
		resultsDto = append(resultsDto, dtos.ResultsFromModel(model))
	}

	return resultsDto, nil
}

func (rs *ResultsService) GetLastByPatientID(ctx context.Context, patientID string) (dtos.ResultsDTO, error) {
	model, err := rs.resultsRepo.GetLastByPatientID(ctx, patientID)
	if err != nil {
		return dtos.ResultsDTO{}, fmt.Errorf("failed to get last result: %w", err)
	}

	return dtos.ResultsFromModel(model), nil
}

func (rs *ResultsService) GetIMC(ctx context.Context, patientID string) (dtos.IMC, error) {
	patientModel, err := rs.patientRepo.GetByID(ctx, patientID)
	if err != nil {
		return dtos.IMC{}, fmt.Errorf("failed to get patient: %w", err)
	}

	entity := patient.NewPatientFromModel(patientModel)
	if err := entity.CalculateIMC(); err != nil {
		return dtos.IMC{}, fmt.Errorf("failed to calculate IMC: %w", err)
	}

	return dtos.IMC(entity.ResultsToModel().Formulas.IMC), nil
}

func (rs *ResultsService) SaveIMC(ctx context.Context, patientID string) (string, error) {
	patientModel, err := rs.patientRepo.GetByID(ctx, patientID)
	if err != nil {
		return "", fmt.Errorf("failed to get patient: %w", err)
	}

	resultModel, err := rs.resultsRepo.GetLastByPatientID(ctx, patientID)
	if err != nil {
		return "", fmt.Errorf("failed to get last result: %w", err)
	}

	entity := patient.NewPatientFromModel(patientModel)
	if err := entity.CalculateIMC(); err != nil {
		return "", fmt.Errorf("failed to calculate IMC: %w", err)
	}

	resultModel.Formulas.IMC = entity.ResultsToModel().Formulas.IMC
	resultModel.ID = primitive.NilObjectID
	resultModel.RecordedAt = time.Now()

	resultID, err := rs.resultsRepo.Create(ctx, resultModel)
	if err != nil {
		return "", fmt.Errorf("failed to save IMC: %w", err)
	}

	return resultID, nil
}

func (rs *ResultsService) GetAdjustedWeight(ctx context.Context, patientID string) (dtos.AdjustedWeightObesity, error) {
	patientModel, err := rs.patientRepo.GetByID(ctx, patientID)
	if err != nil {
		return dtos.AdjustedWeightObesity{}, fmt.Errorf("failed to get patient: %w", err)
	}

	resultsModel, err := rs.resultsRepo.GetLastByPatientID(ctx, patientID)
	if err != nil {
		return dtos.AdjustedWeightObesity{}, fmt.Errorf("failed to get last result: %w", err)
	}

	entity := patient.NewPatientFromModel(patientModel)
	entity.FillResults(resultsModel)

	if err := entity.CalculateIMC(); err != nil {
		return dtos.AdjustedWeightObesity{}, fmt.Errorf("failed to calculate IMC: %w", err)
	}

	if err := entity.CalculateAdjustedWeight(); err != nil {
		return dtos.AdjustedWeightObesity{}, fmt.Errorf("failed to calculate adjusted weight: %w", err)
	}

	return dtos.AdjustedWeightObesity(entity.ResultsToModel().Formulas.AdjustedWeightObesity), nil
}

func (rs *ResultsService) GetPercentageWeightAdequacy(ctx context.Context, patientID string, timeDays string) (dtos.PercentageWeightAdequacy, error) {
	days, err := strconv.Atoi(timeDays)
	if err != nil {
		return dtos.PercentageWeightAdequacy{}, fmt.Errorf("invalid time_days format: %w", err)
	}

	patientModel, err := rs.patientRepo.GetByID(ctx, patientID)
	if err != nil {
		return dtos.PercentageWeightAdequacy{}, fmt.Errorf("failed to get patient: %w", err)
	}
	patientModel.TimeDays = days

	resultsModel, err := rs.resultsRepo.GetLastByPatientID(ctx, patientID)
	if err != nil {
		return dtos.PercentageWeightAdequacy{}, fmt.Errorf("failed to get last result: %w", err)
	}

	entity := patient.NewPatientFromModel(patientModel)
	entity.FillResults(resultsModel)

	if err := entity.CalculatePercentageWeightAdequacy(); err != nil {
		return dtos.PercentageWeightAdequacy{}, fmt.Errorf("failed to calculate percentage weight adequacy: %w", err)
	}

	return dtos.PercentageWeightAdequacy(entity.ResultsToModel().Formulas.PercentageWeightAdequacy), nil
}
