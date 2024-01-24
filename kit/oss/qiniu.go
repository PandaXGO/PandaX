package oss

import (
	"context"
	"io"
	utilFile "pandax/kit/file"

	"github.com/pkg/errors"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type QiniuConfig struct {
	AccessKey     string
	SecretKey     string
	Bucket        string
	PolicyExpires uint64        // 上传凭证的有效时间，单位秒
	Zone          *storage.Zone // 空间所在的机房
	UseHTTPS      bool          // 是否使用https域名
	UseCdnDomains bool          // 是否使用cdn加速域名
	CentralRsHost string        // 中心机房的RsHost，用于list bucket
	Domain        string        // 外链域名
}

type qiniuOss struct {
	config QiniuConfig
}

func NewQnOss(config QiniuConfig) Driver {
	return &qiniuOss{
		config: config,
	}
}

func (q *qiniuOss) Put(objectName, localFileName string) error {
	// 鉴权
	mac := qbox.NewMac(q.config.AccessKey, q.config.SecretKey)
	// 上传策略
	putPolicy := storage.PutPolicy{
		Scope:   q.config.Bucket,
		Expires: q.config.PolicyExpires,
	}
	// 获取上传token
	upToken := putPolicy.UploadToken(mac)

	// 上传Config对象
	cfg := storage.Config{
		Zone:          q.config.Zone,          //指定上传的区域
		UseHTTPS:      q.config.UseHTTPS,      // 是否使用https域名
		UseCdnDomains: q.config.UseCdnDomains, //是否使用CDN上传加速
	}
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 上传文件
	err := formUploader.PutFile(context.Background(), &ret, upToken, objectName, localFileName, nil)
	if err != nil {
		return errors.Wrapf(err, "qiniu oss put file fail")
	}
	return nil
}

func (q *qiniuOss) PutObj(objectName string, file io.Reader) error {
	// 鉴权
	mac := qbox.NewMac(q.config.AccessKey, q.config.SecretKey)
	// 上传策略
	putPolicy := storage.PutPolicy{
		Scope:   q.config.Bucket,
		Expires: q.config.PolicyExpires,
	}
	// 获取上传token
	upToken := putPolicy.UploadToken(mac)

	// 上传Config对象
	cfg := storage.Config{
		Zone:          q.config.Zone,          //指定上传的区域
		UseHTTPS:      q.config.UseHTTPS,      // 是否使用https域名
		UseCdnDomains: q.config.UseCdnDomains, //是否使用CDN上传加速
	}
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 上传文件
	err := formUploader.Put(context.Background(), &ret, upToken, objectName, file, 0, nil)
	if err != nil {
		return errors.Wrapf(err, "qiniu oss put file fail")
	}
	return nil
}

func (q *qiniuOss) Get(objectName, downloadedFileName string) error {
	publicAccessURL := storage.MakePublicURL(q.config.Domain, objectName)
	err := utilFile.DownloadFileWithConcurrency(publicAccessURL, downloadedFileName)
	if err != nil {
		return errors.Wrapf(err, "qiniu oss get file fail")
	}
	return nil
}

func (q *qiniuOss) Del(objectName string) error {
	mac := qbox.NewMac(q.config.AccessKey, q.config.SecretKey)
	cfg := storage.Config{
		Zone:          q.config.Zone,          //指定上传的区域
		UseHTTPS:      q.config.UseHTTPS,      // 是否使用https域名
		UseCdnDomains: q.config.UseCdnDomains, //是否使用CDN上传加速
	}
	bucketManager := storage.NewBucketManager(mac, &cfg)

	err := bucketManager.Delete(q.config.Bucket, objectName)
	if err != nil {
		return errors.Wrapf(err, "qiniu oss del file fail")
	}
	return nil
}
