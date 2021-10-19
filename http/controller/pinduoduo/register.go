package pinduoduo

import (
	"github.com/gin-gonic/gin"
	"github.com/truanter/yizhigo/http/app/error_code"
	"github.com/truanter/yizhigo/http/app/response"
	impl "github.com/truanter/yizhigo/internal/pinduoduo"
)

type Register struct {
	UID string `json:"uid,omitempty" form:"uid" binding:"required"`
}

func (r Register) Post(ctx *gin.Context) {
	if err := ctx.ShouldBindJSON(&r); err != nil {
		response.Error(ctx, error_code.InputError, err.Error())
		return
	}
	resp, err := impl.Register(r.UID, true, true, true, true)
	if err != nil {
		response.Error(ctx, error_code.PddRequestError, err.Error())
		return
	}
	response.SuccessWithData(ctx, resp.RpPromotionUrlGenerateResponse)
}

func (r Register) IsBind(ctx *gin.Context) {
	if err := ctx.ShouldBind(&r); err != nil {
		response.Error(ctx, error_code.InputError, err.Error())
		return
	}
	resp, err := impl.CheckBind(r.UID)
	if err != nil {
		response.Error(ctx, error_code.PddRequestError, err.Error())
		return
	}
	response.SuccessWithData(ctx, resp.AuthorityQueryResponse)
}
