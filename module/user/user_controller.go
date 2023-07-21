package user

import (
	baseController "denitiawan/research-swagger-gomod-gin/common/controller"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

type UserController struct {
	prodcutService UserService
}

func NewUserController(service UserService) *UserController {
	return &UserController{
		prodcutService: service,
	}
}

// @Tags			user
// @Router			/v1/user/save [post]
// @Summary			save
// @Description		save
// @Param			RequestBody body user.UserDto true "UserDto.go"
// @Produce			application/json
// @Success			200 {object} user.UserDto{} "Response Success (UserDto.go)"
func (controller *UserController) Create(ctx *gin.Context) {
	log.Info().Msg("Create")
	baseController.CreateWebResponse(http.StatusOK, "OK", "Create", ctx)
}

// @Tags			user
// @Router			/v1/user/update/{id} [put]
// @Summary			update
// @Description		update
// @Param			id  path int true "id"
// @Param			RequestBody body UserDto true "UserDto.go"
// @Produce			application/json
// @Success			200 {object} user.UserDto{} "Response Success (UserDto.go)"
func (controller *UserController) Update(ctx *gin.Context) {
	log.Info().Msg("Update")
	baseController.CreateWebResponse(http.StatusOK, "OK", "Update", ctx)
}

// @Tags			user
// @Router			/v1/user/delete/{id} [delete]
// @Summary			delete
// @Description		delete
// @Param			id  path int true "id"
// @Produce			application/json
// @Success			200 {object} user.UserDto{} "Response Success (UserDto.go)"
func (controller *UserController) Delete(ctx *gin.Context) {
	log.Info().Msg("Delete")
	baseController.CreateWebResponse(http.StatusOK, "OK", "Delete", ctx)
}

// @Tags			user
// @Router			/v1/user/view/{id} [get]
// @Summary			view
// @Description		view
// @Param			id  path int true "id"
// @Produce			application/json
// @Success			200 {object} user.UserDto{} "Response Success (UserDto.go)"
func (controller *UserController) FindById(ctx *gin.Context) {
	log.Info().Msg("Update")
	baseController.CreateWebResponse(http.StatusOK, "OK", "Update", ctx)

}

// @Tags			user
// @Router			/v1/user/list [get]
// @Summary			list
// @Description		list
// @Produce			application/json
// @Success			200 {object} user.UserDto{} "Response Success (UserDto.go)"
func (controller *UserController) FindAll(ctx *gin.Context) {
	log.Info().Msg("FindAll")
	baseController.CreateWebResponse(http.StatusOK, "OK", "FindAll", ctx)
}

//log.Info().Msg("delete tags")
//tagId := ctx.Param("tagId")
//id, err := strconv.Atoi(tagId)
//helper.ErrorPanic(err)
//controller.prodcutService.Delete(id)
//
//responseBody := dto.WebResponse{
//	Code:   http.StatusOK,
//	Status: "Ok",
//	Data:   nil,
//}
//ctx.Header("Content-Type", "application/json")
//ctx.JSON(http.StatusOK, responseBody)

//log.Info().Msg("findbyid tags")
//tagId := ctx.Param("tagId")
//id, err := strconv.Atoi(tagId)
//helper.ErrorPanic(err)
//
//data := controller.prodcutService.FindById(id)
//
//responseBody := dto.WebResponse{
//	Code:   http.StatusOK,
//	Status: "Ok",
//	Data:   data,
//}
//ctx.Header("Content-Type", "application/json")
//ctx.JSON(http.StatusOK, responseBody)

//log.Info().Msg("create tags")
//createTagsRequest := dto.UserDto{}
//err := ctx.ShouldBindJSON(&createTagsRequest)
//helper.ErrorPanic(err)
//
//controller.prodcutService.Create(createTagsRequest)
//responseBody := dto.WebResponse{
//	Code:   http.StatusOK,
//	Status: "Ok",
//	Data:   nil,
//}
//ctx.Header("Content-Type", "application/json")
//ctx.JSON(http.StatusOK, responseBody)
