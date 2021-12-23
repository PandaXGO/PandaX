package router

import (
	"github.com/gin-gonic/gin"
	api2 "pandax/apps/system/api"
	services2 "pandax/apps/system/services"
	"pandax/base/ctx"
)

func InitDeptRouter(router *gin.RouterGroup) {
	r := &api2.DeptApi{
		DeptApp: services2.SysDeptModelDao,
		RoleApp: services2.SysRoleModelDao,
		UserApp: services2.SysUserModelDao,
	}
	dept := router.Group("dept")

	roleDeptTreSelectLog := ctx.NewLogInfo("获取角色部门树")
	dept.GET("roleDeptTreeSelect/:roleId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(roleDeptTreSelectLog).Handle(r.GetDeptTreeRoleSelect)
	})

	deptTreeLog := ctx.NewLogInfo("获取所有部门树")
	dept.GET("deptTree", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(deptTreeLog).Handle(r.GetDeptTree)
	})

	deptListLog := ctx.NewLogInfo("获取部门列表")
	dept.GET("list", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(deptListLog).Handle(r.GetDeptList)
	})

	deptLog := ctx.NewLogInfo("获取部门信息")
	dept.GET(":deptId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(deptLog).Handle(r.GetDept)
	})

	inertDeptLog := ctx.NewLogInfo("添加部门信息")
	dept.POST("", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(inertDeptLog).Handle(r.InsertDept)
	})

	updateDeptLog := ctx.NewLogInfo("修改部门信息")
	dept.PUT("", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(updateDeptLog).Handle(r.UpdateDept)
	})

	deleteDeptLog := ctx.NewLogInfo("删除部门信息")
	dept.DELETE(":deptId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(deleteDeptLog).Handle(r.DeleteDept)
	})
}
