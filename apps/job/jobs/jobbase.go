package jobs

import (
	"fmt"
	"github.com/kakuilan/kgo"
	"pandax/apps/job/entity"
	"pandax/apps/job/services"
	"pandax/pkg/global"

	logEntity "pandax/apps/job/entity"
	logServices "pandax/apps/job/services"

	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

var timeFormat = "2006-01-02 15:04:05"
var retryCount = 3

var jobList map[string]JobsExec
var lock sync.Mutex

var Crontab = new(cron.Cron)

func InitJob() {
	jobList = map[string]JobsExec{
		"cronDevice":  CronDeviceHandle{},
		"cronProduct": CronProductHandle{},
	}
	// 启动调度任务
	Setup()
}

type JobCore struct {
	InvokeTarget   string
	Name           string
	JobId          string
	EntryId        int
	CronExpression string // 任务表达式
	MisfirePolicy  string
	Args           string
	Content        string
}

type ExecJob struct {
	cron *cron.Cron
	JobCore
}

func (e *ExecJob) Run() {
	startTime := time.Now()
	jobLog := logEntity.JobLog{Name: e.Name, EntryId: e.EntryId, TargetInvoke: e.InvokeTarget, Status: "0"}
	jobLog.Id = kgo.KStr.Uniqid("")
	var obj = jobList[e.InvokeTarget]

	err := CallExec(obj.(JobsExec), e.Args, e.Content)
	if err != nil {
		jobLog.LogInfo = fmt.Sprintf("任务运行错误: %s", err.Error())
		Remove(e.cron, e.EntryId)
	} else {
		latencyTime := time.Now().Sub(startTime)
		jobLog.LogInfo = fmt.Sprintf("任务运行成功，总耗时 %f", latencyTime.Seconds())
	}
	// 执行时间
	logServices.JobLogModelDao.Insert(jobLog)
	// 执行一次
	if e.MisfirePolicy == "1" {
		Remove(e.cron, e.EntryId)
	}
	return
}

func Setup() {
	Crontab = NewWithSeconds()
	err := services.JobModelDao.RemoveAllEntryID()
	if err != nil {
		global.Log.Info(time.Now().Format(timeFormat), " [ERROR] JobCore remove entry_id error", err)
	}
	// 其中任务
	Crontab.Start()
	global.Log.Info(time.Now().Format(timeFormat), " [INFO] JobCore start success.")
	// 关闭任务
	defer Crontab.Stop()
	select {}
}

// AddJob 添加任务
func AddJob(c *cron.Cron, job Job) (int, error) {
	if job == nil {
		return 0, nil
	}
	return job.addJob(c)
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
