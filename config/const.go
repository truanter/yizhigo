package config

var DefaultFavorites = map[string]int64{
	"热门": 2054113014,
	"推荐": 2052602001,
}

var PddCatIDs = map[string]int64{
	"百货": 20100,
	"母婴": 20200,
	"食品": 20300,
	"女装": 20400,
	"电器": 20500,
	"鞋包": 20600,
	"内衣": 20700,
	"美妆": 20800,
	"男装": 20900,
	"水果": 21000,
	"家纺": 21100,
	"文具": 21200,
	"运动": 21300,
	//"虚拟": 21400,
	"汽车": 21500,
	"家装": 21600,
	"家具": 21700,
	"医药": 21800,
}

const (
	MaterialIDGuessYouLike       = int64(6708)
	MaterialIDGetFavorites       = int64(31519)
	MaterialIDGetFavoritesList   = int64(31539)
	MaterialIDGetSimilarGoods    = int64(13256)
	PinduoduoRegisterChannelType = int(10)
	PlatformNameTaobao           = "taobao"
	PlatformPinduoduo            = "pinduoduo"
)
