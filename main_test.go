package main

import (
	"context"
	"fmt"
	"testing"

	// 执行数据迁移
	_ "github.com/gomsa/socialite/database/migrations"
	"github.com/gomsa/socialite/hander"
	mpPB "github.com/gomsa/socialite/proto/miniprogram"
)

func TestAuth(t *testing.T) {

	h := hander.Miniprogram{}
	req := &mpPB.Request{
		Code: `061oCnn32ERnlP0HXjm320vin32oCnn9`,
		Type: `wechat`,
	}
	res := &mpPB.Response{}
	err := h.Auth(context.TODO(), req, res)
	fmt.Println(req, res, err)
	t.Log(req, res, err)
}
