package product

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func APIProduct(db *sql.DB, baseRouter *gin.RouterGroup, validate *validator.Validate) {
	// base
	productRoute := baseRouter.Group("/v1/product")

	// repo,service,controller
	productRepo := NewProductRepoImpl(db)
	productService := NewProductServiceImpl(productRepo, validate)
	productController := NewProductController(productService)

	// endpoint
	productRoute.GET("/list", productController.FindAll)
	productRoute.GET("/view/:Id", productController.FindById)
	productRoute.POST("/save", productController.Create)
	productRoute.PUT("/update/:Id", productController.Update)
	productRoute.DELETE("/delete/:Id", productController.Delete)
}
