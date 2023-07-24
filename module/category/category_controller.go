package category

import (
	baseController "denitiawan/research-swagger-gomod-gin/common/controller"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

type CategoryController struct {
}

func NewCategoryController() *CategoryController {
	return &CategoryController{}
}

// @Tags			category
// @Router			/v1/category/save [post]
// @Summary			save
// @Description		save
// @Param			RequestBody body category.CategoryDto true "CategoryDto.go"
// @Produce			application/json
// @Success			200 {object} category.CategoryDto{} "Response Success (CategoryDto.go)"
func (controller *CategoryController) Create(ctx *gin.Context) {
	log.Info().Msg("Create")

	response := CategoryDto{}
	baseController.OK(http.StatusOK, "", response, "", ctx)
}

// @Tags			category
// @Router			/v1/category/update/{id} [put]
// @Summary			update
// @Description		update
// @Param			id  path int true "id"
// @Param			RequestBody body category.CategoryDto true "CategoryDto.go"
// @Produce			application/json
// @Success			200 {object} category.CategoryDto{} "Response Success (CategoryDto.go)"
func (controller *CategoryController) Update(ctx *gin.Context) {
	log.Info().Msg("Update")

	response := CategoryDto{}
	baseController.OK(http.StatusOK, "", response, "", ctx)
}

// @Tags			category
// @Router			/v1/category/delete/{id} [delete]
// @Summary			delete
// @Description		delete
// @Param			id  path int true "id"
// @Produce			application/json
// @Success			200 {object} category.CategoryDto{} "Response Success (CategoryDto.go)"
func (controller *CategoryController) Delete(ctx *gin.Context) {
	log.Info().Msg("Delete")

	response := CategoryDto{}
	baseController.OK(http.StatusOK, "", response, "", ctx)
}

// @Tags			category
// @Router			/v1/category/view/{id} [get]
// @Summary			view
// @Description		view
// @Param			id  path int true "id"
// @Produce			application/json
// @Success			200 {object} category.CategoryDto{} "Response Success (CategoryDto.go)"
func (controller *CategoryController) FindById(ctx *gin.Context) {
	log.Info().Msg("FindById")

	response := CategoryDto{}
	baseController.OK(http.StatusOK, "", response, "", ctx)
}

// @Tags			category
// @Router			/v1/category/list [get]
// @Summary			list
// @Description		list
// @Produce			application/json
// @Success			200 {object} category.CategoryDto{} "Response Success (CategoryDto.go)"
func (controller *CategoryController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll")

	response := CategoryDto{}
	baseController.OK(http.StatusOK, "", response, "", ctx)
}
