package initialize

import (
	"fmt"

	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/swaggo/gin-swagger/swaggerFiles"

	devRouter "pandax/apps/develop/router"
	jobRouter "pandax/apps/job/router"
	logRouter "pandax/apps/log/router"
	resRouter "pandax/apps/resource/router"
	sysRouter "pandax/apps/system/router"

	"pandax/base/config"
	"pandax/middleware"

	_ "pandax/docs"

	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// server配置
	serverConfig := config.Conf.Server
	gin.SetMode(serverConfig.Model)

	var router = gin.New()
	router.MaxMultipartMemory = 8 << 20

	// 没有路由即 404返回
	router.NoRoute(func(g *gin.Context) {
		g.JSON(http.StatusNotFound, gin.H{"code": 404, "msg": fmt.Sprintf("not found '%s:%s'", g.Request.Method, g.Request.URL.Path)})
	})

	// 设置静态资源
	if staticConfs := serverConfig.Static; staticConfs != nil {
		for _, scs := range *staticConfs {
			router.Static(scs.RelativePath, scs.Root)
		}

	}
	// 设置静态文件
	if staticFileConfs := serverConfig.StaticFile; staticFileConfs != nil {
		for _, sfs := range *staticFileConfs {
			router.StaticFile(sfs.RelativePath, sfs.Filepath)
		}
	}
	// 是否允许跨域
	if serverConfig.Cors {
		router.Use(middleware.Cors())
	}
	// 流量限制
	if serverConfig.Rate.Enable {
		router.Use(middleware.Rate())
	}
	// api接口
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 设置路由组
	sys := router.Group("system")
	{
		sysRouter.InitSysTenantRouter(sys)
		sysRouter.InitSystemRouter(sys)
		sysRouter.InitDeptRouter(sys)
		sysRouter.InitConfigRouter(sys)
		sysRouter.InitApiRouter(sys)
		sysRouter.InitDictRouter(sys)
		sysRouter.InitMenuRouter(sys)
		sysRouter.InitRoleRouter(sys)
		sysRouter.InitPostRouter(sys)
		sysRouter.InitUserRouter(sys)
		sysRouter.InitNoticeRouter(sys)
	}
	// 任务
	job := router.Group("job")
	{
		jobRouter.InitJobRouter(job)
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
	}
	return router
}
