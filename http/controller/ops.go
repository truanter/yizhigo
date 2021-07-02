package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/truanter/yizhigo/config"
	"github.com/truanter/yizhigo/http/app/error_code"
	"github.com/truanter/yizhigo/http/app/response"
)

func OpsButtonText(ctx *gin.Context) {
	d := make(map[string]string)
	err := ctx.BindJSON(&d)
	if err != nil {
		response.Error(ctx, error_code.InputError, err.Error())
		return
	}
	for k, v := range d {
		config.ModifyButtonText(k, v)
	}
	response.Success(ctx)
}

func OpsBlockPlatform(ctx *gin.Context) {
	p := make([]string, 0)
	err := ctx.BindJSON(&p)
	if err != nil {
		response.Error(ctx, error_code.InputError, err.Error())
		return
	}
	config.SetBlockPlatform(p)
	response.Success(ctx)
}

func OpsAddKeyWord(ctx *gin.Context) {
	k := make([]string, 0)
	err := ctx.BindJSON(&k)
	if err != nil {
		response.Error(ctx, error_code.InputError, err.Error())
		return
	}
	config.AddKeyWord(k)
	response.Success(ctx)
}

func MyBigBrotherComing(ctx *gin.Context) {
	config.MyBigBrother(true)
	response.Success(ctx)
}

func MyBigBrotherLeaving(ctx *gin.Context) {
	config.MyBigBrother(false)
	response.Success(ctx)
}
