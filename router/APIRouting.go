package router

import (
	"database/sql"
	"denitiawan/research-swagger-gomod-gin/module/auth"
	"denitiawan/research-swagger-gomod-gin/module/product"
	"denitiawan/research-swagger-gomod-gin/module/user"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func APIRouting(db *sql.DB, router *gin.Engine) {

	// # Validator
	validate := validator.New()

	// # Base Router
	baseRouter := router.Group("")
	APIWelcome(router)

	// # API Registrations
	auth.APIAuth(db, baseRouter, validate)
	user.APIUser(db, baseRouter, validate)
	product.APIProduct(db, baseRouter, validate)

}
