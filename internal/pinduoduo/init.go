package pinduoduo

import (
	"github.com/truanter/pinduoduo-go/pkg/sdk"
	"github.com/truanter/yizhigo/config"
	"github.com/truanter/yizhigo/model/goods"
)

var ddk *sdk.Client
var pid string
var Categories []goods.PddCats

func init() {
	ddk = sdk.NewSdkClient(&sdk.SdkOption{
		IsDebug: false,
	})
	clientID, secret := config.GetPddAuthInfo()
	ddk.SetClientIdAndSecret(clientID, secret)
	ddk.UseDebug()
	pid = config.GetPddPID()
	for k, v := range config.PddCatIDs {
		Categories = append(Categories, goods.PddCats{
			CatName: k,
			CatID:   v,
		})
	}
}
