package main

import (
	// 公共引入
	"github.com/micro/go-log"
	micro "github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"

	// 执行数据迁移
	_ "github.com/gomsa/socialite/database/migrations"
	"github.com/gomsa/socialite/hander"
	"github.com/gomsa/socialite/service"
	mpPB "github.com/gomsa/socialite/proto/miniprogram"
	userPB "github.com/gomsa/socialite/proto/user"
	db "github.com/gomsa/socialite/providers/database"
)

func main() {
	srv := k8s.NewService(
		micro.Name(Conf.Service),
		micro.Version(Conf.Version),
	)
	srv.Init()

	// 用户服务实现
	repo := &service.UserRepository{db.DB}
	userPB.RegisterUsersHandler(srv.Server(), &hander.User{repo})

	// 小程序服务实现
	mpPB.RegisterMiniprogramHandler(srv.Server(), &hander.Miniprogram{Repo:repo})

	// Run the server
	if err := srv.Run(); err != nil {
		log.Log(err)
	}
	log.Log("serviser run ...")
}
