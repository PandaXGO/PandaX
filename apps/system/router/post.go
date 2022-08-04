package router

import (
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"pandax/apps/system/api"
	"pandax/apps/system/entity"
	"pandax/apps/system/services"
)

func InitPostRouter(container *restful.Container) {
	s := &api.PostApi{
		PostApp: services.SysPostModelDao,
		UserApp: services.SysUserModelDao,
		RoleApp: services.SysRoleModelDao,
	}
	ws := new(restful.WebService)
	ws.Path("/system/post").Produces(restful.MIME_JSON)
	tags := []string{"post"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取岗位分页列表").Handle(s.GetPostList)
	}).
		Doc("获取岗位分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("status", "status").DataType("string")).
		Param(ws.QueryParameter("postName", "postName").DataType("string")).
		Param(ws.QueryParameter("postCode", "postCode").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/{postId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取岗位信息").Handle(s.GetPost)
	}).
		Doc("获取岗位信息").
		Param(ws.PathParameter("postId", "Id").DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.SysPost{}).
		Returns(200, "OK", entity.SysPost{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加岗位信息").Handle(s.InsertPost)
	}).
		Doc("添加岗位信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.SysPost{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改岗位信息").Handle(s.UpdatePost)
	}).
		Doc("修改岗位信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.SysPost{}))

	ws.Route(ws.DELETE("/{postId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除岗位信息").Handle(s.DeletePost)
	}).
		Doc("删除岗位信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("postId", "多id 1,2,3").DataType("string")))

	container.Add(ws)
}
