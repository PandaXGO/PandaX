package api

import (
	"errors"
	"fmt"
	entity "pandax/apps/system/entity"
	services "pandax/apps/system/services"
	"pandax/base/biz"
	"pandax/base/casbin"
	"pandax/base/ginx"
	"pandax/base/utils"
	"pandax/pkg/global"
	"strconv"
)

type RoleApi struct {
	RoleApp     services.SysRoleModel
	UserApp     services.SysUserModel
	RoleMenuApp services.SysRoleMenuModel
	RoleDeptApp services.SysRoleDeptModel
}

// @Summary 角色列表数据
// @Description Get JSON
// @Tags 角色
// @Param roleName query string false "roleName"
// @Param status query string false "status"
// @Param roleKey query string false "roleKey"
// @Param pageSize query int false "页条数"
// @Param pageNum query int false "页码"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Router /system/role/rolelist [get]
// @Security
func (r *RoleApi) GetRoleList(rc *ginx.ReqCtx) {
	pageNum := ginx.QueryInt(rc.GinCtx, "pageNum", 1)
	pageSize := ginx.QueryInt(rc.GinCtx, "pageSize", 10)
	status := rc.GinCtx.Query("status")
	roleName := rc.GinCtx.Query("roleName")
	roleKey := rc.GinCtx.Query("roleKey")
	role := entity.SysRole{Status: status, RoleName: roleName, RoleKey: roleKey}

	if !IsTenantAdmin(rc.LoginAccount.TenantId) {
		role.TenantId = rc.LoginAccount.TenantId
	}

	list, total := r.RoleApp.FindListPage(pageNum, pageSize, role)

	rc.ResData = map[string]any{
		"data":     list,
		"total":    total,
		"pageNum":  pageNum,
		"pageSize": pageSize,
	}
}

// @Summary 获取Role数据
// @Description 获取JSON
// @Tags 角色/Role
// @Param roleId path string true "roleId"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": 400, "message": "抱歉未找到相关信息"}"
// @Router /system/role [get]
// @Security X-TOKEN
func (r *RoleApi) GetRole(rc *ginx.ReqCtx) {
	roleId := ginx.PathParamInt(rc.GinCtx, "roleId")
	role := r.RoleApp.FindOne(int64(roleId))
	role.MenuIds = r.RoleApp.GetRoleMeunId(entity.SysRole{RoleId: int64(roleId)})

	rc.ResData = role
}

// @Summary 创建角色
// @Description 获取JSON
// @Tags 角色/Role
// @Accept  application/json
// @Product application/json
// @Param data body entity.SysRole true "data"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "添加失败"}"
// @Router /system/role [post]
func (r *RoleApi) InsertRole(rc *ginx.ReqCtx) {
	var role entity.SysRole
	ginx.BindJsonAndValid(rc.GinCtx, &role)
	role.CreateBy = rc.LoginAccount.UserName
	role.TenantId = rc.LoginAccount.TenantId
	insert := r.RoleApp.Insert(role)
	role.RoleId = insert.RoleId
	r.RoleMenuApp.Insert(insert.RoleId, role.MenuIds)
	//添加权限
	tenantId := strconv.Itoa(int(rc.LoginAccount.TenantId))
	casbin.UpdateCasbin(tenantId, role.RoleKey, role.ApiIds)
}

// @Summary 修改用户角色
// @Description 获取JSON
// @Tags 角色/Role
// @Accept  application/json
// @Product application/json
// @Param data body entity.SysRole true "body"
// @Success 200 {string} string	"{"code": 200, "message": "修改成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "修改失败"}"
// @Router /system/role [put]
func (r *RoleApi) UpdateRole(rc *ginx.ReqCtx) {
	var role entity.SysRole
	ginx.BindJsonAndValid(rc.GinCtx, &role)
	role.UpdateBy = rc.LoginAccount.UserName
	// 修改角色
	r.RoleApp.Update(role)
	// 删除角色的菜单绑定
	r.RoleMenuApp.DeleteRoleMenu(role.RoleId)
	// 添加角色菜单绑定
	r.RoleMenuApp.Insert(role.RoleId, role.MenuIds)
	//修改api权限
	tenantId := strconv.Itoa(int(rc.LoginAccount.TenantId))
	casbin.UpdateCasbin(tenantId, role.RoleKey, role.ApiIds)
}

