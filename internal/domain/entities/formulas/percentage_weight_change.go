package formulas

type PercentageWeightChange struct {
	Classification WeightChangeClassification
	Result         float64
}

type WeightChangeClassification string

const (
	NonSignificantClassification WeightChangeClassification = "non_significant"
	SignificantClassification    WeightChangeClassification = "significant"
	SevereClassification         WeightChangeClassification = "severe"
)

func NewPercentageWeightChange(usualWeight, currentWeight float64, time int) *PercentageWeightChange {
	classification, result := calculatePercentageWeightChange(usualWeight, currentWeight, time)
	return &PercentageWeightChange{
		Classification: classification,
		Result:         result,
	}
}

func calculatePercentageWeightChange(usualWeight, currentWeight float64, time int) (WeightChangeClassification, float64) {
	pupa := (usualWeight - currentWeight) / usualWeight
	result := pupa * 100

	switch time {
	case 7:
		if result >= 1.0 && result <= 2.0 {
			return SignificantClassification, result
		} else if result > 2.0 {
			return SevereClassification, result
		} else {
			return NonSignificantClassification, result
		}
	case 30:
		if result == 5.0 {
			return SignificantClassification, result
		} else if result > 5.0 {
			return SevereClassification, result
		} else {
			return NonSignificantClassification, result
		}
	case 90:
		if result == 7.5 {
			return SignificantClassification, result
		} else if result > 7.5 {
			return SevereClassification, result
		} else {
			return NonSignificantClassification, result
		}
	case 180:
		if result == 10.0 {
			return SignificantClassification, result
		} else if result > 10.0 {
			return SevereClassification, result
		} else {
			return NonSignificantClassification, result
		}
	}

	return "", 0.0
}
