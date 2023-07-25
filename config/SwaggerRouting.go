package config

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// path swagger is customable
// path (/*any) is required for load the html page own by swagger
// http://localhost:8899/nexsoft/doc/api/swagger/index.html
func SwaggerRouting(router *gin.Engine) {
	router.GET("nexsoft/doc/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
