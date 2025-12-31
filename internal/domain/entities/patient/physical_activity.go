package patient

type PhysicalActivityMultipliers struct {
	ChildrenTeenage map[PhysicalActivity]float64
	AdultElderly    map[PhysicalActivity]float64
}

var (
	MalePhysicalActivityMultipliers = PhysicalActivityMultipliers{
		ChildrenTeenage: map[PhysicalActivity]float64{
			PhysicalActivitySedentary:    1.0,
			PhysicalActivityLowActivity:  1.13,
			PhysicalActivityActive:       1.26,
			PhysicalActivityHighActivity: 1.42,
		},
		AdultElderly: map[PhysicalActivity]float64{
			PhysicalActivitySedentary:    1.0,
			PhysicalActivityLowActivity:  1.11,
			PhysicalActivityActive:       1.25,
			PhysicalActivityHighActivity: 1.48,
		},
	}

	FemalePhysicalActivityMultipliers = PhysicalActivityMultipliers{
		ChildrenTeenage: map[PhysicalActivity]float64{
			PhysicalActivitySedentary:    1.0,
			PhysicalActivityLowActivity:  1.16,
			PhysicalActivityActive:       1.31,
			PhysicalActivityHighActivity: 1.56,
		},
		AdultElderly: map[PhysicalActivity]float64{
			PhysicalActivitySedentary:    1.0,
			PhysicalActivityLowActivity:  1.12,
			PhysicalActivityActive:       1.27,
			PhysicalActivityHighActivity: 1.45,
		},
	}
)

func GetPhysicalActivityMultiplier(sex PatientSex, ageClassification PatientAgeClassification, activity PhysicalActivity) float64 {
	var multipliers PhysicalActivityMultipliers

	switch sex {
	case PatientSexMale:
		multipliers = MalePhysicalActivityMultipliers
	case PatientSexFemale:
		multipliers = FemalePhysicalActivityMultipliers
	default:
		return 1.0
	}

	if ageClassification == PatientChildrenTeenage {
		return multipliers.ChildrenTeenage[activity]
	}
	return multipliers.AdultElderly[activity]
}
