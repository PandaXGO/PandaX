package ys

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	//[用户]获取accessToken
	ACCESSTOKEN = "https://open.ys7.com/api/lapp/token/get"
)

type Ys struct {
	AppKey      string
	Secret      string
	AccessToken string
	ExpireTime  int64
}

func (ys *Ys) GetAccessToken() error {
	params := make(map[string]interface{})
	params["appKey"] = ys.AppKey
	params["appSecret"] = ys.Secret
	ac := &AccessToken{}
	_, err := ys.requset("POST", ACCESSTOKEN, params, &ac)
	if err != nil {
		return err
	}
	ys.AccessToken = ac.AccessToken
	ys.ExpireTime = ac.ExpireTime
	return nil
}

func (ys *Ys) requset(method, url string, params map[string]interface{}, data interface{}) (status *Status, err error) {
	defer func() {
		if Rerr := recover(); Rerr != nil {
			err = errors.New("recover error")
			return
		}
	}()
	var r http.Request
	r.ParseForm()
	for k, v := range params {
		r.Form.Add(k, fmt.Sprint(v))
	}
	req, err := http.NewRequest(method, url, strings.NewReader(r.Form.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res = Status{
		Data: data,
	}
	err = json.Unmarshal(buf, &res)
	if err != nil {
		return nil, err
	}
	if res.Code != "200" {
		return status, errors.New(res.Msg)
	}
	return status, nil
}

func (ys *Ys) authorizeRequset(method, url string, params map[string]interface{}, data interface{}) (status *Status, err error) {
	exTime := time.Unix(ys.ExpireTime/1000, 0)
	if exTime.Unix() < time.Now().Unix() || ys.AccessToken == "" {
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