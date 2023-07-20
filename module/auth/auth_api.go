package auth

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func APIAuth(db *sql.DB, basePath *gin.RouterGroup, validate *validator.Validate) {

	// repo,service,controller
	authRepo := NewAuthRepoImpl(db)
	authService := NewAuthServiceImpl(authRepo, validate)
	authController := NewAuthController(authService)

	// endpoint
	authRoute := basePath.Group("/v1/auth")
	authRoute.POST("/login", authController.Login)
}
