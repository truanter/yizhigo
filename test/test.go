package main

import (
	"fmt"
	"github.com/truanter/top-go/pkg/config"
	"github.com/truanter/top-go/pkg/sdk"
	conf "github.com/truanter/yizhigo/config"
)

func main() {
	top := sdk.NewSdkClient(&config.SDKConfig{})
	appKey, secret := conf.GetTbkAuthInfo()
	top.SetKeyAndSecret(appKey, secret)
	top.UseHttps()
	top.SetDebug(!conf.IsProdEnv())
	api := top.GeneralAPI()
	api.SetMethods("taobao.product.get")
	resp, _ := api.Do(map[string]interface{}{
		"fields":     "product_id",
		"product_id": 606208562276,
	})
	fmt.Println(fmt.Sprintf("%v", resp))

}
