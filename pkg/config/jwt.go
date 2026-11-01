package config

import (
	"github.com/PandaXGO/PandaKit/biz"
	"strings"
)

type Jwt struct {
	Key        string `yaml:"key"`
	ExpireTime int64  `yaml:"expire-time"` // 过期时间，单位分钟
}

func (j *Jwt) Valid() {
	key := strings.TrimSpace(j.Key)
	biz.IsTrue(len(key) >= 32 && key != "PandaX", "config.yml之 [jwt.key] 必须是至少32字节的随机密钥")
	biz.IsTrue(j.ExpireTime != 0, "config.yml之 [jwt.expire-time] 不能为空")
}
