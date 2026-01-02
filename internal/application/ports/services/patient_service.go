package portsservices

import (
	"context"

	"github.com/Arthur-Conti/gi_nutri/internal/domain/dtos"
)

type PatientService interface {
	Create(ctx context.Context, patientDto dtos.PatientDTO) (string, error)
	List(ctx context.Context) ([]dtos.PatientDTO, error)
}