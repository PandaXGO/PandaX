package initialize

import (
	"pandax/pkg/global"
	"pandax/pkg/transport"

	devRouter "pandax/apps/develop/router"
	flowRouter "pandax/apps/flow/router"
	jobRouter "pandax/apps/job/router"
	logRouter "pandax/apps/log/router"
	resRouter "pandax/apps/resource/router"
	sysRouter "pandax/apps/system/router"
	visualRouter "pandax/apps/visual/router"

	"pandax/pkg/middleware"
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
		//本地图片上传接口
		sysRouter.InitUploadRouter(container)
	}
	//流程管理
	{
		flowRouter.InitFlowWorkClassifyRouter(container)
		flowRouter.InitFlowWorkInfoRouter(container)
		flowRouter.InitFlowWorkTemplatesRouter(container)
	}
	// 可视化
	{
		visualRouter.InitUploadRouter(container)
		visualRouter.InitRuleChainRouter(container)
		visualRouter.InitVisualScreenGroupRouter(container)
		visualRouter.InitVisualScreenRouter(container)
	}
	// 任务
	{
		jobRouter.InitJobRouter(container)
	}
	//日志系统
	{
		logRouter.InitJobLogRouter(container)
		logRouter.InitOperLogRouter(container)
		logRouter.InitLoginLogRouter(container)
	}
	// 代码生成
	{
		devRouter.InitGenTableRouter(container)
		devRouter.InitGenRouter(container)
	}
	// 资源管理
	{
		resRouter.InitResOssRouter(container)
		resRouter.InitResEmailsRouter(container)
	}

	// api接口
	middleware.SwaggerConfig(container)
	return server
}
