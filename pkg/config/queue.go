package config

type Queue struct {
	Enable bool  `yaml:"  enable"`
	Num    int64 `yaml:"  num"` //并发数
}
