package controller

import (
	"github.com/gin-gonic/gin"
	conf "github.com/truanter/yizhigo/config"
	"github.com/truanter/yizhigo/http/app/error_code"
	"github.com/truanter/yizhigo/http/app/response"
	impl "github.com/truanter/yizhigo/internal/goods"
	"github.com/truanter/yizhigo/model/goods"
	"github.com/truanter/yizhigo/pkg/common"
)

func GetConfig(ctx *gin.Context) {
	response.SuccessWithData(ctx, conf.GetRuntimeConfig())
}

func GetIndex(ctx *gin.Context) {
	pageSizeStr := ctx.Request.URL.Query().Get("page_size")
	pageNOStr := ctx.Request.URL.Query().Get("page_no")
	res, err := impl.GetGuessYouLike(pageSizeStr, pageNOStr)
	if common.IsRuntimeError(err) {
		response.Error(ctx, error_code.InternalError, err.Error())
		return
	}
	favorites := make([]goods.Favorites, 0)
	if pageNOStr == "" || pageNOStr == "1" {
		favorites = impl.GetMySpecialFavorites()
	}
	res["favorites"] = favorites
	response.SuccessWithData(ctx, res)
	return
}

func Search(ctx *gin.Context) {

	q := ctx.Request.URL.Query().Get("q")
	pageSizeStr := ctx.Request.URL.Query().Get("page_size")
	pageNOStr := ctx.Request.URL.Query().Get("page_no")
	isBlock := ctx.Request.URL.Query().Get("is_block")
	res, err := impl.Search(q, pageSizeStr, pageNOStr, isBlock)
	if common.IsRuntimeError(err) {
		response.Error(ctx, error_code.InternalError, err.Error())
		return
	}
	response.SuccessWithData(ctx, res)
}

func GetFavorites(ctx *gin.Context) {
	resp, err := impl.GetAllMyFavorites()
	if common.IsRuntimeError(err) {
		response.Error(ctx, error_code.InternalError, err.Error())
	}
	response.SuccessWithData(ctx, resp)
	return
}

func GetFavoritesLocal(ctx *gin.Context) {
	response.SuccessWithData(ctx, impl.GetMySpecialFavorites())
}

func GetFavoriteList(ctx *gin.Context) {
	FavoritesID := ctx.Request.URL.Query().Get("category_id")
	pageSizeStr := ctx.Request.URL.Query().Get("page_size")
	pageNOStr := ctx.Request.URL.Query().Get("page_no")
	res, err := impl.GetFavoriteList(FavoritesID, pageSizeStr, pageNOStr)
	if common.IsRuntimeError(err) {
		response.Error(ctx, error_code.InternalError, err.Error())
		return
	}
	response.SuccessWithData(ctx, res)
}

func GetSimilarGoods(ctx *gin.Context) {
	itemID := ctx.Request.URL.Query().Get("item_id")
	// itemID, err := strconv.ParseInt(itemIDStr, 10, 64)
	pageSizeStr := ctx.Request.URL.Query().Get("page_size")
	pageNOStr := ctx.Request.URL.Query().Get("page_no")
	res, err := impl.GetSimilarGoods(itemID, pageSizeStr, pageNOStr)
	if common.IsRuntimeError(err) {
		response.Error(ctx, error_code.InternalError, err.Error())
		return
	}
	response.SuccessWithData(ctx, res)
}

func CreateTPWD(ctx *gin.Context) {
	var data map[string]string
	err := ctx.BindJSON(&data)
	if err != nil {
		response.Error(ctx, error_code.InputError, err.Error())
		return
	}
	url := data["url"]

	resp, err := impl.CreateTPWD(url)

	if err != nil {
		response.Error(ctx, error_code.TbkRequestError, err.Error())
		return
	}
	response.SuccessWithData(ctx, resp)
}
