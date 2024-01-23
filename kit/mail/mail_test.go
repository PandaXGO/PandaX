package email

import "testing"


// 自行替换测试账户信息
func TestMail_Email(t *testing.T) {
	ma := Mail{
		Host:     "smtp.163.com",
		Port:     25,
		From:     "x@163.com",
		Nickname: "x",
		Secret:   "x",
		IsSSL:    false,
	}

	email := ma.Email("xx@163.com", "ceshi", "ceshibody")
	t.Log(email)
}
