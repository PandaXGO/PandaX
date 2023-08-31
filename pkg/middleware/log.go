package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/PandaXGO/PandaKit/biz"
	"github.com/PandaXGO/PandaKit/logger"
	"github.com/PandaXGO/PandaKit/restfulx"
	"github.com/PandaXGO/PandaKit/utils"
	"github.com/sirupsen/logrus"
	"reflect"
	"runtime/debug"
)

func LogHandler(rc *restfulx.ReqCtx) error {
	li := rc.LogInfo
	if li == nil {
		return nil
	}

	lfs := logrus.Fields{}
	if la := rc.LoginAccount; la != nil {
		lfs["uid"] = la.UserId
		lfs["uname"] = la.UserName
	}

	req := rc.Request.Request
	lfs[req.Method] = req.URL.Path

	if err := rc.Err; err != nil {
		logger.Log.WithFields(lfs).Error(getErrMsg(rc, err))
		return nil
	}
	logger.Log.WithFields(lfs).Info(getLogMsg(rc))
	return nil
}

func getLogMsg(rc *restfulx.ReqCtx) string {
	msg := rc.LogInfo.Description + fmt.Sprintf(" ->%dms", rc.Timed)
	if !utils.IsBlank(reflect.ValueOf(rc.ReqParam)) {
		rb, _ := json.Marshal(rc.ReqParam)
		msg = msg + fmt.Sprintf("\n--> %s", string(rb))
	}

	// 返回结果不为空，则记录返回结果
	if rc.LogInfo.LogResp && !utils.IsBlank(reflect.ValueOf(rc.ResData)) {
		respB, _ := json.Marshal(rc.ResData)
		msg = msg + fmt.Sprintf("\n<-- %s", string(respB))
	}
	return msg
}

func getErrMsg(rc *restfulx.ReqCtx, err any) string {
	msg := rc.LogInfo.Description
	if !utils.IsBlank(reflect.ValueOf(rc.ReqParam)) {
		rb, _ := json.Marshal(rc.ReqParam)
		msg = msg + fmt.Sprintf("\n--> %s", string(rb))
	}

	var errMsg string
	switch t := err.(type) {
	case *biz.BizError:
		errMsg = fmt.Sprintf("\n<-e errCode: %d, errMsg: %s", t.Code(), t.Error())
	case error:
		errMsg = fmt.Sprintf("\n<-e errMsg: %s\n%s", t.Error(), string(debug.Stack()))
	case string:
		errMsg = fmt.Sprintf("\n<-e errMsg: %s\n%s", t, string(debug.Stack()))
	}
	return (msg + errMsg)
}
