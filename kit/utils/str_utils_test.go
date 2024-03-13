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
	ip := GetRealAddressByIP("58.57.107.34")
	t.Log(ip)
}

func TestDeptPCIds(t *testing.T) {
	split := strings.Split(strings.Trim("/0/2", "/"), "/")
	log.Println("split", split)
	ids := DeptPCIds(split, 2, true)
	t.Log(ids)
}
