package model

import (
	"errors"
	"fmt"
	"time"
)

type RpcPayload struct {
	Method string `json:"method"`
	Params any    `json:"params"`
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
