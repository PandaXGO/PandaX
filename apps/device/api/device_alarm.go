package api

// ==========================================================================
import (
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	"strings"

	"pandax/apps/device/entity"
	"pandax/apps/device/services"
)

type DeviceAlarmApi struct {
	DeviceAlarmApp services.DeviceAlarmModel
}

// GetDeviceAlarmList 告警列表数据
func (p *DeviceAlarmApi) GetDeviceAlarmList(rc *restfulx.ReqCtx) {
	data := entity.DeviceAlarmForm{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	data.DeviceId = restfulx.QueryParam(rc, "deviceId")
	data.Type = restfulx.QueryParam(rc, "type")
	data.Level = restfulx.QueryParam(rc, "level")
	data.State = restfulx.QueryParam(rc, "state")
	data.StartTime = restfulx.QueryParam(rc, "startTime")
	data.EndTime = restfulx.QueryParam(rc, "endTime")

	list, total := p.DeviceAlarmApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

// UpdateDeviceAlarm 修改告警
func (p *DeviceAlarmApi) UpdateDeviceAlarm(rc *restfulx.ReqCtx) {
	var data entity.DeviceAlarm
	restfulx.BindJsonAndValid(rc, &data)
	p.DeviceAlarmApp.Update(data)
}

// DeleteDeviceAlarm 删除告警
func (p *DeviceAlarmApi) DeleteDeviceAlarm(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	ids := strings.Split(id, ",")
	p.DeviceAlarmApp.Delete(ids)
}
