package tool

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"pandax/pkg/global"
	"path"
	"strings"
	"time"
)

type Local struct {
	Path string
}

//@object: *Local
//@function: UploadFile
//@description: 上传文件
//@param: file *multipart.FileHeader
//@return: string, string, error

func (local *Local) UploadFile(file *multipart.FileHeader) (string, string, error) {
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	name = MD5V([]byte(name))
	// 拼接新文件名
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	// 尝试创建此路径
	mkdirErr := os.MkdirAll(local.Path, os.ModePerm)
	if mkdirErr != nil {
		global.Log.Error("function os.MkdirAll() Filed", mkdirErr.Error())
		return "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	// 拼接路径和文件名
	p := local.Path + "/" + filename

	f, openError := file.Open() // 读取文件
	if openError != nil {
		global.Log.Error("function file.Open() Filed", openError.Error())
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	out, createErr := os.Create(p)
	if createErr != nil {
		global.Log.Error("function os.Create() Filed", createErr.Error())
		return "", "", errors.New("function os.Create() Filed, err:" + createErr.Error())
	}
	defer out.Close() // 创建文件 defer 关闭

	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		global.Log.Error("function io.Copy() Filed", copyErr.Error())
		return "", "", errors.New("function io.Copy() Filed, err:" + copyErr.Error())
	}
	return p, filename, nil
}

//@object: *Local
//@function: DeleteFile
//@description: 删除文件
//@param: key string
//@return: error

func (local *Local) DeleteFile(key string) error {
	p := local.Path + "/" + key
	if strings.Contains(p, local.Path) {
		if err := os.Remove(p); err != nil {
			return errors.New("本地文件删除失败, err:" + err.Error())
		}
	}
	return nil
}

func MD5V(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}
