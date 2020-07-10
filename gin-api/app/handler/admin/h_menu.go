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

// MenuSet 注入Menu
var MenuSet = wire.NewSet(wire.Struct(new(Menu), "*"))

// Menu 菜单管理
type Menu struct {
}

// Query 查询数据
func (menuWeb *Menu) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params proto.MenuQueryParam

	if err := ginplus.ParseQuery(c, &params); err != nil {
		ginplus.ResError(c, err)
		return
	}

	orderFields := schema.NewOrderFields(schema.NewOrderField("sequence", schema.OrderByDESC))
	opts := &proto.MenuQueryOptions{}
	copier.Copy(opts, orderFields)

	params.PaginationParam.Pagination = true
	req := &proto.MenuQueryReq{}
	req.MenuQueryOptions = opts
	req.MenuQueryParam = &params
	rsp, err := menuMgrClient.Query(ctx, req)

	if err != nil {
		ginplus.ResError(c, err)
		return
	}

	pbResult := &proto.MenuQueryResult{}
	ptypes.UnmarshalAny(rsp.Items, pbResult)
	result := &schema.MenuQueryResult{}
	copier.Copy(result, pbResult)
	ginplus.ResPage(c, result.Data, result.PageResult)
}

// QueryTree 查询菜单树
func (menuWeb *Menu) QueryTree(c *gin.Context) {
	ctx := c.Request.Context()
	var params proto.MenuQueryParam
	if err := ginplus.ParseQuery(c, &params); err != nil {
		ginplus.ResError(c, err)
		return
	}
	orderFields := schema.NewOrderFields(schema.NewOrderField("sequence", schema.OrderByDESC))
	opts := &proto.MenuQueryOptions{}
	copier.Copy(opts, orderFields)

	params.PaginationParam.Pagination = true
	req := &proto.MenuQueryReq{}
	req.MenuQueryOptions = opts
	req.MenuQueryParam = &params
	rsp, err := menuMgrClient.Query(ctx, req)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	pbResult := &proto.MenuQueryResult{}
	ptypes.UnmarshalAny(rsp.Items, pbResult)
	result := &schema.MenuQueryResult{}
	copier.Copy(result, pbResult)
	ginplus.ResList(c, result.Data.ToTree())
}

// Get 查询指定数据
func (menuWeb *Menu) Get(c *gin.Context) {
	ctx := c.Request.Context()

	rsp, err := menuMgrClient.Get(ctx,
		&proto.MenuReq{
			CurrentID: c.Param("id"),
		})
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, rsp)
}

// Create 创建数据
func (menuWeb *Menu) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var menu proto.Menu
	if err := ginplus.ParseJSON(c, &menu); err != nil {
		ginplus.ResError(c, err)
		return
	}

	menu.Creator = ginplus.GetUserID(c)
	rsp, err := menuMgrClient.Create(ctx, &menu)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, rsp)
}

// Update 更新数据
func (menuWeb *Menu) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var menu proto.Menu
	if err := ginplus.ParseJSON(c, &menu); err != nil {
		ginplus.ResError(c, err)
		return
	}

	_, err := menuMgrClient.Update(ctx,
		&proto.MenuReq{
			CurrentID: c.Param("id"),
			Menu:      &menu,
		})
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}

// Delete 删除数据
func (menuWeb *Menu) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	_, err := menuMgrClient.Delete(ctx,
		&proto.Menu{
			ID: c.Param("id"),
		})
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}

// Enable 启用数据
func (menuWeb *Menu) Enable(c *gin.Context) {
	ctx := c.Request.Context()
	_, err := menuMgrClient.UpdateStatus(ctx, &proto.Menu{
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
func (menuWeb *Menu) Disable(c *gin.Context) {
	ctx := c.Request.Context()
	_, err := menuMgrClient.UpdateStatus(ctx, &proto.Menu{
		ID:     c.Param("id"),
		Status: 2,
	})
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}
