package api

import (
	"github.com/XM-GO/PandaKit/biz"
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	"github.com/kakuilan/kgo"
	"log"
	"pandax/apps/job/api/from"
	"pandax/apps/job/entity"
	"pandax/apps/job/jobs"
	"pandax/apps/job/services"
	"strings"
)

type JobApi struct {
	JobApp services.JobModel
}

func (j *JobApi) CreateJob(rc *restfulx.ReqCtx) {
	var job entity.SysJob
	restfulx.BindQuery(rc, &job)
	job.Id = kgo.KStr.Uniqid("")
	job.Owner = rc.LoginAccount.UserName
	j.JobApp.Insert(job)
}

func (j *JobApi) GetJobList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	jobName := restfulx.QueryParam(rc, "jobName")
	status := restfulx.QueryParam(rc, "status")

	list, total := j.JobApp.FindListPage(pageNum, pageSize, entity.SysJob{JobName: jobName, Status: status})
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageNum),
		Data:     list,
	}
}

func (j *JobApi) GetJob(rc *restfulx.ReqCtx) {
	jobId := restfulx.PathParam(rc, "jobId")
	rc.ResData = j.JobApp.FindOne(jobId)
}

func (l *JobApi) UpdateJob(rc *restfulx.ReqCtx) {
	var job entity.SysJob
	restfulx.BindQuery(rc, &job)
	l.JobApp.Update(job)
}

func (l *JobApi) DeleteJob(rc *restfulx.ReqCtx) {
	jobIds := restfulx.PathParam(rc, "jobId")
	group := strings.Split(jobIds, ",")
	l.JobApp.Delete(group)
}

func (l *JobApi) StopJobForService(rc *restfulx.ReqCtx) {
	jobId := restfulx.PathParam(rc, "jobId")
	job := l.JobApp.FindOne(jobId)
	jobs.Remove(jobs.Crontab, job.EntryId)
}

func (l *JobApi) StartJobForService(rc *restfulx.ReqCtx) {
	jobId := restfulx.PathParam(rc, "jobId")
	job := l.JobApp.FindOne(jobId)

	biz.IsTrue(job.Status == "0", "以关闭的任务不能开启")
	biz.IsTrue(job.EntryId == 0, "任务不能重复启动")

	var err error
	var j = &jobs.ExecJob{}
	j.InvokeTarget = job.TargetInvoke
	j.CronExpression = job.CronExpression
	j.JobId = job.Id
	j.Name = job.JobName
	j.Args = job.TargetArgs
	j.MisfirePolicy = job.MisfirePolicy
	job.EntryId, err = jobs.AddJob(jobs.Crontab, j)
	log.Println(err)
	biz.ErrIsNil(err, "添加JOB失败")

	l.JobApp.Update(*job)
}

func (l *JobApi) UpdateStatusJob(rc *restfulx.ReqCtx) {
	var job from.JobStatus
	restfulx.BindQuery(rc, &job)
	sjob := entity.SysJob{}
	sjob.Id = job.JobId
	sjob.Status = job.Status
	l.JobApp.Update(sjob)
}
