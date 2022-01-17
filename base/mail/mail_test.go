package email

import "testing"

/**
 * @Description 添加qq群467890197 交流学习
 * @Author 熊猫
 * @Date 2022/1/17 10:35
 **/

func TestMail_Email(t *testing.T) {
	ma := Mail{
		Host:     "smtp.163.com",
		Port:     25,
		From:     "18610165312@163.com",
		Nickname: "panda",
		Secret:   "DCXZCAGTCMSEGPZL",
		IsSSL:    false,
	}

	email := ma.Email("18353366911@163.com", "ceshi", "ceshibody")
	t.Log(email)
}
