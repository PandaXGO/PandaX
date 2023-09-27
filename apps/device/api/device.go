package api

// ==========================================================================
// 生成日期：2023-06-30 09:19:43 +0800 CST
// 生成路径: apps/device/api/devices.go
// 生成人：panda
// ==========================================================================
import (
	"encoding/json"
	"fmt"
	"github.com/PandaXGO/PandaKit/biz"
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	"pandax/iothub/client/mqttclient"
	"pandax/iothub/client/tcpclient"
	"pandax/pkg/global"
	"pandax/pkg/shadow"
	"pandax/pkg/tool"
	"strings"
	"time"

	"pandax/apps/device/entity"
	"pandax/apps/device/services"
)

type DeviceApi struct {
	DeviceApp          services.DeviceModel
	DeviceAlarmApp     services.DeviceAlarmModel
	ProductApp         services.ProductModel
	ProductTemplateApp services.ProductTemplateModel
}

func (p *DeviceApi) GetDevicePanel(rc *restfulx.ReqCtx) {
	get, err := global.Cache.ComputeIfAbsent("panel", func(k any) (any, error) {
		var data entity.DeviceTotalOutput
		data.DeviceInfo = p.DeviceApp.FindDeviceCount()
		data.DeviceLinkStatusInfo = p.DeviceApp.FindDeviceCountGroupByLinkStatus()
		data.DeviceCountType = p.DeviceApp.FindDeviceCountGroupByType()
		data.AlarmInfo = p.DeviceAlarmApp.FindAlarmCount()
		data.ProductInfo = p.ProductApp.FindProductCount()
		return data, nil
	})
	biz.ErrIsNil(err, "获取面板数据失败")
	rc.ResData = get
}

