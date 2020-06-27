package injector

import (
	"github.com/casbin/casbin/v2"
	"github.com/google/wire"
	"github.com/ops-cn/go-devops/admin/app/handler"
	"github.com/ops-cn/go-devops/common/auth"
)

// InjectorSet 注入Injector
var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

// Injector 注入器(用于初始化完成之后的引用)
type Injector struct {
	//Engine         *gin.Engine
	Auth           auth.Auther
	CasbinEnforcer *casbin.SyncedEnforcer
	//MenuBll bll.IMenu
	//MenuModel model.IMenu
	LoginAPI *handler.Login
	MenuAPI  *handler.Menu
	RoleAPI  *handler.Role
	UserAPI  *handler.User
}
