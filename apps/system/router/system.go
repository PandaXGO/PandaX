package router

import (
	"github.com/gin-gonic/gin"
	api2 "pandax/apps/system/api"
)

func InitSystemRouter(router *gin.RouterGroup) {
	s := &api2.System{}
	sys := router.Group("")

	{
		sys.GET("", s.ConnectWs)
		sys.GET("server", s.ServerInfo)
	}
}
