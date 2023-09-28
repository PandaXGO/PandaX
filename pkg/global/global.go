package global

import (
	"github.com/PandaXGO/PandaKit/cache"
	"github.com/PandaXGO/PandaKit/rediscli"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"pandax/pkg/config"
	"pandax/pkg/events"
	"pandax/pkg/tdengine"
	"time"
)

var (
	Log     *logrus.Logger // 日志
	Db      *gorm.DB       // gorm
	RedisDb *rediscli.RedisDB
	TdDb    *tdengine.TdEngine
	Conf    *config.Config
)
var EventEmitter = events.EventEmitter{}

// Cache 默认10分钟
var ProductCache = cache.NewTimedCache(cache.NoExpiration, 24*time.Hour)
var SubDeviceCache = cache.NewTimedCache(cache.NoExpiration, 24*time.Hour)
var PanelCache = cache.NewTimedCache(cache.NoExpiration, 600*time.Second)
