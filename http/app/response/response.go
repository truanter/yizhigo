package response

import (
	"github.com/gin-gonic/gin"
	"github.com/truanter/yizhigo/http/app/error_code"
	"net/http"
)

type Response struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func returnJson(ctx *gin.Context, code int, msg string, data interface{}) {
	ctx.AbortWithStatusJSON(http.StatusOK, Response{
		Code: code,
		Msg:  error_code.GetErrorMsg(code, msg),
		Data: data,
	})
}

func Success(ctx *gin.Context) {
	returnJson(ctx, error_code.Success, "", nil)
}

func SuccessWithData(ctx *gin.Context, data interface{}) {
	returnJson(ctx, error_code.Success, "", data)
}

func Error(ctx *gin.Context, code int, msg string) {
	returnJson(ctx, code, msg, nil)
}

func ErrorWithData(ctx *gin.Context, code int, msg string, data interface{}) {
	returnJson(ctx, code, msg, data)
}