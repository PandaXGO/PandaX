package api

import (
	"github.com/PandaXGO/PandaKit/biz"
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	"pandax/apps/job/entity"
	"pandax/apps/job/services"
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

	list, total, err := l.JobLogApp.FindListPage(pageNum, pageSize, job)
	biz.ErrIsNil(err, "查询任务列表失败")
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
	err := l.JobLogApp.Delete(group)
	biz.ErrIsNil(err, "删除失败")
}

// DeleteAll 清空登录日志
func (l *JobLogApi) DeleteAll(rc *restfulx.ReqCtx) {
	err := l.JobLogApp.DeleteAll()
	biz.ErrIsNil(err, "清空失败")
}
