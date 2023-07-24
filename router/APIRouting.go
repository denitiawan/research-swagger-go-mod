package router

import (
	"database/sql"
	"denitiawan/research-swagger-gomod-gin/module/auth"
	"denitiawan/research-swagger-gomod-gin/module/category"
	"denitiawan/research-swagger-gomod-gin/module/product"
	"denitiawan/research-swagger-gomod-gin/module/user"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func APIRouting(db *sql.DB, basePath *gin.RouterGroup) {

	// # Validator
	validate := validator.New()

	// # API Registrations
	auth.APIAuth(db, basePath, validate)
	user.APIUser(db, basePath, validate)
	product.APIProduct(db, basePath, validate)
	category.APICategory(db, basePath, validate)

}
