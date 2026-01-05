package routes

import (
	"github.com/Arthur-Conti/gi_nutri/internal/infra/container"
	resultcontroller "github.com/Arthur-Conti/gi_nutri/internal/infra/http/controllers/result"
	"github.com/Arthur-Conti/gi_nutri/internal/infra/http/middleware"
	"github.com/gin-gonic/gin"
)

func resultRoutes(router *gin.RouterGroup) {
	controller := resultcontroller.NewResultController(container.BaseContainer.Services.ResultService)

	router.Use(middleware.ValidatePatientID())

	router.GET("/", controller.ListHandler)
	router.GET("/last", controller.GetLastByPatientIDHandler)
	router.GET("/imc", controller.GetIMCHandler)
	router.POST("/imc", controller.SaveIMCHandler)
	router.GET("/adjusted-weight", controller.GetAdjustedWeightHandler)
	router.POST("/adjusted-weight", controller.SaveAdjustedWeightHandler)
	router.GET("/percentage-weight-adequacy", 
		middleware.ValidateQueryParam("time_days"),
		controller.GetPercentageWeightAdequacyHandler)
	router.POST("/percentage-weight-adequacy", 
		middleware.ValidateQueryParam("time_days"),
		controller.SavePercentageWeightAdequacyHandler)
	router.GET("/percentage-weight-change",
		middleware.ValidateQueryParam("time_days"),
		controller.GetPercentageWeightChangeHandler)
	router.POST("/percentage-weight-change",
		middleware.ValidateQueryParam("time_days"),
		controller.SavePercentageWeightChangeHandler)
	router.GET("/eer", controller.GetEERHandler)
	router.POST("/eer", controller.SaveEERHandler)
	router.GET("/tmb", controller.GetTMBHandler)
	router.POST("/tmb", controller.SetTMBHandler)
}
