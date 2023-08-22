package entity

type DeviceStatusVo struct {
	Name   string `json:"name"`
	Key    string `json:"key"`
	Type   string `json:"type"`
	Define any    `json:"define"`
	Value  any    `json:"value"`
	Time   any    `json:"time"`
}

type VisualClass struct {
	ClassId string           `json:"classId"`
	Name    string           `json:"name"`
	Attrs   []VisualTwinAttr `json:"attrs"`
}

type VisualTwinAttr struct {
	Key  string `json:"key"`
	Name string `json:"name"`
	Type string `json:"type"`
	Rw   string `json:"rw"` //属性的操作权限
}

type VisualTwin struct {
	TwinId string `json:"twinId"`
	Name   string `json:"name"`
}

// 发送数据
type VisualTwinSendAttrs struct {
	TwinId string                 `json:"twinId"`
	Attrs  map[string]interface{} `json:"attrs"`
}
