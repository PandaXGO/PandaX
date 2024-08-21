package tcpserver

import (
	"context"
	"encoding/hex"
	"fmt"
	"net"
	"pandax/iothub/client/tcpclient"
	"pandax/iothub/hook_message_work"
	"pandax/iothub/netbase"
	"pandax/pkg/global"
	"pandax/pkg/global/model"
	"pandax/pkg/rule_engine/message"
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
		global.Log.Error("IOTHUB TCP服务启动错误", err)
		return
	}
	global.Log.Infof("TCP IOTHUB HOOK Start SUCCESS, Server listen: %s", addr)
	go acceptConnections(server.listener, hs)
}

func acceptConnections(listener *net.TCPListener, hs *hook_message_work.HookService) {
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			global.Log.Error("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn, hs)
	}
}

func handleConnection(conn *net.TCPConn, hs *hook_message_work.HookService) {
	isAuth := false
	etoken := &model.DeviceAuth{}
	defer func() {
		_ = conn.Close()
		if isAuth {
			data := netbase.CreateEvent(message.DisConnectMes, "info", fmt.Sprintf("设备%s断开连接", etoken.Name), etoken)
			go hs.Queue.Queue(data)
		}
		tcpclient.TcpClient.Delete(etoken.DeviceId)
	}()

	for {
		buf := make([]byte, 128)
		n, err := conn.Read(buf)
		if err != nil {
			isAuth = false
			return
		}

		if !isAuth {
			token := string(buf[:n])
			etoken.GetDeviceToken(token)
			auth := netbase.Auth(token)
			if auth {
				global.Log.Infof("TCP协议 设备%s,认证成功", etoken.DeviceId)
				data := netbase.CreateEvent(message.ConnectMes, "info", fmt.Sprintf("设备%s通过TCP协议连接", etoken.Name), etoken)
				go hs.Queue.Queue(data)
				isAuth = true
				tcpclient.TcpClient.Store(etoken.DeviceId, conn)
				sendResponse(conn, "success")
			} else {
				sendResponse(conn, "fail")
			}
		} else {
			hexData := hex.EncodeToString(buf[:n])
			global.Log.Infof("TCP协议 设备%s, 接受消息%s", etoken.DeviceId, hexData)
			ts := time.Now().Format("2006-01-02 15:04:05.000")
			data := &netbase.DeviceEventInfo{
				DeviceId:   etoken.DeviceId,
				DeviceAuth: etoken,
				Type:       message.RowMes,
				Datas:      fmt.Sprintf(`{"ts": "%s","rowdata": "%s"}`, ts, hexData),
			}
			go hs.Queue.Queue(data)
		}
	}
}

func sendResponse(conn *net.TCPConn, message string) {
	_, err := conn.Write([]byte(message))
	if err != nil {
		conn.Close()
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

func isHex(str string) bool {
	_, err := hex.DecodeString(str)
	return err == nil
}
