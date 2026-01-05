package dtos

import (
	"time"

	resultsrepository "github.com/Arthur-Conti/gi_nutri/internal/infra/repositories/results"
)

type ResultsDTO struct {
	ID         string    `json:"id"`
	PatientID  string    `json:"patient_id"`
	RecordedAt time.Time `json:"recorded_at"`
	Measures   Measures  `json:"measures"`
	Formulas   Formulas  `json:"formulas"`
}

type Measures struct {
	HeightCM float64 `json:"height_cm"`
	HeightM  float64 `json:"height_m"`
	Weight   float64 `json:"weight"`
}

type Formulas struct {
	IMC                      IMC                      `json:"imc,omitzero"`
	AdjustedWeightObesity    AdjustedWeightObesity    `json:"adjusted_weight_obesity,omitzero"`
	PercentageWeightAdequacy PercentageWeightAdequacy `json:"percentage_weight_adequacy,omitzero"`
	PercentageWeightChange   PercentageWeightChange   `json:"percentage_weight_change,omitzero"`
	EER                      EER                      `json:"eer,omitzero"`
	TMB                      TMB                      `json:"tmb,omitzero"`
}

type IMC struct {
	Status string  `json:"status"`
	Result float64 `json:"result"`
}

type AdjustedWeightObesity struct {
	IdealWeight float64 `json:"ideal_weight"`
	Result      float64 `json:"result"`
}

type PercentageWeightAdequacy struct {
	Classification string  `json:"classification"`
	Result         float64 `json:"result"`
}

type PercentageWeightChange struct {
	Classification string  `json:"classification"`
	Result         float64 `json:"result"`
}

type EER struct {
	Result float64 `json:"result"`
}

type TMB struct {
	Result float64 `json:"result"`
}

type TMBFormulas struct {
	HarrisBenedict bool
	FAO            bool
	Schofield      bool
	Pocket         bool
	PocketValue    float64
}

func ResultsFromModel(model resultsrepository.ResultsModel) ResultsDTO {
	return ResultsDTO{
		ID:         model.ID.Hex(),
		PatientID:  model.PatientID.Hex(),
		RecordedAt: model.RecordedAt,
		Measures:   Measures(model.Measures),
		Formulas: Formulas{
			IMC:                      IMC(model.Formulas.IMC),
			AdjustedWeightObesity:    AdjustedWeightObesity(model.Formulas.AdjustedWeightObesity),
			PercentageWeightAdequacy: PercentageWeightAdequacy(model.Formulas.PercentageWeightAdequacy),
			PercentageWeightChange:   PercentageWeightChange(model.Formulas.PercentageWeightChange),
			EER:                      EER(model.Formulas.EER),
			TMB:                      TMB(model.Formulas.TMB),
		},
	}
}
