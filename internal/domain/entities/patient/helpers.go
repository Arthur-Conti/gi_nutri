package patient

func ClassifyAge(age int) PatientAgeClassification {
	if age <= 19 {
		return PatientChildrenTeenage
	} else if age >= 20 && age <= 59 {
		return PatientAdult
	} else {
		return PatientElderly
	}
}

func ClassifyAgeSchofield(age int) SchofieldAgeClassification {
	if age < 3 {
		return SchofieldAgeClassificationEarlyKid
	} else if age >= 3 && age < 10 {
		return SchofieldAgeClassificationLateKid
	} else if age >= 10 && age < 18 {
		return SchofieldAgeClassificationTeenage
	} else if age >= 18 && age < 30 {
		return SchofieldAgeClassificationEarlyAdult
	} else if age >= 30 && age < 60 {
		return SchofieldAgeClassificationLateAdult
	} else {
		return SchofieldAgeClassificationElderly
	}
}

func heightMeters(height float64) float64 {
	return height / 100
}
