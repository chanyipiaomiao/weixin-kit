package weixin

import (
	"encoding/json"
	"fmt"

	"github.com/levigross/grequests"
	"github.com/tidwall/gjson"
)

var (

	// WeixinErr 微信错误信息码
	WeixinErr = func(errcode int64, errmsg string) error {
		return fmt.Errorf("weixin return error, errcode: %d, errmsg: %s", errcode, errmsg)
	}
)

// Client 微信
type Client struct {
	AccessTokenAPI string
	APIURL         string
	CorpID         string
	CorpSecret     string
	Message        *Message
}

// GetAccessToken 获取AccessToken
// corpid 每个企业都拥有唯一的corpid，获取此信息可在管理后台“我的企业”－“企业信息”下查看（需要有管理员权限）
// corpsecret 每个应用有独立的secret，所以每个应用的access_token应该分开来获取 在管理后台->“企业应用”->点进应用
func (c *Client) GetAccessToken() (string, error) {

	o := &grequests.RequestOptions{
		Params: map[string]string{
			"corpid":     c.CorpID,
			"corpsecret": c.CorpSecret,
		},
	}

	resp, err := grequests.Get(c.AccessTokenAPI, o)
	if err != nil {
		return "", err
	}

	respJSON := resp.String()
	errcode := gjson.Get(respJSON, "errcode")
	token := gjson.Get(respJSON, "access_token")
	if errcode.Int() == 0 {
		return token.String(), nil
	}
	return "", WeixinErr(errcode.Int(), gjson.Get(respJSON, "errmsg").String())
}

// SendMessage 发送消息
func (c *Client) SendMessage() (bool, error) {
	reqJSON, err := json.Marshal(c.Message)
	if err != nil {
		return false, err
	}
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return false, err
	}
	o := &grequests.RequestOptions{
		Params: map[string]string{
			"access_token": accessToken,
		},
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		JSON: reqJSON,
	}
	resp, err := grequests.Post(c.APIURL, o)
	if err != nil {
		return false, err
	}
	respJSON := resp.String()
	errcode := gjson.Get(respJSON, "errcode")
	errcodeInt := errcode.Int()
	if errcodeInt == 0 {
		return true, nil
	}

	return false, WeixinErr(errcodeInt, gjson.Get(respJSON, "errmsg").String())
}
