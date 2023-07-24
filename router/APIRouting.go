package router

import (
	"context"
	"database/sql"
	"denitiawan/research-swagger-gomod-gin/config"
	"denitiawan/research-swagger-gomod-gin/module/auth"
	"denitiawan/research-swagger-gomod-gin/module/category"
	"denitiawan/research-swagger-gomod-gin/module/product"
	"denitiawan/research-swagger-gomod-gin/module/user"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func APIRouting(appConfig config.AppConfig, db *sql.DB, basePath *gin.RouterGroup) {

	// # context
	context := context.Background()

	// # Validator
	validate := validator.New()

	// # API Registrations
	auth.APIAuth(appConfig, context, db, basePath, validate)
	user.APIUser(appConfig, context, db, basePath, validate)
	product.APIProduct(appConfig, context, db, basePath, validate)
	category.APICategory(appConfig, context, db, basePath, validate)

}
