package config

import "fmt"

type Server struct {
	Port        int            `yaml:"port"`
	Model       string         `yaml:"model"`
	Cors        bool           `yaml:"cors"`
	Rate        *Rate          `yaml:"rate"`
	IsInitTable bool           `yaml:"isInitTable"`
	DbType      string         `yaml:"db-type"`
	ExcelDir    string         `yaml:"excel-dir"`
	Tls         *Tls           `yaml:"tls"`
	Static      *[]*Static     `yaml:"static"`
	StaticFile  *[]*StaticFile `yaml:"static-file"`
}

func (s *Server) GetPort() string {
	return fmt.Sprintf(":%d", s.Port)
}

type Static struct {
	RelativePath string `yaml:"relative-path"`
	Root         string `yaml:"root"`
}

type StaticFile struct {
	RelativePath string `yaml:"relative-path"`
	Filepath     string `yaml:"filepath"`
}

type Tls struct {
	Enable   bool   `yaml:"enable"`    // 是否启用tls
	KeyFile  string `yaml:"key-file"`  // 私钥文件路径
	CertFile string `yaml:"cert-file"` // 证书文件路径
}

type Rate struct {
	Enable  bool    `yaml:"enable"`   // 是否限流
	RateNum float64 `yaml:"rate-num"` // 限流数量
}
