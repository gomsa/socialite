package service

import (
	"fmt"

	srv "github.com/gomsa/socialite/service/miniprogram"
	"github.com/gomsa/tools/env"
)

//Miniprogram 短信发送接口
type Miniprogram interface {
	Code2Session(string) (*srv.Response, error)
}

// NewMiniprogram 创建新的小程序服务
func NewMiniprogram(Type string) (mp Miniprogram, err error) {
	switch Type {
	case "wechat":
		mp = &srv.Wechat{
			AppId:  env.Getenv("MP_WECHAT_APPID", "wx15550c1a89d982c8"),
			Secret: env.Getenv("MP_WECHAT_SECRET", "f9c11f183a5beb592ccd801298ff5533"),
		}
	default:
		return mp, fmt.Errorf("未找 %s 小程序驱动", Type)
	}
	return mp, err
}
