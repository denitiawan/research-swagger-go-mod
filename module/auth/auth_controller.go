package auth

import (
	baseController "denitiawan/research-swagger-gomod-gin/common/controller"
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
	log.Info().Msg("Create")

	// req body
	dto := LoginDto{}
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		baseController.BadRequest(http.StatusBadRequest, "invalid req body", dto, err.Error(), ctx)
		return
	}

	// service
	response := controller.service.Login(dto)
	if response.Error != nil {
		baseController.Error(http.StatusInternalServerError, response.Message, dto, response.ErrorMessage, ctx)
		return
	}

	// ok
	baseController.OK(http.StatusOK, response.Message, response.Data, "", ctx)
}
