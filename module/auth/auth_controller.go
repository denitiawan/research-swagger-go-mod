package auth

import (
	baseController "denitiawan/research-swagger-gomod-gin/common/controller"
	"denitiawan/research-swagger-gomod-gin/common/helper"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

type AuthController struct {
	service AuthService
}

func NewAuthController(service AuthService) *AuthController {
	return &AuthController{
		service: service,
	}
}

func (controller *AuthController) Login(ctx *gin.Context) {
	log.Info().Msg("Login")
	requestDto := LoginDto{}
	err := ctx.ShouldBindJSON(&requestDto)
	helper.ErrorPanic(err)

	data := controller.service.Login(requestDto)
	baseController.CreateWebResponse(http.StatusOK, "OK", data, ctx)
}
