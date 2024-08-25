package api

import (
	"github.com/PandaXGO/PandaKit/biz"
	"github.com/PandaXGO/PandaKit/restfulx"
	"github.com/PandaXGO/PandaKit/utils"
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

	vsg := entity.DeviceGroup{Name: name, Status: status}
	vsg.Id = id
	vsg.RoleId = rc.LoginAccount.RoleId
	vsg.Owner = rc.LoginAccount.UserName
	group, err := p.DeviceGroupApp.SelectDeviceGroup(vsg)
	biz.ErrIsNil(err, "查询设备组失败")
	rc.ResData = group
}

func (p *DeviceGroupApi) GetDeviceGroupTreeLabel(rc *restfulx.ReqCtx) {
	vsg := entity.DeviceGroup{}
	vsg.RoleId = rc.LoginAccount.RoleId
	vsg.Owner = rc.LoginAccount.UserName
	label, err := p.DeviceGroupApp.SelectDeviceGroupLabel(vsg)
	biz.ErrIsNil(err, "查询设备组失败")
	rc.ResData = label
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
		group, err := p.DeviceGroupApp.SelectDeviceGroup(vsg)
		biz.ErrIsNil(err, "查询设备组失败")
		rc.ResData = group
	} else {
		list, err := p.DeviceGroupApp.FindList(vsg)
		biz.ErrIsNil(err, "查询设备组列表失败")
		rc.ResData = list
	}
}

// GetDeviceGroupAllList 查询所有
func (p *DeviceGroupApi) GetDeviceGroupAllList(rc *restfulx.ReqCtx) {
	var vsg entity.DeviceGroup
	vsg.RoleId = rc.LoginAccount.RoleId
	vsg.Owner = rc.LoginAccount.UserName
	list, err := p.DeviceGroupApp.FindList(vsg)
	biz.ErrIsNil(err, "查询设备组列表失败")
	rc.ResData = list
}

// GetDeviceGroup 获取DeviceGroup
func (p *DeviceGroupApi) GetDeviceGroup(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	one, err := p.DeviceGroupApp.FindOne(id)
	biz.ErrIsNil(err, "查询设备组失败")
	rc.ResData = one
}

// InsertDeviceGroup 添加DeviceGroup
func (p *DeviceGroupApi) InsertDeviceGroup(rc *restfulx.ReqCtx) {
	var data entity.DeviceGroup
	restfulx.BindJsonAndValid(rc, &data)
	data.Id = utils.GenerateID("dg")
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
