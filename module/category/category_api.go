package category

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func APICategory(db *sql.DB, basePath *gin.RouterGroup, validate *validator.Validate) {

	// repo,service,controller

	controller := NewCategoryController()

	// endpoint
	productRoute := basePath.Group("/v1/category")
	productRoute.GET("/list", controller.FindAll)
	productRoute.GET("/view/:id", controller.FindById)
	productRoute.POST("/save", controller.Create)
	productRoute.PUT("/update/:id", controller.Update)
	productRoute.DELETE("/delete/:id", controller.Delete)
}
