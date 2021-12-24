package router

import (
	"github.com/gin-gonic/gin"
	"pandax/apps/system/api"
	"pandax/apps/system/services"
	"pandax/base/ctx"
)

func InitNoticeRouter(router *gin.RouterGroup) {
	s := &api.NoticeApi{
		DeptApp:   services.SysDeptModelDao,
		NoticeApp: services.SysNoticeModelDao,
	}
	notice := router.Group("notice")

	noticetList := ctx.NewLogInfo("获取通知分页列表")
	notice.GET("list", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(noticetList).Handle(s.GetNoticeList)
	})

	insertNoticeLog := ctx.NewLogInfo("添加通知信息")
	notice.POST("", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(insertNoticeLog).Handle(s.InsertNotice)
	})

	updateNoticeLog := ctx.NewLogInfo("修改通知信息")
	notice.PUT("", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(updateNoticeLog).Handle(s.UpdateNotice)
	})

	deleteNoticeLog := ctx.NewLogInfo("删除通知信息")
	notice.DELETE(":postId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(deleteNoticeLog).Handle(s.DeleteNotice)
	})
}
