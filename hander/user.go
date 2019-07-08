package hander

import (
	"context"
	"fmt"

	pb "github.com/gomsa/socialite/proto/user"
	"github.com/gomsa/socialite/service"
)

// User 用户结构
type User struct {
	Repo service.URepository
}

// Build 获取用户绑定信息
func (srv *User) Build(ctx context.Context, req *pb.User, res *pb.Response) (err error) {
	users, err := srv.Repo.GetByID(req)
	if err != nil {
		return err
	}
	res.Users = users
	return err
}

// Delete 删除用户用户
func (srv *User) Delete(ctx context.Context, req *pb.User, res *pb.Response) (err error) {
	valid, err := srv.Repo.Delete(req)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("删除用户失败")
	}
	res.Valid = valid
	return err
}
