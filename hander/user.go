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

// List 获取所有用户
func (srv *User) List(ctx context.Context, req *pb.ListQuery, res *pb.Response) (err error) {
	users, err := srv.Repo.List(req)
	total, err := srv.Repo.Total(req)
	if err != nil {
		return err
	}
	res.Total = total
	res.Users = users
	return err
}

// Get 获取用户
func (srv *User) Get(ctx context.Context, req *pb.User, res *pb.Response) (err error) {
	user, err := srv.Repo.Get(req)
	if err != nil {
		return err
	}
	res.User = user
	return err
}

// Create 创建用户
func (srv *User) Create(ctx context.Context, req *pb.User, res *pb.Response) (err error) {
	_, err = srv.Repo.Create(req)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("创建用户失败")
	}
	res.Valid = true
	return err
}

// Update 更新用户
func (srv *User) Update(ctx context.Context, req *pb.User, res *pb.Response) (err error) {
	valid, err := srv.Repo.Update(req)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("更新用户失败")
	}
	res.Valid = valid
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
