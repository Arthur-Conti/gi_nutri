package resultcontroller

import (
	"net/http"

	portsservices "github.com/Arthur-Conti/gi_nutri/internal/application/ports/services"
	"github.com/Arthur-Conti/gi_nutri/internal/domain/dtos"
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

func (rc *ResultController) SaveAdjustedWeightHandler(c *gin.Context) {
	patientID := c.Param("patient_id")

	resultID, err := rc.svc.SaveAdjustedWeight(c.Request.Context(), patientID)
	if err != nil {
		controllers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	controllers.NewSuccessResponse(c, http.StatusCreated, "Adjusted Weight Saved", CreateOrUpdateResponde{ID: resultID})
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

func (rc *ResultController) SavePercentageWeightAdequacyHandler(c *gin.Context) {
	patientID := c.Param("patient_id")
	timeDays := c.Query("time_days")

	id, err := rc.svc.SavePercentageWeightAdequacy(c.Request.Context(), patientID, timeDays)
	if err != nil {
		controllers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	controllers.NewSuccessResponse(c, http.StatusOK, "Percentage Weight Adequacy Saved", CreateOrUpdateResponde{ID: id})
}

func (rc *ResultController) GetPercentageWeightChangeHandler(c *gin.Context) {
	patientID := c.Param("patient_id")
	timeDays := c.Query("time_days")

	wc, err := rc.svc.GetPercentageWeightChange(c.Request.Context(), patientID, timeDays)
	if err != nil {
		controllers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	controllers.NewSuccessResponse(c, http.StatusOK, "Percentage Weight Change Calculated", SchemaPercentageWeight{
		Classification: wc.Classification,
		Result:         wc.Result,
	})
}

func (rc *ResultController) SavePercentageWeightChangeHandler(c *gin.Context) {
	patientID := c.Param("patient_id")
	timeDays := c.Query("time_days")

	id, err := rc.svc.SavePercentageWeightChange(c.Request.Context(), patientID, timeDays)
	if err != nil {
		controllers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	controllers.NewSuccessResponse(c, http.StatusOK, "Percentage Weight Change Saved", CreateOrUpdateResponde{ID: id})
}

func (rc *ResultController) GetEERHandler(c *gin.Context) {
	patientID := c.Param("patient_id")

	eer, err := rc.svc.GetEER(c.Request.Context(), patientID)
	if err != nil {
		controllers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	controllers.NewSuccessResponse(c, http.StatusOK, "EER Calculated", SchemaEER{Result: eer.Result})
}

func (rc *ResultController) SaveEERHandler(c *gin.Context) {
	patientID := c.Param("patient_id")

	id, err := rc.svc.SaveEER(c.Request.Context(), patientID)
	if err != nil {
		controllers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	controllers.NewSuccessResponse(c, http.StatusOK, "EER Saved", CreateOrUpdateResponde{ID: id})
}

func (rc *ResultController) GetTMBHandler(c *gin.Context) {
	patientID := c.Param("patient_id")

	var choice SchemaTMBChoice

	if err := c.ShouldBindQuery(&choice); err != nil {
		controllers.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	eer, err := rc.svc.GetTMB(c.Request.Context(), patientID, dtos.TMBFormulas(choice))
	if err != nil {
		controllers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	controllers.NewSuccessResponse(c, http.StatusOK, "TMB Calculated", SchemaEER{Result: eer.Result})
}

func (rc *ResultController) SetTMBHandler(c *gin.Context) {
	patientID := c.Param("patient_id")

	var choice SchemaTMBChoice

	if err := c.ShouldBindQuery(&choice); err != nil {
		controllers.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := rc.svc.SetTMB(c.Request.Context(), patientID, dtos.TMBFormulas(choice))
	if err != nil {
		controllers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	controllers.NewSuccessResponse(c, http.StatusOK, "TMB Calculated", CreateOrUpdateResponde{ID: id})
}