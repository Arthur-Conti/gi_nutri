package patient

import "github.com/Arthur-Conti/gi_nutri/internal/infra/formulas"

type Patient struct {
	Name              string
	Age               int
	AgeClassification PatientAgeClassification
	HeightCM          float64
	HeightM           float64
	Weight            float64
	UsualWeight       float64
	TimeDays int
	FormulasInfo      FormulasInfo
}

type FormulasInfo struct {
	IMC                      *formulas.IMC
	AdjustedWeight           *formulas.AdjustedWeightObesity
	PercentageWeightAdequacy *formulas.PercentageWeightAdequacy
	PercentageWeightChange   *formulas.PercentageWeightChange
}

type PatientAgeClassification string

const (
	PatientChildrenTeenage PatientAgeClassification = "children/teenage"
	PatientAdult           PatientAgeClassification = "adult"
	PatientElderly         PatientAgeClassification = "elderly"
)

func NewPatient(name string, age, timeDays int, height, weight, usualWeight float64) *Patient {
	return &Patient{
		Name:              name,
		Age:               age,
		AgeClassification: ClassifyAge(age),
		HeightCM:          height,
		HeightM:           HeightMeters(height),
		Weight:            weight,
		TimeDays: timeDays,
		UsualWeight: usualWeight,
	}
}

func (p *Patient) CalculateIMC() {
	p.FormulasInfo.IMC = formulas.NewImc(p.Weight, p.HeightM, string(p.AgeClassification))
}

func (p *Patient) CalculateAdjustedWeight() {
	p.FormulasInfo.AdjustedWeight = formulas.NewAdjustedWeightObesity(p.Weight, string(p.AgeClassification), *p.FormulasInfo.IMC)
}

func (p *Patient) CalculatePercentageWeightAdequacy() {
	p.FormulasInfo.PercentageWeightAdequacy = formulas.NewPercentageWeightAdequacy(p.Weight, p.FormulasInfo.AdjustedWeight.IdealWeight)
}

func (p *Patient) CalculatePercentageWeightChange() {
	p.FormulasInfo.PercentageWeightChange = formulas.NewPercentageWeightChange(p.UsualWeight, p.Weight, p.TimeDays)
}

func ClassifyAge(age int) PatientAgeClassification {
	if age <= 19 {
		return PatientChildrenTeenage
	} else if age >= 20 && age <= 59 {
		return PatientAdult
	} else {
		return PatientElderly
	}
}

func HeightMeters(height float64) float64 {
	return height / 100
}
