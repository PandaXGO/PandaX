package api

import (
	"github.com/XM-GO/PandaKit/biz"
	"github.com/XM-GO/PandaKit/restfulx"
	"github.com/XM-GO/PandaKit/utils"
	"pandax/apps/system/api/vo"
	entity "pandax/apps/system/entity"
	services "pandax/apps/system/services"
)

type MenuApi struct {
	MenuApp services.SysMenuModel
	DeptApp services.SysDeptModel

	RoleMenuApp services.SysRoleMenuModel
	RoleApp     services.SysRoleModel
}

func (m *MenuApi) GetMenuTreeSelect(rc *restfulx.ReqCtx) {
	lable := m.MenuApp.SelectMenuLable(entity.SysMenu{})
	rc.ResData = lable
}

func (m *MenuApi) GetMenuRole(rc *restfulx.ReqCtx) {
	roleKey := restfulx.QueryParam(rc, "roleKey")
	biz.IsTrue(roleKey != "", "请传入角色Key")
	rc.ResData = Build(*m.MenuApp.SelectMenuRole(roleKey))
}

func (m *MenuApi) GetMenuTreeRoleSelect(rc *restfulx.ReqCtx) {
	roleId := restfulx.PathParamInt(rc, "roleId")

	result := m.MenuApp.SelectMenuLable(entity.SysMenu{})
	menuIds := make([]int64, 0)
	if roleId != 0 {
		menuIds = m.RoleApp.GetRoleMeunId(entity.SysRole{RoleId: int64(roleId)})
	}
	rc.ResData = vo.MenuTreeVo{
		Menus:       *result,
		CheckedKeys: menuIds,
	}
}

func (m *MenuApi) GetMenuPaths(rc *restfulx.ReqCtx) {
	roleKey := restfulx.QueryParam(rc, "roleKey")
	biz.IsTrue(roleKey != "", "请传入角色Key")
	rc.ResData = m.RoleMenuApp.GetMenuPaths(entity.SysRoleMenu{RoleName: roleKey})
}

func (m *MenuApi) GetMenuList(rc *restfulx.ReqCtx) {
	menuName := restfulx.QueryParam(rc, "menuName")
	status := restfulx.QueryParam(rc, "status")

	menu := entity.SysMenu{MenuName: menuName, Status: status}
	if menu.MenuName == "" {
		rc.ResData = m.MenuApp.SelectMenu(menu)
	} else {
		rc.ResData = m.MenuApp.FindList(menu)
	}
}

func (m *MenuApi) GetMenu(rc *restfulx.ReqCtx) {
	menuId := restfulx.PathParamInt(rc, "menuId")

	rc.ResData = m.MenuApp.FindOne(int64(menuId))
}

func (m *MenuApi) InsertMenu(rc *restfulx.ReqCtx) {
	var menu entity.SysMenu
	restfulx.BindQuery(rc, &menu)
	menu.CreateBy = rc.LoginAccount.UserName
	m.MenuApp.Insert(menu)
	permis := m.RoleMenuApp.GetPermis(rc.LoginAccount.RoleId)
	menus := m.MenuApp.SelectMenuRole(rc.LoginAccount.RoleKey)
	rc.ResData = vo.MenuPermisVo{
		Menus:       Build(*menus),
		Permissions: permis,
	}
}

func (m *MenuApi) UpdateMenu(rc *restfulx.ReqCtx) {
	var menu entity.SysMenu
	restfulx.BindQuery(rc, &menu)
	menu.UpdateBy = rc.LoginAccount.UserName
	m.MenuApp.Update(menu)
	permis := m.RoleMenuApp.GetPermis(rc.LoginAccount.RoleId)
	menus := m.MenuApp.SelectMenuRole(rc.LoginAccount.RoleKey)
	rc.ResData = vo.MenuPermisVo{
		Menus:       Build(*menus),
		Permissions: permis,
	}
}

func (m *MenuApi) DeleteMenu(rc *restfulx.ReqCtx) {
	menuId := restfulx.PathParam(rc, "menuId")
	m.MenuApp.Delete(utils.IdsStrToIdsIntGroup(menuId))
}
