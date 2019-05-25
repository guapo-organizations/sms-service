package service

import (
	"context"
	sms_service "github.com/guapo-organizations/sms-service/lib/sms"
	"github.com/guapo-organizations/sms-service/proto/sms"
	"log"
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
	log.Println("我操")
	result, err := sms_service.ValidateCode(in.Code, in.Key, in.PublishMode, in.PublishType)
	log.Println("asdfasd")
	log.Println(result)
	if err != nil {
		return nil, err
	}
	log.Println("asdfasd")
	response := new(sms.ValidateCodeResponse)
	response.Result = result
	log.Println(response)
	return response, nil
}
