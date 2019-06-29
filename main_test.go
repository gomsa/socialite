package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/gomsa/socialite/hander"
	mpPB "github.com/gomsa/socialite/proto/miniprogram"
)

func TestAuth(t *testing.T) {

	h := hander.Miniprogram{}
	req := &mpPB.Request{
		Code: `001tiGET1XmBs41UhpET1x3mET1tiGEs`,
		Type: `wechat`,
	}
	res := &mpPB.Response{}
	err := h.Auth(context.TODO(), req, res)
	fmt.Println(req, res, err)
	t.Log(req, res, err)
}
