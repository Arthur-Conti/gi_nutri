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
	GetResults() Results
	PatientToModel() patientrepository.PatientModel
	ResultsToModel() resultsrepository.ResultsModel
	FillResults(model resultsrepository.ResultsModel)
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

func NewPatientFromModel(model patientrepository.PatientModel) Patient {
	opts := PatientOpts{
		Name:             model.Name,
		Age:              model.Age,
		TimeDays:         model.TimeDays,
		Sex:              PatientSex(model.Sex),
		Height:           model.Height,
		Weight:           model.Weight,
		UsualWeight:      model.UsualWeight,
		PhysicalActivity: PhysicalActivity(model.PhysicalActivity),
		IsPregnant:       model.IsPregnant,
		PregnancyInfo:    PregnancyInfo(model.PregnancyInfo),
		IsLactating:      model.IsLactating,
		LactatingInfo:    LactatingInfo(model.LactatingInfo),
	}
	return NewPatient(opts)
}
