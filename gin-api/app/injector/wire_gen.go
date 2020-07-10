// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package injector

import (
	"github.com/ops-cn/go-devops/admin/app/model/impl/gorm/model"
	"github.com/ops-cn/go-devops/admin/app/module/adapter"
	admin2 "github.com/ops-cn/go-devops/gin-api/app/handler/admin"
	"github.com/ops-cn/go-devops/gin-api/app/router"
	"github.com/ops-cn/go-devops/gin-api/app/service/impl/admin"
)

// Injectors from wire.go:

func BuildInjector() (*Injector, func(), error) {
	auther, cleanup, err := InitAuth()
	if err != nil {
		return nil, nil, err
	}
	db, cleanup2, err := InitGormDB()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	role := &model.Role{
		DB: db,
	}
	roleMenu := &model.RoleMenu{
		DB: db,
	}
	menuActionResource := &model.MenuActionResource{
		DB: db,
	}
	user := &model.User{
		DB: db,
	}
	userRole := &model.UserRole{
		DB: db,
	}
	casbinAdapter := &adapter.CasbinAdapter{
		RoleModel:         role,
		RoleMenuModel:     roleMenu,
		MenuResourceModel: menuActionResource,
		UserModel:         user,
		UserRoleModel:     userRole,
	}
	syncedEnforcer, cleanup3, err := InitCasbin(casbinAdapter)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	login := &admin.Login{
		Auth: auther,
	}
	adminLogin := &admin2.Login{
		LoginBll: login,
	}
	menu := &admin2.Menu{}
	adminRole := &admin2.Role{}
	adminUser := &admin2.User{}
	routerRouter := &router.Router{
		Auth:           auther,
		CasbinEnforcer: syncedEnforcer,
		LoginAPI:       adminLogin,
		MenuAPI:        menu,
		RoleAPI:        adminRole,
		UserAPI:        adminUser,
	}
	engine := InitGinEngine(routerRouter)
	injector := &Injector{
		Engine:         engine,
		Auth:           auther,
		CasbinEnforcer: syncedEnforcer,
		MenuWeb:        menu,
	}
	return injector, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
