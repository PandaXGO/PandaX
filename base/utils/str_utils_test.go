package utils

import (
	"log"
	"strings"
	"testing"
)

func TestIdsStrToIdsIntGroup(t *testing.T) {
	group := IdsStrToIdsIntGroup("aaa")
	t.Log(len(group))
}

func TestGetRealAddressByIP(t *testing.T) {
	ip := GetRealAddressByIP("10.42.0.1")
	t.Log(ip)
}

func TestDeptPCIds(t *testing.T) {
	ss := strings.Trim("/0/2/6/4", "/")

	split := strings.Split(ss, "/")
	log.Println("split", split)
	ids := DeptPCIds(split, 4, false)
	t.Log(ids)
}
