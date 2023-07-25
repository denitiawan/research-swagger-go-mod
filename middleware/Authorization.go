package middleware

import (
	"denitiawan/research-swagger-gomod-gin/common/dto"
	"denitiawan/research-swagger-gomod-gin/config"
	"denitiawan/research-swagger-gomod-gin/security"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Authorization(appConfig config.AppConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token string

		// get token from http header
		authorizationHeader := ctx.Request.Header.Get("Authorization")

		// replace prefix on jwt token
		token = strings.Replace(authorizationHeader, "Bearer ", "", 1)

		// validate : token is empty
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				dto.NewWebResponse(http.StatusUnauthorized, "Token is Empty", "", nil))
			return
		}

		// validate : expired time, jwt format, get token value, etc
		jwtValue, err := security.ValidateToken(token, appConfig.TokenSecret)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				dto.NewWebResponse(http.StatusUnauthorized, "Token is invalid", err.Error(), nil))
			return
		}

		// ctx response
		userId, _ := jwtValue.(int64)
		ctx.Set("userId", userId)
		ctx.Next()

	}
}
