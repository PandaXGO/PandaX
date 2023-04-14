package api

import (
	"github.com/XM-GO/PandaKit/biz"
	"github.com/XM-GO/PandaKit/restfulx"
	"net/http"
	"pandax/apps/visual/entity"
	"pandax/apps/visual/services"
	"pandax/pkg/tool"
	"path"
)

type UploadApi struct {
	VisualScreenImageApp services.VisualScreenImageModel
}

func (up *UploadApi) UploadImage(rc *restfulx.ReqCtx) {
	_, fileHeader, err := rc.Request.Request.FormFile("imagefile")
	biz.ErrIsNil(err, "请传入文件")
	local := &tool.Local{Path: "uploads/file"}
	link, fileName, err := local.UploadFile(fileHeader)
	biz.ErrIsNil(err, "文件上传失败")

	up.VisualScreenImageApp.Insert(entity.VisualScreenImage{
		FileName: fileName,
		FilePath: link,
	})
	rc.ResData = map[string]string{"fileName": fileName, "filePath": link}
}

func (up *UploadApi) GetImages(rc *restfulx.ReqCtx) {
	list := up.VisualScreenImageApp.FindList(entity.VisualScreenImage{})
	rc.ResData = list
}

func (up *UploadApi) GetImage(rc *restfulx.ReqCtx) {
	actual := path.Join("uploads/file", restfulx.PathParam(rc, "subpath"))
	http.ServeFile(
		rc.Response.ResponseWriter,
		rc.Request.Request,
		actual)
}

func (up *UploadApi) DeleteImage(rc *restfulx.ReqCtx) {
	fileName := restfulx.QueryParam(rc, "fileName")
	biz.NotEmpty(fileName, "请传要删除的图片名")
	local := &tool.Local{Path: "uploads/file"}
	err := local.DeleteFile(fileName)
	biz.ErrIsNil(err, "文件删除失败")
	up.VisualScreenImageApp.Delete(fileName)
}
