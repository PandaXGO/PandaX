package api

import (
	"github.com/XM-GO/PandaKit/restfulx"
	"github.com/XM-GO/PandaKit/utils"
	"log"
	"pandax/apps/log/entity"
	"pandax/apps/log/services"
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
// @Router /log/logOper/list [get]
// @Security
func (l *LogOperApi) GetOperLogList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	businessType := restfulx.QueryParam(rc, "businessType")
	operName := restfulx.QueryParam(rc, "operName")
	title := restfulx.QueryParam(rc, "title")
	list, total := l.LogOperApp.FindListPage(pageNum, pageSize, entity.LogOper{BusinessType: businessType, OperName: operName, Title: title})
	rc.ResData = map[string]any{
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
// @Router /log/logOper/{operId} [get]
// @Security
func (l *LogOperApi) GetOperLog(rc *restfulx.ReqCtx) {
	operId := restfulx.PathParamInt(rc, "operId")
	rc.ResData = l.LogOperApp.FindOne(int64(operId))
}

// @Summary 批量删除操作日志
// @Description 删除数据
// @Tags 操作日志
// @Param operId path string true "以逗号（,）分割的operId"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "删除失败"}"
// @Router /log/logOper/{operId} [delete]
func (l *LogOperApi) DeleteOperLog(rc *restfulx.ReqCtx) {
	operIds := restfulx.PathParam(rc, "operId")
	group := utils.IdsStrToIdsIntGroup(operIds)
	log.Println("group", group)
	l.LogOperApp.Delete(group)
}

// @Summary 清空操作日志
// @Description 删除数据
// @Tags 操作日志
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "删除失败"}"
// @Router /log/logOper/all [delete]
func (l *LogOperApi) DeleteAll(rc *restfulx.ReqCtx) {
	l.LogOperApp.DeleteAll()
}
