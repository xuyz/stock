package report

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/xuyz/stock/scache"
)

// 发送报警信息到微信企业号鹰眼监控平台

const (
	accessUrlTemplate = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s"
	sendUrlTemplate   = "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s"
)

type corTokenResp struct {
	Errcode     int    `json:"errcode"`      // 43003,
	Errmsg      string `json:"errmsg"`       // "require https"
	AccessToken string `json:"access_token"` // "accesstoken000001",
	ExpiresIn   int    `json:"expires_in"`   // 7200
}

func getWechatCorToken(cropid, appSecret string) (string, error) {
	k := "cgi-bin_gettoken_" + cropid + appSecret
	if v, ok := scache.Get(k); ok {
		return v.(string), nil
	}

	url := fmt.Sprintf(accessUrlTemplate, cropid, appSecret)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	d := &corTokenResp{}
	err = json.NewDecoder(resp.Body).Decode(d)
	if err != nil {
		return "", err
	}

	if d.Errcode != 0 {
		return "", errors.New(d.Errmsg)
	}

	if d.AccessToken == "" {
		return "", errors.New("获取token失败" + d.Errmsg)
	}

	scache.Set(k, d.AccessToken, time.Duration(d.ExpiresIn)*time.Second)

	return d.AccessToken, nil
}

func WechatReport(agentID int64, cropid, appSecret, users, content string) error {
	token, err := getWechatCorToken(cropid, appSecret)
	if err != nil {
		return err
	}

	url := fmt.Sprintf(sendUrlTemplate, token)

	postData := make(map[string]interface{})
	if users == "" {
		postData["touser"] = "@all"
	} else {
		postData["touser"] = users
	}
	postData["msgtype"] = "text"
	postData["agentid"] = agentID
	postData["text"] = map[string]string{"content": content}
	postData["safe"] = 0

	data, err := json.Marshal(postData)
	if err != nil {
		return err
	}

	body := bytes.NewBuffer([]byte(data))

	retstr, err := http.Post(url, "application/json;charset=utf-8", body)

	if err != nil {
		return err
	}
	defer retstr.Body.Close()
	ioutil.ReadAll(retstr.Body)

	return nil
}
