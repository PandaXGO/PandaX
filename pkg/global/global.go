package global

import (
	"github.com/PandaXGO/PandaKit/rediscli"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"pandax/pkg/config"
	"pandax/pkg/events"
	"pandax/pkg/mqtt"
	"pandax/pkg/tdengine"
)

var (
	Log        *logrus.Logger // 日志
	Db         *gorm.DB       // gorm
	RedisDb    *rediscli.RedisDB
	TdDb       *tdengine.TdEngine
	Conf       *config.Config
	MqttClient *mqtt.IothubMqttClient
)
var EventEmitter = events.EventEmitter{}
