package noworker

import (
	"github.com/ops-cn/go-devops/common/config"
	"github.com/ops-cn/go-devops/common/logger"
	"github.com/ops-cn/go-devops/common/noworker/unique"
	"github.com/ops-cn/go-devops/common/trace"
)

var idFunc = func() string {
	return unique.NewSnowflakeID().String()
}

// InitID ...
func InitID() {
	switch config.C.UniqueID.Type {
	case "uuid":
		idFunc = func() string {
			return unique.MustUUID().String()
		}
	case "object":
		idFunc = func() string {
			return unique.NewObjectID().Hex()
		}
	default:
		// Initialize snowflake node
		err := unique.SetSnowflakeNode(config.C.UniqueID.Snowflake.Node, config.C.UniqueID.Snowflake.Epoch)
		if err != nil {
			panic(err)
		}

		logger.SetTraceIDFunc(func() string {
			return unique.NewSnowflakeID().String()
		})

		trace.SetIDFunc(func() string {
			return unique.NewSnowflakeID().String()
		})

		idFunc = func() string {
			return unique.NewSnowflakeID().String()
		}
	}
}

// NewID Create unique id
func NewID() string {
	return idFunc()
}
