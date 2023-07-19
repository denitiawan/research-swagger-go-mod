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

func (controller *UserController) Create(ctx *gin.Context) {
	log.Info().Msg("Create")
	baseController.CreateWebResponse(http.StatusOK, "OK", "Create", ctx)
}

func (controller *UserController) Update(ctx *gin.Context) {
	log.Info().Msg("Update")
	baseController.CreateWebResponse(http.StatusOK, "OK", "Update", ctx)
}

func (controller *UserController) Delete(ctx *gin.Context) {
	log.Info().Msg("Delete")
	baseController.CreateWebResponse(http.StatusOK, "OK", "Delete", ctx)
}

func (controller *UserController) FindById(ctx *gin.Context) {
	log.Info().Msg("Update")
	baseController.CreateWebResponse(http.StatusOK, "OK", "Update", ctx)

}

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
