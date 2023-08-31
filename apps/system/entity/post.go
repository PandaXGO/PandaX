package entity

import "github.com/PandaXGO/PandaKit/model"

type SysPost struct {
	PostId   int64  `gorm:"primary_key;AUTO_INCREMENT" json:"postId"`
	PostName string `gorm:"type:varchar(128);comment:岗位名称" json:"postName"`
	PostCode string `gorm:"type:varchar(128);comment:岗位代码" json:"postCode"`
	Sort     int64  `gorm:"type:int;comment:岗位排序" json:"sort"`
	Status   string `gorm:"type:varchar(1);comment:状态" json:"status"`
	Remark   string `gorm:"type:varchar(255);comment:描述" json:"remark"`
	CreateBy string `gorm:"type:varchar(128);" json:"createBy"`
	UpdateBy string `gorm:"type:varchar(128);" json:"updateBy"`
	model.BaseModel
}
