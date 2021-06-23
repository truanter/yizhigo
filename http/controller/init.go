package controller

import (
	"github.com/truanter/top-go/pkg/config"
	"github.com/truanter/top-go/pkg/sdk"
	conf "github.com/truanter/yizhigo/config"
)

var top *sdk.SDKClient

func init() {
	top = sdk.NewSdkClient(&config.SDKConfig{})
	appKey, secret := conf.GetTbkAuthInfo()
	top.SetKeyAndSecret(appKey, secret)
	top.UseHttps()
	top.SetDebug(!conf.IsProdEnv())
}
