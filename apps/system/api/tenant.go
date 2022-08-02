package api

/**
 * @Description
 * @Author 熊猫
 * @Date 2022/7/14 17:55
 **/
import (
	"pandax/apps/system/entity"
	"pandax/apps/system/services"
	"pandax/base/ginx"
	"pandax/base/utils"
)

type SysTenantsApi struct {
	SysTenantsApp services.SysTenantsModel
}

// @Summary SysTenants列表数据
// @Tags SysTenants
// @Param pageSize query int false "页条数"
// @Param pageNum query int false "页码"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Router /system/tenant/list [get]
// @Security
func (p *SysTenantsApi) GetSysTenantsList(rc *ginx.ReqCtx) {
	data := entity.SysTenants{}
	pageNum := ginx.QueryInt(rc.GinCtx, "pageNum", 1)
	pageSize := ginx.QueryInt(rc.GinCtx, "pageSize", 10)

	list, total := p.SysTenantsApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = map[string]interface{}{
		"data":     list,
		"total":    total,
		"pageNum":  pageNum,
		"pageSize": pageSize,
	}
}

// @Summary SysTenants列表数据
// @Tags SysTenants
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Router /system/tenant/lists [get]
// @Security
func (p *SysTenantsApi) GetSysTenantsAll(rc *ginx.ReqCtx) {
	list := make([]entity.SysTenants, 0)
	if rc.LoginAccount.RoleKey == "admin" {
		data := entity.SysTenants{}
		list = *p.SysTenantsApp.FindList(data)
	} else {
		list = append(list, *p.SysTenantsApp.FindOne(rc.LoginAccount.TenantId))
	}
	rc.ResData = list
}

// @Summary 获取SysTenants
// @Description 获取JSON
// @Tags SysTenants
// @Param tenantId path int true "tenantId"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Router /system/tenant/{tenantId} [get]
// @Security
func (p *SysTenantsApi) GetSysTenants(rc *ginx.ReqCtx) {
	tenantId := ginx.PathParamInt(rc.GinCtx, "tenantId")
	p.SysTenantsApp.FindOne(int64(tenantId))
}

// @Summary 添加SysTenants
// @Description 获取JSON
// @Tags SysTenants
// @Accept  application/json
// @Product application/json
// @Param data body entity.SysTenants true "data"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "添加失败"}"
// @Router /system/tenant [post]
// @Security X-TOKEN
func (p *SysTenantsApi) InsertSysTenants(rc *ginx.ReqCtx) {
	var data entity.SysTenants
	ginx.BindJsonAndValid(rc.GinCtx, &data)

	p.SysTenantsApp.Insert(data)
}

// @Summary 修改SysTenants
// @Description 获取JSON
// @Tags SysTenants
// @Accept  application/json
// @Product application/json
// @Param data body entity.SysTenants true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "添加失败"}"
// @Router /system/tenant [put]
// @Security X-TOKEN
func (p *SysTenantsApi) UpdateSysTenants(rc *ginx.ReqCtx) {
	var data entity.SysTenants
	ginx.BindJsonAndValid(rc.GinCtx, &data)

	p.SysTenantsApp.Update(data)
}

// @Summary 删除SysTenants
// @Description 删除数据
// @Tags SysTenants
// @Param tenantId path string true "tenantId"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "删除失败"}"
// @Router /system/tenant/{tenantId} [delete]
func (p *SysTenantsApi) DeleteSysTenants(rc *ginx.ReqCtx) {
	tenantId := rc.GinCtx.Param("tenantId")
	tenantIds := utils.IdsStrToIdsIntGroup(tenantId)
	p.SysTenantsApp.Delete(tenantIds)
}

// IsTenantAdmin 是否为主租户
func IsTenantAdmin(tenantId int64) bool {
	if tenantId == 1 {
		return true
	}
	return false
}
