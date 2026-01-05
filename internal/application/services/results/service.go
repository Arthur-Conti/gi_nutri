package resultsservice

import (
	"context"
	"fmt"
	"strconv"
	"time"

	portsrepositories "github.com/Arthur-Conti/gi_nutri/internal/application/ports/repositories"
	"github.com/Arthur-Conti/gi_nutri/internal/domain/dtos"
	"github.com/Arthur-Conti/gi_nutri/internal/domain/entities/formulas"
	"github.com/Arthur-Conti/gi_nutri/internal/infra/mappers"
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

	entity := mappers.PatientModelToEntity(patientModel)
	if err := entity.CalculateIMC(); err != nil {
		return dtos.IMC{}, fmt.Errorf("failed to calculate IMC: %w", err)
	}

	results := entity.GetResults()
	if results.FormulasInfo.IMC == nil {
		return dtos.IMC{}, fmt.Errorf("IMC not calculated")
	}

	return dtos.IMC{
		Status: string(results.FormulasInfo.IMC.Status),
		Result: results.FormulasInfo.IMC.Result,
	}, nil
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

	entity := mappers.PatientModelToEntity(patientModel)
	results := mappers.ResultsModelToEntity(resultModel)
	entity.SetResults(results)

	if err := entity.CalculateIMC(); err != nil {
		return "", fmt.Errorf("failed to calculate IMC: %w", err)
	}

	updatedResults := entity.GetResults()
	if updatedResults.FormulasInfo.IMC == nil {
		return "", fmt.Errorf("IMC not calculated")
	}

	resultModel = mappers.ResultsEntityToModel(updatedResults, patientID)
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

	entity := mappers.PatientModelToEntity(patientModel)
	results := mappers.ResultsModelToEntity(resultsModel)
	entity.SetResults(results)

	if err := entity.CalculateIMC(); err != nil {
		return dtos.AdjustedWeightObesity{}, fmt.Errorf("failed to calculate IMC: %w", err)
	}

	if err := entity.CalculateAdjustedWeight(); err != nil {
		return dtos.AdjustedWeightObesity{}, fmt.Errorf("failed to calculate adjusted weight: %w", err)
	}

	updatedResults := entity.GetResults()
	if updatedResults.FormulasInfo.AdjustedWeight == nil {
		return dtos.AdjustedWeightObesity{}, fmt.Errorf("adjusted weight not calculated")
	}

	return dtos.AdjustedWeightObesity{
		IdealWeight: updatedResults.FormulasInfo.AdjustedWeight.IdealWeight,
		Result:      updatedResults.FormulasInfo.AdjustedWeight.Result,
	}, nil
}

func (rs *ResultsService) SaveAdjustedWeight(ctx context.Context, patientID string) (string, error) {
	patientModel, err := rs.patientRepo.GetByID(ctx, patientID)
	if err != nil {
		return "", fmt.Errorf("failed to get patient: %w", err)
	}

	resultsModel, err := rs.resultsRepo.GetLastByPatientID(ctx, patientID)
	if err != nil {
		return "", fmt.Errorf("failed to get last result: %w", err)
	}

	entity := mappers.PatientModelToEntity(patientModel)
	results := mappers.ResultsModelToEntity(resultsModel)
	entity.SetResults(results)

	if err := entity.CalculateIMC(); err != nil {
		return "", fmt.Errorf("failed to calculate IMC: %w", err)
	}

	if err := entity.CalculateAdjustedWeight(); err != nil {
		return "", fmt.Errorf("failed to calculate adjusted weight: %w", err)
	}

	updatedResults := entity.GetResults()
	if updatedResults.FormulasInfo.AdjustedWeight == nil {
		return "", fmt.Errorf("adjusted weight not calculated")
	}

	return rs.resultsRepo.Create(ctx, mappers.ResultsEntityToModel(updatedResults, patientID))
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

	entity := mappers.PatientModelToEntity(patientModel)
	results := mappers.ResultsModelToEntity(resultsModel)
	entity.SetResults(results)

	if err := entity.CalculatePercentageWeightAdequacy(); err != nil {
		return dtos.PercentageWeightAdequacy{}, fmt.Errorf("failed to calculate percentage weight adequacy: %w", err)
	}

	updatedResults := entity.GetResults()
	if updatedResults.FormulasInfo.PercentageWeightAdequacy == nil {
		return dtos.PercentageWeightAdequacy{}, fmt.Errorf("percentage weight adequacy not calculated")
	}

	return dtos.PercentageWeightAdequacy{
		Classification: string(updatedResults.FormulasInfo.PercentageWeightAdequacy.Classification),
		Result:         updatedResults.FormulasInfo.PercentageWeightAdequacy.Result,
	}, nil
}

