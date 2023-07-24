package controller

import (
	"denitiawan/research-swagger-gomod-gin/common/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func createWebResponse(httpStatusCode int, message string, data interface{}, errorMessage string, ctx *gin.Context) {

	// create response body
	responseBody := dto.WebResponse{
		Code:         httpStatusCode,
		Message:      message,
		Data:         data,
		ErrorMessage: errorMessage,
	}

	// header
	ctx.Header("Content-Type", "application/json")

	// write into json
	ctx.JSON(httpStatusCode, responseBody)
}

func OK(httpStatusCode int, message string, data interface{}, errorMessage string, ctx *gin.Context) {
	createWebResponse(http.StatusOK, message, data, errorMessage, ctx)
}

func BadRequest(httpStatusCode int, message string, data interface{}, errorMessage string, ctx *gin.Context) {
	createWebResponse(http.StatusBadRequest, message, data, errorMessage, ctx)
}

func Error(httpStatusCode int, message string, data interface{}, errorMessage string, ctx *gin.Context) {
	createWebResponse(http.StatusInternalServerError, message, data, errorMessage, ctx)
}
