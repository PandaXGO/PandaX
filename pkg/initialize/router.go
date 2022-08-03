package initialize

import (
	"pandax/pkg/global"
	"pandax/pkg/transport"

	sysRouter "pandax/apps/system/router"

	"pandax/pkg/middleware"

	_ "pandax/docs"
)

func InitRouter() *transport.HttpServer {
	// server配置
	serverConfig := global.Conf.Server
	server := transport.NewHttpServer(serverConfig.GetPort())

	container := server.Container
	// 是否允许跨域
	if serverConfig.Cors {
		container.Filter(middleware.Cors(container).Filter)
	}
	// 流量限制
	if serverConfig.Rate.Enable {
		container.Filter(middleware.Rate)
	}
	// 设置路由组
	{
		sysRouter.InitSysTenantRouter(container)
		sysRouter.InitSystemRouter(container)
		sysRouter.InitDeptRouter(container)
		sysRouter.InitConfigRouter(container)
		sysRouter.InitApiRouter(container)
		sysRouter.InitDictRouter(container)
		sysRouter.InitMenuRouter(container)
		sysRouter.InitRoleRouter(container)
		sysRouter.InitPostRouter(container)
		sysRouter.InitUserRouter(container)
		sysRouter.InitNoticeRouter(container)
	}
	/*// 任务
	{
		jobRouter.InitJobRouter()
	}
	//日志系统
	log := router.Group("log")
	{
		logRouter.InitLogRouter(log)
	}
	// 代码生成
	dev := router.Group("develop/code")
	{
		devRouter.InitGenTableRouter(dev)
		devRouter.InitGenRouter(dev)
	}
	// 资源管理
	res := router.Group("resource")
	{
		resRouter.InitResOssRouter(res)
		resRouter.InitResEmailsRouter(res)
	}*/
	// api接口
	middleware.SwaggerConfig(container)
	return server
}
