package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/truanter/yizhigo/router/v1"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{}), gin.Recovery())
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello yizhigo")
	})
	v1.RegisterV1(router)
	return router
}