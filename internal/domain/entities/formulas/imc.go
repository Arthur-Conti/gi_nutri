package formulas

import (
	"math"
)

type IMC struct {
	Status            IMCStatus
	Result            float64
	weight            float64
	heightM           float64
	ageClassification string
}

type IMCStatus string

const (
	StatusMalnutrition IMCStatus = "malnutrition"
	StatusEutrophy     IMCStatus = "eutrophy"
	StatusOverweight   IMCStatus = "overweight"
	StatusObesityI     IMCStatus = "obesityI"
	StatusObesityII    IMCStatus = "obesityII"
	StatusObesityIII   IMCStatus = "obesityIII"
)

func NewImc(weight, heightM float64, ageClassification string) *IMC {
	return &IMC{
		weight:            weight,
		heightM:           heightM,
		ageClassification: ageClassification,
	}
}

func (i *IMC) ReverseImc(desiredResult float64) float64 {
	heightPow := math.Pow(i.heightM, 2)
	return desiredResult * heightPow
}

func (i *IMC) Calculate() {
	heightPow := math.Pow(i.heightM, 2)
	result := i.weight / heightPow

	switch i.ageClassification {
	case "adult":
		if result < 18.5 {
			i.Status = StatusMalnutrition
			i.Result = result
		} else if result >= 18.5 && result <= 24.9 {
			i.Status = StatusEutrophy
			i.Result = result
		} else if result >= 25.0 && result <= 29.9 {
			i.Status = StatusOverweight
			i.Result = result
		} else if result >= 30.0 && result <= 34.9 {
			i.Status = StatusObesityI
			i.Result = result
		} else if result >= 35.0 && result <= 39.9 {
			i.Status = StatusObesityII
			i.Result = result
		} else if result >= 40.0 {
			i.Status = StatusObesityIII
			i.Result = result
		}
	case "elderly":
		if result < 22.0 {
			i.Status = StatusMalnutrition
			i.Result = result
		} else if result >= 22.0 && result <= 26.9 {
			i.Status = StatusEutrophy
			i.Result = result
		} else if result >= 27.0 && result <= 29.9 {
			i.Status = StatusOverweight
			i.Result = result
		} else if result >= 30.0 && result <= 34.9 {
			i.Status = StatusObesityI
			i.Result = result
		} else if result >= 35.0 && result <= 39.9 {
			i.Status = StatusObesityII
			i.Result = result
		} else if result >= 40.0 {
			i.Status = StatusObesityIII
			i.Result = result
		}
	}
}
