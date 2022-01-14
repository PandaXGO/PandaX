package sms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/pkg/errors"
)

type AliConfig struct {
	accessKeyId  string
	accessSecret string
	RegionId     string
}

type AliSms struct {
	config AliConfig
}

func NewAliSms(AliConfig AliConfig) *AliSms {
	return &AliSms{
		config: AliConfig,
	}
}

func (a *AliSms) Send(PhoneNumbers, SignName, TemplateCode, TemplateParam string) error {
	//客户端
	client, err := dysmsapi.NewClientWithAccessKey(a.config.RegionId, a.config.accessKeyId, a.config.accessSecret)
	if err != nil {
		return errors.Wrapf(err, "ali sms client init fail")
	}
	//参数处理
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = PhoneNumbers
	request.SignName = SignName
	request.TemplateCode = TemplateCode
	//json格式
	request.TemplateParam = TemplateParam
	//发送
	_, err = client.SendSms(request)
	if err != nil {
		return errors.Wrapf(err, "ali sms send  fail")
	}
	return nil
}
