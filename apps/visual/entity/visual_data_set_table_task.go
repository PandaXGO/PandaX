// ==========================================================================
// 生成日期：2023-04-10 11:31:34 +0800 CST
// 生成路径: apps/visual/entity/visual_data_set_table_task.go
// 生成人：panda
// ==========================================================================
package entity

import (
	"github.com/XM-GO/PandaKit/model"
)

type VisualDataSetTableTask struct {
	TableId   string `gorm:"table_id;type:varchar(64);comment:表ID" json:"tableId" `              // 表ID
	Id        string `gorm:"primary_key;" json:"id"`                                             // ID
	StartTime int64  `gorm:"start_time;type:bigint;comment:开始时间" json:"startTime" `              // 开始时间
	End       string `gorm:"end;type:varchar(50);comment:结束限制 0 无限制 1 设定结束时间" json:"end" `       // 结束限制 0 无限制 1 设定结束时间
	EndTime   int64  `gorm:"end_time;type:bigint;comment:结束时间" json:"endTime" `                  // 结束时间
	Name      string `gorm:"name;type:varchar(255);comment:任务名称" json:"name" binding:"required"` // 任务名称
	Cron      string `gorm:"cron;type:varchar(255);comment:cron表达式" json:"cron" `                // cron表达式
	Type      string `gorm:"type;type:varchar(50);comment:更新方式：0-全量更新 1-增量更新" json:"type" `      // 更新方式：0-全量更新 1-增量更新
	Rate      string `gorm:"rate;type:varchar(50);comment:执行频率：0 一次性 1 cron" json:"rate" `       // 执行频率：0 一次性 1 cron
	model.BaseModelD
}

func (VisualDataSetTableTask) TableName() string {
	return "visual_data_set_table_task"
}
