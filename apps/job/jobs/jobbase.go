package jobs

import (
	"fmt"
	"pandax/apps/job/entity"
	"pandax/apps/job/services"

	logEntity "pandax/apps/log/entity"
	logServices "pandax/apps/log/services"

	"github.com/XM-GO/PandaKit/httpclient"
	"pandax/pkg/global"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

var timeFormat = "2006-01-02 15:04:05"
var retryCount = 3

var jobList map[string]JobsExec
var lock sync.Mutex

var Crontab = new(cron.Cron)

type JobCore struct {
	InvokeTarget   string
	Name           string
	JobGroup       string
	JobId          int64
	EntryId        int
	CronExpression string // 任务表达式
	MisfirePolicy  string // 错误执行策略
	Args           string
}

// 任务类型 http
type HttpJob struct {
	cron *cron.Cron
	JobCore
}

type ExecJob struct {
	cron *cron.Cron
	JobCore
}

func (e *ExecJob) Run() {
	startTime := time.Now()
	var obj = jobList[e.InvokeTarget]
	if obj == nil {
		if e.MisfirePolicy == "2" {
			Remove(e.cron, e.EntryId)
		}
		return
	}
	err := CallExec(obj.(JobsExec), e.Args)
	if err != nil {
		if e.MisfirePolicy == "2" {
			Remove(e.cron, e.EntryId)
			return
		}
	}
	// 执行时间
	latencyTime := time.Now().Sub(startTime)

	logInfo := fmt.Sprintf("任务运行总耗时 %f", latencyTime.Seconds())
	logServices.LogJobModelDao.Insert(logEntity.LogJob{Name: e.Name, JobGroup: e.JobGroup, EntryId: e.EntryId, InvokeTarget: e.InvokeTarget, LogInfo: logInfo, Status: "0"})
	// 执行一次
	if e.MisfirePolicy == "1" {
		Remove(e.cron, e.EntryId)
	}
	return
}

//http 任务接口
func (h *HttpJob) Run() {

	startTime := time.Now()
	var count = 0
LOOP:
	if count < retryCount {
		_, err := httpclient.NewRequest(h.InvokeTarget).Get().BodyToString()
		if err != nil {
			time.Sleep(time.Duration(count+1) * 5 * time.Second)
			count = count + 1
			goto LOOP
		}
	}
	// 执行时间
	latencyTime := time.Now().Sub(startTime)

	logInfo := fmt.Sprintf("任务运行总耗时 %f", latencyTime.Seconds())
	logServices.LogJobModelDao.Insert(logEntity.LogJob{Name: h.Name, JobGroup: h.JobGroup, EntryId: h.EntryId, InvokeTarget: h.InvokeTarget, LogInfo: logInfo, Status: "0"})
	if h.MisfirePolicy == "1" {
		Remove(h.cron, h.EntryId)
	}
	return
}

func Setup() {
	Crontab = NewWithSeconds()
	// 获取系统job SYSTEM是系统
	jl := services.JobModelDao.FindList(entity.SysJob{JobGroup: "SYSTEM"})
	jobList := *jl
	if len(jobList) == 0 {
		global.Log.Info(time.Now().Format(timeFormat), " [INFO] JobCore total:0")
		return
	}
	err := services.JobModelDao.RemoveAllEntryID()
	if err != nil {
		global.Log.Info(time.Now().Format(timeFormat), " [ERROR] JobCore remove entry_id error", err)
	}
	sysJob := entity.SysJob{}
	for i := 0; i < len(jobList); i++ {
		//去除禁用的
		if jobList[i].Status != "0" {
			continue
		}
		if jobList[i].JobType == "1" {
			j := &HttpJob{}
			j.InvokeTarget = jobList[i].InvokeTarget
			j.CronExpression = jobList[i].CronExpression
			j.JobId = jobList[i].JobId
			j.Name = jobList[i].JobName
			j.JobGroup = jobList[i].JobGroup
			j.MisfirePolicy = jobList[i].MisfirePolicy
			sysJob.EntryId, err = AddJob(Crontab, j)
		} else if jobList[i].JobType == "2" {
			j := &ExecJob{}
			j.InvokeTarget = jobList[i].InvokeTarget
			j.CronExpression = jobList[i].CronExpression
			j.JobId = jobList[i].JobId
			j.Name = jobList[i].JobName
			j.JobGroup = jobList[i].JobGroup
			j.Args = jobList[i].Args
			j.MisfirePolicy = jobList[i].MisfirePolicy
			sysJob.EntryId, err = AddJob(Crontab, j)
		}
		sysJob.JobId = jobList[i].JobId
		services.JobModelDao.Update(sysJob)
	}
	// 其中任务
	Crontab.Start()
	global.Log.Info(time.Now().Format(timeFormat), " [INFO] JobCore start success.")
	// 关闭任务
	defer Crontab.Stop()
	select {}
}

// 添加任务 AddJob(invokeTarget string, jobId int, jobName string, cronExpression string)
func AddJob(c *cron.Cron, job Job) (int, error) {
	if job == nil {
		return 0, nil
	}
	return job.addJob(c)
}

func (h *HttpJob) addJob(c *cron.Cron) (int, error) {
	id, err := c.AddJob(h.CronExpression, h)
	if err != nil {
		return 0, err
	}
	h.cron = c
	EntryId := int(id)
	h.EntryId = EntryId
	return EntryId, nil
}

func (h *ExecJob) addJob(c *cron.Cron) (int, error) {
	id, err := c.AddJob(h.CronExpression, h)
	if err != nil {
		return 0, err
	}
	h.cron = c
	EntryId := int(id)
	h.EntryId = EntryId
	return EntryId, nil
}

// 移除任务(停止任务)
func Remove(c *cron.Cron, entryID int) {
	c.Remove(cron.EntryID(entryID))
	var job entity.SysJob
	job.EntryId = entryID
	services.JobModelDao.RemoveEntryID(entryID)
}

func NewWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}
