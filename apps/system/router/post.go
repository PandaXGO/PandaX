package router

import (
	"github.com/gin-gonic/gin"
	"pandax/apps/system/api"
	"pandax/apps/system/services"
	"pandax/base/ctx"
)

func InitPostRouter(router *gin.RouterGroup) {
	s := &api.PostApi{
		PostApp: services.SysPostModelDao,
		UserApp: services.SysUserModelDao,
		RoleApp: services.SysRoleModelDao,
	}
	post := router.Group("post")

	postList := ctx.NewLogInfo("获取岗位分页列表")
	post.GET("list", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(postList).Handle(s.GetPostList)
	})

	postLog := ctx.NewLogInfo("获取岗位信息")
	post.GET(":postId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(postLog).Handle(s.GetPost)
	})

	insertPostLog := ctx.NewLogInfo("添加岗位信息")
	post.POST("", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(insertPostLog).Handle(s.InsertPost)
	})

	updatePostLog := ctx.NewLogInfo("修改岗位信息")
	post.PUT("", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(updatePostLog).Handle(s.UpdatePost)
	})

	deletePostLog := ctx.NewLogInfo("删除岗位信息")
	post.DELETE(":postId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(deletePostLog).Handle(s.DeletePost)
	})
}
