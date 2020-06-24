package injector

import (
	"github.com/google/wire"
	"github.com/ops-cn/go-devops/admin/app/bll"
)

// InjectorSet 注入Injector
var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

// Injector 注入器(用于初始化完成之后的引用)
type Injector struct {
	//Engine         *gin.Engine
	//Auth           auth.Auther
	//CasbinEnforcer *casbin.SyncedEnforcer
	MenuBll bll.IMenu
}
