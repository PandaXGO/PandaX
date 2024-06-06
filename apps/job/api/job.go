package api

import (
	"pandax/apps/job/api/from"
	"pandax/apps/job/entity"
	"pandax/apps/job/jobs"
	"pandax/apps/job/services"
	"pandax/kit/biz"
	"pandax/kit/model"
	"pandax/kit/restfulx"
	"pandax/kit/utils"
	"strings"
)

type JobApi struct {
	JobApp services.JobModel
}

func (j *JobApi) CreateJob(rc *restfulx.ReqCtx) {
	var job entity.SysJob
	restfulx.BindQuery(rc, &job)
	job.Id = utils.GenerateID()
	job.Owner = rc.LoginAccount.UserName
	job.OrgId = rc.LoginAccount.OrganizationId
	_, err := j.JobApp.Insert(job)
	biz.ErrIsNil(err, "添加任务失败")
}

func (j *JobApi) GetJobList(rc *restfulx.ReqCtx) {
	job := entity.SysJob{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	job.JobName = restfulx.QueryParam(rc, "jobName")
	job.Status = restfulx.QueryParam(rc, "status")

	job.RoleId = rc.LoginAccount.RoleId
	job.Owner = rc.LoginAccount.UserName

	list, total, err := j.JobApp.FindListPage(pageNum, pageSize, job)
	biz.ErrIsNil(err, "查询任务列表失败")
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

func (j *JobApi) GetJob(rc *restfulx.ReqCtx) {
	jobId := restfulx.PathParam(rc, "jobId")
	data, err := j.JobApp.FindOne(jobId)
	biz.ErrIsNil(err, "查询任务失败")
	rc.ResData = data
}

func (l *JobApi) UpdateJob(rc *restfulx.ReqCtx) {
	var job entity.SysJob
	restfulx.BindQuery(rc, &job)
	_, err := l.JobApp.Update(job)
	biz.ErrIsNil(err, "修改任务失败")
}

func (l *JobApi) DeleteJob(rc *restfulx.ReqCtx) {
	jobIds := restfulx.PathParam(rc, "jobId")
	group := strings.Split(jobIds, ",")
	err := l.JobApp.Delete(group)
	biz.ErrIsNil(err, "删除任务失败")
}

func (l *JobApi) StopJobForService(rc *restfulx.ReqCtx) {
	jobId := restfulx.PathParam(rc, "jobId")
	job, err := l.JobApp.FindOne(jobId)
	biz.ErrIsNil(err, "任务不存在")
	jobs.Remove(jobs.Crontab, job.EntryId)
}

func (l *JobApi) StartJobForService(rc *restfulx.ReqCtx) {
	jobId := restfulx.PathParam(rc, "jobId")
	job, err := l.JobApp.FindOne(jobId)
	biz.ErrIsNil(err, "任务不存在")
	biz.IsTrue(job.Status == "0", "以关闭的任务不能开启")
	biz.IsTrue(job.EntryId == 0, "任务不能重复启动")

	var j = &jobs.ExecJob{}
	j.InvokeTarget = job.TargetInvoke
	j.CronExpression = job.CronExpression
	j.JobId = job.Id
	j.Name = job.JobName
	j.Args = job.TargetArgs
	j.MisfirePolicy = job.MisfirePolicy
	j.OrgId = job.OrgId
	j.Owner = job.Owner
	job.EntryId, err = jobs.AddJob(jobs.Crontab, j)
	biz.ErrIsNil(err, "添加任务失败，可能任务表达式错误")

	_, err = l.JobApp.Update(*job)
	biz.ErrIsNil(err, "修改任务失败")
}

func (l *JobApi) UpdateStatusJob(rc *restfulx.ReqCtx) {
	var job from.JobStatus
	restfulx.BindQuery(rc, &job)
	sjob := entity.SysJob{}
	sjob.Id = job.JobId
	sjob.Status = job.Status
	_, err := l.JobApp.Update(sjob)
	biz.ErrIsNil(err, "修改任务状态失败")
}
