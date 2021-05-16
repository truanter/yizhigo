package main

import (
	"github.com/gin-gonic/gin"
	"github.com/truanter/yizhigo/config"
	"github.com/truanter/yizhigo/router"
)

func main() {
	if !config.IsProdEnv() {
		gin.SetMode(gin.DebugMode)
	}
	r := router.InitRouter()
	_ = r.Run(":8000")
}
