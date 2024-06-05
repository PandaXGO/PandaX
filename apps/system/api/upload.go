package api

import (
	"errors"
	"fmt"
	"github.com/kakuilan/kgo"
	"net/http"
	"os"
	"pandax/kit/biz"
	"pandax/kit/oss"
	"pandax/kit/restfulx"
	"pandax/pkg/config"
	"pandax/pkg/global"
	"pandax/pkg/tool"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type UploadApi struct{}

const filePath = "uploads/file"

// UploadImage 图片上传
func (up *UploadApi) UploadImage(rc *restfulx.ReqCtx) {
	_, fileHeader, err := rc.Request.Request.FormFile("file")
	biz.ErrIsNil(err, "请传入文件")
	// 判断上传文件类型，不支持返回
	biz.IsTrue(kgo.KFile.IsImg(fileHeader.Filename), "请传入图片文件")

	local := &tool.Local{Path: filePath}
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

	absFilePath := filepath.Join(filePath, fileName)
	if !strings.HasPrefix(absFilePath, filepath.Clean(filePath)+string(filepath.Separator)) {
		biz.ErrIsNil(errors.New("非法文件路径"), "非法文件路径")
		return
	}

	err := os.Remove(fmt.Sprintf("%s/%s", filePath, fileName))
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
