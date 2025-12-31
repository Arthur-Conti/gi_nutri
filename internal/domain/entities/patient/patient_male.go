package patient

type PatientMale struct {
	BasePatient
}

func NewPatientMale(opts PatientOpts) *PatientMale {
	return &PatientMale{
		BasePatient: *NewBasePatient(opts),
	}
}

func (p *PatientMale) GetPhysicalActivityResult() {
	p.PhysicalActivityResult = GetPhysicalActivityMultiplier(
		p.Sex,
		p.AgeClassification,
		p.PhysicalActivity,
	)
}
