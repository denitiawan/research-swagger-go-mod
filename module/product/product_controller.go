package product

import (
	baseController "denitiawan/research-swagger-gomod-gin/common/controller"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

type ProductController struct {
	prodcutService ProductService
}

func NewProductController(service ProductService) *ProductController {
	return &ProductController{
		prodcutService: service,
	}
}

func (controller *ProductController) Create(ctx *gin.Context) {
	log.Info().Msg("Create")
	baseController.CreateWebResponse(http.StatusOK, "OK", "Create", ctx)
}

func (controller *ProductController) Update(ctx *gin.Context) {
	log.Info().Msg("Update")
	baseController.CreateWebResponse(http.StatusOK, "OK", "Update", ctx)
}

func (controller *ProductController) Delete(ctx *gin.Context) {
	log.Info().Msg("Delete")
	baseController.CreateWebResponse(http.StatusOK, "OK", "Delete", ctx)
}

func (controller *ProductController) FindById(ctx *gin.Context) {
	log.Info().Msg("FindById")
	baseController.CreateWebResponse(http.StatusOK, "OK", "FindById", ctx)
}

func (controller *ProductController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll")
	baseController.CreateWebResponse(http.StatusOK, "OK", "findAll", ctx)
}
