package api

import (
	"fmt"
	"pandax/kit/biz"
	"pandax/kit/oss"
	"pandax/kit/restfulx"
	"pandax/pkg/config"
	"pandax/pkg/global"
	"pandax/pkg/tool"
	"path"
	"time"
)

type UploadApi struct{}

// UploadImage 图片上传
func (up *UploadApi) UploadImage(rc *restfulx.ReqCtx) {
	_, fileHeader, err := rc.Request.Request.FormFile("file")
	fileType := restfulx.QueryParam(rc, "fileType")
	biz.ErrIsNil(err, "请传入文件")
	local := &tool.Local{Path: tool.GetFilePath(fileType)}
	link, fileName, err := local.UploadFile(fileHeader)
	biz.ErrIsNil(err, "文件上传失败")

	rc.ResData = map[string]string{"fileName": fileName, "filePath": link}
}

// UplaodResOsses 上传文件ResOsses
func (p *UploadApi) UplaodToOss(rc *restfulx.ReqCtx) {
	_, handler, _ := rc.Request.Request.FormFile("file")
	yunFileTmpPath := "uploads/" + time.Now().Format("2006-01-02") + "/" + handler.Filename
	// 读取本地文件。
	f, openError := handler.Open()
	biz.ErrIsNil(openError, "function file.Open() Failed")
	config := global.Conf.Oss
	biz.ErrIsNil(NewOss(*config).PutObj(yunFileTmpPath, f), "上传OSS失败")

	rc.ResData = fmt.Sprintf("http://%s/%s/%s", config.Endpoint, config.BucketName, yunFileTmpPath)
}

// subpath 是fileName
func (up *UploadApi) GetImage(rc *restfulx.ReqCtx) {
	fileType := restfulx.QueryParam(rc, "fileType")
	actual := path.Join(tool.GetFilePath(fileType), restfulx.PathParam(rc, "subpath"))

	rc.Download(actual)
}

func (up *UploadApi) DeleteImage(rc *restfulx.ReqCtx) {
	fileName := restfulx.QueryParam(rc, "fileName")
	fileType := restfulx.QueryParam(rc, "fileType")
	biz.NotEmpty(fileName, "请传要删除的图片名")
	local := &tool.Local{Path: tool.GetFilePath(fileType)}
	err := local.DeleteFile(fileName)
	biz.ErrIsNil(err, "文件删除失败")
}

func NewOss(ens config.Oss) oss.Driver {
	return oss.NewMiniOss(oss.MiniOConfig{
		BucketName:      ens.BucketName,
		Endpoint:        ens.Endpoint,
		AccessKeyID:     ens.AccessKey,
		SecretAccessKey: ens.SecretKey,
		UseSSL:          ens.UseSSL,
		Location:        "cn-north-1",
	})
}