func (rs *ResultsService) SavePercentageWeightAdequacy(ctx context.Context, patientID string, timeDays string) (string, error) {
	days, err := strconv.Atoi(timeDays)
	if err != nil {
		return "", fmt.Errorf("invalid time_days format: %w", err)
	}

	patientModel, err := rs.patientRepo.GetByID(ctx, patientID)
	if err != nil {
		return "", fmt.Errorf("failed to get patient: %w", err)
	}
	patientModel.TimeDays = days

	resultsModel, err := rs.resultsRepo.GetLastByPatientID(ctx, patientID)
	if err != nil {
		return "", fmt.Errorf("failed to get last result: %w", err)
	}

	entity := mappers.PatientModelToEntity(patientModel)
	results := mappers.ResultsModelToEntity(resultsModel)
	entity.SetResults(results)

	if err := entity.CalculatePercentageWeightAdequacy(); err != nil {
		return "", fmt.Errorf("failed to calculate percentage weight adequacy: %w", err)
	}

	updatedResults := entity.GetResults()
	if updatedResults.FormulasInfo.PercentageWeightAdequacy == nil {
		return "", fmt.Errorf("percentage weight adequacy not calculated")
	}

	return rs.resultsRepo.Create(ctx, mappers.ResultsEntityToModel(updatedResults, patientID))
}

func (rs *ResultsService) GetPercentageWeightChange(ctx context.Context, patientID string, timeDays string) (dtos.PercentageWeightChange, error) {
	days, err := strconv.Atoi(timeDays)
	if err != nil {
		return dtos.PercentageWeightChange{}, fmt.Errorf("invalid time_days format: %w", err)
	}

	patientModel, err := rs.patientRepo.GetByID(ctx, patientID)
	if err != nil {
		return dtos.PercentageWeightChange{}, fmt.Errorf("failed to get patient: %w", err)
	}
	patientModel.TimeDays = days

	resultsModel, err := rs.resultsRepo.GetLastByPatientID(ctx, patientID)
	if err != nil {
		return dtos.PercentageWeightChange{}, fmt.Errorf("failed to get last result: %w", err)
	}

	entity := mappers.PatientModelToEntity(patientModel)
	results := mappers.ResultsModelToEntity(resultsModel)
	entity.SetResults(results)

	if err := entity.CalculatePercentageWeightChange(); err != nil {
		return dtos.PercentageWeightChange{}, fmt.Errorf("failed to calculate percentage weight change: %w", err)
	}

	updatedResults := entity.GetResults()
	if updatedResults.FormulasInfo.PercentageWeightChange == nil {
		return dtos.PercentageWeightChange{}, fmt.Errorf("percentage weight change not calculated")
	}

	return dtos.PercentageWeightChange{
		Classification: string(updatedResults.FormulasInfo.PercentageWeightChange.Classification),
		Result:         updatedResults.FormulasInfo.PercentageWeightChange.Result,
	}, nil
}

func (rs *ResultsService) SavePercentageWeightChange(ctx context.Context, patientID string, timeDays string) (string, error) {
	days, err := strconv.Atoi(timeDays)
	if err != nil {
		return "", fmt.Errorf("invalid time_days format: %w", err)
	}

	patientModel, err := rs.patientRepo.GetByID(ctx, patientID)
	if err != nil {
		return "", fmt.Errorf("failed to get patient: %w", err)
	}
	patientModel.TimeDays = days

	resultsModel, err := rs.resultsRepo.GetLastByPatientID(ctx, patientID)
	if err != nil {
		return "", fmt.Errorf("failed to get last result: %w", err)
	}

	entity := mappers.PatientModelToEntity(patientModel)
	results := mappers.ResultsModelToEntity(resultsModel)
	entity.SetResults(results)

	if err := entity.CalculatePercentageWeightChange(); err != nil {
		return "", fmt.Errorf("failed to calculate percentage weight change: %w", err)
	}

	updatedResults := entity.GetResults()
	if updatedResults.FormulasInfo.PercentageWeightChange == nil {
		return "", fmt.Errorf("percentage weight change not calculated")
	}

	return rs.resultsRepo.Create(ctx, mappers.ResultsEntityToModel(updatedResults, patientID))
}

