package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/truanter/yizhigo/config"
	"github.com/truanter/yizhigo/http/app/error_code"
	"github.com/truanter/yizhigo/http/app/response"
)

func AccessTokenCheck(ctx *gin.Context) {
	accessToken := ctx.Request.Header.Get("access_token")
	authToken := config.GetAuthToken()

	if authToken == "" {
		ctx.AbortWithStatus(500)
		return
	}

	if accessToken != authToken {
		response.Error(ctx, error_code.PermissionDeny, "")
		return
	}

	ctx.Next()
}
