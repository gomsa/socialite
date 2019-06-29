package hander

import (
	"context"
	"fmt"

	pb "github.com/gomsa/socialite/proto/miniprogram"
	"github.com/gomsa/socialite/service"
	"github.com/micro/go-micro/util/log"
)

// Miniprogram 小程序
type Miniprogram struct {
	Mp service.Miniprogram
}

// Auth 小程序登录授权
func (srv *Miniprogram) Auth(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// 选择小程序驱动
	srv.Mp = service.NewMiniprogram(req.Type)
	// 换取 session
	mp, err := srv.Mp.Code2Session(req.Code)
	if err != nil {
		log.Log(err)
		return err
	}
	fmt.Println(mp)
	return err
}
