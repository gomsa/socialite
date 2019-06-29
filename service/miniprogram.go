package service

import (
	srv "github.com/gomsa/socialite/service/miniprogram"
)

//Miniprogram 短信发送接口
type Miniprogram interface {
	Code2Session(string) (*ResponseMiniprogram, error)
}

// ResponseMiniprogram 小程序返回的数据结构
type ResponseMiniprogram struct {
	Openid  string
	Session string
}

// NewMiniprogram 创建新的小程序服务
func NewMiniprogram(Type string) (mp Miniprogram) {
	switch Type {
	case "wechat":
		mp = &srv.Wechat{
			AppId:  env.Getenv("MP_WECHAT_APPID", ""),
			Secret: env.Getenv("MP_WECHAT_SECRET", ""),
		}
	}
	return mp
}
