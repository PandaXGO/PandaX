package api

import (
	"github.com/PandaXGO/PandaKit/restfulx"
	"pandax/apps/device/entity"
	"pandax/apps/device/services"
	"pandax/pkg/global/model"
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

	vsg := entity.DeviceGroup{Name: name, Status: status}
	vsg.Id = id
	vsg.RoleId = rc.LoginAccount.RoleId
	vsg.Owner = rc.LoginAccount.UserName

	rc.ResData = p.DeviceGroupApp.SelectDeviceGroup(vsg)
}

func (p *DeviceGroupApi) GetDeviceGroupTreeLabel(rc *restfulx.ReqCtx) {
	vsg := entity.DeviceGroup{}
	vsg.RoleId = rc.LoginAccount.RoleId
	vsg.Owner = rc.LoginAccount.UserName

	rc.ResData = p.DeviceGroupApp.SelectDeviceGroupLabel(vsg)
}

func (p *DeviceGroupApi) GetDeviceGroupList(rc *restfulx.ReqCtx) {
	name := restfulx.QueryParam(rc, "name")
	status := restfulx.QueryParam(rc, "status")
	id := restfulx.QueryParam(rc, "id")

	vsg := entity.DeviceGroup{Name: name, Status: status}
	vsg.Id = id

	vsg.RoleId = rc.LoginAccount.RoleId
	vsg.Owner = rc.LoginAccount.UserName
	if vsg.Name == "" {
		rc.ResData = p.DeviceGroupApp.SelectDeviceGroup(vsg)
	} else {
		rc.ResData = p.DeviceGroupApp.FindList(vsg)
	}
}

// GetDeviceGroupAllList 查询所有
func (p *DeviceGroupApi) GetDeviceGroupAllList(rc *restfulx.ReqCtx) {
	var vsg entity.DeviceGroup
	vsg.RoleId = rc.LoginAccount.RoleId
	vsg.Owner = rc.LoginAccount.UserName

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
	data.Id = model.GenerateID()
	data.Owner = rc.LoginAccount.UserName
	data.OrgId = rc.LoginAccount.OrganizationId
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
