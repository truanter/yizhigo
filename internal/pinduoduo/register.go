package pinduoduo

import (
	"github.com/truanter/pinduoduo-go/pkg/models"
	"github.com/truanter/yizhigo/config"
	"github.com/truanter/yizhigo/pkg/common"
)

func Register(uid string, generateQQAPP, generateWeAPP, generateSchemaURL, generateShortURL bool) (models.DDKRpPromotionUrlGenerateResult, error) {
	customParameters := generateCustomParams(uid)
	channelType := config.PinduoduoRegisterChannelType
	resp, err := ddk.DDKRpPromUrlGenerate().Do(models.DDKRpPromUrlGenerateRequest{
		ChannelType:       &channelType,
		CustomParameters:  &customParameters,
		GenerateQQApp:     &generateQQAPP,
		GenerateWeApp:     &generateWeAPP,
		GenerateSchemaUrl: &generateSchemaURL,
		GenerateShortUrl:  &generateShortURL,
		PIDList:           &[]string{pid},
	})
	if err == nil {
		if resp.ErrorResponse.ErrorCode != 0 {
			err = common.NewPddRequestError(resp.ErrorResponse.ErrorCode, resp.ErrorResponse.SubCode, resp.ErrorResponse.ErrorMsg, resp.ErrorResponse.SubMsg, resp.ErrorResponse.RequestID)
		}
	}
	return resp, err
}

func CheckBind(uid string) (resp models.DDKMemberAuthorityQueryResult, err error) {
	customParameters := generateCustomParams(uid)
	resp, err = ddk.DDKMemberAuthorityQuery().Do(models.DDKMemberAuthorityQueryRequest{
		Pid:              &pid,
		CustomParameters: &customParameters,
	})
	if err == nil {
		if resp.ErrorResponse.ErrorCode != 0 {
			err = common.NewPddRequestError(resp.ErrorResponse.ErrorCode, resp.ErrorResponse.SubCode, resp.ErrorResponse.ErrorMsg, resp.ErrorResponse.SubMsg, resp.ErrorResponse.RequestID)
		}
	}
	return
}
