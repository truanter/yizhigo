package goods

import (
	"github.com/truanter/top-go/pkg/config"
	"github.com/truanter/top-go/pkg/sdk"
	conf "github.com/truanter/yizhigo/config"
	"github.com/truanter/yizhigo/model/goods"
)

var top *sdk.SDKClient
var myFavorites []goods.Favorites

func init() {
	// init top client
	top = sdk.NewSdkClient(&config.SDKConfig{})
	appKey, secret := conf.GetTbkAuthInfo()
	top.SetKeyAndSecret(appKey, secret)
	top.UseHttps()
	top.SetDebug(!conf.IsProdEnv())

	// init my favorites
	for k, v := range conf.DefaultFavorites {
		myFavorites = append(myFavorites, goods.Favorites{
			FavoritesID:    v,
			FavoritesTitle: k,
		})
	}
}
