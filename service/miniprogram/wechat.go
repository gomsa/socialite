package miniprogram

import (
	"encoding/json"
	"errors"

	sdk "github.com/bigrocs/wechat"
	"github.com/bigrocs/wechat/requests"
	"github.com/bigrocs/wechat/responses"
)

// Wechat 微信小程序
type Wechat struct {
	AppId  string
	Secret string
}

// Code2Session 使用 code 换取 session
func (srv *Wechat) Code2Session(code string) (res *Response, err error) {
	// 创建连接
	client, err := srv.client()
	if err != nil {
		return res, err
	}
	// 配置参数
	request := srv.request(code)
	// 请求
	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		return res, err
	}
	// 返回数据处理
	res, err = srv.response(response)
	return res, err
}

// client 创建连接
func (srv *Wechat) client() (client *sdk.Client, err error) {
	// 创建连接
	client, err = sdk.NewClient()
	if err != nil {
		return client, err
	}
	client.Credential.Miniprogram.AppId = srv.AppId
	client.Credential.Miniprogram.Secret = srv.Secret
	return client, err
}

// request 构建请求参数
func (srv *Wechat) request(code string) (request *requests.CommonRequest) {
	// 配置参数
	request = requests.NewCommonRequest()
	request.Domain = "miniprogram"
	request.ApiName = "auth.code2Session"
	request.QueryParams["js_code"] = code
	return
}

// response 返回数据处理
func (srv *Wechat) response(response *responses.CommonResponse) (res *Response, err error) {
	// res 返回请求
	r := map[string]string{}
	err = json.Unmarshal([]byte(response.GetHttpContentString()), &r)
	if err != nil {
		return res, err
	}
	if r["errcode"] != "0" {
		return res, errors.New(r["errmsg"])
	}
	res.Openid = r["openid"]
	res.Session = r["session_key"]
	return res, err
}
