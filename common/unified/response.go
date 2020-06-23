package unified

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/micro/go-micro/v2/metadata"
	"net/http"
	"strings"

	"github.com/ops-cn/go-devops/common/errors"
	"github.com/ops-cn/go-devops/common/logger"
	"github.com/ops-cn/go-devops/common/schema"
	"github.com/ops-cn/go-devops/common/util"
)

// 定义上下文中的键
const (
	prefix           = "ops-cn"
	UserIDKey        = prefix + "/user-id"
	ReqBodyKey       = prefix + "/req-body"
	ResBodyKey       = prefix + "/res-body"
	LoggerReqBodyKey = prefix + "/logger-req-body"
)

const (
	BASIC_SCHEMA  string = "Basic "
	BEARER_SCHEMA string = "Bearer "
)

// GetToken 获取用户令牌
func GetToken(ctx context.Context) string {
	md, _ := metadata.FromContext(ctx)
	var token string
	auth := md["Authorization"]
	if auth != "" && strings.HasPrefix(auth, BEARER_SCHEMA) {
		token = auth[len(prefix):]
	}
	return token
}

// GetUserID 获取用户ID
func GetUserID(ctx context.Context) string {
	md, _ := metadata.FromContext(ctx)
	return md[UserIDKey]
}

// SetUserID 设定用户ID
func SetUserID(ctx context.Context, userID string) {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = metadata.Metadata{}
	}
	md[UserIDKey] = userID
}

// ParseJSON 解析请求JSON
func ParseJSON(ctx context.Context, obj interface{}) error {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = metadata.Metadata{}
	}
	m := make(map[string]string)
	json.Unmarshal(obj, m)
	if err := c.ShouldBindJSON(obj); err != nil {
		return errors.Wrap400Response(err, fmt.Sprintf("解析请求参数发生错误 - %s", err.Error()))
	}
	return nil
}

// ParseQuery 解析Query参数
func ParseQuery(c *gin.Context, obj interface{}) error {
	md, ok := metadata.FromContext(c)
	if err := c.ShouldBindQuery(obj); err != nil {
		return errors.Wrap400Response(err, fmt.Sprintf("解析请求参数发生错误 - %s", err.Error()))
	}
	return nil
}

// ParseForm 解析Form请求
func ParseForm(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindWith(obj, binding.Form); err != nil {
		return errors.Wrap400Response(err, fmt.Sprintf("解析请求参数发生错误 - %s", err.Error()))
	}
	return nil
}

// ResOK 响应OK
func ResOK(c *gin.Context) {
	ResSuccess(c, schema.StatusResult{Status: schema.OKStatus})
}

// ResList 响应列表数据
func ResList(c *gin.Context, v interface{}) {
	ResSuccess(c, schema.ListResult{List: v})
}

// ResPage 响应分页数据
func ResPage(c *gin.Context, v interface{}, pr *schema.PaginationResult) {
	list := schema.ListResult{
		List:       v,
		Pagination: pr,
	}
	ResSuccess(c, list)
}

// ResSuccess 响应成功
func ResSuccess(c *gin.Context, v interface{}) {
	ResJSON(c, http.StatusOK, v)
}

// ResJSON 响应JSON数据
func ResJSON(c *gin.Context, status int, v interface{}) {
	buf, err := util.JSONMarshal(v)
	if err != nil {
		panic(err)
	}
	c.Set(ResBodyKey, buf)
	c.Data(status, "application/json; charset=utf-8", buf)
	c.Abort()
}

// ResError 响应错误
func ResError(c *gin.Context, err error, status ...int) {
	ctx := c.Request.Context()
	var res *errors.ResponseError
	if err != nil {
		if e, ok := err.(*errors.ResponseError); ok {
			res = e
		} else {
			res = errors.UnWrapResponse(errors.Wrap500Response(err, "服务器错误"))
		}
	} else {
		res = errors.UnWrapResponse(errors.ErrInternalServer)
	}

	if len(status) > 0 {
		res.StatusCode = status[0]
	}

	if err := res.ERR; err != nil {
		if status := res.StatusCode; status >= 400 && status < 500 {
			logger.StartSpan(ctx).Warnf(err.Error())
		} else if status >= 500 {
			logger.ErrorStack(ctx, err)
		}
	}

	eitem := schema.ErrorItem{
		Code:    res.Code,
		Message: res.Message,
	}
	ResJSON(c, res.StatusCode, schema.ErrorResult{Error: eitem})
}
