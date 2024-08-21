package api

// ==========================================================================
import (
	"encoding/json"
	"pandax/kit/biz"
	"pandax/kit/model"
	"pandax/kit/restfulx"
	devicerpc "pandax/pkg/device_rpc"
	"strings"

	"pandax/apps/device/entity"
	"pandax/apps/device/services"
	"pandax/apps/device/util"
)

type DeviceCmdLogApi struct {
	DeviceCmdLogApp services.DeviceCmdLogModel
	DeviceApp       services.DeviceModel
}

// GetDeviceCmdLogList 告警列表数据
func (p *DeviceCmdLogApi) GetDeviceCmdLogList(rc *restfulx.ReqCtx) {
	data := entity.DeviceCmdLog{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	data.DeviceId = restfulx.QueryParam(rc, "deviceId")
	data.State = restfulx.QueryParam(rc, "state")
	data.Type = restfulx.QueryParam(rc, "type")

	list, total, err := p.DeviceCmdLogApp.FindListPage(pageNum, pageSize, data)
	biz.ErrIsNil(err, "查询告警列表数据失败")
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

// InsertDeviceCmdLog 添加DeviceCmdLog
func (p *DeviceCmdLogApi) InsertDeviceCmdLog(rc *restfulx.ReqCtx) {
	var data entity.DeviceCmdLog
	restfulx.BindJsonAndValid(rc, &data)
	biz.NotEmpty(data.CmdContent, "请设置指令内容")
	//验证指令格式
	ms := make(map[string]interface{})
	err := json.Unmarshal([]byte(data.CmdContent), &ms)
	biz.ErrIsNil(err, "指令格式不正确")
	biz.IsTrue(len(ms) > 0, "指令格式不正确")
	go func() {
		rpc := devicerpc.RpcPayload{
			Method: data.CmdName,
			Params: ms,
		}
		err := util.BuildRunDeviceRpc(data.DeviceId, data.Mode, rpc)
		biz.ErrIsNil(err, "添加指令记录失败")
	}()
}

// DeleteDeviceCmdLog 删除告警
func (p *DeviceCmdLogApi) DeleteDeviceCmdLog(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	ids := strings.Split(id, ",")
	biz.ErrIsNil(p.DeviceCmdLogApp.Delete(ids), "删除指令失败")
}
