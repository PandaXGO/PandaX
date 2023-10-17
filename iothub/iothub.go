package iothub

import (
	"fmt"
	"pandax/iothub/hook_message_work"
	"pandax/iothub/server/emqxserver"
	"pandax/iothub/server/httpserver"
	"pandax/iothub/server/tcpserver"
	updserver "pandax/iothub/server/udpserver"
	"pandax/pkg/global"
)

func InitIothub() {
	service := hook_message_work.NewHookService()
	// 初始化EMQX
	go emqxserver.InitEmqxHook(fmt.Sprintf(":%d", global.Conf.Server.GrpcPort), service)
	// 初始化HTTP
	go httpserver.InitHttpHook(fmt.Sprintf(":%d", global.Conf.Server.HttpPort), service)
	//初始化TCP
	go tcpserver.InitTcpHook(fmt.Sprintf(":%d", global.Conf.Server.TcpPort), service)

	go updserver.InitUdpHook(fmt.Sprintf(":%d", global.Conf.Server.TcpPort), service)
	// 开启线程处理消息
	go service.MessageWork()
}
