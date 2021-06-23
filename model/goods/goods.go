package goods

var goodsMeta = map[string]interface{}{
	"category_id": 0,
	"coupon_share_url": "",
	"coupon_amount": 0,
	"shop_title": "",
	"short_title": "",
	"title": "",
	"small_images": []string{},
	"pict_url": "",
	"zk_final_price": "0",
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