package patient

import "github.com/Arthur-Conti/gi_nutri/internal/infra/formulas"

type PatientOpts struct {
	Name                       string
	Age                        int
	AgeClassification          PatientAgeClassification
	SchofieldAgeClassification SchofieldAgeClassification
	TimeDays                   int
	Sex                        PatientSex
	Height                     float64
	Weight                     float64
	UsualWeight                float64
	PhysicalActivity           PhysicalActivity
	Measures                   Measures
	IsPregnant                 bool
	PregnancyInfo              PregnancyInfo
	IsLactating                bool
	LactatingInfo              LactatingInfo
}

type Measures struct {
	HeightCM float64
	HeightM  float64
	Weight   float64
}

type PregnancyInfo struct {
	Weeks    int
	Quarters int
}

type LactatingInfo struct {
	BabyAgeMonths int
}

type FormulasInfo struct {
	IMC                      *formulas.IMC
	AdjustedWeight           *formulas.AdjustedWeightObesity
	PercentageWeightAdequacy *formulas.PercentageWeightAdequacy
	PercentageWeightChange   *formulas.PercentageWeightChange
	EER                      *formulas.EER
	TMB                      *formulas.TMB
}

type PatientAgeClassification string

const (
	PatientChildrenTeenage PatientAgeClassification = "children/teenage"
	PatientAdult           PatientAgeClassification = "adult"
	PatientElderly         PatientAgeClassification = "elderly"
)

type SchofieldAgeClassification string

const (
	SchofieldAgeClassificationEarlyKid   SchofieldAgeClassification = "early_kid"
	SchofieldAgeClassificationLateKid    SchofieldAgeClassification = "late_kid"
	SchofieldAgeClassificationTeenage    SchofieldAgeClassification = "teenage"
	SchofieldAgeClassificationEarlyAdult SchofieldAgeClassification = "early_adult"
	SchofieldAgeClassificationLateAdult  SchofieldAgeClassification = "late_adult"
	SchofieldAgeClassificationElderly    SchofieldAgeClassification = "elderly"
)

type PatientSex string

const (
	PatientSexMale   PatientSex = "male"
	PatientSexFemale PatientSex = "female"
)

type PhysicalActivity string

const (
	PhysicalActivitySedentary    PhysicalActivity = "sedentary"
	PhysicalActivityLowActivity  PhysicalActivity = "low_activity"
	PhysicalActivityActive       PhysicalActivity = "active"
	PhysicalActivityHighActivity PhysicalActivity = "high_activity"
)
