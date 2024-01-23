package api

import (
	"errors"
	"fmt"
	"pandax/apps/system/api/vo"
	"pandax/apps/system/entity"
	"pandax/apps/system/services"
	"pandax/kit/biz"
	"pandax/kit/restfulx"
	"pandax/kit/utils"
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
	result := m.OrganizationApp.SelectOrganizationLable(organization)

	organizationIds := make([]int64, 0)
	if roleId != 0 {
		organizationIds = m.RoleApp.GetRoleOrganizationId(entity.SysRole{RoleId: int64(roleId)})
	}
	rc.ResData = vo.OrganizationTreeVo{
		Organizations: result,
		CheckedKeys:   organizationIds,
	}
}

func (a *OrganizationApi) GetOrganizationList(rc *restfulx.ReqCtx) {
	//pageNum := restfulx.QueryInt(rc.GinCtx, "pageNum", 1)
	//pageSize := restfulx.QueryInt(rc.GinCtx, "pageSize", 10)
	organizationName := restfulx.QueryParam(rc, "organizationName")
	status := restfulx.QueryParam(rc, "status")
	organizationId := restfulx.QueryInt(rc, "organizationId", 0)
	organization := entity.SysOrganization{OrganizationName: organizationName, Status: status, OrganizationId: int64(organizationId)}

	if organization.OrganizationName == "" {
		rc.ResData = a.OrganizationApp.SelectOrganization(organization)
	} else {
		rc.ResData = a.OrganizationApp.FindList(organization)
	}
}

func (a *OrganizationApi) GetOrdinaryOrganizationList(rc *restfulx.ReqCtx) {
	var organization entity.SysOrganization

	rc.ResData = a.OrganizationApp.FindList(organization)
}

func (a *OrganizationApi) GetOrganizationTree(rc *restfulx.ReqCtx) {
	organizationName := restfulx.QueryParam(rc, "organizationName")
	status := restfulx.QueryParam(rc, "status")
	organizationId := restfulx.QueryInt(rc, "organizationId", 0)
	organization := entity.SysOrganization{OrganizationName: organizationName, Status: status, OrganizationId: int64(organizationId)}

	rc.ResData = a.OrganizationApp.SelectOrganization(organization)
}

func (a *OrganizationApi) GetOrganization(rc *restfulx.ReqCtx) {
	organizationId := restfulx.PathParamInt(rc, "organizationId")
	rc.ResData = a.OrganizationApp.FindOne(int64(organizationId))
}

func (a *OrganizationApi) InsertOrganization(rc *restfulx.ReqCtx) {
	var organization entity.SysOrganization
	restfulx.BindJsonAndValid(rc, &organization)
	organization.CreateBy = rc.LoginAccount.UserName
	a.OrganizationApp.Insert(organization)
}

func (a *OrganizationApi) UpdateOrganization(rc *restfulx.ReqCtx) {
	var organization entity.SysOrganization
	restfulx.BindJsonAndValid(rc, &organization)

	organization.UpdateBy = rc.LoginAccount.UserName
	a.OrganizationApp.Update(organization)
}

func (a *OrganizationApi) DeleteOrganization(rc *restfulx.ReqCtx) {
	organizationId := restfulx.PathParam(rc, "organizationId")
	organizationIds := utils.IdsStrToIdsIntGroup(organizationId)

	deList := make([]int64, 0)
	for _, id := range organizationIds {
		user := entity.SysUser{}
		user.OrganizationId = id
		list := a.UserApp.FindList(user)
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