// @Summary 修改用户角色状态
// @Description 获取JSON
// @Tags 角色/Role
// @Accept  application/json
// @Product application/json
// @Param data body entity.SysRole true "body"
// @Success 200 {string} string	"{"code": 200, "message": "修改成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "修改失败"}"
// @Router /system/role/changeStatus [put]
func (r *RoleApi) UpdateRoleStatus(rc *ginx.ReqCtx) {
	var role entity.SysRole
	ginx.BindJsonAndValid(rc.GinCtx, &role)
	role.UpdateBy = rc.LoginAccount.UserName
	// 修改角色
	r.RoleApp.Update(role)
}

// @Summary 修改用户角色部门
// @Description 获取JSON
// @Tags 角色/Role
// @Accept  application/json
// @Product application/json
// @Param data body entity.SysRole true "body"
// @Success 200 {string} string	"{"code": 200, "message": "修改成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "修改失败"}"
// @Router /system/role/dataScope [put]
func (r *RoleApi) UpdateRoleDataScope(rc *ginx.ReqCtx) {
	var role entity.SysRole
	ginx.BindJsonAndValid(rc.GinCtx, &role)
	role.UpdateBy = rc.LoginAccount.UserName
	// 修改角色
	update := r.RoleApp.Update(role)
	if role.DataScope == "2" {
		// 删除角色的部门绑定
		r.RoleDeptApp.Delete(entity.SysRoleDept{RoleId: update.RoleId})
		// 添加角色部门绑定
		r.RoleDeptApp.Insert(role.RoleId, role.DeptIds)
	}
}

// @Summary 删除用户角色
// @Description 删除数据
// @Tags 角色/Role
// @Param roleId path string true "roleId 多个用，分割"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "删除失败"}"
// @Router /system/role/{roleId} [delete]
func (r *RoleApi) DeleteRole(rc *ginx.ReqCtx) {
	roleId := rc.GinCtx.Param("roleId")
	roleIds := utils.IdsStrToIdsIntGroup(roleId)

	user := entity.SysUser{}
	delList := make([]int64, 0)
	// 判断角色下面是否绑定用户
	for _, rid := range roleIds {
		user.RoleId = rid
		role := r.RoleApp.FindOne(rid)
		list := r.UserApp.FindList(user)
		if len(*list) == 0 {
			delList = append(delList, rid)
			//删除角色绑定api
			casbin.ClearCasbin(0, role.RoleKey)
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

// @Summary 导出角色
// @Description 导出数据
// @Tags 角色
// @Param roleName query string false "roleName"
// @Param status query string false "status"
// @Param roleKey query string false "roleKey"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "删除失败"}"
// @Router /system/dict/type/export [get]
func (p *RoleApi) ExportRole(rc *ginx.ReqCtx) {
	filename := rc.GinCtx.Query("filename")
	status := rc.GinCtx.Query("status")
	roleName := rc.GinCtx.Query("roleName")
	roleKey := rc.GinCtx.Query("roleKey")
	role := entity.SysRole{Status: status, RoleName: roleName, RoleKey: roleKey}
	if !IsTenantAdmin(rc.LoginAccount.TenantId) {
		role.TenantId = rc.LoginAccount.TenantId
	}
	list := p.RoleApp.FindList(role)

	fileName := utils.GetFileName(global.Conf.Server.ExcelDir, filename)
	utils.InterfaceToExcel(*list, fileName)
	rc.Download(fileName)
}
