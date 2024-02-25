package api

import (
	"pandax/apps/log/entity"
	"pandax/apps/log/services"
	"pandax/kit/biz"
	"pandax/kit/model"
	"pandax/kit/restfulx"
	"pandax/kit/utils"
)

type LogLoginApi struct {
	LogLoginApp services.LogLoginModel
}

func (l *LogLoginApi) GetLoginLogList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	loginLocation := restfulx.QueryParam(rc, "loginLocation")
	username := restfulx.QueryParam(rc, "username")
	list, total, err := l.LogLoginApp.FindListPage(pageNum, pageSize, entity.LogLogin{LoginLocation: loginLocation, Username: username})
	biz.ErrIsNil(err, "查询任务日志失败")
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

func (l *LogLoginApi) GetLoginLog(rc *restfulx.ReqCtx) {
	infoId := restfulx.PathParamInt(rc, "infoId")
	data, err := l.LogLoginApp.FindOne(int64(infoId))
	biz.ErrIsNil(err, "查询日志信息失败")
	rc.ResData = data
}

func (l *LogLoginApi) UpdateLoginLog(rc *restfulx.ReqCtx) {
	var log entity.LogLogin
	restfulx.BindQuery(rc, &log)
	_, err := l.LogLoginApp.Update(log)
	biz.ErrIsNil(err, "修改日志信息失败")
}

func (l *LogLoginApi) DeleteLoginLog(rc *restfulx.ReqCtx) {
	infoIds := restfulx.PathParam(rc, "infoId")
	group := utils.IdsStrToIdsIntGroup(infoIds)
	err := l.LogLoginApp.Delete(group)
	biz.ErrIsNil(err, "删除日志失败")
}

func (l *LogLoginApi) DeleteAll(rc *restfulx.ReqCtx) {
	err := l.LogLoginApp.DeleteAll()
	biz.ErrIsNil(err, "清空日志失败")
}
