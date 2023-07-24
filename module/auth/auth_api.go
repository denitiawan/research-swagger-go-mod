package auth

import (
	"context"
	"database/sql"
	"denitiawan/research-swagger-gomod-gin/config"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func APIAuth(appConfig config.AppConfig, ctx context.Context, db *sql.DB, basePath *gin.RouterGroup, validate *validator.Validate) {

	// repo,service,controller
	authRepo := NewAuthRepoImpl(appConfig, ctx, db)
	authService := NewAuthServiceImpl(authRepo, validate)
	authController := NewAuthController(authService)

	// endpoint
	authRoute := basePath.Group("/v1/auth")
	authRoute.POST("/login", authController.Login)
}
