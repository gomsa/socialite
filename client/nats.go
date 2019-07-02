package client

import (
	"os"

	"github.com/gomsa/tools/k8s/client"

	miniprogramPB "github.com/gomsa/socialite/proto/miniprogram"
	userPB "github.com/gomsa/socialite/proto/user"
)

var (
	// Miniprogram 社会登录小程序客户端
	Miniprogram miniprogramPB.MiniprogramClient
	// User 社会登录用户管理客户端
	User userPB.UsersClient
)

func init() {
	srvName := os.Getenv("SOCIALITE_NAME")

	Miniprogram = miniprogramPB.NewMiniprogramClient(srvName, client.DefaultClient)
	User = userPB.NewUsersClient(srvName, client.DefaultClient)
}
