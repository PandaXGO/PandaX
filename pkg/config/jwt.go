package config

import (
	"github.com/XM-GO/PandaKit/biz"
)

type Jwt struct {
	Key        string `yaml:"key"`
	ExpireTime int64  `yaml:"expire-time"` // 过期时间，单位分钟
}

func (j *Jwt) Valid() {
	biz.IsTrue(j.Key != "", "config.yml之 [jwt.key] 不能为空")
	biz.IsTrue(j.ExpireTime != 0, "config.yml之 [jwt.expire-time] 不能为空")
}
