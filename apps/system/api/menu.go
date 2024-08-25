package api

import (
	"github.com/PandaXGO/PandaKit/biz"
	"github.com/PandaXGO/PandaKit/restfulx"
	"github.com/PandaXGO/PandaKit/utils"
	"pandax/apps/system/api/vo"
	entity "pandax/apps/system/entity"
	services "pandax/apps/system/services"
)

type MenuApi struct {
	MenuApp         services.SysMenuModel
	OrganizationApp services.SysOrganizationModel

	RoleMenuApp services.SysRoleMenuModel
	RoleApp     services.SysRoleModel
}

func (m *MenuApi) GetMenuTreeSelect(rc *restfulx.ReqCtx) {
	lable, err := m.MenuApp.SelectMenuLable(entity.SysMenu{})
	biz.ErrIsNil(err, "查询菜单树失败")
	rc.ResData = lable
}

func (m *MenuApi) GetMenuRole(rc *restfulx.ReqCtx) {
	roleKey := restfulx.QueryParam(rc, "roleKey")
	biz.IsTrue(roleKey != "", "请传入角色Key")
	data, err := m.MenuApp.SelectMenuRole(roleKey)
	biz.ErrIsNil(err, "查询角色菜单失败")
	rc.ResData = Build(*data)
}

func (m *MenuApi) GetMenuTreeRoleSelect(rc *restfulx.ReqCtx) {
	roleId := restfulx.PathParamInt(rc, "roleId")
	result, err := m.MenuApp.SelectMenuLable(entity.SysMenu{})
	biz.ErrIsNil(err, "查询菜单树失败")
	menuIds := make([]int64, 0)
	if roleId != 0 {
		menuIds, err = m.RoleApp.GetRoleMeunId(entity.SysRole{RoleId: int64(roleId)})
		biz.ErrIsNil(err, "通过角色查询菜单失败")
	}
	rc.ResData = vo.MenuTreeVo{
		Menus:       *result,
		CheckedKeys: menuIds,
	}
}

func (m *MenuApi) GetMenuPaths(rc *restfulx.ReqCtx) {
	roleKey := restfulx.QueryParam(rc, "roleKey")
	biz.IsTrue(roleKey != "", "请传入角色Key")
	data, err := m.RoleMenuApp.GetMenuPaths(entity.SysRoleMenu{RoleName: roleKey})
	biz.ErrIsNil(err, "查询菜单路径失败")
	rc.ResData = data
}

func (m *MenuApi) GetMenuList(rc *restfulx.ReqCtx) {
	menuName := restfulx.QueryParam(rc, "menuName")
	status := restfulx.QueryParam(rc, "status")

	menu := entity.SysMenu{MenuName: menuName, Status: status}
	if menu.MenuName == "" {
		data, err := m.MenuApp.SelectMenu(menu)
		biz.ErrIsNil(err, "查询菜单列表失败")
		rc.ResData = data
	} else {
		data, err := m.MenuApp.FindList(menu)
		biz.ErrIsNil(err, "查询菜单列表失败")
		rc.ResData = data
	}
}

func (m *MenuApi) GetMenu(rc *restfulx.ReqCtx) {
	menuId := restfulx.PathParamInt(rc, "menuId")
	data, err := m.MenuApp.FindOne(int64(menuId))
	biz.ErrIsNil(err, "查询菜单失败")
	rc.ResData = data
}

func (m *MenuApi) InsertMenu(rc *restfulx.ReqCtx) {
	var menu entity.SysMenu
	restfulx.BindJsonAndValid(rc, &menu)
	menu.CreateBy = rc.LoginAccount.UserName
	m.MenuApp.Insert(menu)
	permis, _ := m.RoleMenuApp.GetPermis(rc.LoginAccount.RoleId)
	menus, _ := m.MenuApp.SelectMenuRole(rc.LoginAccount.RoleKey)
	rc.ResData = vo.MenuPermisVo{
		Menus:       Build(*menus),
		Permissions: permis,
	}
}

func (m *MenuApi) UpdateMenu(rc *restfulx.ReqCtx) {
	var menu entity.SysMenu
	restfulx.BindJsonAndValid(rc, &menu)
	menu.UpdateBy = rc.LoginAccount.UserName
	m.MenuApp.Update(menu)
	permis, _ := m.RoleMenuApp.GetPermis(rc.LoginAccount.RoleId)
	menus, _ := m.MenuApp.SelectMenuRole(rc.LoginAccount.RoleKey)
	rc.ResData = vo.MenuPermisVo{
		Menus:       Build(*menus),
		Permissions: permis,
	}
}

func (m *MenuApi) DeleteMenu(rc *restfulx.ReqCtx) {
	menuId := restfulx.PathParam(rc, "menuId")
	err := m.MenuApp.Delete(utils.IdsStrToIdsIntGroup(menuId))
	biz.ErrIsNil(err, "删除菜单失败")
}
