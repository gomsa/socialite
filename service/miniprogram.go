package service

import (
	srv "github.com/gomsa/socialite/service/miniprogram"
)

//Miniprogram 短信发送接口
type Miniprogram interface {
	Code2Session(string) (*ResponseMiniprogram, error)
}

// NewMiniprogram 创建新的小程序服务
func NewMiniprogram(Type string) (mp srv.Miniprogram) {
	switch Type {
	case "wechat":
		mp = &srv.Wechat{
			AppId:  env.Getenv("MP_WECHAT_APPID", ""),
			Secret: env.Getenv("MP_WECHAT_SECRET", ""),
		}
	}
	return mp
}
