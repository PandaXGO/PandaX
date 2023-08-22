package config

type Taos struct {
	Host     string `yaml:"host"`     // 服务器地址:端口
	Username string `yaml:"username"` // 数据库用户名
	Password string `yaml:"password"` // 数据库密码
	Database string `yaml:"database"`
	Config   string `yaml:"config"`
}
