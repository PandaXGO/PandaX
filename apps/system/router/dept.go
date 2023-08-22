package router

import (
	"github.com/XM-GO/PandaKit/restfulx"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"pandax/apps/system/api"
	"pandax/apps/system/api/vo"
	"pandax/apps/system/entity"
	"pandax/apps/system/services"
)

func InitDeptRouter(container *restful.Container) {
	s := &api.DeptApi{
		DeptApp: services.SysDeptModelDao,
		RoleApp: services.SysRoleModelDao,
		UserApp: services.SysUserModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/system/dept").Produces(restful.MIME_JSON)
	tags := []string{"dept"}

	ws.Route(ws.GET("/roleDeptTreeSelect/{roleId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取角色部门树").Handle(s.GetDeptTreeRoleSelect)
	}).
		Doc("获取角色部门树").
		Param(ws.PathParameter("roleId", "角色Id").DataType("int").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(vo.DeptTreeVo{}).
		Returns(200, "OK", vo.DeptTreeVo{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/deptTree").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取所有部门树").Handle(s.GetDeptTree)
	}).
		Doc("获取所有部门树").
		Param(ws.QueryParameter("deptName", "deptName").DataType("string")).
		Param(ws.QueryParameter("status", "status").DataType("string")).
		Param(ws.QueryParameter("deptId", "deptId").DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]entity.SysDept{}).
		Returns(200, "OK", []entity.SysDept{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取部门列表").Handle(s.GetDeptList)
	}).
		Doc("获取部门列表").
		Param(ws.QueryParameter("deptName", "deptName").DataType("string")).
		Param(ws.QueryParameter("status", "status").DataType("string")).
		Param(ws.QueryParameter("deptId", "deptId").DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]entity.SysDept{}).
		Returns(200, "OK", []entity.SysDept{}))

	ws.Route(ws.GET("/{deptId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取部门信息").Handle(s.GetDept)
	}).
		Doc("获取部门信息").
		Param(ws.PathParameter("deptId", "部门Id").DataType("int").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.SysDept{}). // on the response
		Returns(200, "OK", entity.SysDept{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加部门信息").Handle(s.InsertDept)
	}).
		Doc("添加部门信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.SysDept{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改部门信息").Handle(s.UpdateDept)
	}).
		Doc("修改部门信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.SysDept{}))

	ws.Route(ws.DELETE("/{deptId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除部门信息").Handle(s.DeleteDept)
	}).
		Doc("删除部门信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("deptId", "多id 1,2,3").DataType("int")))

	container.Add(ws)

}
