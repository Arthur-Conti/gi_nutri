package resultcontroller

import (
	"net/http"

	portsservices "github.com/Arthur-Conti/gi_nutri/internal/application/ports/services"
	"github.com/Arthur-Conti/gi_nutri/internal/infra/http/controllers"
	"github.com/gin-gonic/gin"
)

type ResultController struct {
	svc portsservices.ResultsService
}

func NewResultController(svc portsservices.ResultsService) *ResultController {
	return &ResultController{
		svc: svc,
	}
}

func (rc *ResultController) ListHandler(c *gin.Context) {
	patientID := c.Param("patient_id")

	results, err := rc.svc.List(c.Request.Context(), patientID)
	if err != nil {
		controllers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	controllers.NewSuccessResponse(c, http.StatusOK, "Results Listed", results)
}

func (rc *ResultController) GetLastByPatientIDHandler(c *gin.Context) {
	patientID := c.Param("patient_id")

	result, err := rc.svc.GetLastByPatientID(c.Request.Context(), patientID)
	if err != nil {
		controllers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	controllers.NewSuccessResponse(c, http.StatusOK, "Result Found", result)
}

func (rc *ResultController) GetIMCHandler(c *gin.Context) {
	patientID := c.Param("patient_id")

	imc, err := rc.svc.GetIMC(c.Request.Context(), patientID)
	if err != nil {
		controllers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	controllers.NewSuccessResponse(c, http.StatusOK, "IMC Calculated", SchemaIMC{
		Status: imc.Status,
		Result: imc.Result,
	})
}

func (rc *ResultController) SaveIMCHandler(c *gin.Context) {
	patientID := c.Param("patient_id")

	resultID, err := rc.svc.SaveIMC(c.Request.Context(), patientID)
	if err != nil {
		controllers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	controllers.NewSuccessResponse(c, http.StatusCreated, "IMC Saved", CreateOrUpdateResponde{ID: resultID})
}

func (rc *ResultController) GetAdjustedWeightHandler(c *gin.Context) {
	patientID := c.Param("patient_id")

	aw, err := rc.svc.GetAdjustedWeight(c.Request.Context(), patientID)
	if err != nil {
		controllers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	controllers.NewSuccessResponse(c, http.StatusOK, "Adjusted Weight Calculated", SchemaAdjutedWeight{
		IdealWeight: aw.IdealWeight,
		Result:      aw.Result,
	})
}

func (rc *ResultController) GetPercentageWeightAdequacyHandler(c *gin.Context) {
	patientID := c.Param("patient_id")
	timeDays := c.Query("time_days")

	wa, err := rc.svc.GetPercentageWeightAdequacy(c.Request.Context(), patientID, timeDays)
	if err != nil {
		controllers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	controllers.NewSuccessResponse(c, http.StatusOK, "Percentage Weight Adequacy Calculated", SchemaPercentageWeight{
		Classification: wa.Classification,
		Result:         wa.Result,
	})
}
