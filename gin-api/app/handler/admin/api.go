package admin

import (
	"github.com/google/wire"
	"github.com/micro/go-micro/v2/client"
	proto "github.com/ops-cn/go-devops/proto/admin"
)

// APISet 注入api
var APISet = wire.NewSet(
	LoginSet,
	MenuSet,
	RoleSet,
	UserSet,
)

var loginMgrClient proto.LoginMgrService

func Init() {
	loginMgrClient = proto.NewLoginMgrService("ops-cn.admin", client.DefaultClient)
}
