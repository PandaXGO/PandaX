package entity

import "github.com/PandaXGO/PandaKit/model"

type SysNotice struct {
	NoticeId       int64  `json:"noticeId" gorm:"primary_key;AUTO_INCREMENT"`
	Title          string `json:"title" gorm:"type:varchar(128);comment:标题"`
	Content        string `json:"content" gorm:"type:text;comment:标题"`
	NoticeType     string `json:"noticeType" gorm:"type:varchar(1);comment:通知类型"`
	OrganizationId int64  `json:"organizationId" gorm:"type:int;comment:组织Id,组织及子组织"`
	UserName       string `json:"userName" gorm:"type:varchar(64);comment:发布人"`

	OrganizationIds []int64 `json:"organizationIds" gorm:"-"`
	model.BaseModel
}
