package ys

import "testing"

func TestYs_GetDeviceList(t *testing.T) {
	ys := &Ys{
		AppKey:    "",
		Secret:    "",
		IsRAM:     0,
		AccountID: "",
	}
	devices, total, err := ys.GetDeviceList(1, 10)
	if err != nil {
		t.Error(err)
	}
	t.Log(devices)
	t.Log(total)
}
