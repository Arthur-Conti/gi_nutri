package formulas

import (
	"fmt"
)

type TMB struct {
	Result                     float64
	weight                     float64
	height                     float64
	age                        int
	schofieldAgeClassification string
	sex                        string
}

type TMBFormulas struct {
	HarrisBenedict bool
	FAO            bool
	Schofield      bool
	Pocket         bool
	PocketValue    float64
}

func NewTMB(weight, heightCM float64, age int, schofieldAgeClassification, sex string) *TMB {
	return &TMB{
		weight:                     weight,
		height:                     heightCM,
		age:                        age,
		schofieldAgeClassification: schofieldAgeClassification,
		sex:                        sex,
	}
}

func (t *TMB) Calculate(choice TMBFormulas) {
	switch t.sex {
	case "male":
		if choice.HarrisBenedict {
			t.HarrisBenedictMale()
		} else if choice.FAO {
			t.FaoOmsMale()
		} else if choice.Schofield {
			t.SchofieldMale()
		} else if choice.Pocket {
			if choice.PocketValue == 0.0 {
				fmt.Println("Error: To use pocket you must provide the value")
			}
			t.Pocket(choice.PocketValue)
		}
	case "female":
		if choice.HarrisBenedict {
			t.HarrisBenedictFemale()
		} else if choice.FAO {
			t.FaoOmsFemale()
		} else if choice.Schofield {
			t.SchofieldFemale()
		} else if choice.Pocket {
			if choice.PocketValue == 0.0 {
				fmt.Println("Error: To use pocket you must provide the value")
			}
			t.Pocket(choice.PocketValue)
		}
	}
}

func (t *TMB) HarrisBenedictMale() {
	t.Result = 66.5 + (13.75 * t.weight) + (5.0 * t.height) - (6.76 * float64(t.age))
}

func (t *TMB) HarrisBenedictFemale() {
	t.Result = 655.1 + (9.56 * t.weight) + (1.85 * t.height) - (4.68 * float64(t.age))
}

func (t *TMB) FaoOmsMale() {
	t.Result = 15.3*t.weight + 679
}

func (t *TMB) FaoOmsFemale() {
	t.Result = 14.7*t.weight + 496
}

func (t *TMB) SchofieldMale() {
	switch t.schofieldAgeClassification {
	case "early_kid":
		t.Result = 59.5*t.weight - 30.4
	case "late_kid":
		t.Result = 22.7*t.weight + 495
	case "teenage":
		t.Result = 17.5*t.weight + 651
	case "early_adult":
		t.Result = 15.3*t.weight + 679
	case "late_adult":
		t.Result = 11.6*t.weight + 879
	case "elderly":
		t.Result = 11.6*t.weight + 879
	}
}

func (t *TMB) SchofieldFemale() {
	switch t.schofieldAgeClassification {
	case "early_kid":
		t.Result = 58.3*t.weight - 31.01
	case "late_kid":
		t.Result = 22.5*t.weight + 499
	case "teenage":
		t.Result = 12.2*t.weight + 746
	case "early_adult":
		t.Result = 14.7*t.weight + 496
	case "late_adult":
		t.Result = 8.7*t.weight + 829
	case "elderly":
		t.Result = 8.7*t.weight + 829
	}
}

func (t *TMB) Pocket(value float64) {
	t.Result = t.weight * value
}
