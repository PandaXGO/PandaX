package router

import (
	"github.com/gin-gonic/gin"
	"pandax/apps/system/api"
	"pandax/apps/system/services"
	"pandax/base/ginx"
)

func InitApiRouter(router *gin.RouterGroup) {
	s := &api.SystemApiApi{
		ApiApp: services.SysApiModelDao,
	}
	api := router.Group("api")

	api.GET("list", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取api分页列表").Handle(s.GetApiList)
	})

	api.GET("all", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取所有api").Handle(s.GetAllApis)
	})

	api.GET("getPolicyPathByRoleId", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取角色拥有的api权限").Handle(s.GetPolicyPathByRoleId)
	})

	api.GET(":id", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取api信息").Handle(s.GetApiById)
	})

	api.POST("", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("添加api信息").Handle(s.CreateApi)
	})

	api.PUT("", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("修改api信息").Handle(s.UpdateApi)
	})

	api.DELETE(":id", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("删除api信息").Handle(s.DeleteApi)
	})
}
