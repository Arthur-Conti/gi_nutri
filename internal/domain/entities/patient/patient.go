package patient

import (
	"github.com/Arthur-Conti/gi_nutri/internal/domain/dtos"
	patientrepository "github.com/Arthur-Conti/gi_nutri/internal/infra/repositories/patient"
	resultsrepository "github.com/Arthur-Conti/gi_nutri/internal/infra/repositories/results"
)

type Patient interface {
	GetPhysicalActivityResult()
	CalculateIMC() error
	CalculateAdjustedWeight() error
	CalculatePercentageWeightAdequacy() error
	CalculatePercentageWeightChange() error
	CalculateEER() error
	CalculateTMB(useHarrisBenedict, useFao, useSchofield, usePocket bool, pocketValue float64) error
	GetFormulas() FormulasInfo
	PatientToModel() patientrepository.PatientModel
	ResultsToModel() resultsrepository.ResultsModel
}

func NewPatient(opts PatientOpts) Patient {
	var patient Patient

	opts.Measures = Measures{
		HeightCM: opts.Height,
		HeightM:  heightMeters(opts.Height),
		Weight:   opts.Weight,
	}

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
		Measures:         Measures(dto.Measures),
		IsPregnant:       dto.IsPregnant,
		PregnancyInfo:    PregnancyInfo(dto.PregnancyInfo),
		IsLactating:      dto.IsLactating,
		LactatingInfo:    LactatingInfo(dto.LactatingInfo),
	}
	return NewPatient(opts)
}
