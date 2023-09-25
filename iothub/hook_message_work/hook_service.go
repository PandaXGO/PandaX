package hook_message_work

import (
	"pandax/iothub/netbase"
	"pandax/pkg/global"
	"sync"
)

type HookService struct {
	Cache     sync.Map
	Wg        sync.WaitGroup //  优雅关闭
	Ch        chan struct{}  //  并发限制
	MessageCh chan *netbase.DeviceEventInfo
}

func NewHookService() *HookService {
	hs := &HookService{
		Cache:     sync.Map{},
		MessageCh: make(chan *netbase.DeviceEventInfo),
	}
	// 并发限制，代表服务器处理能力
	if global.Conf.Queue.Enable && global.Conf.Queue.Num > 0 {
		hs.Ch = make(chan struct{}, global.Conf.Queue.Num)
	}
	return hs
}
