package controller

import (
	"denitiawan/research-swagger-gomod-gin/common/dto"
	"github.com/gin-gonic/gin"
)

func CreateWebResponse(httpStatusCode int, status string, data interface{}, ctx *gin.Context) {

	// create response body
	responseBody := dto.WebResponse{
		Code:   httpStatusCode,
		Status: status,
		Data:   data,
	}

	// header
	ctx.Header("Content-Type", "application/json")

	// write into json
	ctx.JSON(httpStatusCode, responseBody)
}
