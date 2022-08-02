package router

import (
	"github.com/gin-gonic/gin"
	"pandax/apps/system/api"
	"pandax/apps/system/services"
	"pandax/base/ginx"
)

func InitPostRouter(router *gin.RouterGroup) {
	s := &api.PostApi{
		PostApp: services.SysPostModelDao,
		UserApp: services.SysUserModelDao,
		RoleApp: services.SysRoleModelDao,
	}
	post := router.Group("post")

	post.GET("list", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取岗位分页列表").Handle(s.GetPostList)
	})

	post.GET(":postId", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取岗位信息").Handle(s.GetPost)
	})

	post.POST("", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("添加岗位信息").Handle(s.InsertPost)
	})

	post.PUT("", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("修改岗位信息").Handle(s.UpdatePost)
	})

	post.DELETE(":postId", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("删除岗位信息").Handle(s.DeletePost)
	})
}
