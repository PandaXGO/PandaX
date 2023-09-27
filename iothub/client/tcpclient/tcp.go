package tcpclient

import (
	"encoding/hex"
	"net"
	"pandax/pkg/global"
)

var TcpClient = make(map[string]*net.TCPConn)

func Send(deviceId, msg string) error {
	if conn, ok := TcpClient[deviceId]; ok {
		global.Log.Infof("设备%s, 发送指令%s", deviceId, msg)
		_, err := conn.Write([]byte(msg))
		if err != nil {
			return err
		}
	} else {
		global.Log.Infof("设备%s TCP连接不存在, 发送指令失败", deviceId)
	}
	return nil
}

func SendHex(deviceId, msg string) error {
	if conn, ok := TcpClient[deviceId]; ok {
		global.Log.Infof("设备%s, 发送指令%s", deviceId, msg)
		b, err := hex.DecodeString(msg)
		if err != nil {
			return err
		}
		_, err = conn.Write(b)
		if err != nil {
			return err
		}
	} else {
		global.Log.Infof("设备%s TCP连接不存在, 发送指令失败", deviceId)
	}
	return nil
}
