package user

import (
	"context"
	"database/sql"
	"denitiawan/research-swagger-gomod-gin/config"
	"denitiawan/research-swagger-gomod-gin/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func APIUser(appConfig config.AppConfig, ctx context.Context, db *sql.DB, basePath *gin.RouterGroup, validate *validator.Validate) {

	// repo,service,controller
	userRepo := NewUserRepoImpl(ctx, db)
	userService := NewUserServiceImpl(userRepo, validate)
	userController := NewUserController(userService)

	// endpoint
	userRoute := basePath.Group("/v1/user")
	userRoute.GET("/list", middleware.Authorization(appConfig), userController.FindAll)
	userRoute.GET("/view/:id", middleware.Authorization(appConfig), userController.FindById)
	userRoute.POST("/save", middleware.Authorization(appConfig), userController.Create)
	userRoute.PUT("/update/:id", middleware.Authorization(appConfig), userController.Update)
	userRoute.DELETE("/delete/:id", middleware.Authorization(appConfig), userController.Delete)
}
