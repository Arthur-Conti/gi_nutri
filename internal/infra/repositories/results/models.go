package resultsrepository

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ResultsModel struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	PatientID  primitive.ObjectID `bson:"patient_id"`
	RecordedAt time.Time          `bson:"recorded_at"`
	Measures   Measures           `bson:"measures"`
	Formulas   Formulas           `bson:"formulas"`
}

type Measures struct {
	HeightCM float64 `bson:"height_cm"`
	HeightM  float64 `bson:"height_m"`
	Weight   float64 `bson:"weight"`
}

type Formulas struct {
	IMC                      IMC
	AdjustedWeightObesity    AdjustedWeightObesity
	PercentageWeightAdequacy PercentageWeightAdequacy
	PercentageWeightChange   PercentageWeightChange
	EER                      EER
	TMB                      TMB
}

type IMC struct {
	Status string
	Result float64
}

type AdjustedWeightObesity struct {
	IdealWeight float64
	Result      float64
}

type PercentageWeightAdequacy struct {
	Classification string
	Result         float64
}

type PercentageWeightChange struct {
	Classification string
	Result         float64
}

type EER struct {
	Result float64
}

type TMB struct {
	Result float64
}
