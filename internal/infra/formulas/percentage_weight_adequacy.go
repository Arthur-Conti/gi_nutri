package formulas

type PercentageWeightAdequacy struct {
	Classification WeightAdequacyClassification
	Result         float64
}

type WeightAdequacyClassification string

const (
	SevereMalnutrition   WeightAdequacyClassification = "severe_malnutrition"
	ModerateMalnutrition WeightAdequacyClassification = "moderate_malnutrition"
	SoftMalnutrition     WeightAdequacyClassification = "soft_malnutrition"
	Eutrophy             WeightAdequacyClassification = "eutrophy"
	Overweight           WeightAdequacyClassification = "overweight"
	Obesity              WeightAdequacyClassification = "obesity"
)

func NewPercentageWeightAdequacy(weight, idealWeight float64) *PercentageWeightAdequacy {
	classification, result := calculatePercentageWeightAdequacy(weight, idealWeight)
	return &PercentageWeightAdequacy{
		Classification: classification,
		Result:         result,
	}
}

func calculatePercentageWeightAdequacy(weight, idealWeight float64) (WeightAdequacyClassification, float64) {
	wi := weight / idealWeight
	result := wi * 100

	if result <= 70 {
		return SevereMalnutrition, result
	} else if result >= 70.1 && result <= 80.0 {
		return ModerateMalnutrition, result
	} else if result >= 80.1 && result <= 90.0 {
		return SoftMalnutrition, result
	} else if result >= 90.1 && result <= 110.0 {
		return Eutrophy, result
	} else if result >= 110.1 && result <= 120.0 {
		return Overweight, result
	} else if result > 120.0 {
		return Obesity, result
	}

	return "", 0.0
}
