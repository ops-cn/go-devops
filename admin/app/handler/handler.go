package handler

import "github.com/google/wire"

// BllSet bll注入
var HandlerSet = wire.NewSet(
	LoginSet,
	MenuSet,
	RoleSet,
	UserSet,
)
