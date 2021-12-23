package api

import (
	"log"
	"pandax/apps/log/entity"
	"pandax/apps/log/services"
	"pandax/base/ctx"
	"pandax/base/ginx"
	"pandax/base/utils"
)

type LogOperApi struct {
	LogOperApp services.LogOperModel
}

// @Summary 操作日志列表
// @Description 获取JSON
// @Tags 操作日志
// @Param status query string false "status"
// @Param operName query string false "operName"
// @Param pageSize query int false "页条数"
// @Param pageNum query int false "页码"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Router /system/log/list [get]
// @Security
func (l *LogOperApi) GetOperLogList(rc *ctx.ReqCtx) {
	pageNum := ginx.QueryInt(rc.GinCtx, "pageNum", 1)
	pageSize := ginx.QueryInt(rc.GinCtx, "pageSize", 10)

	businessType := rc.GinCtx.Query("businessType")
	operName := rc.GinCtx.Query("operName")
	title := rc.GinCtx.Query("title")
	list, total := l.LogOperApp.FindListPage(pageNum, pageSize, entity.LogOper{BusinessType: businessType, OperName: operName, Title: title})
	rc.ResData = map[string]interface{}{
		"data":     list,
		"total":    total,
		"pageNum":  pageNum,
		"pageSize": pageSize,
	}
}

// @Summary 通过编码获取操作日志
// @Description 获取JSON
// @Tags 操作日志
// @Param operId path int true "operId"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Router /system/log/logOper/{operId} [get]
// @Security
func (l *LogOperApi) GetOperLog(rc *ctx.ReqCtx) {
	operId := ginx.PathParamInt(rc.GinCtx, rc.GinCtx.Param("operId"))
	rc.ResData = l.LogOperApp.FindOne(int64(operId))
}

// @Summary 批量删除操作日志
// @Description 删除数据
// @Tags 操作日志
// @Param operId path string true "以逗号（,）分割的operId"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "删除失败"}"
// @Router /system/log/logOper/{operId} [delete]
func (l *LogOperApi) DeleteOperLog(rc *ctx.ReqCtx) {
	operIds := rc.GinCtx.Param("operId")
	group := utils.IdsStrToIdsIntGroup(operIds)
	log.Println("group", group)
	l.LogOperApp.Delete(group)
}

// @Summary 清空操作日志
// @Description 删除数据
// @Tags 操作日志
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "删除失败"}"
// @Router /system/log/logOper/all [delete]
func (l *LogOperApi) DeleteAll(rc *ctx.ReqCtx) {
	l.LogOperApp.DeleteAll()
}
