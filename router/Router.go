package router

import (
	"github.com/gin-gonic/gin"
)

func NewRouterInit() *gin.Engine {
	router := gin.Default()
	return router
}
