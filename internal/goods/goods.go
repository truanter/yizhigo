package goods

import (
	"github.com/truanter/top-go/pkg/model"
	"github.com/truanter/yizhigo/config"
	"github.com/truanter/yizhigo/model/goods"
	"github.com/truanter/yizhigo/pkg/common"
	"strconv"
	"strings"
)

func Search(q, pageSizeStr, pageNOStr string) (map[string]interface{}, error) {
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
		return nil, common.NewRuntimeError(err.Error())
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

	if resp.ErrorResponse.Code != 0 {
		return res, common.NewTbkError(resp.ErrorResponse.Code, resp.ErrorResponse.SubCode, resp.ErrorResponse.Msg, resp.ErrorResponse.SubMsg, resp.ErrorResponse.RequestID)
	}
	return res, nil
}

func GetMySpecialFavorites() []goods.Favorites {
	res := make([]goods.Favorites, 0)
	for _, v := range myFavorites {
		var data map[string]interface{}
		var err error
		for i := 0; i < 3; i++ {
			data, err = GetFavoriteList(strconv.FormatInt(v.FavoritesID, 10), "5", "")
			if err == nil {
				break
			}
		}
		if data != nil {
			v.Data = data
			res = append(res, v)
		}
	}
	return res
}

func GetAllMyFavorites() ([]map[string]interface{}, error) {
	resp, err := top.TbkDgOptimusMaterial().Do(model.TbkDgOptimusMaterialRequest{
		PageSize:   100,
		AdzoneID:   config.GetAdzoneID(),
		MaterialID: config.MaterialIDGetFavorites,
	})
	if err != nil {
		return nil, common.NewRuntimeError(err.Error())
	}
	if resp.ErrorResponse.Code != 0 {
		return resp.ResultList, common.NewTbkError(resp.ErrorResponse.Code, resp.ErrorResponse.SubCode, resp.ErrorResponse.Msg, resp.ErrorResponse.SubMsg, resp.ErrorResponse.RequestID)
	}
	return resp.ResultList, nil
}

func GetFavoriteList(favoriteID, pageSizeStr, pageNOStr string) (map[string]interface{}, error) {
	materialID := config.MaterialIDGetFavoritesList
	return generalTbkGoodsList(materialID, favoriteID, pageSizeStr, pageNOStr, 0)
}

func GetGuessYouLike(pageSizeStr, pageNOStr string) (map[string]interface{}, error) {
	materialID := config.MaterialIDGuessYouLike
	return generalTbkGoodsList(materialID, "", pageSizeStr, pageNOStr, 0)
}

func GetSimilarGoods(itemID int64, pageSizeStr, pageNOStr string) (map[string]interface{}, error) {
	materialID := config.MaterialIDGetSimilarGoods
	return generalTbkGoodsList(materialID, "", pageSizeStr, pageNOStr, itemID)
}

func CreateTPWD(url string) (map[string]interface{}, error) {
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
		return nil, common.NewRuntimeError(err.Error())
	}
	if resp.ErrorResponse.Code != 0 {
		return resp.Data, common.NewTbkError(resp.ErrorResponse.Code, resp.ErrorResponse.SubCode, resp.ErrorResponse.Msg, resp.ErrorResponse.SubMsg, resp.ErrorResponse.RequestID)
	}
	return resp.Data, nil
}

func generalTbkGoodsList(materialID int64, favoritesID, pageSizeStr, pageNOStr string, itemID int64) (map[string]interface{}, error) {
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
		MaterialID:  materialID,
		FavoritesID: favoritesID,
		ItemID:      itemID,
	})

	if err != nil {
		return nil, common.NewRuntimeError(err.Error())
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

	if resp.ErrorResponse.Code != 0 {
		return res, common.NewTbkError(resp.ErrorResponse.Code, resp.ErrorResponse.SubCode, resp.ErrorResponse.Msg, resp.ErrorResponse.SubMsg, resp.ErrorResponse.RequestID)
	}
	return res, nil
}
