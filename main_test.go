package main

import (
	"context"
	"fmt"
	"testing"

	// 执行数据迁移
	_ "github.com/gomsa/socialite/database/migrations"
	"github.com/gomsa/socialite/hander"
	mpPB "github.com/gomsa/socialite/proto/miniprogram"
	userPB "github.com/gomsa/socialite/proto/user"
	db "github.com/gomsa/socialite/providers/database"
	"github.com/gomsa/socialite/service"
)

func TestAuth(t *testing.T) {
	repo := &service.UserRepository{db.DB}
	h := hander.Miniprogram{Repo: repo}
	req := &mpPB.Request{
		Code: `061oCnn32ERnlP0HXjm320vin32oCnn9`,
		Type: `wechat`,
	}
	res := &mpPB.Response{}
	err := h.Auth(context.TODO(), req, res)
	// fmt.Println(req, res, err)
	t.Log(req, res, err)
}

func TestBuild(t *testing.T) {
	repo := &service.UserRepository{db.DB}
	h := hander.User{Repo: repo}
	req := &userPB.User{
		Id: `d80d5d6d-c76c-4689-8956-5c84bd7455f4`,
	}
	res := &userPB.Response{}
	err := h.Build(context.TODO(), req, res)
	fmt.Println(req, res, err)
	t.Log(req, res, err)
}