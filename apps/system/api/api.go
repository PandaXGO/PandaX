package api

import (
	"log"
	entity "pandax/apps/system/entity"
	services "pandax/apps/system/services"
	"pandax/base/casbin"
	"pandax/base/ginx"
	"pandax/base/utils"
	"strconv"
)

type SystemApiApi struct {
	ApiApp services.SysApiModel
}

// @Tags SysApi
// @Summary 创建基础api
// @Security X-TOKEN
// @accept application/json
// @Produce application/json
// @Param data body entity.SysApi true "api路径, api中文描述, api组, 方法"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /system/api [post]
func (s *SystemApiApi) CreateApi(rc *ginx.ReqCtx) {
	var api entity.SysApi
	ginx.BindJsonAndValid(rc.GinCtx, &api)
	log.Println(api)
	s.ApiApp.Insert(api)
}

// @Tags SysApi
// @Summary 删除api
// @Security X-TOKEN
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /system/api/{id} [delete]
func (s *SystemApiApi) DeleteApi(rc *ginx.ReqCtx) {
	ids := rc.GinCtx.Param("id")
	s.ApiApp.Delete(utils.IdsStrToIdsIntGroup(ids))
}

// @Tags SysApi
// @Summary 分页获取API列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param path query string false "path"
// @Param description query string false "description"
// @Param method query string false "method"
// @Param apiGroup query string false "apiGroup"
// @Param pageSize query int false "页条数"
// @Param pageNum query int false "页码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /system/api/list [get]
func (s *SystemApiApi) GetApiList(rc *ginx.ReqCtx) {
	pageNum := ginx.QueryInt(rc.GinCtx, "pageNum", 1)
	pageSize := ginx.QueryInt(rc.GinCtx, "pageSize", 10)
	path := rc.GinCtx.Query("path")
	description := rc.GinCtx.Query("description")
	method := rc.GinCtx.Query("method")
	apiGroup := rc.GinCtx.Query("apiGroup")
	api := entity.SysApi{Path: path, Description: description, Method: method, ApiGroup: apiGroup}
	list, total := s.ApiApp.FindListPage(pageNum, pageSize, api)
	rc.ResData = map[string]any{
		"data":     list,
		"total":    total,
		"pageNum":  pageNum,
		"pageSize": pageSize,
	}
}

// @Tags SysApi
// @Summary 根据id获取api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id path int true "根据id获取api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /system/api/{id} [get]
func (s *SystemApiApi) GetApiById(rc *ginx.ReqCtx) {
	id := ginx.QueryInt(rc.GinCtx, "id", 0)
	rc.ResData = s.ApiApp.FindOne(int64(id))

}

// @Tags SysApi
// @Summary 创建基础api
// @Security X-TOKEN
// @accept application/json
// @Produce application/json
// @Param data body entity.SysApi true "api路径, api中文描述, api组, 方法"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /api/api [put]
func (s *SystemApiApi) UpdateApi(rc *ginx.ReqCtx) {
	var api entity.SysApi
	ginx.BindJsonAndValid(rc.GinCtx, &api)
	s.ApiApp.Update(api)
}

// @Tags SysApi
// @Summary 获取所有的Api 不分页
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /system/api/all [get]
func (s *SystemApiApi) GetAllApis(rc *ginx.ReqCtx) {
	rc.ResData = s.ApiApp.FindList(entity.SysApi{})
}

// @Tags SysApi
// @Summary 获取Api权限列表
// @Security X-TOKEN
// @accept application/json
// @Produce application/json
// @Param roleKey query string true "权限id, 权限模型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /casbin/getPolicyPathByRoleId [get]
func (s *SystemApiApi) GetPolicyPathByRoleId(rc *ginx.ReqCtx) {
	roleKey := rc.GinCtx.Query("roleKey")
	tenantId := strconv.Itoa(int(rc.LoginAccount.TenantId))
	rc.ResData = casbin.GetPolicyPathByRoleId(tenantId, roleKey)
}
