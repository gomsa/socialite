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
		Code: `13954386521`,
		Type: `register_verify`,
	}
	res := &mpPB.Response{}
	err := h.tAuth(context.TODO(), req, res)
	fmt.Println(req, res, err)
	t.Log(req, res, err)
}
