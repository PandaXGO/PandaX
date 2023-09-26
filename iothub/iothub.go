package iothub

import (
	"pandax/iothub/hook_message_work"
	"pandax/iothub/server/emqxserver"
	"pandax/iothub/server/httpserver"
	"pandax/iothub/server/tcpserver"
)

func InitIothub() {
	service := hook_message_work.NewHookService()
	// 初始化EMQX
	go emqxserver.InitEmqxHook("", service)
	// 初始化HTTP
	go httpserver.InitHttpHook("", service)
	//初始化TCP
	go tcpserver.InitTcpHook("", service)
	// 开启线程处理消息
	go service.MessageWork()
}
