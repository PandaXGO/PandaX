package entity

import (
	"github.com/PandaXGO/PandaKit/model"
)

type LogOper struct {
	OperId       int64  `json:"operId" gorm:"primary_key;AUTO_INCREMENT"` //主键
	Title        string `json:"title" gorm:"type:varchar(128);comment:操作的模块"`
	BusinessType string `json:"businessType" gorm:"type:varchar(1);comment:0其它 1新增 2修改 3删除"`
	Method       string `json:"method" gorm:"type:varchar(255);comment:请求方法"`
	OperName     string `json:"operName" gorm:"type:varchar(255);comment:操作人员"`
	OperUrl      string `json:"operUrl" gorm:"type:varchar(255);comment:操作url"`
	OperIp       string `json:"operIp" gorm:"type:varchar(255);comment:操作IP"`
	OperLocation string `json:"operLocation" gorm:"type:varchar(255);comment:操作地点"`
	OperParam    string `json:"operParam" gorm:"type:varchar(255);comment:请求参数"` //
	Status       string `json:"status" gorm:"type:varchar(1);comment:0=正常,1=异常"`
	model.BaseModel
}
