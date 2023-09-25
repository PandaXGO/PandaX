package iothub

import (
	"pandax/iothub/hook_message_work"
	"pandax/iothub/server/emqxserver"
	"pandax/iothub/server/httpserver"
)

func InitIothub() {
	service := hook_message_work.NewHookService()
	// 初始化EMQX
	emqxserver.InitEmqxHook("", service)
	// 初始化HTTP
	httpserver.InitHttpHook("", service)
	//初始化TCP

	// 开启线程处理消息
	go service.MessageWork()
}
