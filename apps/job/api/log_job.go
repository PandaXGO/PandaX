package api

import (
	"pandax/apps/job/entity"
	"pandax/apps/job/services"
	"pandax/kit/model"
	"pandax/kit/restfulx"
	"strings"
)

type JobLogApi struct {
	JobLogApp services.JobLogModel
}

// GetJobLogList Job日志列表
func (l *JobLogApi) GetJobLogList(rc *restfulx.ReqCtx) {
	job := entity.JobLog{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	job.Name = restfulx.QueryParam(rc, "name")
	job.Status = restfulx.QueryParam(rc, "status")

	job.RoleId = rc.LoginAccount.RoleId
	job.Owner = rc.LoginAccount.UserName

	list, total := l.JobLogApp.FindListPage(pageNum, pageSize, job)
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

// DeleteJobLog 批量删除登录日志
func (l *JobLogApi) DeleteJobLog(rc *restfulx.ReqCtx) {
	logIds := restfulx.PathParam(rc, "id")
	group := strings.Split(logIds, ",")
	l.JobLogApp.Delete(group)
}

// DeleteAll 清空登录日志
func (l *JobLogApi) DeleteAll(rc *restfulx.ReqCtx) {
	l.JobLogApp.DeleteAll()
}
