package portsservices

import (
	"context"

	"github.com/Arthur-Conti/gi_nutri/internal/domain/dtos"
)

type ResultsService interface {
	List(ctx context.Context, patientID string) ([]dtos.ResultsDTO, error)
	GetLastByPatientID(ctx context.Context, patientID string) (dtos.ResultsDTO, error)
	GetIMC(ctx context.Context, patientID string) (dtos.IMC, error)
	SaveIMC(ctx context.Context, patientID string) (string, error)
	GetAdjustedWeight(ctx context.Context, patientID string) (dtos.AdjustedWeightObesity, error)
	SaveAdjustedWeight(ctx context.Context, patientID string) (string, error)
	GetPercentageWeightAdequacy(ctx context.Context, patientID string, timeDays string) (dtos.PercentageWeightAdequacy, error)
	SavePercentageWeightAdequacy(ctx context.Context, patientID string, timeDays string) (string, error)
	GetPercentageWeightChange(ctx context.Context, patientID string, timeDays string) (dtos.PercentageWeightChange, error)
	SavePercentageWeightChange(ctx context.Context, patientID string, timeDays string) (string, error)
	GetEER(ctx context.Context, patientID string) (dtos.EER, error)
	SaveEER(ctx context.Context, patientID string) (string, error)
	GetTMB(ctx context.Context, patientID string, choice dtos.TMBFormulas) (dtos.TMB, error)
	SetTMB(ctx context.Context, patientID string, choice dtos.TMBFormulas) (string, error)

}