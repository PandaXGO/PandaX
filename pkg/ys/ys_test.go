package ys

import "testing"

func TestYs_GetDeviceList(t *testing.T) {
	ys := &Ys{
		AppKey:    "",
		Secret:    "",
		IsRAM:     0,
		AccountID: "",
	}
	devices, total, err := ys.GetDeviceList(0, 10)
	if err != nil {
		t.Error(err)
	}
	t.Log(devices)
	t.Log(total)
}

func TestYs_GetDeviceChannelList(t *testing.T) {
	ys := &Ys{
		AppKey:    "",
		Secret:    "",
		IsRAM:     0,
		AccountID: "",
	}
	chans, err := ys.GetDeviceChannelList("BA1996108")
	if err != nil {
		t.Error(err)
	}
	t.Log(chans)
}

func TestYs_GetDeviceDeviceLiveAddress(t *testing.T) {
	ys := &Ys{
		AppKey:    "2cc6a5edcee046c1b8bc9a857d67a287",
		Secret:    "9eb8f595dc02859c91a5d7d0593f8a07",
		IsRAM:     0,
		AccountID: "",
	}
	live, err := ys.GetDeviceLiveAddress("BA1996108", 1)
	if err != nil {
		t.Error(err)
	}
	t.Log(live)
}
