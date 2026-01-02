package patientcontroller

import (
	"net/http"

	portsservices "github.com/Arthur-Conti/gi_nutri/internal/application/ports/services"
	"github.com/Arthur-Conti/gi_nutri/internal/domain/dtos"
	"github.com/Arthur-Conti/gi_nutri/internal/infra/http/controllers"
	"github.com/gin-gonic/gin"
)

type PatientController struct {
	svc portsservices.PatientService
}

func NewPatientController(svc portsservices.PatientService) *PatientController {
	return &PatientController{
		svc: svc,
	}
}

func (pc *PatientController) CreateHandler(c *gin.Context) {
	var input SchemaCreate

	if err := c.ShouldBindJSON(&input); err != nil {
		controllers.NewErrorResponse(c, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	id, err := pc.svc.Create(c.Request.Context(), dtos.PatientDTO{
		Name:             input.Name,
		Age:              input.Age,
		Sex:              input.Sex,
		UsualWeight:      input.UsualWeight,
		PhysicalActivity: input.PhysicalActivity,
		Measures: dtos.PatientMeasures{
			HeightCM: input.Measures.Height,
			Weight:   input.Measures.Weight,
		},
		IsPregnant:    input.IsPregnant,
		PregnancyInfo: dtos.PregnancyInfo(input.PregnancyInfo),
		IsLactating:   input.IsLactating,
		LactatingInfo: dtos.LactatingInfo(input.LactatingInfo),
	})
	if err != nil {
		controllers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	controllers.NewSuccessResponse(c, http.StatusCreated, "Patient Created", CreateOrUpdateResponde{ID: id})
}

func (pc *PatientController) ListHandler(c *gin.Context) {
	patients, err := pc.svc.List(c.Request.Context())
	if err != nil {
		controllers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	controllers.NewSuccessResponse(c, http.StatusOK, "Patients Listed", patients)
}