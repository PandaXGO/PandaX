package oss

import (
	"github.com/pkg/errors"
	"io"

	aliOssSdk "github.com/aliyun/aliyun-oss-go-sdk/oss"
)

//oss 上传配置
type AliConfig struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Bucket    string `json:"bucket"`
	Endpoint  string `json:"endpoint"`
}

//oss 根据参数来创建 Bucket
func (c *AliConfig) CreateBucket() (bucket *aliOssSdk.Bucket, err error) {
	// Endpoint以杭州为例，其它Region请按实际情况填写。
	endpoint := c.Endpoint
	// 阿里云主账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM账号进行API访问或日常运维，请登录 https://ram.console.aliyun.com 创建RAM账号。
	accessKeyId := c.AccessKey
	accessKeySecret := c.SecretKey
	bucketName := c.Bucket
	// 创建OSSClient实例。
	ossClient, err := aliOssSdk.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		return nil, errors.Wrapf(err, "创建 aliyun OSSClient实例失败")
	}
	// 获取存储空间。
	bucket, err = ossClient.Bucket(bucketName)
	if err != nil {
		return nil, errors.Wrapf(err, "获取 aliyun OSS 存储空间失败")
	}
	return
}

//oss 上传客户端
type aliOss struct {
	bucket *aliOssSdk.Bucket
}

func NewAliOss(c AliConfig) Driver {
	bucket, err := c.CreateBucket()
	if err != nil {
		panic(any(err))
	}
	return &aliOss{
		bucket: bucket,
	}
}

// Put 上传
func (c *aliOss) Put(objectName string, localFileName string) error {
	err := c.bucket.PutObjectFromFile(objectName, localFileName)
	if err != nil {
		return errors.Wrapf(err, "put oss file fail")
	}
	return nil
}

func (c *aliOss) PutObj(objectName string, file io.Reader) error {
	err := c.bucket.PutObject(objectName, file)
	if err != nil {
		return errors.Wrapf(err, "put oss file fail")
	}
	return nil
}

// Get 下载
func (c *aliOss) Get(objectName, downloadedFileName string) error {
	err := c.bucket.GetObjectToFile(objectName, downloadedFileName)
	if err != nil {
		return errors.Wrapf(err, "get oss file fail")
	}
	return nil
}

// Del 删除
func (c *aliOss) Del(objectName string) error {
	// 删除文件。
	err := c.bucket.DeleteObject(objectName)
	if err != nil {
		return errors.Wrapf(err, "del oss file fail")
	}
	return nil
}
