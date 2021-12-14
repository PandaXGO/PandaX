package initialize

import (
	"fmt"
	"pandax/base/config"
	"pandax/base/middleware"
	sysRouter "pandax/system/router"

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

	// 设置路由组
	sys := router.Group("system")
	{
		sysRouter.InitSystemRouter(sys)
		sysRouter.InitDeptRouter(sys)
		sysRouter.InitConfigRouter(sys)
		sysRouter.InitApiRouter(sys)
		sysRouter.InitDictRouter(sys)
		sysRouter.InitLogRouter(sys)
		sysRouter.InitMenuRouter(sys)
		sysRouter.InitRoleRouter(sys)
		sysRouter.InitPostRouter(sys)
		sysRouter.InitUserRouter(sys)

	}
	router.Group("job")
	{

	}
	router.Group("log")
	{

	}
	return router
}
