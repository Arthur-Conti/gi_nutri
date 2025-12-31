package formulas

type EER struct {
	Result            float64
	age               int
	ageClassification string
	sex               string
	weight            float64
	height            float64
	physicalActivity  float64
}

func NewEER(
	age int, 
	ageClassification, sex string, 
	weight, heightM, physicalActivity float64, 
	) *EER {
	return &EER{
		age:               age,
		ageClassification: ageClassification,
		sex:               sex,
		weight:            weight,
		height:            heightM,
		physicalActivity:  physicalActivity,
	}
}

func (e *EER) Calculate() {
	switch e.ageClassification {
	case "children/teenage":
		e.CalculateEERChildrenTeenage()
	case "adult":
		e.CalculateEERAdult()
	case "elderly":
		e.CalculateEERElderly()
	}
}

func (e *EER) CalculateEERAdult() {
	switch e.sex {
	case "female":
		wh := (9.36 * e.weight) + (726 * e.height)
		e.Result = 354 - (6.91 * float64(e.age)) + e.physicalActivity*wh
	case "male":
		wh := (15.91 * e.weight) + (539.6 * e.height)
		e.Result = 662 - (9.53 * float64(e.age)) + e.physicalActivity*wh
	}
}

func (e *EER) CalculateEERChildrenTeenage() {
	switch e.sex {
	case "female":
		wh := (10.0 * e.weight) + (934 * e.height)
		e.Result = 135.3 - (30.8 * float64(e.age)) + e.physicalActivity*wh + 20
	case "male":
		wh := (26.7 * e.weight) + (903 * e.height)
		e.Result = 88.5 - (61.9 * float64(e.age)) + e.physicalActivity*wh + 20
	}
}

func (e *EER) CalculateEERElderly() {
	switch e.sex {
	case "female":
		wh := (9.36 * e.weight) + (726 * e.height)
		e.Result = 354 - (6.91 * float64(e.age)) + e.physicalActivity*wh
	case "male":
		wh := (15.91 * e.weight) + (539.6 * e.height)
		e.Result = 662 - (9.53 * float64(e.age)) + e.physicalActivity*wh
	}
}

func (e *EER) BaseFormula() {
	switch e.ageClassification {
	case "children/teenage":

	default:

	}
}
