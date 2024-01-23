package restfulx

type LogInfo struct {
	LogResp     bool   // 是否记录返回结果
	Description string // 请求描述
}

func NewLogInfo(description string) *LogInfo {
	return &LogInfo{Description: description, LogResp: false}
}

func (i *LogInfo) WithLogResp(logResp bool) *LogInfo {
	i.LogResp = logResp
	return i
}
