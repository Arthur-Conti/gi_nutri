package formulas

import "fmt"

type JET struct {
	Result                     float64
	weight                     float64
	height                     float64
	age                        int
	schofieldAgeClassification string
	sex                        string
}

func NewJET(weight, heightCM float64, age int, schofieldAgeClassification, sex string) *JET {
	return &JET{
		weight:                     weight,
		height:                     heightCM,
		age:                        age,
		schofieldAgeClassification: schofieldAgeClassification,
		sex:                        sex,
	}
}

func (j *JET) Calculate(useHarrisBenedict, useFao, useSchofield, usePocket bool, pocketValue float64) {
	switch j.sex {
	case "male":
		if useHarrisBenedict {
			j.HarrisBenedictMale()
		} else if useFao {
			j.FaoOmsMale()
		} else if useSchofield {
			j.SchofieldMale()
		} else if usePocket {
			if pocketValue == 0.0 {
				fmt.Println("Error: To use pocket you must provide the value")
				return
			}
			j.Pocket(pocketValue)
		}
	case "female":
		if useHarrisBenedict {
			j.HarrisBenedictFemale()
		} else if useFao {
			j.FaoOmsFemale()
		} else if useSchofield {
			j.SchofieldFemale()
		} else if usePocket {
			if pocketValue == 0.0 {
				fmt.Println("Error: To use pocket you must provide the value")
				return
			}
			j.Pocket(pocketValue)
		}
	}
}

func (j *JET) HarrisBenedictMale() {
	j.Result = 66.5 + (13.75 * j.weight) + (5.0 * j.height) - (6.76 * float64(j.age))
}

func (j *JET) HarrisBenedictFemale() {
	j.Result = 655.1 + (9.56 * j.weight) + (1.85 * j.height) - (4.68 * float64(j.age))
}

func (j *JET) FaoOmsMale() {
	j.Result = 15.3*j.weight + 679
}

func (j *JET) FaoOmsFemale() {
	j.Result = 14.7*j.weight + 496
}

func (j *JET) SchofieldMale() {
	fmt.Println(j.schofieldAgeClassification)
	switch j.schofieldAgeClassification {
	case "early_kid":
		j.Result = 59.5*j.weight - 30.4
	case "late_kid":
		j.Result = 22.7*j.weight + 495
	case "teenage":
		j.Result = 17.5*j.weight + 651
	case "early_adult":
		j.Result = 15.3*j.weight + 679
	case "late_adult":
		j.Result = 11.6*j.weight + 879
	case "elderly":
		j.Result = 11.6*j.weight + 879
	}
}

func (j *JET) SchofieldFemale() {
	switch j.schofieldAgeClassification {
	case "early_kid":
		j.Result = 58.3*j.weight - 31.01
	case "late_kid":
		j.Result = 22.5*j.weight + 499
	case "teenage":
		j.Result = 12.2*j.weight + 746
	case "early_adult":
		j.Result = 14.7*j.weight + 496
	case "late_adult":
		j.Result = 8.7*j.weight + 829
	case "elderly":
		j.Result = 8.7*j.weight + 829
	}
}

func (j *JET) Pocket(value float64) {
	j.Result = j.weight * value
}
