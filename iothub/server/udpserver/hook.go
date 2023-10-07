package updserver

import (
	"context"
	"encoding/hex"
	"log"
	"net"
	"pandax/iothub/hook_message_work"
	"pandax/iothub/netbase"
	"pandax/pkg/global"
	"pandax/pkg/rule_engine/message"
)

type HookUdpService struct {
	HookService *hook_message_work.HookService
	conn        *net.UDPConn
	addr        *net.UDPAddr
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
	for {
		n, client, err := server.listener.ReadFromUDP(buffer)
		if err != nil {
			global.Log.Error("Error accepting connection:", err)
			continue
		}
		hhs := &HookUdpService{
			HookService: hs,
			conn:        server.listener,
			addr:        client,
		}
		go hhs.hook(buffer[:n])

	}
}

func (hhs *HookUdpService) hook(data []byte) {
	log.Println("udp msg", string(data))
	hhs.Send("success")
}

func (hhs *HookUdpService) Send(message string) error {
	return hhs.SendBytes([]byte(message))
}

func (hhs *HookUdpService) SendHex(msg string) error {
	b, err := hex.DecodeString(msg)
	if err != nil {
		return err
	}
	return hhs.SendBytes(b)
}

func (hhs *HookUdpService) SendBytes(msg []byte) error {
	_, err := hhs.conn.WriteToUDP(msg, hhs.addr)
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
