package udpclient

import (
	"encoding/hex"
	"net"
	"pandax/pkg/global"
	"sync"
)

type UdpClientT struct {
	Conn *net.UDPConn
	Addr *net.UDPAddr
}

var UdpClient sync.Map

func Send(deviceId, msg string) error {
	if conn, ok := UdpClient.Load(deviceId); ok {
		conn := conn.(*UdpClientT)
		global.Log.Infof("设备%s, 发送指令%s", deviceId, msg)
		_, err := conn.Conn.WriteToUDP([]byte(msg), conn.Addr)
		if err != nil {
			return err
		}
	} else {
		global.Log.Infof("设备%s TCP连接不存在, 发送指令失败", deviceId)
	}
	return nil
}

func SendHex(deviceId, msg string) error {
	if conn, ok := UdpClient.Load(deviceId); ok {
		conn := conn.(*UdpClientT)
		global.Log.Infof("设备%s, 发送指令%s", deviceId, msg)
		b, err := hex.DecodeString(msg)
		if err != nil {
			return err
		}
		_, err = conn.Conn.WriteToUDP(b, conn.Addr)
		if err != nil {
			return err
		}
	} else {
		global.Log.Infof("设备%s TCP连接不存在, 发送指令失败", deviceId)
	}
	return nil
}
