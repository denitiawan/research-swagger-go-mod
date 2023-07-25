package user

import (
	baseController "denitiawan/research-swagger-gomod-gin/common/controller"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

type UserController struct {
	service UserService
}

func NewUserController(service UserService) *UserController {
	return &UserController{
		service: service,
	}
}

// @Tags			user
// @Router			/v1/user/save [post]
// @Summary			save
// @Description		save
// @Param			RequestBody body user.UserDto true "UserDto.go"
// @Param			Authorization header string true "Authorization"
// @Produce			application/json
// @Success			200 {object} user.UserDto{} "Response Success (UserDto.go)"
func (controller *UserController) Create(ctx *gin.Context) {
	log.Info().Msg("Create")

	// req body
	dto := UserDto{}
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		baseController.BadRequest(http.StatusBadRequest, "invalid req body", dto, err.Error(), ctx)
		return
	}

	// service
	response := controller.service.Create(dto)
	if response.Error != nil {
		baseController.Error(http.StatusInternalServerError, response.Message, dto, response.ErrorMessage, ctx)
		return
	}

	// ok
	baseController.OK(http.StatusOK, response.Message, response.Data, "", ctx)
}

// @Tags			user
// @Router			/v1/user/update/{id} [put]
// @Summary			update
// @Description		update
// @Param			Authorization header string true "Authorization"
// @Param			id  path int true "id"
// @Param			RequestBody body UserDto true "UserDto.go"
// @Produce			application/json
// @Success			200 {object} user.UserDto{} "Response Success (UserDto.go)"
func (controller *UserController) Update(ctx *gin.Context) {
	log.Info().Msg("Update")

	// path
	param := ctx.Param("id")
	id, _ := strconv.ParseInt(param, 0, 64)

	// req body
	dto := UserDto{}
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		baseController.BadRequest(http.StatusBadRequest, "invalid req body", dto, err.Error(), ctx)
		return
	}

	// service
	response := controller.service.Update(id, dto)
	if response.Error != nil {
		baseController.Error(http.StatusInternalServerError, response.Message, param, response.ErrorMessage, ctx)
		return
	}

	// ok
	baseController.OK(http.StatusOK, response.Message, response.Data, "", ctx)
}

// @Tags			user
// @Router			/v1/user/delete/{id} [delete]
// @Summary			delete
// @Description		delete
// @Param			Authorization header string true "Authorization"
// @Param			id  path int true "id"
// @Produce			application/json
// @Success			200 {object} user.UserDto{} "Response Success (UserDto.go)"
func (controller *UserController) Delete(ctx *gin.Context) {
	log.Info().Msg("Delete")

	param := ctx.Param("id")
	id, _ := strconv.ParseInt(param, 0, 64)

	// service
	response := controller.service.Delete(id)
	if response.Error != nil {
		baseController.Error(http.StatusInternalServerError, response.Message, param, response.ErrorMessage, ctx)
		return
	}

	// ok
	baseController.OK(http.StatusOK, response.Message, response.Data, "", ctx)
}

// @Tags			user
// @Router			/v1/user/view/{id} [get]
// @Summary			view
// @Description		view
// @Param			Authorization header string true "Authorization"
// @Param			id  path int true "id"
// @Produce			application/json
// @Success			200 {object} user.UserDto{} "Response Success (UserDto.go)"
func (controller *UserController) FindById(ctx *gin.Context) {
	log.Info().Msg("FindById")

	//-----------log.Info().Msg("FindById")
	param := ctx.Param("id")
	id, _ := strconv.ParseInt(param, 0, 64)

	// service
	response := controller.service.FindById(id)
	if response.Error != nil {
		baseController.Error(http.StatusInternalServerError, response.Message, param, response.ErrorMessage, ctx)
		return
	}

	// ok
	baseController.OK(http.StatusOK, response.Message, response.Data, "", ctx)

}

// @Tags			user
// @Router			/v1/user/list [get]
// @Summary			list
// @Description		list
// @Param			Authorization header string true "Authorization"
// @Produce			application/json
// @Success			200 {object} user.UserDto{} "Response Success (UserDto.go)"
func (controller *UserController) FindAll(ctx *gin.Context) {
	log.Info().Msg("FindAll")

	// service
	response := controller.service.FindAll()
	if response.Error != nil {
		baseController.Error(http.StatusInternalServerError, response.Message, []User{}, response.ErrorMessage, ctx)
		return
	}

	// ok
	baseController.OK(http.StatusOK, response.Message, response.Data, "", ctx)
}
