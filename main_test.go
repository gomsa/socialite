package main

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/gomsa/socialite/database/migrations"
	"github.com/gomsa/socialite/hander"
	mpPB "github.com/gomsa/socialite/proto/miniprogram"
)

func TestAuth(t *testing.T) {

	h := hander.Miniprogram{}
	req := &mpPB.Request{
		Addressee: `13954386521`,
		Event:     `register_verify`,
		Type:      `sms`,
		QueryParams: map[string]string{
			`code`: `123456`,
		},
	}
	res := &mpPB.Response{}
	err := h.ProcessCommonRequest(context.TODO(), req, res)
	fmt.Println(req, res, err)
	t.Log(req, res, err)
}
