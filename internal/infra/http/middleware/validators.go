package middleware

import (
	"net/http"

	"github.com/Arthur-Conti/gi_nutri/internal/infra/http/controllers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ValidatePatientID valida se o patient_id é um ObjectID válido
func ValidatePatientID() gin.HandlerFunc {
	return func(c *gin.Context) {
		patientID := c.Param("patient_id")
		if patientID == "" {
			controllers.NewErrorResponse(c, http.StatusBadRequest, "patient_id is required")
			c.Abort()
			return
		}

		if _, err := primitive.ObjectIDFromHex(patientID); err != nil {
			controllers.NewErrorResponse(c, http.StatusBadRequest, "invalid patient_id format")
			c.Abort()
			return
		}

		c.Next()
	}
}

// ValidateQueryParam valida se um query parameter existe
func ValidateQueryParam(paramName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		value := c.Query(paramName)
		if value == "" {
			controllers.NewErrorResponse(c, http.StatusBadRequest, paramName+" query parameter is required")
			c.Abort()
			return
		}
		c.Next()
	}
}

