package api

import (
	"github.com/XM-GO/PandaKit/restfulx"
	"github.com/XM-GO/PandaKit/utils"
	"pandax/apps/log/entity"
	"pandax/apps/log/services"
)

type LogLoginApi struct {
	LogLoginApp services.LogLoginModel
}

// @Summary 登录日志列表
// @Description 获取JSON
// @Tags 登录日志
// @Param status query string false "status"
// @Param username query string false "username"
// @Param pageSize query int false "页条数"
// @Param pageNum query int false "页码"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Router /log/logLogin/list [get]
// @Security
func (l *LogLoginApi) GetLoginLogList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	loginLocation := restfulx.QueryParam(rc, "loginLocation")
	username := restfulx.QueryParam(rc, "username")
	list, total := l.LogLoginApp.FindListPage(pageNum, pageSize, entity.LogLogin{LoginLocation: loginLocation, Username: username})
	rc.ResData = map[string]any{
		"data":     list,
		"total":    total,
		"pageNum":  pageNum,
		"pageSize": pageSize,
	}
}

// @Summary 通过编码获取登录日志
// @Description 获取JSON
// @Tags 登录日志
// @Param infoId path int true "infoId"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Router /log/logLogin/{infoId} [get]
// @Security
func (l *LogLoginApi) GetLoginLog(rc *restfulx.ReqCtx) {
	infoId := restfulx.PathParamInt(rc, "infoId")
	rc.ResData = l.LogLoginApp.FindOne(int64(infoId))
}

// @Summary 修改登录日志
// @Description 获取JSON
// @Tags 登录日志
// @Accept  application/json
// @Product application/json
// @Param data body entity.LogLogin true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "添加失败"}"
// @Router /log/logLogin [put]
// @Security X-TOKEN
func (l *LogLoginApi) UpdateLoginLog(rc *restfulx.ReqCtx) {
	var log entity.LogLogin
	restfulx.BindQuery(rc, &log)
	l.LogLoginApp.Update(log)
}

// @Summary 批量删除登录日志
// @Description 删除数据
// @Tags 登录日志
// @Param infoId path string true "以逗号（,）分割的infoId"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "删除失败"}"
// @Router /log/logLogin/{infoId} [delete]
func (l *LogLoginApi) DeleteLoginLog(rc *restfulx.ReqCtx) {
	infoIds := restfulx.PathParam(rc, "infoId")
	group := utils.IdsStrToIdsIntGroup(infoIds)
	l.LogLoginApp.Delete(group)
}

// @Summary 清空登录日志
// @Description 删除数据
// @Tags 登录日志
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "删除失败"}"
// @Router /log/logLogin/all [delete]
func (l *LogLoginApi) DeleteAll(rc *restfulx.ReqCtx) {
	l.LogLoginApp.DeleteAll()
}
