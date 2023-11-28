package hook_message_work

import (
	"context"
	"encoding/json"
	"github.com/golang-queue/queue"
	"github.com/golang-queue/queue/core"
	"pandax/iothub/netbase"
	"pandax/pkg/global"
	"sync"
)

type HookService struct {
	Cache     sync.Map
	Queue     *queue.Queue
	Ch        chan struct{}  //  并发限制
	Wg        sync.WaitGroup //  优雅关闭
	MessageCh chan *netbase.DeviceEventInfo
}

func NewHookService() *HookService {
	hs := &HookService{
		Cache:     sync.Map{},
		Ch:        make(chan struct{}, global.Conf.Queue.ChNum),
		MessageCh: make(chan *netbase.DeviceEventInfo, global.Conf.Queue.TaskNum),
	}
	pool := queue.NewPool(int(global.Conf.Queue.QueuePool), queue.WithFn(func(ctx context.Context, m core.QueuedMessage) error {
		v, ok := m.(*netbase.DeviceEventInfo)
		if !ok {
			if err := json.Unmarshal(m.Bytes(), &v); err != nil {
				return err
			}
		}
		hs.MessageCh <- v
		return nil
	}))
	hs.Queue = pool
	return hs
}
