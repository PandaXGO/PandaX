package initialize

import (
	"pandax/apps/job/jobs"
	ruleRouter "pandax/apps/rule/router"
	"pandax/pkg/global"
	"pandax/pkg/transport"

	devRouter "pandax/apps/develop/router"
	deviceRouter "pandax/apps/device/router"
	jobRouter "pandax/apps/job/router"
	logRouter "pandax/apps/log/router"
	sysRouter "pandax/apps/system/router"
	"pandax/pkg/middleware"
)

func InitRouter() *transport.HttpServer {
	// server配置
	serverConfig := global.Conf.Server
	server := transport.NewHttpServer(serverConfig.GetPort())

	container := server.Container
	// 防止XSS
	container.Filter(middleware.EscapeHTML)
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
	//日志系统
	{
		logRouter.InitOperLogRouter(container)
		logRouter.InitLoginLogRouter(container)
	}
	// 代码生成
	{
		devRouter.InitGenTableRouter(container)
		devRouter.InitGenRouter(container)
	}
	//设备
	{
		deviceRouter.InitProductCategoryRouter(container)
		deviceRouter.InitDeviceGroupRouter(container)
		deviceRouter.InitProductRouter(container)
		deviceRouter.InitProductTemplateRouter(container)
		deviceRouter.InitProductOtaRouter(container)
		deviceRouter.InitDeviceRouter(container)
		deviceRouter.InitDeviceAlarmRouter(container)
		deviceRouter.InitDeviceCmdLogRouter(container)
	}
	{
		jobRouter.InitJobRouter(container)
		jobRouter.InitJobLogRouter(container)
	}
	{
		ruleRouter.InitRuleChainRouter(container)
	}
	// api接口
	middleware.SwaggerConfig(container)
	// 开启调度任务
	go func() {
		jobs.InitJob()
	}()

	return server
}
