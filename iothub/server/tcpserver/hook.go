package tcpserver

import (
	"context"
	"encoding/hex"
	"github.com/emicklei/go-restful/v3"
	"log"
	"net"
	"pandax/iothub/hook_message_work"
	"pandax/iothub/netbase"
	"pandax/pkg/global"
	"pandax/pkg/rule_engine/message"
	"strings"
	"time"
)

type HookTcpService struct {
	HookService *hook_message_work.HookService
	keepAlive   int64
	conn        *net.TCPConn
}

func InitTcpHook(addr string, hs *hook_message_work.HookService) {
	hhs := &HookTcpService{
		HookService: hs,
		keepAlive:   20,
	}
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
			conn.SetReadDeadline(time.Now().Add(20 * time.Second))
			hhs.conn = conn
			go hhs.hook()

		}
	}()
}

// 获取token进行认证
func basicAuthenticate(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	path := req.Request.URL.Path
	log.Println(path)
	split := strings.Split(path, "/")
	log.Println(split)
	chain.ProcessFilter(req, resp)
}

func (hhs *HookTcpService) hook() {
	isAuth := false
	for {
		buf := make([]byte, 128)
		n := 0
		n, err := hhs.conn.Read(buf)
		if err != nil {
			// 断开连接 掉线
			log.Println("断开连接")
			_ = hhs.conn.Close()
			isAuth = false
			return
		}
		if !isAuth {
			token := string(buf[:n])
			log.Println(token)
			isAuth = true
		} else {
			hexData := hex.EncodeToString(buf[:n])
			log.Println(hexData)
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
