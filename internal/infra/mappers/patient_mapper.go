package mappers

import (
	"time"

	"github.com/Arthur-Conti/gi_nutri/internal/domain/entities/formulas"
	"github.com/Arthur-Conti/gi_nutri/internal/domain/entities/patient"
	patientrepository "github.com/Arthur-Conti/gi_nutri/internal/infra/repositories/patient"
	resultsrepository "github.com/Arthur-Conti/gi_nutri/internal/infra/repositories/results"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PatientEntityToModel converte uma entidade Patient para PatientModel
func PatientEntityToModel(p patient.Patient) patientrepository.PatientModel {
	data := p.GetData()
	return patientrepository.PatientModel{
		Name:                       data.Name,
		Age:                        data.Age,
		AgeClassification:          string(data.AgeClassification),
		SchofieldAgeClassification: string(data.SchofieldAgeClassification),
		TimeDays:                   data.TimeDays,
		Sex:                        string(data.Sex),
		Height:                     data.Height,
		Weight:                     data.Weight,
		UsualWeight:                data.UsualWeight,
		PhysicalActivity:           string(data.PhysicalActivity),
		IsPregnant:                 data.IsPregnant,
		PregnancyInfo:              patientrepository.PregnancyInfo(data.PregnancyInfo),
		IsLactating:                data.IsLactating,
		LactatingInfo:              patientrepository.LactatingInfo(data.LactatingInfo),
	}
}

// PatientModelToEntity converte um PatientModel para entidade Patient
func PatientModelToEntity(model patientrepository.PatientModel) patient.Patient {
	opts := patient.PatientOpts{
		Name:             model.Name,
		Age:              model.Age,
		TimeDays:         model.TimeDays,
		Sex:              patient.PatientSex(model.Sex),
		Height:           model.Height,
		Weight:           model.Weight,
		UsualWeight:      model.UsualWeight,
		PhysicalActivity: patient.PhysicalActivity(model.PhysicalActivity),
		IsPregnant:       model.IsPregnant,
		PregnancyInfo:    patient.PregnancyInfo(model.PregnancyInfo),
		IsLactating:      model.IsLactating,
		LactatingInfo:    patient.LactatingInfo(model.LactatingInfo),
	}
	return patient.NewPatient(opts)
}

// ResultsEntityToModel converte Results da entidade para ResultsModel
func ResultsEntityToModel(results patient.Results, patientID string) resultsrepository.ResultsModel {
	var objectID primitive.ObjectID
	if patientID != "" {
		objectID, _ = primitive.ObjectIDFromHex(patientID)
	}

	measures := resultsrepository.Measures{
		HeightCM: results.Measures.HeightCM,
		HeightM:  results.Measures.HeightM,
		Weight:   results.Measures.Weight,
	}

	var imc resultsrepository.IMC
	if results.FormulasInfo.IMC != nil {
		imc = resultsrepository.IMC{
			Status: string(results.FormulasInfo.IMC.Status),
			Result: results.FormulasInfo.IMC.Result,
		}
	}

	var adjustedWeight resultsrepository.AdjustedWeightObesity
	if results.FormulasInfo.AdjustedWeight != nil {
		adjustedWeight = resultsrepository.AdjustedWeightObesity{
			IdealWeight: results.FormulasInfo.AdjustedWeight.IdealWeight,
			Result:      results.FormulasInfo.AdjustedWeight.Result,
		}
	}

	var percentageWeightAdequacy resultsrepository.PercentageWeightAdequacy
	if results.FormulasInfo.PercentageWeightAdequacy != nil {
		percentageWeightAdequacy = resultsrepository.PercentageWeightAdequacy{
			Classification: string(results.FormulasInfo.PercentageWeightAdequacy.Classification),
			Result:         results.FormulasInfo.PercentageWeightAdequacy.Result,
		}
	}

	var percentageWeightChange resultsrepository.PercentageWeightChange
	if results.FormulasInfo.PercentageWeightChange != nil {
		percentageWeightChange = resultsrepository.PercentageWeightChange{
			Classification: string(results.FormulasInfo.PercentageWeightChange.Classification),
			Result:         results.FormulasInfo.PercentageWeightChange.Result,
		}
	}

	var eer resultsrepository.EER
	if results.FormulasInfo.EER != nil {
		eer = resultsrepository.EER{
			Result: results.FormulasInfo.EER.Result,
		}
	}

	var tmb resultsrepository.TMB
	if results.FormulasInfo.TMB != nil {
		tmb = resultsrepository.TMB{
			Result: results.FormulasInfo.TMB.Result,
		}
	}

	formulas := resultsrepository.Formulas{
		IMC:                      imc,
		AdjustedWeightObesity:    adjustedWeight,
		PercentageWeightAdequacy: percentageWeightAdequacy,
		PercentageWeightChange:   percentageWeightChange,
		EER:                      eer,
		TMB:                      tmb,
	}

	return resultsrepository.ResultsModel{
		PatientID:  objectID,
		RecordedAt: time.Now(),
		Measures:   measures,
		Formulas:   formulas,
	}
}

// ResultsModelToEntity converte ResultsModel para Results da entidade
func ResultsModelToEntity(model resultsrepository.ResultsModel) patient.Results {
	var imc *formulas.IMC
	if model.Formulas.IMC.Result != 0 || model.Formulas.IMC.Status != "" {
		imc = &formulas.IMC{
			Status: formulas.IMCStatus(model.Formulas.IMC.Status),
			Result: model.Formulas.IMC.Result,
		}
	}

	var adjustedWeight *formulas.AdjustedWeightObesity
	if model.Formulas.AdjustedWeightObesity.Result != 0 || model.Formulas.AdjustedWeightObesity.IdealWeight != 0 {
		adjustedWeight = &formulas.AdjustedWeightObesity{
			IdealWeight: model.Formulas.AdjustedWeightObesity.IdealWeight,
			Result:      model.Formulas.AdjustedWeightObesity.Result,
		}
	}

	var percentageWeightAdequacy *formulas.PercentageWeightAdequacy
	if model.Formulas.PercentageWeightAdequacy.Result != 0 || model.Formulas.PercentageWeightAdequacy.Classification != "" {
		percentageWeightAdequacy = &formulas.PercentageWeightAdequacy{
			Classification: formulas.WeightAdequacyClassification(model.Formulas.PercentageWeightAdequacy.Classification),
			Result:         model.Formulas.PercentageWeightAdequacy.Result,
		}
	}

	var percentageWeightChange *formulas.PercentageWeightChange
	if model.Formulas.PercentageWeightChange.Result != 0 || model.Formulas.PercentageWeightChange.Classification != "" {
		percentageWeightChange = &formulas.PercentageWeightChange{
			Classification: formulas.WeightChangeClassification(model.Formulas.PercentageWeightChange.Classification),
			Result:         model.Formulas.PercentageWeightChange.Result,
		}
	}

	var eer *formulas.EER
	if model.Formulas.EER.Result != 0 {
		eer = &formulas.EER{
			Result: model.Formulas.EER.Result,
		}
	}

	var tmb *formulas.TMB
	if model.Formulas.TMB.Result != 0 {
		tmb = &formulas.TMB{
			Result: model.Formulas.TMB.Result,
		}
	}

	formulasInfo := patient.FormulasInfo{
		IMC:                      imc,
		AdjustedWeight:           adjustedWeight,
		PercentageWeightAdequacy: percentageWeightAdequacy,
		PercentageWeightChange:   percentageWeightChange,
		EER:                      eer,
		TMB:                      tmb,
	}

	return patient.Results{
		ResultsID:    model.ID.Hex(),
		Measures:     patient.Measures(model.Measures),
		FormulasInfo: formulasInfo,
	}
}
