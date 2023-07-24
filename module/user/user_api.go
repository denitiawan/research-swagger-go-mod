package user

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func APIUser(db *sql.DB, basePath *gin.RouterGroup, validate *validator.Validate) {

	context := context.Background()

	// repo,service,controller
	userRepo := NewUserRepoImpl(context, db)
	userService := NewUserServiceImpl(userRepo, validate)
	userController := NewUserController(userService)

	// endpoint
	userRoute := basePath.Group("/v1/user")
	userRoute.GET("/list", userController.FindAll)
	userRoute.GET("/view/:id", userController.FindById)
	userRoute.POST("/save", userController.Create)
	userRoute.PUT("/update/:id", userController.Update)
	userRoute.DELETE("/delete/:id", userController.Delete)
}
