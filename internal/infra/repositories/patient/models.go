package patientrepository

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PatientModel struct {
	ID                         primitive.ObjectID `bson:"_id,omitempty"`
	Name                       string             `bson:"name"`
	Age                        int                `bson:"age"`
	AgeClassification          string             `bson:"age_classification"`
	SchofieldAgeClassification string             `bson:"schofield_age_classification"`
	Sex                        string             `bson:"sex"`
	TimeDays                   int                `bson:"time_days"`
	Height                     float64            `bson:"height"`
	Weight                     float64            `bson:"weight"`
	UsualWeight                float64            `bson:"usual_weight"`
	PhysicalActivity           string             `bson:"physical_activity"`
	IsPregnant                 bool               `bson:"is_pregnant"`
	PregnancyInfo              PregnancyInfo      `bson:"pregnancy_info"`
	IsLactating                bool               `bson:"is_lactating"`
	LactatingInfo              LactatingInfo      `bson:"lactating_info"`
	CreatedAt                  time.Time          `bson:"created_at"`
	UpdatedAt                  time.Time          `bson:"updated_at"`
	Deleted                    bool               `bson:"deleted"`
	Finished                   bool               `bson:"finished"`
}

type PregnancyInfo struct {
	Weeks    int `bson:"weeks"`
	Quarters int `bson:"quarters"`
}

type LactatingInfo struct {
	BabyAgeMonths int `bson:"baby_age_months"`
}
