package pinduoduo

import (
	"github.com/gin-gonic/gin"
	"github.com/truanter/yizhigo/http/app/error_code"
	"github.com/truanter/yizhigo/http/app/response"
	impl "github.com/truanter/yizhigo/internal/pinduoduo"
)

type GoodsController struct {
	UID       string `json:"uid" form:"uid" binding:"required"`
	Q         string `form:"q"`
	IsBlock   string `json:"is_block" form:"is_block"`
	PageSize  string `form:"page_size"`
	PageNO    string `form:"page_no"`
	ListID    string `form:"list_id"`
	SearchID  string `form:"search_id"`
	GoodsSign string `form:"goods_sign"`
	Count     int    `form:"count"`
}

func (g *GoodsController) paramsCheck(ctx *gin.Context) {
	if err := ctx.ShouldBind(&g); err != nil {
		response.Error(ctx, error_code.InputError, err.Error())
		return
	}
}

func (g GoodsController) Index(ctx *gin.Context) {
	if err := ctx.ShouldBind(&g); err != nil {
		response.Error(ctx, error_code.InputError, err.Error())
		return
	}
	resp, err := impl.GuessYouLike(g.UID, g.ListID, g.PageSize, g.PageNO)
	if err != nil {
		response.Error(ctx, error_code.PddRequestError, err.Error())
		return
	}
	response.SuccessWithData(ctx, resp)
	return
}

func (g GoodsController) GetCategories(ctx *gin.Context) {
	response.SuccessWithData(ctx, impl.Categories)
	return
}
func (g GoodsController) Search(ctx *gin.Context) {
	if err := ctx.ShouldBind(&g); err != nil {
		response.Error(ctx, error_code.InputError, err.Error())
		return
	}

	resp, err := impl.Search(g.UID, g.Q, g.ListID, g.PageSize, g.PageNO, g.IsBlock)
	if err != nil {
		response.Error(ctx, error_code.PddRequestError, err.Error())
		return
	}
	response.SuccessWithData(ctx, resp)
	return
}

func (g GoodsController) Detail(ctx *gin.Context) {
	if err := ctx.ShouldBind(&g); err != nil {
		response.Error(ctx, error_code.InputError, err.Error())
		return
	}
	if g.GoodsSign == "" {
		response.Error(ctx, error_code.InputError, "goods_sign未录入")
		return
	}

	resp, err := impl.GoodsDetail(g.UID, g.GoodsSign, g.SearchID)
	if err != nil {
		response.Error(ctx, error_code.PddRequestError, err.Error())
		return
	}
	response.SuccessWithData(ctx, resp)
}
func (g GoodsController) GetSimilar(ctx *gin.Context) {
	if err := ctx.ShouldBind(&g); err != nil {
		response.Error(ctx, error_code.InputError, err.Error())
		return
	}
	if g.GoodsSign == "" {
		response.Error(ctx, error_code.InputError, "goods_sign未录入")
		return
	}
	resp, err := impl.SimilarGoods(g.UID, g.GoodsSign, g.Count)
	if err != nil {
		response.Error(ctx, error_code.PddRequestError, err.Error())
		return
	}
	response.SuccessWithData(ctx, resp)
}

func (g GoodsController) Promotion(ctx *gin.Context) {
	if err := ctx.ShouldBind(&g); err != nil {
		response.Error(ctx, error_code.InputError, err.Error())
		return
	}
	if g.GoodsSign == "" {
		response.Error(ctx, error_code.InputError, "goods_sign未录入")
		return
	}
	resp, err := impl.GenerateGoodsURL(g.UID, g.GoodsSign, g.SearchID, true, true)
	if err != nil {
		response.Error(ctx, error_code.PddRequestError, err.Error())
		return
	}
	response.SuccessWithData(ctx, resp)
	return
}
