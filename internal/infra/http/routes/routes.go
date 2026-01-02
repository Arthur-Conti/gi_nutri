package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/Arthur-Conti/gi_nutri/internal/infra/http/middleware"
)

func RegisterRoutes(server *gin.Engine) {
	serviceName := "nutri"

	// Middleware global de tratamento de erros
	server.Use(middleware.ErrorHandler())

	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	patientRouter := server.Group(serviceName + "/patient")
	patientRoutes(patientRouter)

	resultRouter := server.Group(serviceName + "/patient/:patient_id/results")
	resultRoutes(resultRouter)
}