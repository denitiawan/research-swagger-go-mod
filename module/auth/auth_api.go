package auth

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func APIAuth(db *sql.DB, baseRouter *gin.RouterGroup, validate *validator.Validate) {
	// base
	authRoute := baseRouter.Group("/v1/auth")

	// repo,service,controller
	authRepo := NewAuthRepoImpl(db)
	authService := NewAuthServiceImpl(authRepo, validate)
	authController := NewAuthController(authService)

	// endpoint
	authRoute.POST("/login", authController.Login)
}
