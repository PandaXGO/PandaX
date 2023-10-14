package api

import (
	"github.com/PandaXGO/PandaKit/biz"
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	"log"
	"pandax/apps/job/api/from"
	"pandax/apps/job/entity"
	"pandax/apps/job/jobs"
	"pandax/apps/job/services"
	"pandax/pkg/global_model"
	"strings"
)

type JobApi struct {
	JobApp services.JobModel
}

func (j *JobApi) CreateJob(rc *restfulx.ReqCtx) {
	var job entity.SysJob
	restfulx.BindQuery(rc, &job)
	job.Id = global_model.GenerateID()
	job.Owner = rc.LoginAccount.UserName
	job.OrgId = rc.LoginAccount.OrganizationId
	j.JobApp.Insert(job)
}

func (j *JobApi) GetJobList(rc *restfulx.ReqCtx) {
	job := entity.SysJob{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	job.JobName = restfulx.QueryParam(rc, "jobName")
	job.Status = restfulx.QueryParam(rc, "status")

	job.RoleId = rc.LoginAccount.RoleId
	job.Owner = rc.LoginAccount.UserName

	list, total := j.JobApp.FindListPage(pageNum, pageSize, job)
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
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
	j.OrgId = job.OrgId
	j.Owner = job.Owner
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
