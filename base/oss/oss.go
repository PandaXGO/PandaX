package oss

import "io"

// Driver oss驱动接口定义
type Driver interface {
	//上传
	Put(objectName, localFileName string) error

	PutObj(objectName string, file io.Reader) error
	//下载
	Get(objectName, downloadedFileName string) error
	//删除
	Del(objectName string) error
}
