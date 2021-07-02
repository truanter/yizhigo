package config

var virtualProductKeyWords = &[]string{
	"充值",
	"券",
	"话费",
	"课程",
	"电子版",
	"账号",
	"帐号",
	"网盘",
	"链接",
}

var blockSearchPlatform = &[]string{
	"ios",
}

type bigBrother struct {
	IsHere *bool `json:"is_big_brother_here"`
}

var isHere = false
var bb = &bigBrother{&isHere}

var buttonText = &map[string]string{
	"detail_has_coupon":    "复制标题",
	"detail_no_coupon":     "复制标题",
	"detail_after_confirm": "标题已复制，可在微信全网搜索查看哦",
	"detail_modal_title":   "",
	"detail_modal_confirm": "复制信息",
}

var runtimeConfig = &map[string]interface{}{
	"block_search_platform": blockSearchPlatform,
	"key_words":             virtualProductKeyWords,
	"button_text":           buttonText,
	"is_big_brother_here":   bb.IsHere,
}

func SetBlockPlatform(newPlatform []string) {
	*blockSearchPlatform = newPlatform
}

func AddKeyWord(words []string) {
	for _, v := range words {
		*virtualProductKeyWords = append(*virtualProductKeyWords, v)
	}
}

func ClareKeyWord() {
	virtualProductKeyWords = &[]string{}
}

func MyBigBrother(isHere bool) {
	*bb.IsHere = isHere
}

func ModifyButtonText(k, v string) {
	d := *buttonText
	d[k] = v
}

func GetRuntimeConfig() map[string]interface{} {
	return *runtimeConfig
}

func GetVirtualProductKeyWords() []string {
	return *virtualProductKeyWords
}
