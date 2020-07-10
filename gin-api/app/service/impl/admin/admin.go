package admin

import "github.com/google/wire"

// ServiceSet 注入
var ServiceSet = wire.NewSet(
	LoginSet,
)
