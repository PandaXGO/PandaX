package router

import (
	"github.com/gin-gonic/gin"
	"pandax/apps/system/api"
	"pandax/apps/system/services"
	"pandax/base/ginx"
)

func InitDeptRouter(router *gin.RouterGroup) {
	r := &api.DeptApi{
		DeptApp: services.SysDeptModelDao,
		RoleApp: services.SysRoleModelDao,
		UserApp: services.SysUserModelDao,
	}
	dept := router.Group("dept")

	dept.GET("roleDeptTreeSelect/:roleId", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取角色部门树").Handle(r.GetDeptTreeRoleSelect)
	})

	dept.GET("deptTree", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取所有部门树").Handle(r.GetDeptTree)
	})

	dept.GET("list", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取部门列表").Handle(r.GetDeptList)
	})

	dept.GET(":deptId", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取部门信息").Handle(r.GetDept)
	})

	dept.POST("", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("添加部门信息").Handle(r.InsertDept)
	})

	dept.PUT("", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("修改部门信息").Handle(r.UpdateDept)
	})

	dept.DELETE(":deptId", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("删除部门信息").Handle(r.DeleteDept)
	})
}
