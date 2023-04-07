package entity

import "github.com/XM-GO/PandaKit/model"

type VisualDataSource struct {
	model.BaseModel
	SourceId      string `gorm:"source_id;comment:数据源Id"   json:"sourceId"`                                       // 数据源Id
	SourceType    string `gorm:"source_type;type:varchar(50);comment:数据源类型"         json:"sourceType"`            // 数据源类型
	SourceName    string `gorm:"source_name;type:varchar(50);comment:数据源名称"      json:"sourceName"`               // 原名称
	SourceComment string `gorm:"source_comment;type:varchar(50);comment:数据源描述"             json:"source_comment"` // 描述
	Status        string `gorm:"status;type:varchar(1);comment:数据源状态" json:"status"`
	Configuration string `gorm:"configuration;type:text;comment:详细信息" json:"configuration"`
	CreateBy      int64  `gorm:"api" json:"createBy"` //创建人ID

}

type VisualDb struct {
	DbIp         string `gorm:"db_ip" json:"dbIp"`
	DbPort       string `gorm:"db_port" json:"dbPort"`
	DbName       string `gorm:"db_name" json:"dbName"`
	DbUsername   string `gorm:"db_username" json:"dbUsername"`
	DbPassword   string `gorm:"db_password" json:"dbPassword"`
	DbJointParam string `gorm:"db_joint_param" json:"dbJointParam"` //额外的链接参数
}

type VisualApi struct {
	Method      string                 `gorm:"method" json:"method"`
	url         string                 `gorm:"url" json:"url"`
	Headers     map[string]interface{} `gorm:"headers" json:"headers"`
	RequestBody string                 `gorm:"db_username" json:"dbUsername"`
	Auth        string                 `gorm:"db_password" json:"dbPassword"`
}

func (VisualDataSource) TableName() string {
	return "visual_data_source"
}
