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

// @Tags			product
// @Router			/v1/product/save [post]
// @Summary			save
// @Description		save
// @Param			RequestBody body product.ProductDto true "ProductDto.go"
// @Produce			application/json
// @Success			200 {object} product.ProductDto{} "Response Success (ProductDto.go)"
func (controller *ProductController) Create(ctx *gin.Context) {
	log.Info().Msg("Create")
	baseController.CreateWebResponse(http.StatusOK, "OK", "Create", ctx)
}

// @Tags			product
// @Router			/v1/product/update/{id} [put]
// @Summary			update
// @Description		update
// @Param			RequestBody body product.ProductDto true "ProductDto.go"
// @Produce			application/json
// @Success			200 {object} product.ProductDto{} "Response Success (ProductDto.go)"
func (controller *ProductController) Update(ctx *gin.Context) {
	log.Info().Msg("Update")
	baseController.CreateWebResponse(http.StatusOK, "OK", "Update", ctx)
}

// @Tags			product
// @Router			/v1/product/delete/{id} [delete]
// @Summary			delete
// @Description		delete
// @Param			RequestBody body product.ProductDto true "ProductDto.go"
// @Produce			application/json
// @Success			200 {object} product.ProductDto{} "Response Success (ProductDto.go)"
func (controller *ProductController) Delete(ctx *gin.Context) {
	log.Info().Msg("Delete")
	baseController.CreateWebResponse(http.StatusOK, "OK", "Delete", ctx)
}

// @Tags			product
// @Router			/v1/product/view/{id} [get]
// @Summary			view
// @Description		view
// @Param			RequestBody body product.ProductDto true "ProductDto.go"
// @Produce			application/json
// @Success			200 {object} product.ProductDto{} "Response Success (ProductDto.go)"
func (controller *ProductController) FindById(ctx *gin.Context) {
	log.Info().Msg("FindById")
	baseController.CreateWebResponse(http.StatusOK, "OK", "FindById", ctx)
}

// @Tags			product
// @Router			/v1/product/list [get]
// @Summary			list
// @Description		list
// @Param			RequestBody body product.ProductDto true "ProductDto.go"
// @Produce			application/json
// @Success			200 {object} product.ProductDto{} "Response Success (ProductDto.go)"
func (controller *ProductController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll")
	baseController.CreateWebResponse(http.StatusOK, "OK", "findAll", ctx)
}
