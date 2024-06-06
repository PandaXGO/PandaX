package api

import (
	"errors"
	"fmt"
	entity "pandax/apps/system/entity"
	services "pandax/apps/system/services"
	"pandax/kit/biz"
	"pandax/kit/casbin"
	"pandax/kit/model"
	"pandax/kit/restfulx"
	"pandax/kit/utils"
	"pandax/pkg/global"
)

type RoleApi struct {
	RoleApp             services.SysRoleModel
	UserApp             services.SysUserModel
	RoleMenuApp         services.SysRoleMenuModel
	OrganizationApp     services.SysOrganizationModel
	RoleOrganizationApp services.SysRoleOrganizationModel
}

// GetRoleList角色列表数据
func (r *RoleApi) GetRoleList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	status := restfulx.QueryParam(rc, "status")
	roleName := restfulx.QueryParam(rc, "roleName")
	roleKey := restfulx.QueryParam(rc, "roleKey")
	role := entity.SysRole{Status: status, RoleName: roleName, RoleKey: roleKey}
	list, total, err := r.RoleApp.FindListPage(pageNum, pageSize, role)
	biz.ErrIsNil(err, "查询角色分页列表失败")
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

// GetRole 获取Role数据
func (r *RoleApi) GetRole(rc *restfulx.ReqCtx) {
	roleId := restfulx.PathParamInt(rc, "roleId")
	role, err := r.RoleApp.FindOne(int64(roleId))
	biz.ErrIsNil(err, "查询角色失败")
	menuIds, err := r.RoleApp.GetRoleMeunId(entity.SysRole{RoleId: int64(roleId)})
	biz.ErrIsNil(err, "查询角色菜单失败")
	role.MenuIds = menuIds
	rc.ResData = role
}

// InsertRole 创建角色
func (r *RoleApi) InsertRole(rc *restfulx.ReqCtx) {
	var role entity.SysRole
	restfulx.BindJsonAndValid(rc, &role)
	role.CreateBy = rc.LoginAccount.UserName
	if role.DataScope == "" {
		role.DataScope = "0"
	}
	// 添加角色对应的菜单
	insert, err := r.RoleApp.Insert(role)
	biz.ErrIsNil(err, "添加角色失败")
	role.RoleId = insert.RoleId
	r.RoleMenuApp.Insert(insert.RoleId, role.MenuIds)
	//添加权限
	ca := casbin.CasbinService{ModelPath: global.Conf.Casbin.ModelPath}
	ca.UpdateCasbin(role.RoleKey, role.ApiIds)
}

// UpdateRole 修改用户角色
func (r *RoleApi) UpdateRole(rc *restfulx.ReqCtx) {
	var role entity.SysRole
	restfulx.BindJsonAndValid(rc, &role)
	role.UpdateBy = rc.LoginAccount.UserName
	// 修改角色
	_, err := r.RoleApp.Update(role)
	biz.ErrIsNil(err, "修改角色失败")
	// 删除角色的菜单绑定
	r.RoleMenuApp.DeleteRoleMenu(role.RoleId)
	// 添加角色菜单绑定
	err = r.RoleMenuApp.Insert(role.RoleId, role.MenuIds)
	biz.ErrIsNil(err, "添加角色菜单绑定失败")
	//修改api权限
	ca := casbin.CasbinService{ModelPath: global.Conf.Casbin.ModelPath}
	ca.UpdateCasbin(role.RoleKey, role.ApiIds)
}

// UpdateRoleStatus 修改用户角色状态
func (r *RoleApi) UpdateRoleStatus(rc *restfulx.ReqCtx) {
	var role entity.SysRole
	restfulx.BindJsonAndValid(rc, &role)
	role.UpdateBy = rc.LoginAccount.UserName
	// 修改角色
	_, err := r.RoleApp.Update(role)
	biz.ErrIsNil(err, "修改角色失败")
}

// UpdateRoleDataScope 修改用户角色组织
func (r *RoleApi) UpdateRoleDataScope(rc *restfulx.ReqCtx) {
	var role entity.SysRole
	restfulx.BindJsonAndValid(rc, &role)
	role.UpdateBy = rc.LoginAccount.UserName
	// 修改角色
	update, err := r.RoleApp.Update(role)
	biz.ErrIsNil(err, "修改角色失败")
	go func() {
		if role.DataScope != entity.SELFDATASCOPE {
			organizationIds := make([]int64, 0)
			if role.DataScope == entity.ALLDATASCOPE {
				orgs, err := r.OrganizationApp.FindList(entity.SysOrganization{})
				if err != nil {
					return
				}
				for _, organization := range *orgs {
					organizationIds = append(organizationIds, organization.OrganizationId)
				}
			}
			if role.DataScope == entity.DIYDATASCOPE {
				organizationIds = role.OrganizationIds
			}
			if role.DataScope == entity.ORGDATASCOPE {
				organizationIds = append(organizationIds, rc.LoginAccount.OrganizationId)
			}
			if role.DataScope == entity.ORGALLDATASCOPE {
				//organizationIds = append(organizationIds, rc.LoginAccount.OrganizationId)
				orgIds, err := r.OrganizationApp.SelectOrganizationIds(entity.SysOrganization{OrganizationId: rc.LoginAccount.OrganizationId})
				if err != nil {
					return
				}
				organizationIds = orgIds
			}
			// 删除角色的组织绑定
			r.RoleOrganizationApp.Delete(entity.SysRoleOrganization{RoleId: update.RoleId})
			// 添加角色组织绑定
			r.RoleOrganizationApp.Insert(role.RoleId, organizationIds)
		}
	}()

}

// DeleteRole 删除用户角色
func (r *RoleApi) DeleteRole(rc *restfulx.ReqCtx) {
	roleId := restfulx.PathParam(rc, "roleId")
	roleIds := utils.IdsStrToIdsIntGroup(roleId)

	user := entity.SysUser{}
	delList := make([]int64, 0)
	// 判断角色下面是否绑定用户
	for _, rid := range roleIds {
		user.RoleId = rid
		role, _ := r.RoleApp.FindOne(rid)
		list, _ := r.UserApp.FindList(user)
		if len(*list) == 0 {
			delList = append(delList, rid)
			//删除角色绑定api
			ca := casbin.CasbinService{ModelPath: global.Conf.Casbin.ModelPath}
			ca.ClearCasbin(0, role.RoleKey)
		} else {
			global.Log.Info(fmt.Sprintf("role:%d 存在用户无法删除", rid))
		}
	}
	if len(delList) == 0 {
		biz.ErrIsNil(errors.New("所有角色都已绑定用户无法删除"), "所有角色都已绑定用户，无法删除")
	}
	r.RoleApp.Delete(delList)
	r.RoleMenuApp.DeleteRoleMenus(delList)
}

// ExportRole 导出角色
func (p *RoleApi) ExportRole(rc *restfulx.ReqCtx) {
	filename := restfulx.QueryParam(rc, "filename")
	status := restfulx.QueryParam(rc, "status")
	roleName := restfulx.QueryParam(rc, "roleName")
	roleKey := restfulx.QueryParam(rc, "roleKey")
	role := entity.SysRole{Status: status, RoleName: roleName, RoleKey: roleKey}

	list, err := p.RoleApp.FindList(role)
	biz.ErrIsNil(err, "查询角色列表失败")
	fileName := utils.GetFileName(global.Conf.Server.ExcelDir, filename)
	utils.InterfaceToExcel(*list, fileName)
	rc.Download(fileName)
}
