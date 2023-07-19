package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func APIWelcome(router *gin.Engine) {
	router.GET("",
		func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "Welcome")
		})
}
