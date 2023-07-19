package user

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func APIUser(db *sql.DB, baseRouter *gin.RouterGroup, validate *validator.Validate) {
	// base
	userRoute := baseRouter.Group("/v1/user")

	// repo,service,controller
	userRepo := NewUserRepoImpl(db)
	userService := NewUserServiceImpl(userRepo, validate)
	userController := NewUserController(userService)

	// endpoint
	userRoute.GET("/list", userController.FindAll)
	userRoute.GET("/view/:Id", userController.FindById)
	userRoute.POST("/save", userController.Create)
	userRoute.PUT("/update/:Id", userController.Update)
	userRoute.DELETE("/delete/:Id", userController.Delete)
}
