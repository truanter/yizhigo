package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterV1(router *gin.Engine) {
	goV1 := router.Group("/go/v1")
	{
		goV1.GET("/hello", func(ctx *gin.Context) {
			name :=	ctx.DefaultQuery("name", "yizhigo")
			ctx.String(http.StatusOK, fmt.Sprintf("hello, %s", name))
		})
	}
}
