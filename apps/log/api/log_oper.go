package api

import (
	"pandax/apps/log/entity"
	"pandax/apps/log/services"
	"pandax/kit/model"
	"pandax/kit/restfulx"
	"pandax/kit/utils"
)

type LogOperApi struct {
	LogOperApp services.LogOperModel
}

func (l *LogOperApi) GetOperLogList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	businessType := restfulx.QueryParam(rc, "businessType")
	operName := restfulx.QueryParam(rc, "operName")
	title := restfulx.QueryParam(rc, "title")
	list, total := l.LogOperApp.FindListPage(pageNum, pageSize, entity.LogOper{BusinessType: businessType, OperName: operName, Title: title})
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

func (l *LogOperApi) GetOperLog(rc *restfulx.ReqCtx) {
	operId := restfulx.PathParamInt(rc, "operId")
	rc.ResData = l.LogOperApp.FindOne(int64(operId))
}

func (l *LogOperApi) DeleteOperLog(rc *restfulx.ReqCtx) {
	operIds := restfulx.PathParam(rc, "operId")
	group := utils.IdsStrToIdsIntGroup(operIds)
	l.LogOperApp.Delete(group)
}

func (l *LogOperApi) DeleteAll(rc *restfulx.ReqCtx) {
	l.LogOperApp.DeleteAll()
}
