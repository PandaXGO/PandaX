package global

// 告警等级
const (
	CRITICAL      = "CRITICAL"      // 危险
	MAJOR         = "MAJOR"         // 重要
	MINOR         = "MINOR"         // 次要
	WARNING       = "WARNING"       // 警告
	INDETERMINATE = "INDETERMINATE" // 不确定
)

// 告警状态
const (
	ALARMING  = "0" // 告警中
	CONFIRMED = "1" // 已确认
	CLEARED   = "2" // 已清除
	CLOSED    = "3" // 已关闭
)

// 设备状态
const (
	INACTIVE = "inactive" //未激活
	ONLINE   = "online"   //在线
	OFFLINE  = "offline"  // 离线
)

// 设备类型
const (
	DIRECT   = "direct"   //直连设备
	GATEWAY  = "gateway"  //网关设备
	GATEWAYS = "gatewayS" // 网关子设备
	MONITOR  = "monitor"  // 监控设备
)

// 设备命令状态
const (
	CMDSUCCESS = "0" //执行成功
	CMDFAIL    = "1" //执行失败
	CMDRUNNING = "2" // 执行中
)
