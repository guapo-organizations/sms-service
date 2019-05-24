package service

import (
	"context"
	sms_service "github.com/guapo-organizations/sms-service/lib/sms"
	"github.com/guapo-organizations/sms-service/proto/sms"
)

type SmsService struct {
}

//发送验证码
func (this *SmsService) SendEmailCode(ctx context.Context, in *sms.SendEmailCodeRequest) (*sms.SendResultResponse, error) {
	_, err := sms_service.SendEmailCode(in.Email)

	if err != nil {
		return nil, err
	}

	response := new(sms.SendResultResponse)

	return response, nil
}

//验证码校验
func (this *SmsService) ValidateCode(ctx context.Context, in *sms.ValidateCodeRequest) (*sms.ValidateCodeResponse, error) {

	result, err := sms_service.ValidateCode(in.Key, in.PublishMode, in.PublishType)

	if err != nil {
		return nil, err
	}

	response := new(sms.ValidateCodeResponse)
	response.Result = result

	return response, nil
}
