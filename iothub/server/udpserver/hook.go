package updserver

import (
	"context"
	"encoding/hex"
	"fmt"
	"net"
	udpclient "pandax/iothub/client/updclient"
	"pandax/iothub/hook_message_work"
	"pandax/iothub/netbase"
	"pandax/pkg/global"
	"pandax/pkg/global/model"
	"pandax/pkg/rule_engine/message"
	"time"
)

type HookUdpService struct {
	HookService *hook_message_work.HookService
	conn        *net.UDPConn
}

func InitUdpHook(addr string, hs *hook_message_work.HookService) {
	server := NewUdpServer(addr)
	err := server.Start(context.TODO())
	if err != nil {
		global.Log.Error("IOTHUB UDP服务启动错误", err)
		return
	} else {
		global.Log.Infof("UDP IOTHUB HOOK Start SUCCESS, Server listen: %s", addr)
	}
	buffer := make([]byte, 1024)
	authMap := make(map[string]bool)
	etoken := &model.DeviceAuth{}
	hhs := &HookUdpService{
		HookService: hs,
		conn:        server.listener,
	}
	for {
		n, client, err := server.listener.ReadFromUDP(buffer)
		if err != nil {
			global.Log.Error("Error accepting connection:", err)
			_ = server.listener.Close()
			//设置断开连接
			if isAuth, ok := authMap[client.AddrPort().String()]; ok && isAuth {
				data := netbase.CreateConnectionInfo(message.DisConnectMes, "udp", client.IP.String(), client.AddrPort().String(), etoken)
				hhs.HookService.MessageCh <- data
			}
			delete(udpclient.UdpClient, etoken.DeviceId)
			delete(authMap, client.AddrPort().String())
			continue
		}
		if isAuth, ok := authMap[client.AddrPort().String()]; ok && isAuth {
			data := &netbase.DeviceEventInfo{
				DeviceId:   etoken.DeviceId,
				DeviceAuth: etoken,
			}

			hexData := hex.EncodeToString(buffer[:n])
			global.Log.Infof("UDP协议 设备%s, 接受消息%s", etoken.DeviceId, hexData)
			ts := time.Now().Format("2006-01-02 15:04:05.000")
			data.Type = message.RowMes
			data.Datas = fmt.Sprintf(`{"ts": "%s","rowdata": "%s"}`, ts, hexData)

			// etoken中添加设备标识
			hhs.HookService.MessageCh <- data
		} else {
			token := string(buffer[:n])
			etoken.GetDeviceToken(token)
			auth := netbase.Auth(token)
			// 认证成功，创建连接记录
			if auth {
				global.Log.Infof("UDP协议 设备%s,认证成功", etoken.DeviceId)
				data := netbase.CreateConnectionInfo(message.ConnectMes, "udp", client.IP.String(), client.AddrPort().String(), etoken)
				hhs.HookService.MessageCh <- data
				authMap[client.AddrPort().String()] = true
				udpclient.UdpClient[etoken.DeviceId] = &udpclient.UdpClientT{
					Conn: server.listener,
					Addr: client,
				}
				hhs.Send(client, "success")
			} else {
				hhs.Send(client, "fail")
			}
		}
	}
}

func (hhs *HookUdpService) Send(addr *net.UDPAddr, message string) error {
	return hhs.SendBytes(addr, []byte(message))
}

func (hhs *HookUdpService) SendHex(addr *net.UDPAddr, msg string) error {
	b, err := hex.DecodeString(msg)
	if err != nil {
		return err
	}
	return hhs.SendBytes(addr, b)
}

func (hhs *HookUdpService) SendBytes(addr *net.UDPAddr, msg []byte) error {
	_, err := hhs.conn.WriteToUDP(msg, addr)
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
