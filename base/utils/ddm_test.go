package utils

import (
	"testing"
)

/**
 * @Description 添加qq群467890197 交流学习
 * @Author 熊猫
 * @Date 2022/1/17 10:01
 **/

func TestDdmKey(t *testing.T) {
	aa := "fsdfsf535343sdfsdf3"
	bb := "23423423@qq.com"

	key := DdmKey(aa)
	t.Log(key)
	ddmKey := IsDdmKey(key)
	t.Log(ddmKey)

	mail := DdmMail(bb)
	t.Log(mail)
	ismail := ISDdmMail(mail)
	t.Log(ismail)
}
