package config

var virtualProductKeyWords = []string{
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

var runtimeConfig = map[string]interface{}{
	"block_search_platform": []string{"ios"},
	"key_words":             virtualProductKeyWords,
}

func GetRuntimeConfig() map[string]interface{} {
	return runtimeConfig
}

func GetVirtualProductKeyWords() []string {
	return virtualProductKeyWords
}
