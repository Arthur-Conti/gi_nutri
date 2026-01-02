package middleware

import (
	"net/http"

	"github.com/Arthur-Conti/gi_nutri/internal/infra/configs"
	"github.com/Arthur-Conti/gi_nutri/internal/infra/http/controllers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// ErrorHandler middleware para tratar erros de forma centralizada
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Verifica se há erros
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			// Verifica se é um AppError
			if appErr, ok := err.(*configs.AppError); ok {
				controllers.NewErrorResponse(c, appErr.Code, appErr.Message)
				return
			}

			// Trata erros do MongoDB
			if err == mongo.ErrNoDocuments {
				controllers.NewErrorResponse(c, http.StatusNotFound, "resource not found")
				return
			}

			// Erro genérico
			controllers.NewErrorResponse(c, http.StatusInternalServerError, "internal server error")
		}
	}
}

