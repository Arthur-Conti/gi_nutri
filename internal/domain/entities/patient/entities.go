package patient

import "github.com/Arthur-Conti/gi_nutri/internal/domain/entities/formulas"

type PatientOpts struct {
	Name             string
	Age              int
	TimeDays         int
	Sex              PatientSex
	Height           float64
	Weight           float64
	UsualWeight      float64
	PhysicalActivity PhysicalActivity
	IsPregnant       bool
	PregnancyInfo    PregnancyInfo
	IsLactating      bool
	LactatingInfo    LactatingInfo
}

type Results struct {
	ResultsID    string
	Measures     Measures
	FormulasInfo FormulasInfo
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

// PatientData contém os dados do paciente para uso em mappers
// Esta estrutura permite que a camada de infraestrutura acesse os dados
// sem conhecer a estrutura interna da entidade
type PatientData struct {
	ID                         string
	Name                       string
	Age                        int
	AgeClassification          PatientAgeClassification
	SchofieldAgeClassification SchofieldAgeClassification
	Sex                        PatientSex
	Height                     float64
	Weight                     float64
	UsualWeight                float64
	TimeDays                   int
	PhysicalActivity           PhysicalActivity
	PhysicalActivityResult     float64
	IsPregnant                 bool
	PregnancyInfo              PregnancyInfo
	IsLactating                bool
	LactatingInfo              LactatingInfo
	Results                    Results
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
