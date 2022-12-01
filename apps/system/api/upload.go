package api

import (
	"fmt"
	"github.com/XM-GO/PandaKit/biz"
	filek "github.com/XM-GO/PandaKit/file"
	"github.com/XM-GO/PandaKit/restfulx"
	"github.com/kakuilan/kgo"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

type UploadApi struct{}

const filePath = "uploads/file"

// UploadImage 图片上传
func (up *UploadApi) UploadImage(rc *restfulx.ReqCtx) {
	_, fileHeader, err := rc.Request.Request.FormFile("imagefile")
	biz.ErrIsNil(err, "请传入文件")
	ext := path.Ext(fileHeader.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(fileHeader.Filename, ext)
	name = kgo.KStr.Md5(name, 32)
	// 拼接新文件名
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	filek.SaveUploadedFile(fileHeader, fmt.Sprintf("%s/%s", filePath, filename))
	biz.ErrIsNil(err, "文件上传失败")
	rc.ResData = map[string]string{"fileName": name}
}

func (up *UploadApi) GetImage(rc *restfulx.ReqCtx) {
	actual := path.Join(filePath, restfulx.PathParam(rc, "subpath"))
	http.ServeFile(
		rc.Response.ResponseWriter,
		rc.Request.Request,
		actual)
}

func (up *UploadApi) DeleteImage(rc *restfulx.ReqCtx) {
	fileName := restfulx.QueryParam(rc, "fileName")
	biz.NotEmpty(fileName, "请传要删除的图片名")
	err := os.Remove(fmt.Sprintf("%s/%s", filePath, fileName))
	biz.ErrIsNil(err, "文件删除失败")
}
