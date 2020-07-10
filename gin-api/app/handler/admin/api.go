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
var menuMgrClient proto.MenuMgrService
var roleMgrClient proto.RoleMgrService
var userMgrClient proto.UserMgrService

func Init() {

	loginMgrClient = proto.NewLoginMgrService("ops-cn.admin", client.DefaultClient)
	menuMgrClient = proto.NewMenuMgrService("ops-cn.admin", client.DefaultClient)
	roleMgrClient = proto.NewRoleMgrService("ops-cn.admin", client.DefaultClient)
	userMgrClient = proto.NewUserMgrService("ops-cn.admin", client.DefaultClient)
}
