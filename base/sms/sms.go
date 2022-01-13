package sms

type Sms interface {
	GetBusiness(key string) Sms
	Send(phone []string, templateParam interface{}) error
}

func NewDefaultSms(use string) Sms {
	switch use {
	case "AliYun":
		return NewAliYun()
	default:
		panic("sms driver err")
	}
}
