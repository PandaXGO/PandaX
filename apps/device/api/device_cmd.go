package api

// ==========================================================================
import (
	"github.com/PandaXGO/PandaKit/biz"
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	"github.com/kakuilan/kgo"
	"pandax/pkg/global"
	"pandax/pkg/mqtt"
	"strings"
	"time"

	"pandax/apps/device/entity"
	"pandax/apps/device/services"
)

type DeviceCmdLogApi struct {
	DeviceCmdLogApp services.DeviceCmdLogModel
}

// GetDeviceCmdLogList 告警列表数据
func (p *DeviceCmdLogApi) GetDeviceCmdLogList(rc *restfulx.ReqCtx) {
	data := entity.DeviceCmdLog{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	data.DeviceId = restfulx.QueryParam(rc, "deviceId")
	data.State = restfulx.QueryParam(rc, "state")
	data.Type = restfulx.QueryParam(rc, "type")

	list, total := p.DeviceCmdLogApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageNum),
		Data:     list,
	}
}

// InsertDeviceCmdLog 添加DeviceCmdLog
func (p *DeviceCmdLogApi) InsertDeviceCmdLog(rc *restfulx.ReqCtx) {
	var data entity.DeviceCmdLog
	restfulx.BindJsonAndValid(rc, &data)
	data.Id = kgo.KStr.Uniqid("cmd_")
	data.State = "2"
	data.RequestTime = time.Now().Format("2006-01-02 15:04:05")
	err := p.DeviceCmdLogApp.Insert(data)
	biz.ErrIsNil(err, "添加指令记录失败")
	// 下发指令
	var rpc = &mqtt.RpcRequest{Client: global.MqttClient, Mode: "single"}
	rpc.GetRequestId()
	_, err = rpc.RequestCmd(mqtt.RpcPayload{Method: data.CmdName, Params: data.CmdContent})
	if err != nil {
		global.Log.Error("指令下发失败")
	}
}

// DeleteDeviceCmdLog 删除告警
func (p *DeviceCmdLogApi) DeleteDeviceCmdLog(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	ids := strings.Split(id, ",")
	p.DeviceCmdLogApp.Delete(ids)
}
