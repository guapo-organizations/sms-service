package service

import (
	"context"
	"github.com/guapo-organizations/sms-service/proto/sms"
)

type SmsService struct {
}

//发送验证码
func (this *SmsService) SendEmailCode(ctx context.Context, in *sms.SendEmailCodeRequest) (*sms.SendResultResponse, error) {
	return nil, nil
}
