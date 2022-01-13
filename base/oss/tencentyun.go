package oss

import (
	"context"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
	"github.com/tencentyun/cos-go-sdk-v5"
)

//oss 上传配置
type TencentConfig struct {
	SecretID  string
	SecretKey string
	bucket    string
}

func (t *TencentConfig) CreateClient() *cos.Client {
	// 将 examplebucket-1250000000 和 COS_REGION 修改为真实的信息
	u, _ := url.Parse(t.bucket)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  t.SecretID,
			SecretKey: t.SecretKey,
		},
	})
	return c
}

type tencentOss struct {
	client *cos.Client
}

func NewTencentOss(c TencentConfig) *tencentOss {
	client := c.CreateClient()
	return &tencentOss{
		client: client,
	}
}

func (t *tencentOss) Put(objectName, localFileName string) error {
	// 通过本地文件上传对象
	_, err := t.client.Object.PutFromFile(context.Background(), objectName, localFileName, nil)
	if err != nil {
		return errors.Wrapf(err, "tencentOss put fail")
	}
	return nil
}

func (t *tencentOss) Get(objectName, downloadedFileName string) error {
	// 获取对象到本地文件
	_, err := t.client.Object.GetToFile(context.Background(), objectName, downloadedFileName, nil)
	if err != nil {
		return errors.Wrapf(err, "tencentOss get fail")
	}
	return nil
}

func (t *tencentOss) Del(objectName string) error {
	_, err := t.client.Object.Delete(context.Background(), objectName)
	if err != nil {
		return errors.Wrapf(err, "tencentOss del fail")
	}
	return nil
}
