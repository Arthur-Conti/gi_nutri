package routes

import (
	"github.com/Arthur-Conti/gi_nutri/internal/infra/container"
	patientcontroller "github.com/Arthur-Conti/gi_nutri/internal/infra/http/controllers/patient"
	"github.com/gin-gonic/gin"
)

func patientRoutes(router *gin.RouterGroup) {
	controller := patientcontroller.NewPatientController(container.BaseContainer.Services.PatientService)

	router.POST("/", controller.CreateHandler)
	router.GET("/", controller.ListHandler)
	// router.GET("/:project_id", controller.GetHandler)
	// router.PUT("/:project_id", controller.UpdateHandler)
	// router.DELETE("/:project_id", controller.DeleteHandler)
	// router.PATCH("/:project_id", controller.ChangeStatusHandler)
}