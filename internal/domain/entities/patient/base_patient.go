package patient

import (
	"fmt"
	"strings"

	"github.com/Arthur-Conti/gi_nutri/internal/domain/entities/formulas"
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

func (p *BasePatient) GetData() PatientData {
	return PatientData{
		ID:                         p.ID,
		Name:                       p.Name,
		Age:                        p.Age,
		AgeClassification:          p.AgeClassification,
		SchofieldAgeClassification: p.SchofieldAgeClassification,
		Sex:                        p.Sex,
		Height:                     p.Height,
		Weight:                     p.Weight,
		UsualWeight:                p.UsualWeight,
		TimeDays:                   p.TimeDays,
		PhysicalActivity:           p.PhysicalActivity,
		PhysicalActivityResult:     p.PhysicalActivityResult,
		IsPregnant:                 false, // Será sobrescrito pelas implementações específicas
		PregnancyInfo:              PregnancyInfo{},
		IsLactating:                false, // Será sobrescrito pelas implementações específicas
		LactatingInfo:              LactatingInfo{},
		Results:                    p.Results,
	}
}

func (p *BasePatient) SetResults(results Results) {
	p.Results = results
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

func (p *BasePatient) CalculateTMB(choices formulas.TMBFormulas) error {
	if p.Results.Measures.Weight <= 0 {
		return NewValidationError("Weight", "Weight must be greater than 0")
	}
	if p.Results.Measures.HeightCM <= 0 {
		return NewValidationError("HeightCM", "Height must be greater than 0")
	}

	if choices.Pocket && choices.PocketValue == 0 {
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
	p.Results.FormulasInfo.TMB.Calculate(choices)
	return nil
}

func (p *BasePatient) GetPhysicalActivityResult() {
	p.PhysicalActivityResult = 1.0
}

