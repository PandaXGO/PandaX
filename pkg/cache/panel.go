package cache

import (
	"github.com/PandaXGO/PandaKit/cache"
	"time"
)

var PanelCache = cache.NewTimedCache(cache.NoExpiration, 600*time.Second)
