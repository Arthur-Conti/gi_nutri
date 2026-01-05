package mappers

import (
	"time"

	"github.com/Arthur-Conti/gi_nutri/internal/domain/dtos"
	"github.com/Arthur-Conti/gi_nutri/internal/domain/entities/patient"
)

// PatientEntityToDTO converte uma entidade Patient para PatientDTO
func PatientEntityToDTO(p patient.Patient, id string, createdAt, updatedAt time.Time, finished bool) dtos.PatientDTO {
	data := p.GetData()
	
	return dtos.PatientDTO{
		ID:                         id,
		Name:                       data.Name,
		Age:                        data.Age,
		AgeClassification:          string(data.AgeClassification),
		SchofieldAgeClassification: string(data.SchofieldAgeClassification),
		TimeDays:                   data.TimeDays,
		Sex:                        string(data.Sex),
		UsualWeight:                data.UsualWeight,
		PhysicalActivity:           string(data.PhysicalActivity),
		Measures: dtos.PatientMeasures{
			HeightCM: data.Height,
			HeightM:  data.Results.Measures.HeightM,
			Weight:   data.Weight,
		},
		IsPregnant:    data.IsPregnant,
		PregnancyInfo: dtos.PregnancyInfo(data.PregnancyInfo),
		IsLactating:   data.IsLactating,
		LactatingInfo: dtos.LactatingInfo(data.LactatingInfo),
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
		Finished:      finished,
	}
}

// ResultsEntityToDTO converte Results da entidade para ResultsDTO
func ResultsEntityToDTO(results patient.Results, patientID string, recordedAt time.Time) dtos.ResultsDTO {
	return dtos.ResultsDTO{
		ID:         results.ResultsID,
		PatientID:  patientID,
		RecordedAt: recordedAt,
		Measures:   dtos.Measures(results.Measures),
		Formulas:   formulasInfoToDTO(results.FormulasInfo),
	}
}

func formulasInfoToDTO(formulas patient.FormulasInfo) dtos.Formulas {
	var imc dtos.IMC
	if formulas.IMC != nil {
		imc = dtos.IMC{
			Status: string(formulas.IMC.Status),
			Result: formulas.IMC.Result,
		}
	}

	var adjustedWeight dtos.AdjustedWeightObesity
	if formulas.AdjustedWeight != nil {
		adjustedWeight = dtos.AdjustedWeightObesity{
			IdealWeight: formulas.AdjustedWeight.IdealWeight,
			Result:      formulas.AdjustedWeight.Result,
		}
	}

	var percentageWeightAdequacy dtos.PercentageWeightAdequacy
	if formulas.PercentageWeightAdequacy != nil {
		percentageWeightAdequacy = dtos.PercentageWeightAdequacy{
			Classification: string(formulas.PercentageWeightAdequacy.Classification),
			Result:         formulas.PercentageWeightAdequacy.Result,
		}
	}

	var percentageWeightChange dtos.PercentageWeightChange
	if formulas.PercentageWeightChange != nil {
		percentageWeightChange = dtos.PercentageWeightChange{
			Classification: string(formulas.PercentageWeightChange.Classification),
			Result:         formulas.PercentageWeightChange.Result,
		}
	}

	var eer dtos.EER
	if formulas.EER != nil {
		eer = dtos.EER{
			Result: formulas.EER.Result,
		}
	}

	var tmb dtos.TMB
	if formulas.TMB != nil {
		tmb = dtos.TMB{
			Result: formulas.TMB.Result,
		}
	}

	return dtos.Formulas{
		IMC:                      imc,
		AdjustedWeightObesity:    adjustedWeight,
		PercentageWeightAdequacy: percentageWeightAdequacy,
		PercentageWeightChange:   percentageWeightChange,
		EER:                      eer,
		TMB:                      tmb,
	}
}

