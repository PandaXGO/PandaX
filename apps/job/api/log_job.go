package api

import (
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	"github.com/PandaXGO/PandaKit/utils"
	"pandax/apps/job/entity"
	"pandax/apps/job/services"
)

type JobLogApi struct {
	JobLogApp services.JobLogModel
}

// GetJobLogList Job日志列表
func (l *JobLogApi) GetJobLogList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	name := restfulx.QueryParam(rc, "name")
	status := restfulx.QueryParam(rc, "status")

	list, total := l.JobLogApp.FindListPage(pageNum, pageSize, entity.JobLog{Name: name, Status: status})
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageNum),
		Data:     list,
	}
}

// DeleteJobLog 批量删除登录日志
func (l *JobLogApi) DeleteJobLog(rc *restfulx.ReqCtx) {
	logIds := restfulx.QueryParam(rc, "logId")
	group := utils.IdsStrToIdsIntGroup(logIds)
	l.JobLogApp.Delete(group)
}

// DeleteAll 清空登录日志
func (l *JobLogApi) DeleteAll(rc *restfulx.ReqCtx) {
	l.JobLogApp.DeleteAll()
}
