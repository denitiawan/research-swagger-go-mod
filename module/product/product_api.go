package product

import (
	"context"
	"database/sql"
	"denitiawan/research-swagger-gomod-gin/config"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func APIProduct(appConfig config.AppConfig, ctx context.Context, db *sql.DB, basePath *gin.RouterGroup, validate *validator.Validate) {

	// repo,service,controller
	productRepo := NewProductRepoImpl(db)
	productService := NewProductServiceImpl(productRepo, validate)
	productController := NewProductController(productService)

	// endpoint
	productRoute := basePath.Group("/v1/product")
	productRoute.GET("/list", productController.FindAll)
	productRoute.GET("/view/:id", productController.FindById)
	productRoute.POST("/save", productController.Create)
	productRoute.PUT("/update/:id", productController.Update)
	productRoute.DELETE("/delete/:id", productController.Delete)
}
