package api

import (
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	"github.com/XM-GO/PandaKit/utils"
	"pandax/apps/log/entity"
	"pandax/apps/log/services"
)

type LogJobApi struct {
	LogJobApp services.LogJobModel
}

// GetJobLogList Job日志列表
func (l *LogJobApi) GetJobLogList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	name := restfulx.QueryParam(rc, "name")
	jobGroup := restfulx.QueryParam(rc, "jobGroup")
	status := restfulx.QueryParam(rc, "status")

	list, total := l.LogJobApp.FindListPage(pageNum, pageSize, entity.LogJob{Name: name, JobGroup: jobGroup, Status: status})
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageNum),
		Data:     list,
	}
}

// DeleteJobLog 批量删除登录日志
func (l *LogJobApi) DeleteJobLog(rc *restfulx.ReqCtx) {
	logIds := restfulx.QueryParam(rc, "logId")
	group := utils.IdsStrToIdsIntGroup(logIds)
	l.LogJobApp.Delete(group)
}

// DeleteAll 清空登录日志
func (l *LogJobApi) DeleteAll(rc *restfulx.ReqCtx) {
	l.LogJobApp.DeleteAll()
}
