package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoutes(server *gin.Engine) {
	serviceName := "assistant"

	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	patientRouter := server.Group(serviceName + "/patient")
	patientRoutes(patientRouter)

	// resultsRouter := server.Group(serviceName + "/projects/:patient_id/results")
	// resultsRouter(resultsRouter)
}