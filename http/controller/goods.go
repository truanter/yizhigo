package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/truanter/top-go/pkg/model"
	"github.com/truanter/yizhigo/config"
	"github.com/truanter/yizhigo/http/app/error_code"
	"github.com/truanter/yizhigo/http/app/response"
	"github.com/truanter/yizhigo/model/goods"
	"strconv"
	"strings"
)

func Search(ctx *gin.Context) {

	q := ctx.Request.URL.Query().Get("q")
	pageSizeStr := ctx.Request.URL.Query().Get("page_size")
	pageNOStr := ctx.Request.URL.Query().Get("page_no")
	pageSize := 20
	pageNO := 1

	if pageSizeStr != "" {
		pageSize, _ = strconv.Atoi(pageSizeStr)
	}

	if pageNOStr != "" {
		pageNO, _ = strconv.Atoi(pageNOStr)
	}

	resp, err := top.TbkDgMaterialOptional().Do(model.TbkDgMaterialOptionalRequest{
		PageSize:  int64(pageSize),
		PageNO:    int64(pageNO),
		AdzoneID:  config.GetAdzoneID(),
		Q:         q,
		HasCoupon: true,
	})
	if err != nil {
		response.Error(ctx, error_code.TbkRequestError, err.Error())
		return
	}
	if resp.ErrorResponse.Code != 0 {
		response.Error(ctx, error_code.TbkRequestError, fmt.Sprintf("code: %d, sub_code: %s, msg: %s, sub_msg: %s", resp.ErrorResponse.Code, resp.ErrorResponse.SubCode, resp.ErrorResponse.Msg, resp.ErrorResponse.SubMsg))
		return
	}
	goodsList := make([]map[string]interface{}, 0)
	for _, v := range resp.ResultList {
		goodsList = append(goodsList, goods.Parse(v))
	}
	res := map[string]interface{}{
		"list":            goodsList,
		"total":           resp.TotalResults,
		"page_result_key": resp.PageResultKey,
	}
	response.SuccessWithData(ctx, res)
	return
}

func GetFavorites(ctx *gin.Context) {
	MaterialID := int64(31519)
	resp, err := top.TbkDgOptimusMaterial().Do(model.TbkDgOptimusMaterialRequest{
		PageSize:   100,
		AdzoneID:   config.GetAdzoneID(),
		MaterialID: MaterialID,
	})
	if err != nil {
		response.Error(ctx, error_code.TbkRequestError, err.Error())
		return
	}
	if resp.ErrorResponse.Code != 0 {
		response.Error(ctx, error_code.TbkRequestError, fmt.Sprintf("code: %d, sub_code: %s, msg: %s, sub_msg: %s", resp.ErrorResponse.Code, resp.ErrorResponse.SubCode, resp.ErrorResponse.Msg, resp.ErrorResponse.SubMsg))
		return
	}
	response.SuccessWithData(ctx, resp.ResultList)
	return
}

func GetFavoriteList(ctx *gin.Context) {
	MaterialID := int64(31539)
	FavoritesID := ctx.Request.URL.Query().Get("category_id")
	pageSizeStr := ctx.Request.URL.Query().Get("page_size")
	pageNOStr := ctx.Request.URL.Query().Get("page_no")
	pageSize := 20
	pageNO := 1

	if pageSizeStr != "" {
		pageSize, _ = strconv.Atoi(pageSizeStr)
	}

	if pageNOStr != "" {
		pageNO, _ = strconv.Atoi(pageNOStr)
	}

	resp, err := top.TbkDgOptimusMaterial().Do(model.TbkDgOptimusMaterialRequest{
		PageSize:    int64(pageSize),
		PageNO:      int64(pageNO),
		AdzoneID:    config.GetAdzoneID(),
		MaterialID:  MaterialID,
		FavoritesID: FavoritesID,
	})
	if err != nil {
		response.Error(ctx, error_code.TbkRequestError, err.Error())
		return
	}
	if resp.ErrorResponse.Code != 0 {
		response.Error(ctx, error_code.TbkRequestError, fmt.Sprintf("code: %d, sub_code: %s, msg: %s, sub_msg: %s", resp.ErrorResponse.Code, resp.ErrorResponse.SubCode, resp.ErrorResponse.Msg, resp.ErrorResponse.SubMsg))
		return
	}

	goodsList := make([]map[string]interface{}, 0)
	for _, v := range resp.ResultList {
		goodsList = append(goodsList, goods.Parse(v))
	}
	res := map[string]interface{}{
		"list":       goodsList,
		"total":      resp.TotalCount,
		"is_default": resp.IsDefault,
	}
	response.SuccessWithData(ctx, res)
	return
}

func CreateTPWD(ctx *gin.Context) {
	var data map[string]string
	err := ctx.BindJSON(&data)
	if err != nil {
		response.Error(ctx, error_code.InputError, err.Error())
		return
	}
	url := data["url"]

	if !strings.HasPrefix(url, "http") {
		if strings.HasPrefix(url, "//") {
			url = "https:" + url
		} else if strings.HasPrefix(url, "/") {
			url = "https:/" + url
		} else {
			url = "https://" + url
		}
	}

	resp, err := top.TbkTpwdCreate().Do(model.TbkTpwdCreateRequest{Url: url})
	if err != nil {
		response.Error(ctx, error_code.TbkRequestError, err.Error())
		return
	}
	if resp.ErrorResponse.Code != 0 {
		response.Error(ctx, error_code.TbkRequestError, fmt.Sprintf("code: %d, sub_code: %s, msg: %s, sub_msg: %s", resp.ErrorResponse.Code, resp.ErrorResponse.SubCode, resp.ErrorResponse.Msg, resp.ErrorResponse.SubMsg))
		return
	}
	response.SuccessWithData(ctx, resp.Data)
	return
}