// GetDeviceList Device列表数据
func (p *DeviceApi) GetDeviceList(rc *restfulx.ReqCtx) {
	data := entity.Device{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	data.Name = restfulx.QueryParam(rc, "name")
	data.Status = restfulx.QueryParam(rc, "status")
	data.Pid = restfulx.QueryParam(rc, "pid")
	data.Gid = restfulx.QueryParam(rc, "gid")
	data.DeviceType = restfulx.QueryParam(rc, "deviceType")
	data.ParentId = restfulx.QueryParam(rc, "parentId")
	data.LinkStatus = restfulx.QueryParam(rc, "linkStatus")

	data.RoleId = rc.LoginAccount.RoleId
	data.Owner = rc.LoginAccount.UserName
	list, total := p.DeviceApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

// GetDeviceListAll Device获取所有
func (p *DeviceApi) GetDeviceListAll(rc *restfulx.ReqCtx) {
	data := entity.Device{}
	data.Name = restfulx.QueryParam(rc, "name")
	data.Status = restfulx.QueryParam(rc, "status")
	data.Pid = restfulx.QueryParam(rc, "pid")
	data.DeviceType = restfulx.QueryParam(rc, "deviceType")
	data.RoleId = rc.LoginAccount.RoleId
	data.Owner = rc.LoginAccount.UserName

	list := p.DeviceApp.FindList(data)
	rc.ResData = list
}

// GetDevice 获取Device
func (p *DeviceApi) GetDevice(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	rc.ResData = p.DeviceApp.FindOne(id)
}

// GetDeviceStatus 获取Device状态信息
func (p *DeviceApi) GetDeviceStatus(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	classify := restfulx.QueryParam(rc, "classify")
	device := p.DeviceApp.FindOne(id)
	template := p.ProductTemplateApp.FindList(entity.ProductTemplate{Classify: classify, Pid: device.Pid})

	res := make([]entity.DeviceStatusVo, 0)
	getDevice := shadow.InitDeviceShadow(device.Name, device.Pid)
	rs := make(map[string]shadow.DevicePoint)
	if classify == global.TslAttributesType {
		rs = getDevice.AttributesPoints
	}
	if classify == global.TslTelemetryType {
		rs = getDevice.TelemetryPoints
	}
	for _, tel := range *template {
		sdv := entity.DeviceStatusVo{
			Name:   tel.Name,
			Key:    tel.Key,
			Type:   tel.Type,
			Define: tel.Define,
		}
		if v, ok := rs[strings.ToLower(tel.Key)]; ok {
			sdv.Value = v.Value
			sdv.Time = v.UpdatedAt
		} else {
			if classify == global.TslAttributesType {
				if value, ok := tel.Define["default_value"]; ok {
					sdv.Value = value
				}
			}
		}
		res = append(res, sdv)
	}

	rc.ResData = res
}

// GetDeviceTelemetryHistory 获取Device属性的遥测历史
func (p *DeviceApi) GetDeviceTelemetryHistory(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	key := restfulx.QueryParam(rc, "key")
	startTime := restfulx.QueryParam(rc, "startTime")
	endTime := restfulx.QueryParam(rc, "endTime")
	limit := restfulx.QueryInt(rc, "limit", 1000)
	device := p.DeviceApp.FindOne(id)
	sql := `select ts,? from ? where ts > '?' and ts < '?' and ? is not null ORDER BY ts DESC LIMIT ? `
	rs, err := global.TdDb.GetAll(sql, key, fmt.Sprintf("%s_telemetry", strings.ToLower(device.Name)), startTime, endTime, key, limit)
	biz.ErrIsNilAppendErr(err, "查询设备属性的遥测历史失败")
	rc.ResData = rs
}

// 下发设备属性
func (p *DeviceApi) DownAttribute(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	key := restfulx.QueryParam(rc, "key")
	value := restfulx.QueryParam(rc, "value")
	one := p.DeviceApp.FindOne(id)
	if one.Product.ProtocolName == global.TCPProtocol {
		err := tcpclient.Send(id, value)
		biz.ErrIsNil(err, "属性下发失败")
	}
	if one.Product.ProtocolName == global.MQTTProtocol {
		// 下发指令
		contentMap := map[string]interface{}{
			key: value,
		}
		content, _ := json.Marshal(contentMap)
		var rpc = &mqttclient.RpcRequest{Client: mqttclient.MqttClient, Mode: "single"}
		rpc.GetRequestId()
		err := rpc.RequestAttributes(mqttclient.RpcPayload{Params: string(content)})
		biz.ErrIsNil(err, "属性下发失败")
	}
}

// InsertDevice 添加Device
func (p *DeviceApi) InsertDevice(rc *restfulx.ReqCtx) {
	var data entity.Device
	restfulx.BindJsonAndValid(rc, &data)
	data.Owner = rc.LoginAccount.UserName
	data.OrgId = rc.LoginAccount.OrganizationId
	list := p.DeviceApp.FindList(entity.Device{Name: data.Name})
	biz.IsTrue(!(list != nil && len(*list) > 0), fmt.Sprintf("名称%s已存在，设置其他命名", data.Name))
	data.Id = tool.GenerateID()
	data.LinkStatus = global.INACTIVE
	data.LastAt = time.Now()
	p.DeviceApp.Insert(data)
}

// UpdateDevice 修改Device
func (p *DeviceApi) UpdateDevice(rc *restfulx.ReqCtx) {
	var data entity.Device
	restfulx.BindJsonAndValid(rc, &data)

	p.DeviceApp.Update(data)
}

// DeleteDevice 删除Device
func (p *DeviceApi) DeleteDevice(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	ids := strings.Split(id, ",")
	p.DeviceApp.Delete(ids)
}

func (p *DeviceApi) ScreenTwinData(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	classId := restfulx.QueryParam(rc, "classId")
	if classId == "" {
		vp := make([]entity.VisualClass, 0)
		list := p.ProductApp.FindList(entity.Product{})
		for _, pro := range *list {
			data := p.ProductTemplateApp.FindListAttrs(entity.ProductTemplate{Pid: pro.Id})
			vta := make([]entity.VisualTwinAttr, 0)
			for _, attr := range *data {
				twinAttr := entity.VisualTwinAttr{
					Key:  attr.Key,
					Name: attr.Name,
					Type: attr.Type,
				}
				if attr.Classify == "attributes" {
					if rw, ok := attr.Define["rw"].(string); ok {
						twinAttr.Rw = rw
					} else {
						twinAttr.Rw = "r"
					}
				} else {
					twinAttr.Rw = "r"
				}
				vta = append(vta, twinAttr)
			}
			vp = append(vp, entity.VisualClass{
				ClassId: pro.Id,
				Name:    pro.Name,
				Attrs:   vta,
			})
		}
		rc.ResData = vp
	} else {
		device := entity.Device{Pid: classId, RoleId: rc.LoginAccount.RoleId}
		device.Owner = rc.LoginAccount.UserName
		findList, _ := p.DeviceApp.FindListPage(pageNum, pageSize, device)
		vt := make([]entity.VisualTwin, 0)
		for _, device := range *findList {
			vt = append(vt, entity.VisualTwin{
				TwinId: device.Id,
				Name:   device.Name + "-" + device.Alias,
			})
		}
		rc.ResData = vt
	}
}

func (p *DeviceApi) DeviceAllotOrg(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	orgId := restfulx.QueryInt(rc, "orgId", 0)
	biz.IsTrue(orgId != 0, "请选择组织")

	device := entity.Device{}
	device.Id = id
	device.OrgId = int64(orgId)
	p.DeviceApp.Update(device)
}
