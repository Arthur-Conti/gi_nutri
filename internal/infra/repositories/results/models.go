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
	IMC                      IMC                      `bson:"imc,omitempty"`
	AdjustedWeightObesity    AdjustedWeightObesity    `bson:"adjusted_weight_obesity,omitempty"`
	PercentageWeightAdequacy PercentageWeightAdequacy `bson:"percentage_weight_adequacy,omitempty"`
	PercentageWeightChange   PercentageWeightChange   `bson:"percentage_weight_change,omitempty"`
	EER                      EER                      `bson:"eer,omitempty"`
	TMB                      TMB                      `bson:"tmb,omitempty"`
}

type IMC struct {
	Status string  `bson:"status"`
	Result float64 `bson:"result"`
}

type AdjustedWeightObesity struct {
	IdealWeight float64 `bson:"ideal_weight"`
	Result      float64 `bson:"result"`
}

type PercentageWeightAdequacy struct {
	Classification string  `bson:"classfication"`
	Result         float64 `bson:"result"`
}

type PercentageWeightChange struct {
	Classification string  `bson:"classification"`
	Result         float64 `bson:"result"`
}

type EER struct {
	Result float64 `bson:"result"`
}

type TMB struct {
	Result float64 `bson:"result"`
}
