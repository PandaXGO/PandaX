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
	"github.com/kakuilan/kgo"
	"pandax/pkg/global"
	"pandax/pkg/mqtt"
	"strings"
	"time"

	"pandax/apps/device/entity"
	"pandax/apps/device/services"
)

type DeviceApi struct {
	DeviceApp          services.DeviceModel
	ProductApp         services.ProductModel
	ProductTemplateApp services.ProductTemplateModel
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

	list, total := p.DeviceApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageNum),
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
	rs, err := global.TdDb.GetOne(fmt.Sprintf(`select * from %s_%s ORDER BY ts DESC LIMIT 1`, device.Name, classify))
	biz.ErrIsNil(err, "查询设备状态信息失败")
	for _, tel := range *template {
		sdv := entity.DeviceStatusVo{
			Name:   tel.Name,
			Key:    tel.Key,
			Type:   tel.Type,
			Define: tel.Define,
			Time:   rs["ts"],
		}
		if v, ok := rs[tel.Key]; ok {
			sdv.Value = v
		} else {
			sdv.Value = tel.Define["default_value"]
		}
		res = append(res, sdv)
	}

	rc.ResData = res
}

// 下发设备属性
func (p *DeviceApi) DownAttribute(rc *restfulx.ReqCtx) {
	//id := restfulx.PathParam(rc, "id")
	key := restfulx.QueryParam(rc, "key")
	value := restfulx.QueryParam(rc, "value")
	// 下发指令
	contentMap := map[string]interface{}{
		key: value,
	}
	content, _ := json.Marshal(contentMap)
	var rpc = &mqtt.RpcRequest{Client: global.MqttClient, Mode: "single"}
	rpc.GetRequestId()
	err := rpc.RequestAttributes(mqtt.RpcPayload{Params: string(content)})
	biz.ErrIsNil(err, "属性下发失败")
}

// InsertDevice 添加Device
func (p *DeviceApi) InsertDevice(rc *restfulx.ReqCtx) {
	var data entity.Device
	restfulx.BindJsonAndValid(rc, &data)
	list := p.DeviceApp.FindList(entity.Device{Name: data.Name})
	biz.IsTrue(!(list != nil && len(*list) > 0), fmt.Sprintf("名称%s已存在，设置其他命名", data.Name))
	data.Id = kgo.KStr.Uniqid("d_")
	data.Owner = rc.LoginAccount.UserName
	data.LinkStatus = global.INACTIVE
	data.LastAt = time.Now()
	p.DeviceApp.Insert(data)
	// 视频设备不创建超级表
	if data.DeviceType != global.MONITOR {
		createDeviceTable(data.Pid, data.Name)
	}
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
	for _, id := range ids {
		list := p.DeviceApp.FindOne(id)
		// 删除表
		if list.DeviceType != global.MONITOR {
			deleteDeviceTable(list.Name)
		}
	}
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
		findList, _ := p.DeviceApp.FindListPage(pageNum, pageSize, entity.Device{Pid: classId})
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

func createDeviceTable(productId, device string) {
	err := global.TdDb.CreateTable(productId+"_"+entity.ATTRIBUTES_TSL, device+"_"+entity.ATTRIBUTES_TSL)
	biz.ErrIsNil(err, "创建时序属性表失败")
	err = global.TdDb.CreateTable(productId+"_"+entity.TELEMETRY_TSL, device+"_"+entity.TELEMETRY_TSL)
	biz.ErrIsNil(err, "创建时序遥测表失败")
}

func deleteDeviceTable(device string) {
	err := global.TdDb.DropTable(device + "_" + entity.ATTRIBUTES_TSL)
	biz.ErrIsNil(err, "删除时序属性表失败")
	err = global.TdDb.DropTable(device + "_" + entity.TELEMETRY_TSL)
	biz.ErrIsNil(err, "删除时序遥测表失败")
}
