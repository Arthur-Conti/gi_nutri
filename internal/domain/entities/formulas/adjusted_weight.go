package formulas

type AdjustedWeightObesity struct {
	Result      float64
	IdealWeight float64
	imc         IMC
}

func NewAdjustedWeightObesity(weight float64, ageClassification string, imc IMC) *AdjustedWeightObesity {
	idealWeight := calculateIdealWeightObesity(imc, ageClassification)
	return &AdjustedWeightObesity{
		Result:      calculateAdjustedWeightObesity(weight, idealWeight),
		IdealWeight: idealWeight,
	}
}

func calculateAdjustedWeightObesity(weight, idealWeight float64) float64 {
	wi := weight - idealWeight
	return (wi * 0.33) + idealWeight
}

func calculateIdealWeightObesity(imc IMC, ageClassification string) float64 {
	switch ageClassification {
	case "children/teenage":
		return 0.0
	case "adult":
		return imc.ReverseImc(25.0)
	case "elderly":
		return imc.ReverseImc(27.0)
	}
	return 0.0
}
