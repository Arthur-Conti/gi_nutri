package patientcontroller

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
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	CreatedAt   time.Time `json:"created_at"`
	Status      string    `json:"status"`
}
