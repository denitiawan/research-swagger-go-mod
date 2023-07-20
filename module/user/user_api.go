package user

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func APIUser(db *sql.DB, basePath *gin.RouterGroup, validate *validator.Validate) {

	// repo,service,controller
	userRepo := NewUserRepoImpl(db)
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
