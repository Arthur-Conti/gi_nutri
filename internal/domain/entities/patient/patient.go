package patient

import (
	"github.com/Arthur-Conti/gi_nutri/internal/domain/dtos"
	"github.com/Arthur-Conti/gi_nutri/internal/domain/entities/formulas"
)

type Patient interface {
	GetPhysicalActivityResult()
	CalculateIMC() error
	CalculateAdjustedWeight() error
	CalculatePercentageWeightAdequacy() error
	CalculatePercentageWeightChange() error
	CalculateEER() error
	CalculateTMB(choice formulas.TMBFormulas) error
	GetResults() Results
	GetData() PatientData
	SetResults(results Results)
}

func NewPatient(opts PatientOpts) Patient {
	var patient Patient

	switch opts.Sex {
	case PatientSexMale:
		patient = NewPatientMale(opts)
	case PatientSexFemale:
		patient = NewPatientFemale(opts)
	}

	if opts.IsPregnant {
		patient = NewPatientPregnant(opts)
	}

	patient.GetPhysicalActivityResult()
	return patient
}

func NewPatientFromDTO(dto dtos.PatientDTO) Patient {
	opts := PatientOpts{
		Name:             dto.Name,
		Age:              dto.Age,
		Sex:              PatientSex(dto.Sex),
		Height:           dto.Measures.HeightCM,
		Weight:           dto.Measures.Weight,
		UsualWeight:      dto.UsualWeight,
		PhysicalActivity: PhysicalActivity(dto.PhysicalActivity),
		IsPregnant:       dto.IsPregnant,
		PregnancyInfo:    PregnancyInfo(dto.PregnancyInfo),
		IsLactating:      dto.IsLactating,
		LactatingInfo:    LactatingInfo(dto.LactatingInfo),
	}
	return NewPatient(opts)
}
