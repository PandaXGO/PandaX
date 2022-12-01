package router

import (
	"github.com/XM-GO/PandaKit/restfulx"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"pandax/apps/system/api"
)

func InitUploadRouter(container *restful.Container) {
	s := &api.UploadApi{}
	ws := new(restful.WebService)
	ws.Path("/upload").Produces(restful.MIME_JSON)
	tags := []string{"upload"}

	ws.Route(ws.POST("/up").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("上传图片").Handle(s.UploadImage)
	}).
		Doc("上传图片").
		Param(ws.FormParameter("imagefile", "文件")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", map[string]string{}))

	ws.Route(ws.GET("/get/{subpath}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithNeedToken(false).WithNeedCasbin(false).WithLog("获取图片").Handle(s.GetImage)
	}).
		Doc("获取图片").
		Param(ws.PathParameter("subpath", "文件名")).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.DELETE("/delete").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除图片").Handle(s.DeleteImage)
	}).
		Doc("删除图片").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("fileName", "文件名称").DataType("string")))

	container.Add(ws)
}
