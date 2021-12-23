package utils

import "testing"

func TestIdsStrToIdsIntGroup(t *testing.T) {
	group := IdsStrToIdsIntGroup("aaa")
	t.Log(len(group))
}

func TestGetRealAddressByIP(t *testing.T) {
	ip := GetRealAddressByIP("10.42.0.1")
	t.Log(ip)
}
