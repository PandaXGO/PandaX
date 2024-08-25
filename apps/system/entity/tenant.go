package entity

import (
	"github.com/PandaXGO/PandaKit/model"
	"time"
)

/**
 * @Description
 * @Author 熊猫
 * @Date 2022/7/14 16:14
 **/

type SysTenants struct {
	model.BaseAutoModel
	TenantName string    `json:"tenantName" gorm:"type:varchar(255);comment:租户名"`
	Remark     string    `json:"remark" gorm:"type:varchar(255);comment:备注"`
	ExpireTime time.Time `json:"expireTime" gorm:"comment:过期时间"`
}

func (SysTenants) TableName() string {
	return "sys_tenants"
}
