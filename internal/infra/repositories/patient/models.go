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
	UsualWeight                float64            `bson:"usual_weight"`
	PhysicalActivity           string             `bson:"physical_activity"`
	CreatedAt                  time.Time          `bson:"created_at"`
	UpdatedAt                  time.Time          `bson:"updated_at"`
	Deleted                    bool               `bson:"deleted"`
	Finished                   bool               `bson:"finished"`
}

type PatientMeasures struct {
	HeightCM float64 `bson:"height_cm"`
	HeightM  float64 `bson:"height_m"`
	Weight   float64 `bson:"weight"`
}
