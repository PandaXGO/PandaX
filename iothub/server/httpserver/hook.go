package httpserver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"io"
	"net"
	"net/http"
	"pandax/iothub/hook_message_work"
	"pandax/iothub/netbase"
	"pandax/pkg/global"
	"pandax/pkg/rule_engine/message"
	"pandax/pkg/tool"
	"sync"
	"time"
)

type HookHttpService struct {
	HookService *hook_message_work.HookService
}

var (
	activeConnections sync.Map
)

func InitHttpHook(addr string, hs *hook_message_work.HookService) {
	server := NewHttpServer(addr)
	service := &HookHttpService{
		HookService: hs,
	}
	container := server.Container
	ws := new(restful.WebService)
	ws.Path("/api/v1").Produces(restful.MIME_JSON)
	ws.Route(ws.POST("/{token}/{pathType}").Filter(basicAuthenticate).To(service.hook))
	container.Add(ws)

	server.srv.ConnState = func(conn net.Conn, state http.ConnState) {
		// 断开连接
		switch state {
		case http.StateHijacked, http.StateClosed:
			etoken, _ := activeConnections.Load(conn.RemoteAddr())
			data := netbase.CreateConnectionInfo(message.DisConnectMes, "http", conn.RemoteAddr().String(), conn.RemoteAddr().String(), etoken.(*tool.DeviceAuth))
			service.HookService.MessageCh <- data
			activeConnections.Delete(conn.RemoteAddr())
		}
	}
	err := server.Start(context.TODO())
	if err != nil {
		global.Log.Error("IOTHUB HTTP服务启动错误", err)
	} else {
		global.Log.Infof("HTTP IOTHUB HOOK Start SUCCESS,Server listen: %s", addr)
	}
}

// 获取token进行认证
func basicAuthenticate(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	token := req.PathParameter("token")
	auth := netbase.Auth(token)
	if !auth {
		resp.Write([]byte("认证错误"))
		return
	}
	chain.ProcessFilter(req, resp)
}

func (hhs *HookHttpService) hook(req *restful.Request, resp *restful.Response) {
	token := req.PathParameter("token")
	pathType := req.PathParameter("pathType")
	if token == "" || pathType == "" {
		resp.Write([]byte("路径未识别token，或上报类型"))
		return
	}
	var upData map[string]interface{}
	err := req.ReadEntity(&upData)
	if err != nil {
		resp.Write([]byte("解析上报参数失败"))
		return
	}
	etoken := &tool.DeviceAuth{}
	etoken.GetDeviceToken(token)
	_, ok := activeConnections.Load(req.Request.RemoteAddr)
	// 是否需要添加设备上线通知
	if !ok {
		activeConnections.Store(req.Request.RemoteAddr, etoken)
		data := netbase.CreateConnectionInfo(message.ConnectMes, "http", req.Request.RemoteAddr, req.Request.RemoteAddr, etoken)
		hhs.HookService.MessageCh <- data
	}
	marshal, _ := json.Marshal(upData)
	data := &netbase.DeviceEventInfo{
		Datas:      string(marshal),
		DeviceAuth: etoken,
		DeviceId:   etoken.DeviceId,
	}
	switch pathType {
	case Row:
		ts := time.Now().Format("2006-01-02 15:04:05.000")
		data.Type = message.RowMes
		data.Datas = fmt.Sprintf(`{"ts": "%s","rowdata": "%s"}`, ts, data.Datas)
	case Telemetry:
		telemetryData := netbase.UpdateDeviceTelemetryData(data.Datas)
		if telemetryData == nil {
			resp.Write([]byte("解析遥测失败"))
			return
		}
		bytes, _ := json.Marshal(telemetryData)
		data.Type = message.TelemetryMes
		data.Datas = string(bytes)
	case Attributes:
		attributesData := netbase.UpdateDeviceAttributesData(data.Datas)
		if attributesData == nil {
			resp.Write([]byte("解析属性失败"))
			return
		}
		bytes, _ := json.Marshal(attributesData)
		data.Datas = string(bytes)
		data.Type = message.AttributesMes
	default:
		resp.Write([]byte("路径上报类型错误"))
		return
	}
	hhs.HookService.MessageCh <- data
	io.WriteString(resp, "ok")
}
