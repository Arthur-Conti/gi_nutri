package resultcontroller

import "time"

type SchemaCreate struct {
	Name             string        `mandatory:"true" json:"name"`
	Age              int           `mandatory:"true" json:"age"`
	Sex              string        `mandatory:"true" json:"sex"`
	UsualWeight      float64       `mandatory:"false" json:"usual_weight"`
	PhysicalActivity string        `mandatory:"false" json:"physical_activity"`
	Measures         Measures      `mandatory:"true" json:"measures"`
	IsPregnant       bool          `mandatory:"true" json:"is_pregnant"`
	PregnancyInfo    PregnancyInfo `mandatory:"false" json:"pregnancy_info"`
	IsLactating      bool          `mandatory:"true" json:"is_lactating"`
	LactatingInfo    LactatingInfo `mandatory:"false" json:"lactating_info"`
}

type Measures struct {
	Height float64 `mandatory:"true" json:"height"`
	Weight float64 `mandatory:"true" json:"weight"`
}

type PregnancyInfo struct {
	Weeks    int
	Quarters int
}

type LactatingInfo struct {
	BabyAgeMonths int
}

type SchemaUpdate struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	CreatedAt   time.Time `json:"created_at"`
	Status      string    `json:"status"`
}

type SchemaChangeStatus struct {
	Status string `mandatory:"true" json:"status"`
}

type CreateOrUpdateResponde struct {
	ID string `json:"id"`
}

type SchemaList struct {
	ID                         string        `json:"id"`
	Name                       string        `json:"name"`
	Age                        int           `json:"age"`
	AgeClassification          string        `json:"age_classification"`
	SchofieldAgeClassification string        `json:"schofield_age_classification"`
	Sex                        string        `json:"sex"`
	UsualWeight                float64       `json:"usual_weight"`
	PhysicalActivity           string        `json:"physical_activity"`
	IsPregnant                 bool          `json:"is_pregnant"`
	PregnancyInfo              PregnancyInfo `json:"pregnancy_info"`
	IsLactating                bool          `json:"is_lactating"`
	LactatingInfo              LactatingInfo `json:"lactating_info"`
	CreatedAt                  time.Time     `json:"created_at"`
	UpdatedAt                  time.Time     `json:"updated_at"`
	Finished                   bool          `json:"finished"`
}

type SchemaIMC struct {
	Status string  `json:"status"`
	Result float64 `json:"result"`
}

type SchemaAdjutedWeight struct {
	Result      float64 `json:"result"`
	IdealWeight float64 `json:"ideal_weight"`
}

type SchemaPercentageWeight struct {
	Classification string  `json:"classification"`
	Result         float64 `json:"result"`
}

type SchemaEER struct {
	Result float64 `json:"result"`
}

type SchemaTMBChoice struct {
	HarrisBenedict bool    `form:"harris_benedict"`
	FAO            bool    `form:"fao"`
	Schofield      bool    `form:"schofield"`
	Pocket         bool    `form:"pocket"`
	PocketValue    float64 `form:"pocket_value"`
}
