package httpserver

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"
	"time"

	"pandax/iothub/hook_message_work"
	"pandax/iothub/netbase"
	"pandax/pkg/global"
	"pandax/pkg/global/model"
	"pandax/pkg/rule_engine/message"

	"github.com/emicklei/go-restful/v3"
)

type HookHttpService struct {
	HookService *hook_message_work.HookService
}

type ConnectionManager struct {
	activeConnections sync.Map
}

type PathType string

const (
	AEPTYPE    PathType = "AEP"
	ONENETTYPE PathType = "ONENET"
)

const (
	AuthError         = "认证错误"
	UnknownPath       = "路径未识别token，或上报类型"
	ParseParamError   = "解析上报参数失败"
	ParseTelemetryErr = "解析遥测失败"
	ParseAttrErr      = "解析属性失败"
	UnknownPathType   = "路径上报类型错误"
)

var connectionManager ConnectionManager

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

	connectionManager = ConnectionManager{
		activeConnections: sync.Map{},
	}

	server.srv.ConnState = func(conn net.Conn, state http.ConnState) {
		// 断开连接
		switch state {
		case http.StateHijacked, http.StateClosed:
			etoken, _ := connectionManager.activeConnections.Load(conn.RemoteAddr().String())
			if etoken != nil {
				connectionManager.RemoveConnection(conn.RemoteAddr().String(), etoken.(*model.DeviceAuth), service.HookService)
			}
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
		resp.Write([]byte(AuthError))
		return
	}
	chain.ProcessFilter(req, resp)
}

func (hhs *HookHttpService) hook(req *restful.Request, resp *restful.Response) {
	token := req.PathParameter("token")
	pathType := req.PathParameter("pathType")
	platformType := req.QueryParameter("platformType")
	if token == "" || pathType == "" {
		resp.Write([]byte(UnknownPath))
		return
	}
	var upData map[string]interface{}
	err := req.ReadEntity(&upData)
	if err != nil {
		resp.Write([]byte(ParseParamError))
		return
	}
	etoken := &model.DeviceAuth{}
	etoken.GetDeviceToken(token)
	_, ok := connectionManager.activeConnections.Load(req.Request.RemoteAddr)
	// 是否需要添加设备上线通知
	if !ok {
		connectionManager.AddConnection(req.Request.RemoteAddr, etoken, hhs.HookService)
	}
	var data *netbase.DeviceEventInfo
	if platformType == string(AEPTYPE) {

	}
	marshal, _ := json.Marshal(upData)
	data = &netbase.DeviceEventInfo{
		Datas:      string(marshal),
		DeviceAuth: etoken,
		DeviceId:   etoken.DeviceId,
	}
	switch pathType {
	case "row":
		ts := time.Now().Format("2006-01-02 15:04:05.000")
		data.Type = message.RowMes
		data.Datas = fmt.Sprintf(`{"ts": "%s","rowdata": "%s"}`, ts, data.Datas)
	case "telemetry":
		telemetryData := netbase.UpdateDeviceTelemetryData(data.Datas)
		if telemetryData == nil {
			resp.Write([]byte(ParseTelemetryErr))
			return
		}
		bytes, _ := json.Marshal(telemetryData)
		data.Type = message.TelemetryMes
		data.Datas = string(bytes)
	case "attributes":
		attributesData := netbase.UpdateDeviceAttributesData(data.Datas)
		if attributesData == nil {
			resp.Write([]byte(ParseAttrErr))
			return
		}
		bytes, _ := json.Marshal(attributesData)
		data.Datas = string(bytes)
		data.Type = message.AttributesMes
	default:
		resp.Write([]byte(UnknownPathType))
		return
	}
	go hhs.HookService.Queue.Queue(data)
	io.WriteString(resp, "ok")
}

func (cm *ConnectionManager) AddConnection(addr string, etoken *model.DeviceAuth, service *hook_message_work.HookService) {
	cm.activeConnections.Store(addr, etoken)
	data := netbase.CreateConnectionInfo(message.ConnectMes, "http", addr, addr, etoken)
	go service.Queue.Queue(data)
}

func (cm *ConnectionManager) RemoveConnection(addr string, etoken *model.DeviceAuth, service *hook_message_work.HookService) {
	data := netbase.CreateConnectionInfo(message.DisConnectMes, "http", addr, addr, etoken)
	cm.activeConnections.Delete(addr)
	go service.Queue.Queue(data)
}
