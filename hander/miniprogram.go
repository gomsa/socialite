package hander

import (
	"context"
	"fmt"

	pb "github.com/gomsa/socialite/proto/miniprogram"
	userPD "github.com/gomsa/socialite/proto/user"
	"github.com/gomsa/socialite/service"
	"github.com/micro/go-micro/util/log"
)

// Miniprogram 小程序
type Miniprogram struct {
	Repo service.URepository
	Mp   service.Miniprogram
}

// Auth 小程序登录授权
func (srv *Miniprogram) Auth(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// 选择小程序驱动
	srv.Mp, err = service.NewMiniprogram(req.Type)
	if err != nil {
		log.Log(err)
		return err
	}
	// 换取 session
	mp, err := srv.Mp.Code2Session(req.Code)
	if err != nil {
		log.Log(err)
		return err
	}
	u := &userPD.User{
		Origin:  req.Type,
		Openid:  mp.Openid,
		Session: mp.Session,
	}
	fmt.Println(u)
	user := &userPD.User{}
	if srv.Repo.Exist(u) {
		user, err = srv.Repo.Get(u)
		if err != nil {
			log.Log(err)
			return err
		}
	} else {
		user, err = srv.Repo.Create(u)
		if err != nil {
			log.Log(err)
			return err
		}
	}
	// 根据 id 去获取 token
	fmt.Println(user.Id)
	fmt.Println(mp)
	return err
}
