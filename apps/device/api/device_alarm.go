package api

// ==========================================================================
import (
	"github.com/PandaXGO/PandaKit/biz"
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	"strings"

	"pandax/apps/device/entity"
	"pandax/apps/device/services"
)

type DeviceAlarmApi struct {
	DeviceAlarmApp services.DeviceAlarmModel
}

func (p *DeviceAlarmApi) GetDeviceAlarmPanel(rc *restfulx.ReqCtx) {
	panel, _ := p.DeviceAlarmApp.FindAlarmPanel()
	rc.ResData = panel
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

	data.RoleId = rc.LoginAccount.RoleId
	data.Owner = rc.LoginAccount.UserName

	list, total, err := p.DeviceAlarmApp.FindListPage(pageNum, pageSize, data)
	biz.ErrIsNil(err, "设备查询失败")
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
	err := p.DeviceAlarmApp.Update(data)
	biz.ErrIsNil(err, "修改告警失败")
}

// DeleteDeviceAlarm 删除告警
func (p *DeviceAlarmApi) DeleteDeviceAlarm(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	ids := strings.Split(id, ",")
	err := p.DeviceAlarmApp.Delete(ids)
	biz.ErrIsNil(err, "删除告警失败")
}
