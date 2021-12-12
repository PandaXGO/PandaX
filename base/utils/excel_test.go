package utils

import (
	"testing"
)

type User struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

func TestExportExcel(t *testing.T) {
	us := make([]User, 0)
	us = append(us, User{Name: "张三", Age: 12})
	us = append(us, User{Name: "李四", Age: 23})
	name := GetFileName("./", "字典")
	t.Log(name)
	InterfaceToExcel(us, name)
}
