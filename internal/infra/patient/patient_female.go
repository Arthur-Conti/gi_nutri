package patient

type PatientFemale struct {
	BasePatient
	IsPregnant    bool
	PregnancyInfo PregnancyInfo
	IsLactating   bool
	LactatingInfo LactatingInfo
}

func NewPatientFemale(opts PatientOpts) *PatientFemale {
	return &PatientFemale{
		BasePatient:   *NewBasePatient(opts),
		IsPregnant:    opts.IsPregnant,
		PregnancyInfo: opts.PregnancyInfo,
		IsLactating:   opts.IsLactating,
		LactatingInfo: opts.LactatingInfo,
	}
}

func (p *PatientFemale) GetPhysicalActivityResult() {
	p.PhysicalActivityResult = GetPhysicalActivityMultiplier(
		p.Sex,
		p.AgeClassification,
		p.PhysicalActivity,
	)
}
