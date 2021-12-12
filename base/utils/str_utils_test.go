package utils

import "testing"

func TestIdsStrToIdsIntGroup(t *testing.T) {
	group := IdsStrToIdsIntGroup("aaa")
	t.Log(len(group))
}
