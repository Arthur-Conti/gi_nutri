package dtos

import (
	"time"

	patientrepository "github.com/Arthur-Conti/gi_nutri/internal/infra/repositories/patient"
)

type PatientDTO struct {
	ID                         string          `json:"id"`
	Name                       string          `json:"name"`
	Age                        int             `json:"age"`
	AgeClassification          string          `json:"age_classification"`
	SchofieldAgeClassification string          `json:"schofield_age_classification"`
	TimeDays                   int             `json:"time_days,omitzero"`
	Sex                        string          `json:"sex"`
	UsualWeight                float64         `json:"usual_weight"`
	PhysicalActivity           string          `json:"physical_activity"`
	Measures                   PatientMeasures `json:"patient_measures,omitzero"`
	IsPregnant                 bool            `json:"is_pregnant"`
	PregnancyInfo              PregnancyInfo   `json:"pregnancy_info,omitzero"`
	IsLactating                bool            `json:"is_lactating"`
	LactatingInfo              LactatingInfo   `json:"lactating_info,omitzero"`
	CreatedAt                  time.Time       `json:"created_at"`
	UpdatedAt                  time.Time       `json:"updated_at"`
	Deleted                    bool            `json:"deleted,omitempty"`
	Finished                   bool            `json:"finished"`
}

type PatientMeasures struct {
	HeightCM float64 `json:"height_cm"`
	HeightM  float64 `json:"height_m"`
	Weight   float64 `json:"weight"`
}

type PregnancyInfo struct {
	Weeks    int `json:"weeks"`
	Quarters int `json:"quarters"`
}

type LactatingInfo struct {
	BabyAgeMonths int `json:"baby_age_months"`
}

func PatientFromModel(model patientrepository.PatientModel) PatientDTO {
	return PatientDTO{
		ID:                         model.ID.Hex(),
		Name:                       model.Name,
		Age:                        model.Age,
		AgeClassification:          model.AgeClassification,
		SchofieldAgeClassification: model.SchofieldAgeClassification,
		Sex:                        model.Sex,
		UsualWeight:                model.UsualWeight,
		PhysicalActivity:           model.PhysicalActivity,
		IsPregnant:                 model.IsPregnant,
		PregnancyInfo:              PregnancyInfo(model.PregnancyInfo),
		IsLactating:                model.IsLactating,
		LactatingInfo:              LactatingInfo(model.LactatingInfo),
		CreatedAt:                  model.CreatedAt,
		UpdatedAt:                  model.UpdatedAt,
		Finished:                   model.Finished,
	}
}
