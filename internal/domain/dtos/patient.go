package dtos

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PatientDTO struct {
	ID                         primitive.ObjectID `json:"_id,omitempty"`
	Name                       string             `json:"name"`
	Age                        int                `json:"age"`
	AgeClassification          string             `json:"age_classification"`
	SchofieldAgeClassification string             `json:"schofield_age_classification"`
	Sex                        string             `json:"sex"`
	UsualWeight                float64            `json:"usual_weight"`
	PhysicalActivity           string             `json:"physical_activity"`
	Measures                   PatientMeasures    `json:"patient_measures"`
	IsPregnant                 bool               `json:"is_pregnant"`
	PregnancyInfo              PregnancyInfo      `json:"pregnancy_info"`
	IsLactating                bool               `json:"is_lactating"`
	LactatingInfo              LactatingInfo      `json:"lactating_info"`
	CreatedAt                  time.Time          `json:"created_at"`
	UpdatedAt                  time.Time          `json:"updated_at"`
	Deleted                    bool               `json:"deleted"`
	Finished                   bool               `json:"finished"`
}

type PatientMeasures struct {
	HeightCM float64 `json:"height_cm"`
	HeightM  float64 `json:"height_m"`
	Weight   float64 `json:"weight"`
}

type PregnancyInfo struct {
	Weeks    int
	Quarters int
}

type LactatingInfo struct {
	BabyAgeMonths int
}
