package tcpserver

import (
	"context"
	"encoding/hex"
	"fmt"
	"net"
	"pandax/iothub/hook_message_work"
	"pandax/iothub/netbase"
	"pandax/pkg/global"
	"pandax/pkg/rule_engine/message"
	"pandax/pkg/tool"
	"time"
)

type HookTcpService struct {
	HookService *hook_message_work.HookService
	conn        *net.TCPConn
}

func InitTcpHook(addr string, hs *hook_message_work.HookService) {
	server := NewTcpServer(addr)
	err := server.Start(context.TODO())
	if err != nil {
		global.Log.Error("IOTHUB HTTP服务启动错误", err)
		return
	} else {
		global.Log.Infof("TCP IOTHUB HOOK Start SUCCESS, Server listen: %s", addr)
	}
	go func() {
		for {
			conn, err := server.listener.AcceptTCP()
			if err != nil {
				global.Log.Error("Error accepting connection:", err)
				continue
			}
			conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
			hhs := &HookTcpService{
				HookService: hs,
				conn:        conn,
			}
			go hhs.hook()

		}
	}()
}

func (hhs *HookTcpService) hook() {
	isAuth := false
	etoken := &tool.DeviceAuth{}
	for {
		buf := make([]byte, 128)
		n := 0
		n, err := hhs.conn.Read(buf)
		if err != nil {
			_ = hhs.conn.Close()
			//设置断开连接
			if isAuth {
				data := netbase.CreateConnectionInfo(message.DisConnectMes, "tcp", hhs.conn.RemoteAddr().String(), hhs.conn.RemoteAddr().String(), etoken)
				hhs.HookService.MessageCh <- data
			}
			isAuth = false
			return
		}
		if !isAuth {
			token := string(buf[:n])
			etoken.GetDeviceToken(token)
			auth := netbase.Auth(token)
			// 认证成功，创建连接记录
			if auth {
				data := netbase.CreateConnectionInfo(message.ConnectMes, "tcp", hhs.conn.RemoteAddr().String(), hhs.conn.RemoteAddr().String(), etoken)
				hhs.HookService.MessageCh <- data
				isAuth = true
				hhs.Send("success")
			} else {
				hhs.Send("fail")
			}
		} else {
			hexData := hex.EncodeToString(buf[:n])
			ts := time.Now().Format("2006-01-02 15:04:05.000")
			data := &netbase.DeviceEventInfo{
				DeviceId:   etoken.DeviceId,
				DeviceAuth: etoken,
				Type:       message.RowMes,
			}
			data.Datas = fmt.Sprintf(`{"ts": "%s","rowdata": "%s"}`, ts, hexData)
			// etoken中添加设备标识
			hhs.HookService.MessageCh <- data
		}
	}

}

func (hhs *HookTcpService) Send(message string) error {
	return hhs.SendBytes([]byte(message))
}

func (hhs *HookTcpService) SendHex(msg string) error {
	b, err := hex.DecodeString(msg)
	if err != nil {
		return err
	}
	return hhs.SendBytes(b)
}

func (hhs *HookTcpService) SendBytes(msg []byte) error {
	_, err := hhs.conn.Write(msg)
	if err != nil {
		hhs.conn.Close()
		data := &netbase.DeviceEventInfo{
			DeviceId: "",
			Datas:    "",
			Type:     message.ConnectMes,
		}
		hhs.HookService.MessageCh <- data
	}
	return err
}
