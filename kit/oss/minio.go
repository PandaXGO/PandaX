package oss

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pkg/errors"
	"io"
	"log"
)

//oss 上传配置
type MiniOConfig struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	BucketName      string
	Location        string
	UseSSL          bool
	TargetFilePath  string
}

func (m *MiniOConfig) CreateClient() *minio.Client {
	client, err := minio.New(m.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(m.AccessKeyID, m.SecretAccessKey, ""),
		Secure: m.UseSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}
	isExists, err := client.BucketExists(context.Background(), m.BucketName)
	if err != nil {
		log.Println(err)
		return nil
	}

	if !isExists {
		err = client.MakeBucket(context.Background(), m.BucketName, minio.MakeBucketOptions{Region: m.Location})
		if err != nil {
			log.Println(err)
			return nil
		}
	}
	return client
}

type minioOss struct {
	client     *minio.Client
	bucketName string
}

func NewMiniOss(c MiniOConfig) Driver {
	client := c.CreateClient()
	return &minioOss{
		client:     client,
		bucketName: c.BucketName,
	}
}

// contentType 需要上传的文件类型 例如application/zip
func (m *minioOss) Put(objectName, localFileName string) error {
	contentType, err := GetFileContentType(localFileName)
	if err != nil {
		return err
	}
	_, err = m.client.FPutObject(context.Background(), m.bucketName, objectName, localFileName, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return errors.Wrapf(err, "put oss file fail")
	}
	return nil
}

func (m *minioOss) PutObj(objectName string, file io.Reader) error {
	_, err := m.client.PutObject(context.Background(), m.bucketName, objectName, file, -1, minio.PutObjectOptions{})
	if err != nil {
		return errors.Wrapf(err, "put oss file fail")
	}
	return nil
}

func (m *minioOss) Get(objectName, downloadedFileName string) error {
	err := m.client.FGetObject(context.Background(), m.bucketName, objectName, downloadedFileName, minio.GetObjectOptions{})
	if err != nil {
		return errors.Wrapf(err, "get oss file fail")
	}
	return nil
}

func (m *minioOss) Del(objectName string) error {
	err := m.client.RemoveObject(context.Background(), m.bucketName, objectName, minio.RemoveObjectOptions{GovernanceBypass: true})
	if err != nil {
		return errors.Wrapf(err, "qiniu oss del file fail")
	}
	return nil
}
