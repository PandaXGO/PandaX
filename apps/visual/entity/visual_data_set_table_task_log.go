// ==========================================================================
// 生成日期：2023-04-10 11:31:34 +0800 CST
// 生成路径: apps/visual/entity/visual_data_set_table_task_log.go
// 生成人：panda
// ==========================================================================
package entity

import "github.com/XM-GO/PandaKit/model"

type VisualDataSetTableTaskLog struct {
	TaskId     string `gorm:"task_id;type:varchar(50);comment:任务ID" json:"taskId" `                  // 任务ID
	Id         string `gorm:"primary_key;" json:"id"`                                                // ID
	Status     string `gorm:"status;type:varchar(50);comment:执行状态" json:"status" binding:"required"` // 执行状态
	CreateTime int64  `gorm:"create_time;type:bigint;comment:创建时间" json:"createTime" `               // 创建时间
	Info       string `gorm:"info;type:text;comment:错误信息" json:"info" `                              // 错误信息
	EndTime    int64  `gorm:"end_time;type:bigint;comment:结束时间" json:"endTime" `                     // 结束时间
	TableId    string `gorm:"table_id;type:varchar(50);comment:表ID" json:"tableId" `                 // 表ID
	model.BaseModelD
}

func (VisualDataSetTableTaskLog) TableName() string {
	return "visual_data_set_table_task_log"
}
