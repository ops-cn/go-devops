// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package injector

import (
	api "github.com/ops-cn/go-devops/gin-api/app/handler/admin"
	// "github.com/ops-cn/go-devops/admin/app/api/mock"
	"github.com/google/wire"
	"github.com/ops-cn/go-devops/gin-api/app/module/adapter"
	"github.com/ops-cn/go-devops/gin-api/app/router"
	service "github.com/ops-cn/go-devops/gin-api/app/service/impl/admin"

	// mongoModel "github.com/ops-cn/go-devops/admin/app/model/impl/mongo/model"
	gormModel "github.com/ops-cn/go-devops/admin/app/model/impl/gorm/model"
)

// BuildInjector 生成注入器
func BuildInjector() (*Injector, func(), error) {
	// 默认使用gorm存储注入，这里可使用 InitMongoDB & mongoModel.ModelSet 替换为 gorm 存储
	wire.Build(
		InitGormDB,
		gormModel.ModelSet,
		// InitMongoDB,
		// mongoModel.ModelSet,
		InitAuth,
		InitCasbin,
		InitGinEngine,
		service.ServiceSet,
		api.APISet,
		// mock.MockSet,
		router.RouterSet,
		adapter.CasbinAdapterSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
