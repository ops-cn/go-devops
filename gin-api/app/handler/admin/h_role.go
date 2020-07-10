package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"github.com/ops-cn/go-devops/common/ginplus"
	"github.com/ops-cn/go-devops/common/schema"
	proto "github.com/ops-cn/go-devops/proto/admin"
)

// RoleSet 注入Role
var RoleSet = wire.NewSet(wire.Struct(new(Role), "*"))

// Role 角色管理
type Role struct {
}

func (roleWeb *Role) QueryFunc(c *gin.Context) (*schema.RoleQueryResult, error) {
	ctx := c.Request.Context()
	var params proto.RoleQueryParam

	if err := ginplus.ParseQuery(c, &params); err != nil {
		ginplus.ResError(c, err)
		return nil, err
	}

	orderFields := schema.NewOrderFields(schema.NewOrderField("sequence", schema.OrderByDESC))
	opts := &proto.RoleQueryOptions{}
	copier.Copy(opts, orderFields)

	params.PaginationParam.Pagination = true
	req := &proto.RoleQueryReq{}
	req.RoleQueryOptions = opts
	req.RoleQueryParam = &params
	rsp, err := roleMgrClient.Query(ctx, req)

	if err != nil {
		ginplus.ResError(c, err)
		return nil, err
	}

	pbResult := &proto.RoleQueryResult{}
	ptypes.UnmarshalAny(rsp.Items, pbResult)
	result := &schema.RoleQueryResult{}
	copier.Copy(result, pbResult)
	return result, nil
}

// Query 查询数据
func (roleWeb *Role) Query(c *gin.Context) {
	result, err := roleWeb.QueryFunc(c)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResPage(c, result.Data, result.PageResult)
}

// QuerySelect 查询选择数据
func (roleWeb *Role) QuerySelect(c *gin.Context) {
	result, err := roleWeb.QueryFunc(c)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResList(c, result.Data)
}

// Get 查询指定数据
func (roleWeb *Role) Get(c *gin.Context) {
	ctx := c.Request.Context()
	rsp, err := roleMgrClient.Get(ctx, &proto.RoleReq{
		CurrentID: c.Param("id"),
	})
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, rsp)
}

// Create 创建数据
func (roleWeb *Role) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var role proto.Role
	if err := ginplus.ParseJSON(c, &role); err != nil {
		ginplus.ResError(c, err)
		return
	}

	role.Creator = ginplus.GetUserID(c)
	rsp, err := roleMgrClient.Create(ctx, &role)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, rsp)
}

// Update 更新数据
func (roleWeb *Role) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var role proto.Role
	if err := ginplus.ParseJSON(c, &role); err != nil {
		ginplus.ResError(c, err)
		return
	}

	_, err := roleMgrClient.Update(ctx, &proto.RoleReq{
		CurrentID: c.Param("id"),
		Role:      &role,
	})
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}

// Delete 删除数据
func (roleWeb *Role) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	_, err := roleMgrClient.Delete(ctx, &proto.Role{
		ID: c.Param("id"),
	})
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}

// Enable 启用数据
func (roleWeb *Role) Enable(c *gin.Context) {
	ctx := c.Request.Context()
	_, err := roleMgrClient.UpdateStatus(ctx, &proto.Role{
		ID:     c.Param("id"),
		Status: 1,
	})
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}

// Disable 禁用数据
func (roleWeb *Role) Disable(c *gin.Context) {
	ctx := c.Request.Context()
	_, err := roleMgrClient.UpdateStatus(ctx, &proto.Role{
		ID:     c.Param("id"),
		Status: 2,
	})
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}
