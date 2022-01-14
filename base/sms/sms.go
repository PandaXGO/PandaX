package sms

type Sms interface {
	Send(PhoneNumbers, SignName, TemplateCode, TemplateParam string) error
}

func NewDefaultSms(use string) Sms {
	switch use {
	case "AliYun":
		return NewAliSms(AliConfig{})
	default:
		panic("sms driver err")
	}
}
