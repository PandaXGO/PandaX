package global

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"pandax/pkg/config"
)

var (
	Log  *logrus.Logger // 日志
	Db   *gorm.DB       // gorm
	Conf *config.Config
)
