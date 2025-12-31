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
	StatusMalnutrition IMCStatus = "Desnutrição"
	StatusEutrophy     IMCStatus = "Eutrofia"
	StatusOverweight   IMCStatus = "Sobrepeso"
	StatusObesityI     IMCStatus = "Obesidade I"
	StatusObesityII    IMCStatus = "Obesidade II"
	StatusObesityIII   IMCStatus = "Obesidade III"
)

func NewImc(weight, heightM float64, ageClassification string) *IMC {
	status, result := calculateIMC(weight, heightM, ageClassification)
	return &IMC{
		Status:            status,
		Result:            result,
		weight:            weight,
		heightM:           heightM,
		ageClassification: ageClassification,
	}
}

func (i *IMC) ReverseImc(desiredResult float64) float64 {
	heightPow := math.Pow(i.heightM, 2)
	return desiredResult * heightPow
}

func calculateIMC(weight, height float64, ageClassification string) (IMCStatus, float64) {
	heightPow := math.Pow(height, 2)
	result := weight / heightPow

	switch ageClassification {
	case "adult":
		if result < 18.5 {
			return StatusMalnutrition, result
		} else if result >= 18.5 && result <= 24.9 {
			return StatusEutrophy, result
		} else if result >= 25.0 && result <= 29.9 {
			return StatusOverweight, result
		} else if result >= 30.0 && result <= 34.9 {
			return StatusObesityI, result
		} else if result >= 35.0 && result <= 39.9 {
			return StatusObesityII, result
		} else if result >= 40.0 {
			return StatusObesityIII, result
		}
	case "elderly":
		if result < 22.0 {
			return StatusMalnutrition, result
		} else if result >= 22.0 && result <= 26.9 {
			return StatusEutrophy, result
		} else if result >= 27.0 && result <= 29.9 {
			return StatusOverweight, result
		} else if result >= 30.0 && result <= 34.9 {
			return StatusObesityI, result
		} else if result >= 35.0 && result <= 39.9 {
			return StatusObesityII, result
		} else if result >= 40.0 {
			return StatusObesityIII, result
		}
	}
	return "", 0.0
}
