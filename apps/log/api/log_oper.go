package api

import (
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	"github.com/XM-GO/PandaKit/utils"
	"log"
	"pandax/apps/log/entity"
	"pandax/apps/log/services"
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
		PageSize: int64(pageNum),
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
	log.Println("group", group)
	l.LogOperApp.Delete(group)
}

func (l *LogOperApi) DeleteAll(rc *restfulx.ReqCtx) {
	l.LogOperApp.DeleteAll()
}
