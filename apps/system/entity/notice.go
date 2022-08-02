package entity

import "github.com/XM-GO/PandaKit/model"

type SysNotice struct {
	NoticeId   int64  `json:"noticeId" gorm:"primary_key;AUTO_INCREMENT"`
	Title      string `json:"title" gorm:"type:varchar(128);comment:标题"`
	Content    string `json:"content" gorm:"type:text;comment:标题"`
	NoticeType string `json:"noticeType" gorm:"type:varchar(1);comment:通知类型"`
	DeptId     int64  `json:"deptId" gorm:"type:int;comment:部门Id,部门及子部门"`
	UserName   string `json:"userName" gorm:"type:varchar(64);comment:发布人"`

	DeptIds []int64 `json:"deptIds" gorm:"-"`
	model.BaseModel
}
