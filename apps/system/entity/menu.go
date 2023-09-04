package entity

import "github.com/PandaXGO/PandaKit/model"

type SysMenu struct {
	MenuId      int64     `json:"menuId" gorm:"primary_key;AUTO_INCREMENT"`
	MenuName    string    `json:"menuName" gorm:"type:varchar(128);"`
	Title       string    `json:"title" gorm:"type:varchar(64);"`
	ParentId    int64     `json:"parentId" gorm:"type:int;"`
	Sort        int64     `json:"sort" gorm:"type:int;"`
	Icon        string    `json:"icon" gorm:"type:varchar(128);"`
	Path        string    `json:"path" gorm:"type:varchar(128);"`
	Component   string    `json:"component" gorm:"type:varchar(255);"` // 组件路径
	IsIframe    string    `json:"isIframe" gorm:"type:varchar(1);"`    //是否为内嵌
	IsLink      string    `json:"isLink" gorm:"type:varchar(255);"`    //是否超链接菜单
	MenuType    string    `json:"menuType" gorm:"type:varchar(1);"`    //菜单类型（M目录 C菜单 F按钮）
	IsHide      string    `json:"isHide" gorm:"type:varchar(1);"`      //显示状态（0显示 1隐藏）
	IsKeepAlive string    `json:"isKeepAlive" gorm:"type:varchar(1);"` //是否缓存组件状态（0是 1否）
	IsAffix     string    `json:"isAffix" gorm:"type:varchar(1);"`     //是否固定在 tagsView 栏上（0是 1否）
	Permission  string    `json:"permission" gorm:"type:varchar(32);"` //权限标识
	Status      string    `json:"status" gorm:"type:varchar(1);`       // 菜单状态（0正常 1停用）
	CreateBy    string    `json:"createBy" gorm:"type:varchar(128);"`
	UpdateBy    string    `json:"updateBy" gorm:"type:varchar(128);"`
	Remark      string    `json:"remark"  gorm:"type:varchar(256);` // 备注
	Children    []SysMenu `json:"children" gorm:"-"`
	model.BaseModel
}

type MenuLable struct {
	MenuId   int64       `json:"menuId" gorm:"-"`
	MenuName string      `json:"menuName" gorm:"-"`
	Children []MenuLable `json:"children" gorm:"-"`
}

type MenuRole struct {
	SysMenu
	IsSelect bool `json:"is_select" gorm:"-"`
}
