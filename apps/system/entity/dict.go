package entity

import "github.com/PandaXGO/PandaKit/model"

type SysDictData struct {
	DictCode  int64  `json:"dictCode" gorm:"primary_key;AUTO_INCREMENT"`
	DictSort  int    `json:"dictSort" gorm:"type:int;comment:排序"`
	DictLabel string `json:"dictLabel" gorm:"type:varchar(64);comment:标签"`
	DictValue string `json:"dictValue" gorm:"type:varchar(64);comment:值"`
	DictType  string `json:"dictType" gorm:"type:varchar(64);comment:字典类型"`
	Status    string `json:"status" gorm:"type:varchar(1);comment:状态（0正常 1停用）"`
	CssClass  string `json:"cssClass" gorm:"type:varchar(128);comment:CssClass"`
	ListClass string `json:"listClass" gorm:"type:varchar(128);comment:ListClass"`
	IsDefault string `json:"isDefault" gorm:"type:varchar(8);comment:IsDefault"`
	CreateBy  string `json:"createBy"`
	UpdateBy  string `json:"updateBy"`
	Remark    string `json:"remark" gorm:"type:varchar(256);comment:备注"`
	model.BaseModel
}

type SysDictType struct {
	DictId   int64  `json:"dictId" gorm:"primary_key;AUTO_INCREMENT"`
	DictName string `json:"dictName" gorm:"type:varchar(64);comment:名称"`
	DictType string `json:"dictType" gorm:"type:varchar(64);comment:类型"`
	Status   string `json:"status" gorm:"type:varchar(1);comment:状态"`
	CreateBy string `json:"createBy"`
	UpdateBy string `json:"updateBy"`
	Remark   string `json:"remark" gorm:"type:varchar(256);comment:备注"`
	model.BaseModel
}
