package ys

import (
	"errors"
	"fmt"
	"github.com/PandaXGO/PandaKit/httpclient"
	"log"
	"strings"
	"time"
)

const (
	MASTERACC = 0 //主账号
	RAMACC    = 1 //子账号
	//[用户]获取accessToken
	ACCESSTOKEN = "https://open.ys7.com/api/lapp/token/get"
	RAMTOKENGET = "https://open.ys7.com/api/lapp/ram/token/get" //获取子账户AccessToken
)

type Ys struct {
	AppKey      string
	Secret      string
	IsRAM       int
	AccountID   string
	AccessToken string
	ExpireTime  int64
}

func (ys *Ys) GetAccessToken() error {
	params := make(map[string]interface{})
	params["appKey"] = ys.AppKey
	params["appSecret"] = ys.Secret
	ac := &AccessToken{}
	_, err := ys.requset("POST", ACCESSTOKEN, params, ac)
	if err != nil {
		return err
	}
	if ys.IsRAM == MASTERACC {
		ys.AccessToken = ac.AccessToken
		ys.ExpireTime = ac.ExpireTime
	} else {
		ys.AccessToken = ac.AccessToken
		ac, err = ys.RAMGetAccessToken(ys.AccountID)
		if err != nil {
			ys.AccessToken = ""
			return err
		}
		ys.AccessToken = ac.AccessToken
		ys.ExpireTime = ac.ExpireTime
	}

	return nil
}

// RAMGetAccessToken 获取B模式子账户accessToken
func (ys *Ys) RAMGetAccessToken(accountID string) (ac *AccessToken, err error) {
	params := make(map[string]interface{})
	params["accountId"] = accountID
	params["accessToken"] = ys.AccessToken
	ac = &AccessToken{}
	_, err = ys.requset("POST", RAMTOKENGET, params, ac)
	if err != nil {
		return nil, err
	}
	log.Println(*ac)
	return ac, nil
}

func (ys *Ys) requset(method, url string, params map[string]interface{}, data interface{}) (status *Status, err error) {
	defer func() {
		if Rerr := recover(); Rerr != nil {
			err = errors.New("recover error")
			return
		}
	}()
	ps := make([]string, 0)
	for k, v := range params {
		ps = append(ps, fmt.Sprintf("%s=%v", k, v))
	}
	status = &Status{
		Data: data,
	}
	err = httpclient.NewRequest(url).Timeout(60).PostParams(strings.Join(ps, "&")).BodyToObj(status)
	if err != nil {
		return nil, err
	}
	if status.Code != "200" {
		return status, errors.New(status.Msg)
	}
	return status, nil
}

func (ys *Ys) authorizeRequset(method, url string, params map[string]interface{}, data interface{}) (status *Status, err error) {
	exTime := time.Unix(ys.ExpireTime/1000, 0)
	if exTime.Unix() < time.Now().Unix() {
		ys.GetAccessToken()
	}
	defer func() {
		if Rerr := recover(); Rerr != nil {
			err = errors.New("recover error")
			return
		}
	}()
	params["accessToken"] = ys.AccessToken
	status, err = ys.requset(method, url, params, data)
	return
}
