package devicerpc

import (
	"errors"
	"fmt"
	"time"

	"github.com/kakuilan/kgo"
)

type RpcPayload struct {
	Method string `json:"method"`
	Params any    `json:"params"`
}

func (rp *RpcPayload) ToMap() map[string]any {
	data, err := kgo.KConv.Struct2Map(rp, "")
	if err != nil {
		return nil
	}
	return data
}

// GetRequestResult 处理设备端请求服务端方法
func (rpc RpcPayload) GetRequestResult() (string, error) {
	//TODO 此处处理设备的请求参数逻辑
	//自己定义请求逻辑
	if rpc.Params == "getCurrentTime" {
		unix := time.Now().Unix()
		msg := fmt.Sprintf("%d", unix)
		return msg, nil
	}
	// 获取属性 ...
	return "", errors.New("未获取到请求方法")
}