func (rs *ResultsService) GetEER(ctx context.Context, patientID string) (dtos.EER, error) {
	patientModel, err := rs.patientRepo.GetByID(ctx, patientID)
	if err != nil {
		return dtos.EER{}, fmt.Errorf("failed to get patient: %w", err)
	}

	resultsModel, err := rs.resultsRepo.GetLastByPatientID(ctx, patientID)
	if err != nil {
		return dtos.EER{}, fmt.Errorf("failed to get last result: %w", err)
	}

	entity := mappers.PatientModelToEntity(patientModel)
	results := mappers.ResultsModelToEntity(resultsModel)
	entity.SetResults(results)

	if err := entity.CalculateEER(); err != nil {
		return dtos.EER{}, fmt.Errorf("failed to calculate eer: %w", err)
	}

	updatedResults := entity.GetResults()
	if updatedResults.FormulasInfo.PercentageWeightChange == nil {
		return dtos.EER{}, fmt.Errorf("eer not calculated")
	}

	return dtos.EER{
		Result: updatedResults.FormulasInfo.EER.Result,
	}, nil
}

func (rs *ResultsService) SaveEER(ctx context.Context, patientID string) (string, error) {
	patientModel, err := rs.patientRepo.GetByID(ctx, patientID)
	if err != nil {
		return "", fmt.Errorf("failed to get patient: %w", err)
	}

	resultsModel, err := rs.resultsRepo.GetLastByPatientID(ctx, patientID)
	if err != nil {
		return "", fmt.Errorf("failed to get last result: %w", err)
	}

	entity := mappers.PatientModelToEntity(patientModel)
	results := mappers.ResultsModelToEntity(resultsModel)
	entity.SetResults(results)

	if err := entity.CalculateEER(); err != nil {
		return "", fmt.Errorf("failed to calculate eer: %w", err)
	}

	updatedResults := entity.GetResults()
	if updatedResults.FormulasInfo.PercentageWeightChange == nil {
		return "", fmt.Errorf("eer not calculated")
	}

	return rs.resultsRepo.Create(ctx, mappers.ResultsEntityToModel(updatedResults, patientID))
}

func (rs *ResultsService) GetTMB(ctx context.Context, patientID string, choice dtos.TMBFormulas) (dtos.TMB, error) {
	patientModel, err := rs.patientRepo.GetByID(ctx, patientID)
	if err != nil {
		return dtos.TMB{}, fmt.Errorf("failed to get patient: %w", err)
	}

	resultsModel, err := rs.resultsRepo.GetLastByPatientID(ctx, patientID)
	if err != nil {
		return dtos.TMB{}, fmt.Errorf("failed to get last result: %w", err)
	}

	entity := mappers.PatientModelToEntity(patientModel)
	results := mappers.ResultsModelToEntity(resultsModel)
	entity.SetResults(results)

	if err := entity.CalculateTMB(formulas.TMBFormulas(choice)); err != nil {
		return dtos.TMB{}, fmt.Errorf("failed to calculate tmb: %w", err)
	}

	updatedResults := entity.GetResults()
	if updatedResults.FormulasInfo.PercentageWeightChange == nil {
		return dtos.TMB{}, fmt.Errorf("tmb not calculated")
	}

	return dtos.TMB{
		Result: updatedResults.FormulasInfo.TMB.Result,
	}, nil
}

func (rs *ResultsService) SetTMB(ctx context.Context, patientID string, choice dtos.TMBFormulas) (string, error) {
	patientModel, err := rs.patientRepo.GetByID(ctx, patientID)
	if err != nil {
		return "", fmt.Errorf("failed to get patient: %w", err)
	}

	resultsModel, err := rs.resultsRepo.GetLastByPatientID(ctx, patientID)
	if err != nil {
		return "", fmt.Errorf("failed to get last result: %w", err)
	}

	entity := mappers.PatientModelToEntity(patientModel)
	results := mappers.ResultsModelToEntity(resultsModel)
	entity.SetResults(results)

	if err := entity.CalculateTMB(formulas.TMBFormulas(choice)); err != nil {
		return "", fmt.Errorf("failed to calculate tmb: %w", err)
	}

	updatedResults := entity.GetResults()
	if updatedResults.FormulasInfo.PercentageWeightChange == nil {
		return "", fmt.Errorf("tmb not calculated")
	}

	return rs.resultsRepo.Create(ctx, mappers.ResultsEntityToModel(updatedResults, patientID))
}
