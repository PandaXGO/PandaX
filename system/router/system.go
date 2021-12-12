package router

import (
	"github.com/gin-gonic/gin"
	"pandax/system/api"
)

func InitSystemRouter(router *gin.RouterGroup) {
	s := &api.System{}
	sys := router.Group("")

	{
		sys.GET("", s.ConnectWs)
		sys.GET("server", s.ServerInfo)
	}
}
