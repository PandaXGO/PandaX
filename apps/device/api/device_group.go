package api

import (
	"github.com/XM-GO/PandaKit/restfulx"
	"github.com/kakuilan/kgo"
	"pandax/apps/device/entity"
	"pandax/apps/device/services"
	"strings"
)

type DeviceGroupApi struct {
	DeviceGroupApp services.DeviceGroupModel
}

// GetDeviceGroupTree DeviceGroup 树
func (p *DeviceGroupApi) GetDeviceGroupTree(rc *restfulx.ReqCtx) {
	name := restfulx.QueryParam(rc, "name")
	status := restfulx.QueryParam(rc, "status")
	id := restfulx.QueryParam(rc, "id")

	sg := entity.DeviceGroup{Name: name, Status: status}
	sg.Id = id
	rc.ResData = p.DeviceGroupApp.SelectDeviceGroup(sg)
}

func (p *DeviceGroupApi) GetDeviceGroupTreeLabel(rc *restfulx.ReqCtx) {
	sg := entity.DeviceGroup{}
	rc.ResData = p.DeviceGroupApp.SelectDeviceGroupLabel(sg)
}

func (p *DeviceGroupApi) GetDeviceGroupList(rc *restfulx.ReqCtx) {
	name := restfulx.QueryParam(rc, "name")
	status := restfulx.QueryParam(rc, "status")
	id := restfulx.QueryParam(rc, "id")

	sg := entity.DeviceGroup{Name: name, Status: status}
	sg.Id = id
	if sg.Name == "" {
		rc.ResData = p.DeviceGroupApp.SelectDeviceGroup(sg)
	} else {
		rc.ResData = p.DeviceGroupApp.FindList(sg)
	}
}

// GetDeviceGroupAllList 查询所有
func (p *DeviceGroupApi) GetDeviceGroupAllList(rc *restfulx.ReqCtx) {
	var vsg entity.DeviceGroup
	rc.ResData = p.DeviceGroupApp.FindList(vsg)
}

// GetDeviceGroup 获取DeviceGroup
func (p *DeviceGroupApi) GetDeviceGroup(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	rc.ResData = p.DeviceGroupApp.FindOne(id)
}

// InsertDeviceGroup 添加DeviceGroup
func (p *DeviceGroupApi) InsertDeviceGroup(rc *restfulx.ReqCtx) {
	var data entity.DeviceGroup
	restfulx.BindJsonAndValid(rc, &data)
	data.Id = kgo.KStr.Uniqid("dg_")
	data.Owner = rc.LoginAccount.UserName
	p.DeviceGroupApp.Insert(data)
}

// UpdateDeviceGroup 修改DeviceGroup
func (p *DeviceGroupApi) UpdateDeviceGroup(rc *restfulx.ReqCtx) {
	var data entity.DeviceGroup
	restfulx.BindJsonAndValid(rc, &data)

	p.DeviceGroupApp.Update(data)
}

// DeleteDeviceGroup 删除DeviceGroup
func (p *DeviceGroupApi) DeleteDeviceGroup(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	ids := strings.Split(id, ",")
	p.DeviceGroupApp.Delete(ids)
}
