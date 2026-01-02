package patient

import (
	"fmt"
	"strings"

	"github.com/Arthur-Conti/gi_nutri/internal/domain/entities/formulas"
	patientrepository "github.com/Arthur-Conti/gi_nutri/internal/infra/repositories/patient"
	resultsrepository "github.com/Arthur-Conti/gi_nutri/internal/infra/repositories/results"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BasePatient struct {
	ID                         string
	Name                       string
	Age                        int
	AgeClassification          PatientAgeClassification
	SchofieldAgeClassification SchofieldAgeClassification
	Sex                        PatientSex
	Height                     float64
	Weight                     float64
	UsualWeight                float64
	TimeDays                   int
	PhysicalActivity           PhysicalActivity
	PhysicalActivityResult     float64
	Results                    Results
}

func NewBasePatient(opts PatientOpts) *BasePatient {
	return &BasePatient{
		Name:                       opts.Name,
		Age:                        opts.Age,
		AgeClassification:          ClassifyAge(opts.Age),
		SchofieldAgeClassification: ClassifyAgeSchofield(opts.Age),
		Sex:                        opts.Sex,
		TimeDays:                   opts.TimeDays,
		PhysicalActivity:           opts.PhysicalActivity,
		Height:                     opts.Height,
		Weight:                     opts.Weight,
		UsualWeight:                opts.UsualWeight,
		Results: Results{
			Measures: Measures{
				HeightCM: opts.Height,
				HeightM:  heightMeters(opts.Height),
				Weight:   opts.Weight,
			},
		},
	}
}

func (p *BasePatient) GetResults() Results {
	return p.Results
}

func (p *BasePatient) CalculateIMC() error {
	if p.Results.Measures.Weight <= 0 {
		return NewValidationError("Weight", "Weight must be greater than 0")
	}
	if p.Results.Measures.HeightM <= 0 {
		return NewValidationError("HeightM", "Height must be greater than 0")
	}

	p.Results.FormulasInfo.IMC = formulas.NewImc(
		p.Results.Measures.Weight,
		p.Results.Measures.HeightM,
		string(p.AgeClassification),
	)
	p.Results.FormulasInfo.IMC.Calculate()

	return nil
}

func (p *BasePatient) CalculateAdjustedWeight() error {
	if p.Results.FormulasInfo.IMC == nil {
		return NewFormulaDependencyError(
			"AdjustedWeight",
			"IMC",
			"Calculate IMC first",
		)
	}

	if !strings.HasPrefix(string(p.Results.FormulasInfo.IMC.Status), "obesity") {
		return NewValidationError(
			"IMC.Status",
			"AdjustedWeight formula is only applicable for obese patients",
		)
	}

	fmt.Printf("%+v\n", *p.Results.FormulasInfo.IMC)

	p.Results.FormulasInfo.AdjustedWeight = formulas.NewAdjustedWeightObesity(
		p.Results.Measures.Weight,
		string(p.AgeClassification),
		*p.Results.FormulasInfo.IMC,
	)
	return nil
}

func (p *BasePatient) CalculatePercentageWeightAdequacy() error {
	if p.Results.FormulasInfo.AdjustedWeight == nil {
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

	if p.Results.FormulasInfo.AdjustedWeight.IdealWeight == 0 {
		return NewFormulaDependencyError(
			"PercentageWeightAdequacy",
			"AdjustedWeight",
			"Calculate AdjutedWeight first",
		)
	}

	p.Results.FormulasInfo.PercentageWeightAdequacy = formulas.NewPercentageWeightAdequacy(
		p.Results.Measures.Weight,
		p.Results.FormulasInfo.AdjustedWeight.IdealWeight,
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

	p.Results.FormulasInfo.PercentageWeightChange = formulas.NewPercentageWeightChange(
		p.UsualWeight,
		p.Results.Measures.Weight,
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

	p.Results.FormulasInfo.EER = formulas.NewEER(
		p.Age,
		string(p.AgeClassification),
		string(p.Sex),
		p.Results.Measures.Weight,
		p.Results.Measures.HeightM,
		p.PhysicalActivityResult,
	)
	p.Results.FormulasInfo.EER.Calculate()
	return nil
}

func (p *BasePatient) CalculateTMB(useHarrisBenedict, useFao, useSchofield, usePocket bool, pocketValue float64) error {
	if p.Results.Measures.Weight <= 0 {
		return NewValidationError("Weight", "Weight must be greater than 0")
	}
	if p.Results.Measures.HeightCM <= 0 {
		return NewValidationError("HeightCM", "Height must be greater than 0")
	}

	if usePocket && pocketValue == 0 {
		return NewValidationError(
			"pocketValue",
			"pocketValue must be provided when using pocket method",
		)
	}

	p.Results.FormulasInfo.TMB = formulas.NewTMB(
		p.Results.Measures.Weight,
		p.Results.Measures.HeightCM,
		p.Age,
		string(p.SchofieldAgeClassification),
		string(p.Sex),
	)
	p.Results.FormulasInfo.TMB.Calculate(useHarrisBenedict, useFao, useSchofield, usePocket, pocketValue)
	return nil
}

func (p *BasePatient) GetPhysicalActivityResult() {
	p.PhysicalActivityResult = 1.0
}

func (p *BasePatient) PatientToModel() patientrepository.PatientModel {
	return patientrepository.PatientModel{
		Name:                       p.Name,
		Age:                        p.Age,
		AgeClassification:          string(p.AgeClassification),
		SchofieldAgeClassification: string(p.SchofieldAgeClassification),
		TimeDays:                   p.TimeDays,
		Sex:                        string(p.Sex),
		Height:                     p.Height,
		Weight:                     p.Weight,
		UsualWeight:                p.UsualWeight,
		PhysicalActivity:           string(p.PhysicalActivity),
	}
}

func (p *BasePatient) ResultsToModel() resultsrepository.ResultsModel {
	patientID, _ := primitive.ObjectIDFromHex(p.ID)

	measures := resultsrepository.Measures{
		HeightCM: p.Results.Measures.HeightCM,
		HeightM:  p.Results.Measures.HeightM,
		Weight:   p.Results.Measures.Weight,
	}

	var imc resultsrepository.IMC
	if p.Results.FormulasInfo.IMC != nil {
		imc = resultsrepository.IMC{
			Status: string(p.Results.FormulasInfo.IMC.Status),
			Result: p.Results.FormulasInfo.IMC.Result,
		}
	}

	var adjustedWeight resultsrepository.AdjustedWeightObesity
	if p.Results.FormulasInfo.AdjustedWeight != nil {
		adjustedWeight = resultsrepository.AdjustedWeightObesity{
			IdealWeight: p.Results.FormulasInfo.AdjustedWeight.IdealWeight,
			Result:      p.Results.FormulasInfo.AdjustedWeight.Result,
		}
	}

	var percentageWeightAdequacy resultsrepository.PercentageWeightAdequacy
	if p.Results.FormulasInfo.PercentageWeightAdequacy != nil {
		percentageWeightAdequacy = resultsrepository.PercentageWeightAdequacy{
			Classification: string(p.Results.FormulasInfo.PercentageWeightAdequacy.Classification),
			Result:         p.Results.FormulasInfo.PercentageWeightAdequacy.Result,
		}
	}

	var percentageWeightChange resultsrepository.PercentageWeightChange
	if p.Results.FormulasInfo.PercentageWeightChange != nil {
		percentageWeightChange = resultsrepository.PercentageWeightChange{
			Classification: string(p.Results.FormulasInfo.PercentageWeightChange.Classification),
			Result:         p.Results.FormulasInfo.PercentageWeightChange.Result,
		}
	}

	var eer resultsrepository.EER
	if p.Results.FormulasInfo.EER != nil {
		eer = resultsrepository.EER{
			Result: p.Results.FormulasInfo.EER.Result,
		}
	}

	var tmb resultsrepository.TMB
	if p.Results.FormulasInfo.TMB != nil {
		tmb = resultsrepository.TMB{
			Result: p.Results.FormulasInfo.TMB.Result,
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
		PatientID: patientID,
		Measures:  measures,
		Formulas:  formulas,
	}
}

func (p *BasePatient) FillResults(model resultsrepository.ResultsModel) {
	imc := formulas.IMC{
		Status: formulas.IMCStatus(model.Formulas.IMC.Status),
		Result: model.Formulas.IMC.Result,
	}

	adjustedWeight := formulas.AdjustedWeightObesity{
		IdealWeight: model.Formulas.AdjustedWeightObesity.IdealWeight,
		Result:      model.Formulas.AdjustedWeightObesity.Result,
	}

	percentageWeightAdequacy := formulas.PercentageWeightAdequacy{
		Classification: formulas.WeightAdequacyClassification(model.Formulas.PercentageWeightAdequacy.Classification),
		Result:         model.Formulas.PercentageWeightAdequacy.Result,
	}

	percentageWeightChange := formulas.PercentageWeightChange{
		Classification: formulas.WeightChangeClassification(model.Formulas.PercentageWeightChange.Classification),
		Result:         model.Formulas.PercentageWeightChange.Result,
	}

	eer := formulas.EER{
		Result: model.Formulas.EER.Result,
	}

	tmb := formulas.TMB{
		Result: model.Formulas.TMB.Result,
	}

	formulas := FormulasInfo{
		IMC:                      &imc,
		AdjustedWeight:           &adjustedWeight,
		PercentageWeightAdequacy: &percentageWeightAdequacy,
		PercentageWeightChange:   &percentageWeightChange,
		EER:                      &eer,
		TMB:                      &tmb,
	}

	p.Results = Results{
		ResultsID:    model.ID.Hex(),
		Measures:     Measures(model.Measures),
		FormulasInfo: formulas,
	}
}
