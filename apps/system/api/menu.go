package api

import (
	"log"
	entity2 "pandax/apps/system/entity"
	services2 "pandax/apps/system/services"
	"pandax/base/biz"
	"pandax/base/ctx"
	"pandax/base/ginx"
)

type MenuApi struct {
	MenuApp services2.SysMenuModel
	DeptApp services2.SysDeptModel

	RoleMenuApp services2.SysRoleMenuModel
	RoleApp     services2.SysRoleModel
}

// @Summary 获取菜单树
// @Description 获取JSON
// @Tags 菜单
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "添加失败"}"
// @Router /system/menu/menuTreSelect [get]
// @Security X-TOKEN
func (m *MenuApi) GetMenuTreeSelect(rc *ctx.ReqCtx) {
	lable := m.MenuApp.SelectMenuLable(entity2.SysMenu{})
	rc.ResData = lable
}

// @Summary 登陆成功获取的路由，根据角色名称获取菜单列表数据（左菜单使用）
// @Description 获取JSON
// @Tags 菜单
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": 400, "message": "抱歉未找到相关信息"}"
// @Router /system/menu/menuRole [get]
// @Security X-TOKEN
func (m *MenuApi) GetMenuRole(rc *ctx.ReqCtx) {
	rc.ResData = Build(*m.MenuApp.SelectMenuRole(rc.LoginAccount.RoleKey))
}

// @Summary 获取角色的菜单树
// @Description 获取JSON
// @Tags 菜单
// @Param roleId path int false "roleId"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": 400, "message": "抱歉未找到相关信息"}"
// @Router /system/menu/menuTreRoleSelect/{roleId} [get]
// @Security X-TOKEN
func (m *MenuApi) GetMenuTreeRoleSelect(rc *ctx.ReqCtx) {
	roleId := ginx.PathParamInt(rc.GinCtx, "roleId")

	result := m.MenuApp.SelectMenuLable(entity2.SysMenu{})
	menuIds := make([]int64, 0)
	if roleId != 0 {
		menuIds = m.RoleApp.GetRoleMeunId(entity2.SysRole{RoleId: int64(roleId)})
	}
	rc.ResData = map[string]interface{}{
		"menus":       result,
		"checkedKeys": menuIds,
	}
}

// @Summary 获取角色对应的菜单路径数组
// @Description 获取JSON
// @Tags 菜单
// @Param roleKey query string true "roleKey"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": 400, "message": "抱歉未找到相关信息"}"
// @Router /system/menu/menuPaths [get]
// @Security X-TOKEN
func (m *MenuApi) GetMenuPaths(rc *ctx.ReqCtx) {
	roleKey := rc.GinCtx.Query("roleKey")
	biz.IsTrue(roleKey != "", "请传入角色Key")
	rc.ResData = m.RoleMenuApp.GetMenuPaths(entity2.SysRoleMenu{RoleName: roleKey})
}

// @Summary Menu列表数据
// @Description 获取JSON
// @Tags 菜单
// @Param menuName query string false "menuName"
// @Param IsHide query string false "IsHide"
// @Param title query string false "title"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": 400, "message": "抱歉未找到相关信息"}"
// @Router /system/menu/menuList [get]
// @Security Bearer
func (m *MenuApi) GetMenuList(rc *ctx.ReqCtx) {
	menuName := rc.GinCtx.Query("menuName")
	status := rc.GinCtx.Query("status")

	menu := entity2.SysMenu{MenuName: menuName, Status: status}
	log.Println(menuName)
	if menu.MenuName == "" {
		rc.ResData = m.MenuApp.SelectMenu(menu)
	} else {
		rc.ResData = m.MenuApp.FindList(menu)
	}
}

// @Summary Menu列表数据
// @Description 获取JSON
// @Tags 菜单
// @Param menuId path string true "menuId"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": 400, "message": "抱歉未找到相关信息"}"
// @Router /system/menu/{menuId} [get]
// @Security Bearer
func (m *MenuApi) GetMenu(rc *ctx.ReqCtx) {
	menuId := ginx.PathParamInt(rc.GinCtx, "menuId")

	rc.ResData = m.MenuApp.FindOne(int64(menuId))
}

// @Summary 创建菜单
// @Description 获取JSON
// @Tags 菜单
// @Param data body entity.SysMenu true "data"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /system/menu [post]
// @Security X-TOKEN
func (m *MenuApi) InsertMenu(rc *ctx.ReqCtx) {
	var menu entity2.SysMenu
	ginx.BindJsonAndValid(rc.GinCtx, &menu)
	menu.CreateBy = rc.LoginAccount.UserName
	m.MenuApp.Insert(menu)
	permis := m.RoleMenuApp.GetPermis(rc.LoginAccount.RoleId)
	menus := m.MenuApp.SelectMenuRole(rc.LoginAccount.RoleKey)
	rc.ResData = map[string]interface{}{
		"permissions": permis,
		"menus":       Build(*menus),
	}
}

// @Summary 修改菜单
// @Description 获取JSON
// @Tags 菜单
// @Param data body entity.SysMenu true "body"
// @Success 200 {string} string	"{"code": 200, "message": "修改成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "修改失败"}"
// @Router /system/menu [put]
// @Security X-TOKEN
func (m *MenuApi) UpdateMenu(rc *ctx.ReqCtx) {
	var menu entity2.SysMenu
	ginx.BindJsonAndValid(rc.GinCtx, &menu)
	menu.UpdateBy = rc.LoginAccount.UserName
	m.MenuApp.Update(menu)
	permis := m.RoleMenuApp.GetPermis(rc.LoginAccount.RoleId)
	menus := m.MenuApp.SelectMenuRole(rc.LoginAccount.RoleKey)
	rc.ResData = map[string]interface{}{
		"permissions": permis,
		"menus":       Build(*menus),
	}
}

// @Summary 删除菜单
// @Description 删除数据
// @Tags 菜单
// @Param menuId path int true "menuId"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "删除失败"}"
// @Router /system/menu/{menuId} [delete]
func (m *MenuApi) DeleteMenu(rc *ctx.ReqCtx) {
	menuId := ginx.PathParamInt(rc.GinCtx, "menuId")
	m.MenuApp.Delete([]int64{int64(menuId)})
}
