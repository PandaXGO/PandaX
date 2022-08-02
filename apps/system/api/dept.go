package api

import (
	"errors"
	"fmt"
	entity "pandax/apps/system/entity"
	services "pandax/apps/system/services"
	"pandax/base/biz"
	"pandax/base/ginx"
	"pandax/base/utils"
	"pandax/pkg/global"
)

type DeptApi struct {
	DeptApp services.SysDeptModel
	UserApp services.SysUserModel
	RoleApp services.SysRoleModel
}

// @Summary 获取角色的部门树
// @Description 获取JSON
// @Tags 菜单
// @Param roleId path int false "roleId"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": 400, "message": "抱歉未找到相关信息"}"
// @Router /system/menu/menuTreRoleSelect/{roleId} [get]
// @Security X-TOKEN
func (m *DeptApi) GetDeptTreeRoleSelect(rc *ginx.ReqCtx) {
	roleId := ginx.PathParamInt(rc.GinCtx, "roleId")
	var dept entity.SysDept
	if !IsTenantAdmin(rc.LoginAccount.TenantId) {
		dept.TenantId = rc.LoginAccount.TenantId
	}
	result := m.DeptApp.SelectDeptLable(dept)

	deptIds := make([]int64, 0)
	if roleId != 0 {
		deptIds = m.RoleApp.GetRoleDeptId(entity.SysRole{RoleId: int64(roleId)})
	}
	rc.ResData = map[string]any{
		"depts":       result,
		"checkedKeys": deptIds,
	}
}

// @Summary 部门列表数据
// @Description 分页列表
// @Tags 部门
// @Param deptName query string false "deptName"
// @Param status query string false "status"
// @Param deptId query int false "deptId"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Router /system/dept/deptList [get]
// @Security
func (a *DeptApi) GetDeptList(rc *ginx.ReqCtx) {
	//pageNum := ginx.QueryInt(rc.GinCtx, "pageNum", 1)
	//pageSize := ginx.QueryInt(rc.GinCtx, "pageSize", 10)
	deptName := rc.GinCtx.Query("deptName")
	status := rc.GinCtx.Query("status")
	deptId := ginx.QueryInt(rc.GinCtx, "deptId", 0)
	dept := entity.SysDept{DeptName: deptName, Status: status, DeptId: int64(deptId)}
	if !IsTenantAdmin(rc.LoginAccount.TenantId) {
		dept.TenantId = rc.LoginAccount.TenantId
	}
	if dept.DeptName == "" {
		rc.ResData = a.DeptApp.SelectDept(dept)
	} else {
		rc.ResData = a.DeptApp.FindList(dept)
	}
}

// @Summary 所有部门列表数据
// @Description 部门列表
// @Tags 部门
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Router /system/dept/ordinaryDeptLis [get]
// @Security
func (a *DeptApi) GetOrdinaryDeptList(rc *ginx.ReqCtx) {
	var dept entity.SysDept
	if !IsTenantAdmin(rc.LoginAccount.TenantId) {
		dept.TenantId = rc.LoginAccount.TenantId
	}

	rc.ResData = a.DeptApp.FindList(dept)
}

// @Summary 所有部门树数据
// @Description 部门树列表
// @Tags 部门
// @Param deptName query string false "deptName"
// @Param status query string false "status"
// @Param deptId query int false "deptId"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Router /system/dept/deptTree [get]
// @Security
func (a *DeptApi) GetDeptTree(rc *ginx.ReqCtx) {
	deptName := rc.GinCtx.Query("deptName")
	status := rc.GinCtx.Query("status")
	deptId := ginx.QueryInt(rc.GinCtx, "deptId", 0)
	dept := entity.SysDept{DeptName: deptName, Status: status, DeptId: int64(deptId)}
	if !IsTenantAdmin(rc.LoginAccount.TenantId) {
		dept.TenantId = rc.LoginAccount.TenantId
	}
	rc.ResData = a.DeptApp.SelectDept(dept)
}

// @Summary 部门数据
// @Description 获取JSON
// @Tags 部门
// @Param deptId path string false "deptId"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Router /system/dept/{deptId} [get]
// @Security
func (a *DeptApi) GetDept(rc *ginx.ReqCtx) {
	deptId := ginx.PathParamInt(rc.GinCtx, "deptId")
	rc.ResData = a.DeptApp.FindOne(int64(deptId))
}

// @Summary 添加部门
// @Description 获取JSON
// @Tags 部门
// @Accept  application/json
// @Product application/json
// @Param data body entity.SysDept true "data"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "添加失败"}"
// @Router /system/dept [post]
// @Security Bearer
func (a *DeptApi) InsertDept(rc *ginx.ReqCtx) {
	var dept entity.SysDept
	g := rc.GinCtx
	ginx.BindJsonAndValid(g, &dept)
	dept.TenantId = rc.LoginAccount.TenantId
	dept.CreateBy = rc.LoginAccount.UserName
	a.DeptApp.Insert(dept)
}

// @Summary 修改部门
// @Description 获取JSON
// @Tags 部门
// @Accept  application/json
// @Product application/json
// @Param data body entity.SysDept true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /system/dept [put]
// @Security Bearer
func (a *DeptApi) UpdateDept(rc *ginx.ReqCtx) {
	var dept entity.SysDept
	g := rc.GinCtx
	ginx.BindJsonAndValid(g, &dept)

	dept.UpdateBy = rc.LoginAccount.UserName
	a.DeptApp.Update(dept)
}

// @Summary 删除部门
// @Description 删除数据
// @Tags 部门
// @Param deptId path string true "deptId, 逗号隔开"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "删除失败"}"
// @Router /system/dept/{deptId} [delete]
func (a *DeptApi) DeleteDept(rc *ginx.ReqCtx) {
	deptId := rc.GinCtx.Param("deptId")
	deptIds := utils.IdsStrToIdsIntGroup(deptId)

	deList := make([]int64, 0)
	for _, id := range deptIds {
		user := entity.SysUser{}
		user.DeptId = id
		list := a.UserApp.FindList(user)
		if len(*list) == 0 {
			deList = append(deList, id)
		} else {
			global.Log.Info(fmt.Sprintf("dictId: %d 存在用户绑定无法删除", id))
		}
	}
	if len(deList) == 0 {
		biz.ErrIsNil(errors.New("所有部门都已绑定用户无法删除"), "所有部门都已绑定用户，无法删除")
	}
	a.DeptApp.Delete(deList)
}
