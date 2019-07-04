package hander

import (
	"context"
	"math/rand"
	"time"

	"github.com/gomsa/user/client"
	userSrvPB "github.com/gomsa/user/proto/user"
	authSrvPB "github.com/gomsa/user/proto/auth"

	pb "github.com/gomsa/socialite/proto/miniprogram"
	userPB "github.com/gomsa/socialite/proto/user"
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
	err = srv.initMp(req.Type)
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
	u := &userPB.User{
		Origin:  req.Type,
		Openid:  mp.Openid,
		Session: mp.Session,
	}
	err = srv.getUser(ctx, u)
	if err != nil {
		log.Log(err)
		return err
	}
	// 获取 token
	auth := &authSrvPB.User{
		Id: u.Id,
	}
	token, err := client.Auth.AuthById(ctx, auth)
	if err != nil {
		log.Log(err)
		return err
	}
	// 根据 id 去获取 token
	res.Token = token.Token
	res.Valid = token.Valid
	return err
}

func (srv *Miniprogram) initMp(Type string) (err error){
	// 选择小程序驱动
	srv.Mp, err = service.NewMiniprogram(Type)
	return err
}

func (srv *Miniprogram) getUser(ctx context.Context,u *userPB.User) (err error){
	if srv.Repo.Exist(u) {
		// 获取 user
		u, err = srv.Repo.Get(u)
		if err != nil {
			return err
		}
	} else {
		// 无用户先用过用户服务创建用户
		user := &userSrvPB.User{
			Origin:u.Origin,
			Password: srv.getRandomString(16),	// 密码默认为 16 位随机数
		}
		userRes, err := client.User.Create(ctx, user)
		if err != nil {
			return err
		}
		u.Id = userRes.User.Id
		// 创建社会用户
		u, err = srv.Repo.Create(u)
		if err != nil {
			return err
		}
	}
	return err
}


// getRandomString 生成随机字符串
func (srv *Miniprogram) getRandomString(length int64) string{
   str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
   bytes := []byte(str)
   result := []byte{}
   r := rand.New(rand.NewSource(time.Now().UnixNano()))
   for i := 0; int64(i) < length; i++ {
      result = append(result, bytes[r.Intn(len(bytes))])
   }
   return string(result)
}