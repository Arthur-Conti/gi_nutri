package patient

type Patient interface {
	GetPhysicalActivityResult()
	CalculateIMC() error
	CalculateAdjustedWeight() error
	CalculatePercentageWeightAdequacy() error
	CalculatePercentageWeightChange() error
	CalculateEER() error
	CalculateTMB(useHarrisBenedict, useFao, useSchofield, usePocket bool, pocketValue float64) error
	GetFormulas() FormulasInfo
}

func NewPatient(opts PatientOpts) Patient {
	var patient Patient

	opts.AgeClassification = ClassifyAge(opts.Age)
	opts.SchofieldAgeClassification = ClassifyAgeSchofield(opts.Age)
	opts.Measures = Measures{
		HeightCM: opts.Height,
		HeightM:  heightMeters(opts.Height),
		Weight:   opts.Weight,
	}

	switch opts.Sex {
	case PatientSexMale:
		patient = NewPatientMale(opts)
	case PatientSexFemale:
		patient = NewPatientFemale(opts)
	}

	if opts.IsPregnant {
		patient = NewPatientPregnant(opts)
	}

	patient.GetPhysicalActivityResult()
	return patient
}
