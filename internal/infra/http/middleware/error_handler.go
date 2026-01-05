package middleware

import (
	"net/http"

	"github.com/Arthur-Conti/gi_nutri/internal/infra/configs"
	"github.com/Arthur-Conti/gi_nutri/internal/infra/http/controllers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			if appErr, ok := err.(*configs.AppError); ok {
				controllers.NewErrorResponse(c, appErr.Code, appErr.Message)
				return
			}

			if err == mongo.ErrNoDocuments {
				controllers.NewErrorResponse(c, http.StatusNotFound, "resource not found")
				return
			}

			controllers.NewErrorResponse(c, http.StatusInternalServerError, "internal server error")
		}
	}
}

