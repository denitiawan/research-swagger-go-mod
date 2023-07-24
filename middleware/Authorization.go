package middleware

import (
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

		// replace prefix on token jwt
		token = strings.Replace(authorizationHeader, "Bearer ", "", 1)

		// validation if token is empty
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "fail",
				"message": "You are not logged in",
			})
			return
		}

		// get userId from  jwt token
		jwtValue, err := security.ValidateToken(token, appConfig.TokenSecret)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "fail",
				"message": err.Error(),
			})
			return
		}

		// ctx response
		userId, _ := jwtValue.(int64)
		ctx.Set("userId", userId)
		ctx.Next()

		// repo
		//authRepo := auth.NewAuthRepoImpl(appConfig, ctx, db)

		// find user by id
		//userId, _ := jwtValue.(int64)
		//res := authRepo.FindById(userId)
		//if res.Error != nil {
		//	ctx.AbortWithStatusJSON(
		//		http.StatusForbidden,
		//		gin.H{
		//			"status":  "fail",
		//			"message": "the user belonging to this token no logger exists",
		//		},
		//	)
		//	return
		//}

		// set context, this context will used on every controller when authorization success
		//userModel, _ := res.Data.(user.User)
		//ctx.Set("currentUser", userModel.Username)
		//ctx.Next()

	}
}
