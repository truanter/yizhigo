package goods

import "github.com/truanter/pinduoduo-go/pkg/models"

var goodsMeta = map[string]interface{}{
	"category_id":         0,
	"coupon_share_url":    "",
	"coupon_amount":       0,
	"shop_title":          "",
	"short_title":         "",
	"title":               "",
	"small_images":        []string{},
	"pict_url":            "",
	"zk_final_price":      "0",
	"volume":              0,
	"item_id":             0,
	"coupon_remain_count": 0,
	"click_url":           "",
	"url":                 "",
}

func Parse(raw map[string]interface{}) map[string]interface{} {
	goods := make(map[string]interface{})
	for k, dv := range goodsMeta {
		if v, ok := raw[k]; ok {
			goods[k] = v
		} else {
			goods[k] = dv
		}
	}
	return goods
}

func getMinPrice(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}

type YiZhiGood struct {
	Name              string   `json:"name"`
	Description       string   `json:"description"`
	ShopID            int64    `json:"shop_id,omitempty"`
	ShopName          string   `json:"shop_name"`
	ShopLogo          string   `json:"shop_logo"`
	Tag               string   `json:"tag"`
	TagID             int64    `json:"tag_id,omitempty"`
	ImageUrl          string   `json:"image_url"`
	ImageThumbnailUrl string   `json:"image_thumbnail_url,omitempty"`
	GalleryUrls       []string `json:"gallery_urls,omitempty"`
	Coupon            float32  `json:"coupon"`
	CouponTotal       int      `json:"coupon_total,omitempty"`
	CouponRemain      int      `json:"coupon_remain,omitempty"`
	Price             float32  `json:"price"`
	Tags              []string `json:"tags,omitempty"`
	ServiceTags       []int    `json:"service_tags,omitempty"`
	SearchID          string   `json:"search_id,omitempty"`
	CatIDs            []int64  `json:"cat_ids,omitempty"`
	BrandName         string   `json:"brand_name"`
	GoodsSign         string   `json:"goods_sign,omitempty"`
	GoodsID           int64    `json:"goods_id,omitempty"`
	Volume            string   `json:"volume,omitempty"`
	VideoUrls         []string `json:"video_urls,omitempty"`
	HasCoupon         bool     `json:"has_coupon,omitempty"`
}

type PinduoduoGoodsList struct {
	List     []YiZhiGood `json:"list"`
	ListID   string      `json:"list_id"`
	SearchID string      `json:"search_id"`
	Total    int         `json:"total"`
	IsBlock  bool        `json:"is_block,omitempty"`
}

func TransferPinduoduo(result models.DDKGoodsSearchInfo) YiZhiGood {
	return YiZhiGood{
		Name:              result.GoodsName,
		Description:       result.GoodsDesc,
		ShopName:          result.MallName,
		Tag:               result.OptName,
		TagID:             result.OptID,
		ImageUrl:          result.GoodsImageUrl,
		ImageThumbnailUrl: result.GoodsThumbnailUrl,
		Coupon:            float32(result.CouponDiscount) / 100,
		CouponTotal:       result.CouponTotalQuantity,
		CouponRemain:      result.CouponRemainQuantity,
		Price:             float32(getMinPrice(result.MinGroupPrice, result.MinNormalPrice)) / 100,
		ServiceTags:       result.ServiceTags,
		SearchID:          result.SearchID,
		CatIDs:            result.CatIds,
		GoodsSign:         result.GoodsSign,
		Volume:            result.SalesTip,
		HasCoupon:         result.HasCoupon,
		BrandName:         result.BrandName,
	}
}

func TransferPinduoduoDetail(result models.GoodsDetailInfo) YiZhiGood {
	return YiZhiGood{
		Name:              result.GoodsName,
		Description:       result.GoodsDesc,
		ShopName:          result.MallName,
		Tag:               result.OptName,
		TagID:             result.OptID,
		ImageUrl:          result.GoodsImageUrl,
		ImageThumbnailUrl: result.GoodsThumbnailUrl,
		Coupon:            float32(result.CouponDiscount) / 100,
		CouponTotal:       result.CouponTotalQuantity,
		CouponRemain:      result.CouponRemainQuantity,
		Price:             float32(getMinPrice(result.MinGroupPrice, result.MinNormalPrice)) / 100,
		ServiceTags:       result.ServiceTags,
		CatIDs:            result.CatIds,
		GoodsSign:         result.GoodsSign,
		Volume:            result.SalesTip,
		HasCoupon:         result.HasCoupon,
		BrandName:         result.BrandName,
		ShopLogo:          result.MallImgUrl,
		Tags:              result.UnifiedTags,
		VideoUrls:         result.VideoUrls,
		ShopID:            result.MallID,
		GalleryUrls:       result.GoodsGalleryUrls,
	}
}

func TransferPinduoduoRecommendGoods(result models.GoodsRecommendInfo) YiZhiGood {
	return YiZhiGood{
		Name:              result.GoodsName,
		Description:       result.GoodsDesc,
		ShopName:          result.MallName,
		Tag:               result.OptName,
		TagID:             result.OptID,
		ImageUrl:          result.GoodsImageURL,
		ImageThumbnailUrl: result.GoodsThumbnailUrl,
		Coupon:            float32(result.CouponDiscount / 100),
		CouponTotal:       int(result.CouponTotalQuantity),
		CouponRemain:      int(result.CouponRemainQuantity),
		Price:             float32(getMinPrice(result.MinGroupPrice, result.MinNormalPrice)) / 100,
		CatIDs:            result.CatIds,
		GoodsSign:         result.GoodsSign,
		Volume:            result.SalesTip,
		HasCoupon:         result.HasCoupon,
		BrandName:         result.BrandName,
		Tags:              result.UnifiedTags,
		ShopID:            result.MallID,
	}
}
