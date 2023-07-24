package auth

import (
	baseController "denitiawan/research-swagger-gomod-gin/common/controller"
	"denitiawan/research-swagger-gomod-gin/common/error"
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

// @Tags			auth
// @Router			/v1/auth/login [post]
// @Summary			login
// @Description		Get JWT token for access all APIs
// @Param			RequestBody body LoginDto true "LoginDto.go"
// @Produce			application/json
// @Success			200 {object} user.UserDto{} "Response Success (UserDto.go)"
func (controller *AuthController) Login(ctx *gin.Context) {
	log.Info().Msg("Login")
	requestDto := LoginDto{}
	err := ctx.ShouldBindJSON(&requestDto)
	error.ErrorPanic(err)

	data := controller.service.Login(requestDto)
	baseController.OK(http.StatusOK, "", data, "", ctx)
}
