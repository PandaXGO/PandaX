package api

import (
	"log"
	entity2 "pandax/apps/system/entity"
	services2 "pandax/apps/system/services"
	"pandax/base/casbin"
	"pandax/base/ctx"
	"pandax/base/ginx"
	"pandax/base/utils"
)

type SystemApiApi struct {
	ApiApp services2.SysApiModel
}

// @Tags SysApi
// @Summary 创建基础api
// @Security X-TOKEN
// @accept application/json
// @Produce application/json
// @Param data body entity.SysApi true "api路径, api中文描述, api组, 方法"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /system/api [post]
func (s *SystemApiApi) CreateApi(rc *ctx.ReqCtx) {
	var api entity2.SysApi
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
func (s *SystemApiApi) DeleteApi(rc *ctx.ReqCtx) {
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
func (s *SystemApiApi) GetApiList(rc *ctx.ReqCtx) {
	pageNum := ginx.QueryInt(rc.GinCtx, "pageNum", 1)
	pageSize := ginx.QueryInt(rc.GinCtx, "pageSize", 10)
	path := rc.GinCtx.Query("path")
	description := rc.GinCtx.Query("description")
	method := rc.GinCtx.Query("method")
	apiGroup := rc.GinCtx.Query("apiGroup")
	api := entity2.SysApi{Path: path, Description: description, Method: method, ApiGroup: apiGroup}
	list, total := s.ApiApp.FindListPage(pageNum, pageSize, api)
	rc.ResData = map[string]interface{}{
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
// @Param data body request.GetById true "根据id获取api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /system/api/{id} [get]
func (s *SystemApiApi) GetApiById(rc *ctx.ReqCtx) {
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
func (s *SystemApiApi) UpdateApi(rc *ctx.ReqCtx) {
	var api entity2.SysApi
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
func (s *SystemApiApi) GetAllApis(rc *ctx.ReqCtx) {
	rc.ResData = s.ApiApp.FindList(entity2.SysApi{})
}

// @Tags SysApi
// @Summary 获取Api权限列表
// @Security X-TOKEN
// @accept application/json
// @Produce application/json
// @Param data body request.CasbinInReceive true "权限id, 权限模型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /casbin/getPolicyPathByRoleId [get]
func (s *SystemApiApi) GetPolicyPathByRoleId(rc *ctx.ReqCtx) {
	roleKey := rc.GinCtx.Query("roleKey")
	rc.ResData = casbin.GetPolicyPathByRoleId(roleKey)
}
