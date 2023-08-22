package config

type Oss struct {
	Endpoint   string `yaml:"endpoint"`
	AccessKey  string `yaml:"accessKey"`
	SecretKey  string `yaml:"secretKey"`
	BucketName string `yaml:"bucketName"`
	UseSSL     bool   `yaml:"useSSL"`
}
