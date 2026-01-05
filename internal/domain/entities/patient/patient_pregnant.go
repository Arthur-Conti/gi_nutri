package patient

type PatientPregnant struct {
	BasePatient
	IsPregnant    bool
	PregnancyInfo PregnancyInfo
}

func NewPatientPregnant(opts PatientOpts) *PatientPregnant {
	return &PatientPregnant{
		BasePatient:   *NewBasePatient(opts),
		IsPregnant:    opts.IsPregnant,
		PregnancyInfo: opts.PregnancyInfo,
	}
}

func (p *PatientPregnant) GetPhysicalActivityResult() {
	p.PhysicalActivityResult = GetPhysicalActivityMultiplier(
		PatientSexFemale, 
		p.AgeClassification,
		p.PhysicalActivity,
	)
}

func (p *PatientPregnant) GetData() PatientData {
	data := p.BasePatient.GetData()
	data.IsPregnant = p.IsPregnant
	data.PregnancyInfo = p.PregnancyInfo
	return data
}
