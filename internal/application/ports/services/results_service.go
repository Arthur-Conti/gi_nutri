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
	GetPercentageWeightAdequacy(ctx context.Context, patientID string, timeDays string) (dtos.PercentageWeightAdequacy, error)
}