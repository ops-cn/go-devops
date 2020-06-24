package mock

import (
	"github.com/google/wire"
	"github.com/ops-cn/go-devops/gin-api/app/handler/admin"
)

// MockSet 注入mock
var MockSet = wire.NewSet(
	LoginSet,
	MenuSet,
	RoleSet,
	admin.UserSet,
)
