package patient

import (
	"strings"

	"github.com/Arthur-Conti/gi_nutri/internal/domain/entities/formulas"
	patientrepository "github.com/Arthur-Conti/gi_nutri/internal/infra/repositories/patient"
	resultsrepository "github.com/Arthur-Conti/gi_nutri/internal/infra/repositories/results"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BasePatient struct {
	ID                         string
	ResultsID                  string
	Name                       string
	Age                        int
	AgeClassification          PatientAgeClassification
	SchofieldAgeClassification SchofieldAgeClassification
	Sex                        PatientSex
	UsualWeight                float64
	TimeDays                   int
	PhysicalActivity           PhysicalActivity
	PhysicalActivityResult     float64
	Measures                   Measures
	FormulasInfo               FormulasInfo
}

func NewBasePatient(opts PatientOpts) *BasePatient {
	return &BasePatient{
		Name:                       opts.Name,
		Age:                        opts.Age,
		AgeClassification:          ClassifyAge(opts.Age),
		SchofieldAgeClassification: ClassifyAgeSchofield(opts.Age),
		Sex:                        opts.Sex,
		Measures:                   opts.Measures,
		TimeDays:                   opts.TimeDays,
		PhysicalActivity:           opts.PhysicalActivity,
		UsualWeight:                opts.UsualWeight,
	}
}

func (p *BasePatient) GetFormulas() FormulasInfo {
	return p.FormulasInfo
}

func (p *BasePatient) CalculateIMC() error {
	if p.Measures.Weight <= 0 {
		return NewValidationError("Weight", "Weight must be greater than 0")
	}
	if p.Measures.HeightM <= 0 {
		return NewValidationError("HeightM", "Height must be greater than 0")
	}

	p.FormulasInfo.IMC = formulas.NewImc(
		p.Measures.Weight,
		p.Measures.HeightM,
		string(p.AgeClassification),
	)
	p.FormulasInfo.IMC.Calculate()
	return nil
}

func (p *BasePatient) CalculateAdjustedWeight() error {
	if p.FormulasInfo.IMC == nil {
		return NewFormulaDependencyError(
			"AdjustedWeight",
			"IMC",
			"Calculate IMC first",
		)
	}

	if !strings.HasPrefix(string(p.FormulasInfo.IMC.Status), "obesity") {
		return NewValidationError(
			"IMC.Status",
			"AdjustedWeight formula is only applicable for obese patients",
		)
	}

	p.FormulasInfo.AdjustedWeight = formulas.NewAdjustedWeightObesity(
		p.Measures.Weight,
		string(p.AgeClassification),
		*p.FormulasInfo.IMC,
	)
	return nil
}

func (p *BasePatient) CalculatePercentageWeightAdequacy() error {
	if p.FormulasInfo.AdjustedWeight == nil {
		return NewFormulaDependencyError(
			"PercentageWeightAdequacy",
			"AdjustedWeight",
			"Calculate AdjustedWeight first",
		)
	}

	if p.TimeDays == 0 {
		return NewValidationError(
			"TimeDays",
			"TimeDays is required to calculate PercentageWeightAdequacy",
		)
	}

	p.FormulasInfo.PercentageWeightAdequacy = formulas.NewPercentageWeightAdequacy(
		p.Measures.Weight,
		p.FormulasInfo.AdjustedWeight.IdealWeight,
	)
	return nil
}

func (p *BasePatient) CalculatePercentageWeightChange() error {
	if p.UsualWeight == 0 {
		return NewValidationError(
			"UsualWeight",
			"UsualWeight is required to calculate PercentageWeightChange",
		)
	}

	if p.TimeDays == 0 {
		return NewValidationError(
			"TimeDays",
			"TimeDays is required to calculate PercentageWeightChange",
		)
	}

	p.FormulasInfo.PercentageWeightChange = formulas.NewPercentageWeightChange(
		p.UsualWeight,
		p.Measures.Weight,
		p.TimeDays,
	)
	return nil
}

func (p *BasePatient) CalculateEER() error {
	if p.PhysicalActivityResult == 0 {
		return NewValidationError(
			"PhysicalActivityResult",
			"PhysicalActivityResult must be calculated first. Call GetPhysicalActivityResult()",
		)
	}

	p.FormulasInfo.EER = formulas.NewEER(
		p.Age,
		string(p.AgeClassification),
		string(p.Sex),
		p.Measures.Weight,
		p.Measures.HeightM,
		p.PhysicalActivityResult,
	)
	p.FormulasInfo.EER.Calculate()
	return nil
}

func (p *BasePatient) CalculateTMB(useHarrisBenedict, useFao, useSchofield, usePocket bool, pocketValue float64) error {
	if p.Measures.Weight <= 0 {
		return NewValidationError("Weight", "Weight must be greater than 0")
	}
	if p.Measures.HeightCM <= 0 {
		return NewValidationError("HeightCM", "Height must be greater than 0")
	}

	if usePocket && pocketValue == 0 {
		return NewValidationError(
			"pocketValue",
			"pocketValue must be provided when using pocket method",
		)
	}

	p.FormulasInfo.TMB = formulas.NewTMB(
		p.Measures.Weight,
		p.Measures.HeightCM,
		p.Age,
		string(p.SchofieldAgeClassification),
		string(p.Sex),
	)
	p.FormulasInfo.TMB.Calculate(useHarrisBenedict, useFao, useSchofield, usePocket, pocketValue)
	return nil
}

func (p *BasePatient) GetPhysicalActivityResult() {
	p.PhysicalActivityResult = 1.0
}

func (p *BasePatient) PatientToModel() patientrepository.PatientModel {
	id, _ := primitive.ObjectIDFromHex(p.ID)
	return patientrepository.PatientModel{
		ID:                         id,
		Name:                       p.Name,
		Age:                        p.Age,
		AgeClassification:          string(p.AgeClassification),
		SchofieldAgeClassification: string(p.SchofieldAgeClassification),
		Sex:                        string(p.Sex),
		UsualWeight:                p.UsualWeight,
		PhysicalActivity:           string(p.PhysicalActivity),
	}
}

func (p *BasePatient) ResultsToModel() resultsrepository.ResultsModel {
	resultsID, _ := primitive.ObjectIDFromHex(p.ResultsID)
	patientID, _ := primitive.ObjectIDFromHex(p.ID)
	return resultsrepository.ResultsModel{
		ID:        resultsID,
		PatientID: patientID,
		Measures: resultsrepository.Measures{
			HeightCM: p.Measures.HeightCM,
			HeightM:  p.Measures.HeightM,
			Weight:   p.Measures.Weight,
		},
		Formulas: resultsrepository.Formulas{
			IMC: resultsrepository.IMC{
				Status: string(p.FormulasInfo.IMC.Status),
				Result: p.FormulasInfo.IMC.Result,
			},
			AdjustedWeightObesity: resultsrepository.AdjustedWeightObesity{
				IdealWeight: p.FormulasInfo.AdjustedWeight.IdealWeight,
				Result:      p.FormulasInfo.AdjustedWeight.Result,
			},
			PercentageWeightAdequacy: resultsrepository.PercentageWeightAdequacy{
				Classification: string(p.FormulasInfo.PercentageWeightAdequacy.Classification),
				Result:         p.FormulasInfo.PercentageWeightAdequacy.Result,
			},
			PercentageWeightChange: resultsrepository.PercentageWeightChange{
				Classification: string(p.FormulasInfo.PercentageWeightChange.Classification),
				Result:         p.FormulasInfo.PercentageWeightChange.Result,
			},
			EER: resultsrepository.EER{
				Result: p.FormulasInfo.EER.Result,
			},
			TMB: resultsrepository.TMB{
				Result: p.FormulasInfo.TMB.Result,
			},
		},
	}
}
