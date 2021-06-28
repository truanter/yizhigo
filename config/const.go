package config

var DefaultFavorites = map[string]int64{
	"热门": 2054113014,
	"推荐": 2052602001,
}

const (
	MaterialIDGuessYouLike     = int64(6708)
	MaterialIDGetFavorites     = int64(31519)
	MaterialIDGetFavoritesList = int64(31539)
	MaterialIDGetSimilarGoods  = int64(13256)
)
