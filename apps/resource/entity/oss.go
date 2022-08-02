package entity

import "github.com/XM-GO/PandaKit/model"

/**
 * @Description 添加qq群467890197 交流学习
 * @Author 熊猫
 * @Date 2022/1/13 14:47
 **/

type ResOss struct {
	OssId      int64  `json:"ossId" gorm:"primaryKey;AUTO_INCREMENT;comment:主键编码"`
	Category   string `json:"category" grom:"type:varchar(1);comment:种类"` // 0 阿里云 1.七牛云 2. 腾讯云
	AppId      string `json:"appId" gorm:"type:varchar(128);comment:AppId"`
	AccessKey  string `json:"accessKey" gorm:"type:varchar(128);comment:accessKey"`
	SecretKey  string `json:"secretKey" gorm:"type:varchar(128);comment:secretKey"`
	BucketName string `json:"bucketName" gorm:"type:varchar(128);comment:bucketName"`
	Endpoint   string `json:"endpoint" gorm:"type:varchar(128);comment:endpoint"`
	OssCode    string `json:"ossCode" gorm:"type:varchar(128);comment:ossCode"`
	Region     string `json:"region" gorm:"type:varchar(128);comment:地址"`
	Remark     string `json:"remark" gorm:"type:varchar(128);comment:说明"`
	Status     string `json:"status" gorm:"type:varchar(1);comment:状态"` // 0.启用 1.禁止
	model.BaseModel
}
