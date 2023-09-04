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

type DeptApi struct {
	DeptApp services.SysDeptModel
	UserApp services.SysUserModel
	RoleApp services.SysRoleModel
}

func (m *DeptApi) GetDeptTreeRoleSelect(rc *restfulx.ReqCtx) {
	roleId := restfulx.PathParamInt(rc, "roleId")
	var dept entity.SysDept
	result := m.DeptApp.SelectDeptLable(dept)

	deptIds := make([]int64, 0)
	if roleId != 0 {
		deptIds = m.RoleApp.GetRoleDeptId(entity.SysRole{RoleId: int64(roleId)})
	}
	rc.ResData = vo.DeptTreeVo{
		Depts:       result,
		CheckedKeys: deptIds,
	}
}

func (a *DeptApi) GetDeptList(rc *restfulx.ReqCtx) {
	//pageNum := restfulx.QueryInt(rc.GinCtx, "pageNum", 1)
	//pageSize := restfulx.QueryInt(rc.GinCtx, "pageSize", 10)
	deptName := restfulx.QueryParam(rc, "deptName")
	status := restfulx.QueryParam(rc, "status")
	deptId := restfulx.QueryInt(rc, "deptId", 0)
	dept := entity.SysDept{DeptName: deptName, Status: status, DeptId: int64(deptId)}

	if dept.DeptName == "" {
		rc.ResData = a.DeptApp.SelectDept(dept)
	} else {
		rc.ResData = a.DeptApp.FindList(dept)
	}
}

func (a *DeptApi) GetOrdinaryDeptList(rc *restfulx.ReqCtx) {
	var dept entity.SysDept

	rc.ResData = a.DeptApp.FindList(dept)
}

func (a *DeptApi) GetDeptTree(rc *restfulx.ReqCtx) {
	deptName := restfulx.QueryParam(rc, "deptName")
	status := restfulx.QueryParam(rc, "status")
	deptId := restfulx.QueryInt(rc, "deptId", 0)
	dept := entity.SysDept{DeptName: deptName, Status: status, DeptId: int64(deptId)}

	rc.ResData = a.DeptApp.SelectDept(dept)
}

func (a *DeptApi) GetDept(rc *restfulx.ReqCtx) {
	deptId := restfulx.PathParamInt(rc, "deptId")
	rc.ResData = a.DeptApp.FindOne(int64(deptId))
}

func (a *DeptApi) InsertDept(rc *restfulx.ReqCtx) {
	var dept entity.SysDept
	restfulx.BindJsonAndValid(rc, &dept)
	dept.CreateBy = rc.LoginAccount.UserName
	a.DeptApp.Insert(dept)
}

func (a *DeptApi) UpdateDept(rc *restfulx.ReqCtx) {
	var dept entity.SysDept
	restfulx.BindJsonAndValid(rc, &dept)

	dept.UpdateBy = rc.LoginAccount.UserName
	a.DeptApp.Update(dept)
}

func (a *DeptApi) DeleteDept(rc *restfulx.ReqCtx) {
	deptId := restfulx.PathParam(rc, "deptId")
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
