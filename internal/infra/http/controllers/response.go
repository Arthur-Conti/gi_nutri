package controllers

import "github.com/gin-gonic/gin"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewSuccessResponse(ctx *gin.Context, status int, message string, data interface{}) {
	ctx.JSON(status, Response{
		Status:  status,
		Message: message,
		Data:    data,
	})
}

func NewErrorResponse(ctx *gin.Context, status int, message string) {
	ctx.JSON(status, Response{
		Status:  status,
		Message: message,
		Data:    nil,
	})
}