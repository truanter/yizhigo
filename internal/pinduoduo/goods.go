package pinduoduo

import (
	"fmt"
	"github.com/truanter/pinduoduo-go/pkg/models"
	"github.com/truanter/yizhigo/model/goods"
	"github.com/truanter/yizhigo/pkg/common"
	"github.com/truanter/yizhigo/pkg/util"
	"strconv"
)

func Search(uid, keyWord, listID, pageSizeStr, pageNOStr, isBlock string) (goods.PinduoduoGoodsList, error) {
	cp := generateCustomParams(uid)
	res := goods.PinduoduoGoodsList{
		List:     make([]goods.YiZhiGood, 0),
		ListID:   "",
		SearchID: "",
		Total:    0,
	}

	if util.IsBlock(keyWord, isBlock) {
		res.IsBlock = true
		return res, nil
	}
	pageSize := 20
	pageNO := 1

	if pageSizeStr != "" {
		pageSize, _ = strconv.Atoi(pageSizeStr)
	}

	if pageNOStr != "" {
		pageNO, _ = strconv.Atoi(pageNOStr)
	}
	withCoupon := true
	isBrand := true
	resp, err := ddk.DDKGoodsSearch().Do(models.DDKGoodsSearchRequest{
		Keyword:          &keyWord,
		Page:             &pageNO,
		PageSize:         &pageSize,
		WithCoupon:       &withCoupon,
		ListId:           &listID,
		CustomParameters: &cp,
		Pid:              &pid,
		IsBrandGoods:     &isBrand,
	})
	if err != nil {
		return res, err
	}

	if resp.BaseResponse.ErrorResponse.ErrorCode != 0 {
		return res, nil
	}

	goodsList := make([]goods.YiZhiGood, 0)
	for _, v := range resp.GoodsSearchResponse.GoodsList {
		goodsList = append(goodsList, goods.TransferPinduoduo(v))
	}
	res.List = goodsList
	res.SearchID = resp.GoodsSearchResponse.SearchID
	res.Total = resp.GoodsSearchResponse.TotalCount
	return res, nil
}

func GoodsDetail(uid, goodSign, searchID string) (goods.YiZhiGood, error) {
	res := new(goods.YiZhiGood)
	customParams := generateCustomParams(uid)
	resp, err := ddk.DDKGoodsDetail().Do(models.DDKGoodsDetailRequest{
		Pid:              &pid,
		CustomParameters: &customParams,
		SearchId:         &searchID,
		GoodsSign:        &goodSign,
	})

	if err != nil {
		return *res, err
	}

	if resp.ErrorResponse.ErrorCode != 0 {
		err = common.NewPddRequestError(resp.ErrorResponse.ErrorCode, resp.ErrorResponse.SubCode, resp.ErrorResponse.ErrorMsg, resp.ErrorResponse.SubMsg, resp.ErrorResponse.RequestID)
		return *res, err
	}
	*res = goods.TransferPinduoduoDetail(resp.GoodsDetailResponse.GoodsDetails[0])
	return *res, nil
}

func GenerateGoodsURL(uid, goodSign, searchID string, generateQQApp, generateWeApp bool) (models.DDKGoodsPromotionUrlGenerateInfo, error) {
	res := new(models.DDKGoodsPromotionUrlGenerateInfo)
	generateShortURL := true
	customParams := generateCustomParams(uid)
	resp, err := ddk.DDKGoodsPromotionUrlGenerate().Do(models.DDKGoodsPromotionUrlGenerateRequest{
		PID:              &pid,
		GoodsSignList:    &[]string{goodSign},
		GenerateQqApp:    &generateQQApp,
		GenerateWeApp:    &generateWeApp,
		SearchId:         &searchID,
		GenerateShortUrl: &generateShortURL,
		CustomParameters: &customParams,
	})
	if err != nil {
		return *res, err
	}
	if resp.ErrorResponse.ErrorCode != 0 {
		err = common.NewPddRequestError(resp.ErrorResponse.ErrorCode, resp.ErrorResponse.SubCode, resp.ErrorResponse.ErrorMsg, resp.ErrorResponse.SubMsg, resp.ErrorResponse.RequestID)
		return *res, err
	}
	*res = resp.GoodsPromotionUrlGenerateResponse.GoodsPromotionUrlList[0]
	return *res, nil
}

func GuessYouLike(uid, listID, pageSizeStr, pageNOStr string) (goods.PinduoduoGoodsList, error) {
	pageSize := 20
	pageNO := 1

	if pageSizeStr != "" {
		pageSize, _ = strconv.Atoi(pageSizeStr)
	}

	if pageNOStr != "" {
		pageNO, _ = strconv.Atoi(pageNOStr)
	}

	offset := (pageNO - 1) * pageSize
	return recommendGet(uid, listID, 0, 4, make([]string, 0), pageSize, offset)
}

func SimilarGoods(uid, goodsSign string, count int) ([]goods.YiZhiGood, error) {
	if count == 0 {
		count = 4
	}
	resp, err := recommendGet(uid, "", 0, 3, []string{goodsSign}, count, 0)
	if err != nil {
		return make([]goods.YiZhiGood, 0), err
	}
	return resp.List, nil
}

func recommendGet(uid, listID string, catID int64, channelType int, goodsSignList []string, limit, offset int) (res goods.PinduoduoGoodsList, err error) {
	cp := generateCustomParams(uid)
	resp, err := ddk.DDKGoodsRecommendGet().Do(models.DDKGoodsRecommendRequest{
		CatID:            &catID,
		CustomParameters: &cp,
		GoodsSignList:    &goodsSignList,
		ListID:           &listID,
		PID:              &pid,
		Offset:           &offset,
		Limit:            &limit,
		ChannelType:      &channelType,
	})
	if err != nil {
		return res, err
	}
	goodsList := make([]goods.YiZhiGood, 0)

	if resp.ErrorResponse.ErrorCode != 0 {
		err = common.NewPddRequestError(resp.ErrorResponse.ErrorCode, resp.ErrorResponse.SubCode, resp.ErrorResponse.ErrorMsg, resp.ErrorResponse.SubMsg, resp.ErrorResponse.RequestID)
		return res, err
	}
	for _, v := range resp.GoodsBasicDetailResponse.List {
		goodsList = append(goodsList, goods.TransferPinduoduoRecommendGoods(v))
	}
	res.List = goodsList
	res.SearchID = resp.GoodsBasicDetailResponse.SearchID
	res.Total = resp.GoodsBasicDetailResponse.Total
	res.ListID = resp.GoodsBasicDetailResponse.ListID
	return res, nil
}

func generateCustomParams(uid string) string {
	return fmt.Sprintf(`{"uid":%s}`, uid)
}
