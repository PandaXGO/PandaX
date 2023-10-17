package global

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"pandax/pkg/config"
	"pandax/pkg/events"
	"pandax/pkg/tdengine"
)

var (
	Log  *logrus.Logger // 日志
	Db   *gorm.DB       // gorm
	TdDb *tdengine.TdEngine
	Conf *config.Config
)
var EventEmitter = events.EventEmitter{}
