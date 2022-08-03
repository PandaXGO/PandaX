package api

import (
	"github.com/XM-GO/PandaKit/biz"
	"github.com/XM-GO/PandaKit/restfulx"
	"github.com/XM-GO/PandaKit/utils"
	"pandax/apps/job/api/from"
	"pandax/apps/job/entity"
	"pandax/apps/job/jobs"
	"pandax/apps/job/services"
)

type JobApi struct {
	JobApp services.JobModel
}

func (j *JobApi) CreateJob(rc *restfulx.ReqCtx) {
	var job entity.SysJob
	restfulx.BindQuery(rc, &job)

	job.CreateBy = rc.LoginAccount.UserName
	j.JobApp.Insert(job)
}

func (j *JobApi) GetJobList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	jobName := restfulx.QueryParam(rc, "jobName")
	jobGroup := restfulx.QueryParam(rc, "jobGroup")
	status := restfulx.QueryParam(rc, "status")

	list, total := j.JobApp.FindListPage(pageNum, pageSize, entity.SysJob{JobName: jobName, JobGroup: jobGroup, Status: status})
	rc.ResData = map[string]any{
		"data":     list,
		"total":    total,
		"pageNum":  pageNum,
		"pageSize": pageSize,
	}
}

func (j *JobApi) GetJob(rc *restfulx.ReqCtx) {
	jobId := restfulx.PathParamInt(rc, "jobId")
	rc.ResData = j.JobApp.FindOne(int64(jobId))
}

func (l *JobApi) UpdateJob(rc *restfulx.ReqCtx) {
	var job entity.SysJob
	restfulx.BindQuery(rc, &job)
	l.JobApp.Update(job)
}

func (l *JobApi) DeleteJob(rc *restfulx.ReqCtx) {
	jobIds := restfulx.PathParam(rc, "jobId")
	group := utils.IdsStrToIdsIntGroup(jobIds)
	l.JobApp.Delete(group)
}

func (l *JobApi) StopJobForService(rc *restfulx.ReqCtx) {
	jobId := restfulx.PathParamInt(rc, "jobId")
	job := l.JobApp.FindOne(int64(jobId))
	jobs.Remove(jobs.Crontab, job.EntryId)
}

func (l *JobApi) StartJobForService(rc *restfulx.ReqCtx) {
	jobId := restfulx.PathParamInt(rc, "jobId")
	job := l.JobApp.FindOne(int64(jobId))

	biz.IsTrue(job.Status == "0", "以关闭的任务不能开启")
	biz.IsTrue(job.EntryId == 0, "任务不能重复启动")

	var err error
	if job.JobType == "1" {
		var j = &jobs.HttpJob{}
		j.InvokeTarget = job.InvokeTarget
		j.CronExpression = job.CronExpression
		j.JobId = job.JobId
		j.Name = job.JobName
		j.JobGroup = job.JobGroup
		j.MisfirePolicy = job.MisfirePolicy
		job.EntryId, err = jobs.AddJob(jobs.Crontab, j)
		biz.ErrIsNil(err, "添加JOB失败")
	} else {
		var j = &jobs.ExecJob{}
		j.InvokeTarget = job.InvokeTarget
		j.CronExpression = job.CronExpression
		j.JobId = job.JobId
		j.Name = job.JobName
		j.JobGroup = job.JobGroup
		j.Args = job.Args
		j.MisfirePolicy = job.MisfirePolicy
		job.EntryId, err = jobs.AddJob(jobs.Crontab, j)
		biz.ErrIsNil(err, "添加JOB失败")
	}
	l.JobApp.Update(*job)
}

func (l *JobApi) UpdateStatusJob(rc *restfulx.ReqCtx) {
	var job from.JobStatus
	restfulx.BindQuery(rc, &job)

	l.JobApp.Update(entity.SysJob{JobId: job.JobId, Status: job.Status})
}
