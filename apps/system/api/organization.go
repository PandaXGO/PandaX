package api

import (
	"errors"
	"fmt"
	"github.com/PandaXGO/PandaKit/biz"
	"github.com/PandaXGO/PandaKit/restfulx"
	"github.com/PandaXGO/PandaKit/utils"
	"pandax/apps/system/api/vo"
	"pandax/apps/system/entity"
	"pandax/apps/system/services"
	"pandax/pkg/global"
)

type OrganizationApi struct {
	OrganizationApp services.SysOrganizationModel
	UserApp         services.SysUserModel
	RoleApp         services.SysRoleModel
}

func (m *OrganizationApi) GetOrganizationTreeRoleSelect(rc *restfulx.ReqCtx) {
	roleId := restfulx.PathParamInt(rc, "roleId")
	var organization entity.SysOrganization
	result, err := m.OrganizationApp.SelectOrganizationLable(organization)
	biz.ErrIsNil(err, "查询组织树失败")
	organizationIds := make([]int64, 0)
	if roleId != 0 {
		organizationIds, err = m.RoleApp.GetRoleOrganizationId(entity.SysRole{RoleId: int64(roleId)})
		biz.ErrIsNil(err, "查询角色组织失败")
	}
	rc.ResData = vo.OrganizationTreeVo{
		Organizations: result,
		CheckedKeys:   organizationIds,
	}
}

func (a *OrganizationApi) GetOrganizationList(rc *restfulx.ReqCtx) {
	organizationName := restfulx.QueryParam(rc, "organizationName")
	status := restfulx.QueryParam(rc, "status")
	organizationId := restfulx.QueryInt(rc, "organizationId", 0)
	organization := entity.SysOrganization{OrganizationName: organizationName, Status: status, OrganizationId: int64(organizationId)}

	if organization.OrganizationName == "" {
		data, err := a.OrganizationApp.SelectOrganization(organization)
		biz.ErrIsNil(err, "查询组织树失败")
		rc.ResData = data
	} else {
		data, err := a.OrganizationApp.FindList(organization)
		biz.ErrIsNil(err, "查询组织列表失败")
		rc.ResData = data
	}
}

func (a *OrganizationApi) GetOrdinaryOrganizationList(rc *restfulx.ReqCtx) {
	var organization entity.SysOrganization
	data, err := a.OrganizationApp.FindList(organization)
	biz.ErrIsNil(err, "查询组织列表失败")
	rc.ResData = data
}

func (a *OrganizationApi) GetOrganizationTree(rc *restfulx.ReqCtx) {
	organizationName := restfulx.QueryParam(rc, "organizationName")
	status := restfulx.QueryParam(rc, "status")
	organizationId := restfulx.QueryInt(rc, "organizationId", 0)
	organization := entity.SysOrganization{OrganizationName: organizationName, Status: status, OrganizationId: int64(organizationId)}

	data, err := a.OrganizationApp.SelectOrganization(organization)
	biz.ErrIsNil(err, "查询组织树失败")
	rc.ResData = data
}

func (a *OrganizationApi) GetOrganization(rc *restfulx.ReqCtx) {
	organizationId := restfulx.PathParamInt(rc, "organizationId")
	data, err := a.OrganizationApp.FindOne(int64(organizationId))
	biz.ErrIsNil(err, "查询组织失败")
	rc.ResData = data
}

func (a *OrganizationApi) InsertOrganization(rc *restfulx.ReqCtx) {
	var organization entity.SysOrganization
	restfulx.BindJsonAndValid(rc, &organization)
	organization.CreateBy = rc.LoginAccount.UserName
	_, err := a.OrganizationApp.Insert(organization)
	biz.ErrIsNil(err, "添加组织失败")
}

func (a *OrganizationApi) UpdateOrganization(rc *restfulx.ReqCtx) {
	var organization entity.SysOrganization
	restfulx.BindJsonAndValid(rc, &organization)

	organization.UpdateBy = rc.LoginAccount.UserName
	err := a.OrganizationApp.Update(organization)
	biz.ErrIsNil(err, "修改组织失败")
}

func (a *OrganizationApi) DeleteOrganization(rc *restfulx.ReqCtx) {
	organizationId := restfulx.PathParam(rc, "organizationId")
	organizationIds := utils.IdsStrToIdsIntGroup(organizationId)

	deList := make([]int64, 0)
	for _, id := range organizationIds {
		user := entity.SysUser{}
		user.OrganizationId = id
		list, err := a.UserApp.FindList(user)
		if err != nil {
			continue
		}
		if len(*list) == 0 {
			deList = append(deList, id)
		} else {
			global.Log.Info(fmt.Sprintf("dictId: %d 存在用户绑定无法删除", id))
		}
	}
	if len(deList) == 0 {
		biz.ErrIsNil(errors.New("所有组织都已绑定用户无法删除"), "所有组织都已绑定用户，无法删除")
	}
	a.OrganizationApp.Delete(deList)
}
