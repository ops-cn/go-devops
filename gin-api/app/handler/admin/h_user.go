package admin

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/copier"
	proto "github.com/ops-cn/go-devops/proto/admin"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/ops-cn/go-devops/common/errors"
	"github.com/ops-cn/go-devops/common/ginplus"
	"github.com/ops-cn/go-devops/common/schema"
)

// UserSet 注入User
var UserSet = wire.NewSet(wire.Struct(new(User), "*"))

// User 用户管理
type User struct {
}

// Query 查询数据
func (userWeb *User) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params proto.UserQueryParam
	if err := ginplus.ParseQuery(c, &params); err != nil {
		ginplus.ResError(c, err)
		return
	}
	if v := c.Query("roleIDs"); v != "" {
		params.RoleIDs = strings.Split(v, ",")
	}
	orderFields := schema.NewOrderFields(schema.NewOrderField("sequence", schema.OrderByDESC))
	opts := &proto.UserQueryOptions{}
	copier.Copy(opts, orderFields)

	params.PaginationParam.Pagination = true
	req := &proto.UserQueryReq{}
	req.UserQueryOptions = opts
	req.UserQueryParam = &params
	rsp, err := userMgrClient.Query(ctx, req)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	pbResult := &proto.UserQueryResult{}
	ptypes.UnmarshalAny(rsp.Items, pbResult)
	result := &schema.UserQueryResult{}

	ginplus.ResPage(c, result.Data, result.PageResult)
}

// Get 查询指定数据
func (userWeb *User) Get(c *gin.Context) {
	ctx := c.Request.Context()
	rsp, err := userMgrClient.Get(ctx, &proto.UserReq{
		CurrentID: c.Param("id"),
	})
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	pbUser := &proto.User{}
	ptypes.UnmarshalAny(rsp.Items, pbUser)
	user := schema.User{}
	copier.Copy(user, pbUser)
	ginplus.ResSuccess(c, user.CleanSecure())
}

// Create 创建数据
func (userWeb *User) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var user proto.User
	if err := ginplus.ParseJSON(c, &user); err != nil {
		ginplus.ResError(c, err)
		return
	} else if user.Password == "" {
		ginplus.ResError(c, errors.New400Response("密码不能为空"))
		return
	}

	user.Creator = ginplus.GetUserID(c)
	rsp, err := userMgrClient.Create(ctx, &user)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, rsp)
}

// Update 更新数据
func (userWeb *User) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var user proto.User
	if err := ginplus.ParseJSON(c, &user); err != nil {
		ginplus.ResError(c, err)
		return
	}

	_, err := userMgrClient.Update(ctx, &proto.UserReq{
		CurrentID: c.Param("id"),
		User:      &user,
	})
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}

// Delete 删除数据
func (userWeb *User) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	_, err := userMgrClient.Delete(ctx, &proto.User{
		ID: c.Param("id"),
	})
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}

// Enable 启用数据
func (userWeb *User) Enable(c *gin.Context) {
	ctx := c.Request.Context()
	_, err := userMgrClient.UpdateStatus(ctx, &proto.User{
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
func (userWeb *User) Disable(c *gin.Context) {
	ctx := c.Request.Context()
	_, err := userMgrClient.UpdateStatus(ctx, &proto.User{
		ID:     c.Param("id"),
		Status: 2,
	})
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}
