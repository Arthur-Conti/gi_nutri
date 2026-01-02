package routes

import (
	"github.com/Arthur-Conti/gi_nutri/internal/infra/container"
	resultcontroller "github.com/Arthur-Conti/gi_nutri/internal/infra/http/controllers/result"
	"github.com/Arthur-Conti/gi_nutri/internal/infra/http/middleware"
	"github.com/gin-gonic/gin"
)

func resultRoutes(router *gin.RouterGroup) {
	controller := resultcontroller.NewResultController(container.BaseContainer.Services.ResultService)

	// Aplica validação de patient_id em todas as rotas
	router.Use(middleware.ValidatePatientID())

	router.GET("/", controller.ListHandler)
	router.GET("/last", controller.GetLastByPatientIDHandler)
	router.GET("/imc", controller.GetIMCHandler)
	router.POST("/imc", controller.SaveIMCHandler)
	router.GET("/adjusted-weight", controller.GetAdjustedWeightHandler)
	router.GET("/percentage-weight-adequacy", 
		middleware.ValidateQueryParam("time_days"),
		controller.GetPercentageWeightAdequacyHandler)
}
