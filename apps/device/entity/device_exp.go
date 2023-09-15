package entity

import "time"

// 说明
// 设备上报属性，遥测，连接事件存储到时序数据库中
// 其他存在MySQL中

// DeviceAlarm 设备告警表 需要更改告警状态不能存在时序数据库中
type DeviceAlarm struct {
	Id        string    `json:"id" gorm:"primary_key;"`
	OrgId     int64     `json:"orgId"  gorm:"type:int;comment:机构ID"`
	Owner     string    `json:"owner"  gorm:"type:varchar(64);comment:创建者,所有者"`
	Time      time.Time `gorm:"comment:告警时间" json:"time"`
	Name      string    `gorm:"type:varchar(64);comment:告警名称" json:"name"`
	DeviceId  string    `gorm:"type:varchar(64);comment:所属设备" json:"deviceId"`
	ProductId string    `gorm:"type:varchar(64);comment:所属产品" json:"productId"`
	Type      string    `gorm:"type:varchar(64);comment:告警类型" json:"type"`
	Level     string    `gorm:"type:varchar(64);comment:告警级别" json:"level"` // 危险 重要 次要 警告 不确定
	State     string    `gorm:"type:varchar(1);comment:告警状态" json:"state"`  // 告警中 0 已确认 1 已清除 2 已关闭 3
	Details   string    `gorm:"type:varchar(255);comment:告警详情" json:"details"`

	RoleId int64 `gorm:"-"` // 角色数据权限
}

type DeviceAlarmForm struct {
	DeviceAlarm
	StartTime string `gorm:"-" json:"startTime"`
	EndTime   string `gorm:"-" json:"endTime"`
}

type DeviceCmdLog struct {
	Id              string `json:"id" gorm:"primary_key;"`
	DeviceId        string `gorm:"type:varchar(64);comment:所属设备" json:"deviceId"`
	CmdName         string `gorm:"type:varchar(64);comment:命令名称" json:"cmdName"`
	CmdContent      string `gorm:"type:longtext;comment:命令内容" json:"cmdContent"`
	State           string `gorm:"type:varchar(1);comment:命令状态" json:"state"`
	Type            string `gorm:"type:varchar(1);comment:命令类型" json:"type"` // 0 自定义 1 命令
	ResponseContent string `gorm:"type:longtext;comment:响应内容" json:"responseContent"`
	RequestTime     string `gorm:"comment:命令下发时间" json:"requestTime"`
	ResponseTime    string `gorm:"comment:命令响应时间" json:"responseTime"`
}

func (DeviceCmdLog) TableName() string {
	return "device_cmd_log"
}

// DeviceStatistics 设备统计表
// 每小时、每天、每周、每月的平均值、最大值、最小值等。可以使用定时任务或实时计算框架来更新数据统计表。
type DeviceStatistics struct {
	Time time.Time `gorm:"comment:时间" json:"time"`
}
