package config

type Queue struct {
	QueuePool int64 `yaml:"queue-pool"` //消息队列池
	TaskNum   int64 `yaml:"task-num"`   //任务队列数
	ChNum     int64 `yaml:"ch-num"`     //并发数
}
