package api

import (
	"pandax/apps/job/api/from"
	"pandax/apps/job/entity"
	"pandax/apps/job/jobs"
	"pandax/apps/job/services"
	"pandax/base/biz"
	"pandax/base/ginx"
	"pandax/base/utils"
)

type JobApi struct {
	JobApp services.JobModel
}

// @Summary 添加任务
// @Description 获取JSON
// @Tags 任务
// @Accept  application/json
// @Product application/json
// @Param data body entity.SysJob true "data"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "添加失败"}"
// @Router /job/job [post]
// @Security X-TOKEN
func (j *JobApi) CreateJob(rc *ginx.ReqCtx) {
	var job entity.SysJob
	ginx.BindJsonAndValid(rc.GinCtx, &job)

	job.CreateBy = rc.LoginAccount.UserName
	j.JobApp.Insert(job)
}

// @Summary job列表
// @Description 获取JSON
// @Tags 任务
// @Param status query string false "status"
// @Param jobName query string false "jobName"
// @Param jobGroup query string false "jobGroup"
// @Param pageSize query int false "页条数"
// @Param pageNum query int false "页码"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Router /job/list [get]
// @Security
func (j *JobApi) GetJobList(rc *ginx.ReqCtx) {
	pageNum := ginx.QueryInt(rc.GinCtx, "pageNum", 1)
	pageSize := ginx.QueryInt(rc.GinCtx, "pageSize", 10)
	jobName := rc.GinCtx.Query("jobName")
	jobGroup := rc.GinCtx.Query("jobGroup")
	status := rc.GinCtx.Query("status")

	list, total := j.JobApp.FindListPage(pageNum, pageSize, entity.SysJob{JobName: jobName, JobGroup: jobGroup, Status: status})
	rc.ResData = map[string]any{
		"data":     list,
		"total":    total,
		"pageNum":  pageNum,
		"pageSize": pageSize,
	}
}

// @Summary 获取一个job
// @Description 获取JSON
// @Tags 任务
// @Param jobId path int true "jobId"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Router /job/{jobId} [get]
// @Security
func (j *JobApi) GetJob(rc *ginx.ReqCtx) {
	jobId := ginx.PathParamInt(rc.GinCtx, rc.GinCtx.Param("jobId"))
	rc.ResData = j.JobApp.FindOne(int64(jobId))
}

// @Summary 修改JOB
// @Description 获取JSON
// @Tags 任务
// @Accept  application/json
// @Product application/json
// @Param data body entity.SysJob true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "添加失败"}"
// @Router /job [put]
// @Security X-TOKEN
func (l *JobApi) UpdateJob(rc *ginx.ReqCtx) {
	var job entity.SysJob
	ginx.BindJsonAndValid(rc.GinCtx, &job)
	l.JobApp.Update(job)
}

// @Summary 批量删除JOB
// @Description 删除数据
// @Tags 任务
// @Param infoId path string true "以逗号（,）分割的infoId"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "删除失败"}"
// @Router /job/{jobId} [delete]
func (l *JobApi) DeleteJob(rc *ginx.ReqCtx) {
	jobIds := rc.GinCtx.Param("jobId")
	group := utils.IdsStrToIdsIntGroup(jobIds)
	l.JobApp.Delete(group)
}

// @Summary 停止JOB
// @Description 停止Job
// @Tags 任务
// @Param jobId path int true "jobId"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "删除失败"}"
// @Router /job/stop/{jobId} [get]
func (l *JobApi) StopJobForService(rc *ginx.ReqCtx) {
	jobId := ginx.PathParamInt(rc.GinCtx, "jobId")
	job := l.JobApp.FindOne(int64(jobId))
	jobs.Remove(jobs.Crontab, job.EntryId)
}

// @Summary 开始JOB
// @Description 开始Job
// @Tags 任务
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "删除失败"}"
// @Router /job/stop/{jobId} [get]
func (l *JobApi) StartJobForService(rc *ginx.ReqCtx) {
	jobId := ginx.PathParamInt(rc.GinCtx, "jobId")
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

// @Summary 修改JOB状态
// @Description 获取JSON
// @Tags 任务
// @Accept  application/json
// @Product application/json
// @Param data body from.JobStatus true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "添加失败"}"
// @Router /job/changeStatus [put]
// @Security X-TOKEN
func (l *JobApi) UpdateStatusJob(rc *ginx.ReqCtx) {
	var job from.JobStatus
	ginx.BindJsonAndValid(rc.GinCtx, &job)

	l.JobApp.Update(entity.SysJob{JobId: job.JobId, Status: job.Status})
}
