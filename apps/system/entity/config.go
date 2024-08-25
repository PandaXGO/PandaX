package entity

import "github.com/PandaXGO/PandaKit/model"

type SysConfig struct {
	ConfigId    int64  `json:"configId" gorm:"primaryKey;AUTO_INCREMENT;comment:主键编码"`
	ConfigName  string `json:"configName" gorm:"type:varchar(128);comment:ConfigName"`
	ConfigKey   string `json:"configKey" gorm:"type:varchar(128);comment:ConfigKey"`
	ConfigValue string `json:"configValue" gorm:"type:varchar(255);comment:ConfigValue"`
	ConfigType  string `json:"configType" gorm:"type:varchar(64);comment:是否系统内置0，1"`
	IsFrontend  string `json:"isFrontend" gorm:"type:varchar(1);comment:是否前台"`
	Remark      string `json:"remark" gorm:"type:varchar(128);comment:Remark"`
	model.BaseModel
}
