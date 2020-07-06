package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/wire"
	"github.com/ops-cn/go-devops/common/config"
	"github.com/ops-cn/go-devops/common/errors"
	"github.com/ops-cn/go-devops/common/ginplus"
	"github.com/ops-cn/go-devops/common/logger"
	"github.com/ops-cn/go-devops/common/schema"
	"github.com/ops-cn/go-devops/common/thirdparty/captcha"
	"github.com/ops-cn/go-devops/gin-api/app/service/admin"
	proto "github.com/ops-cn/go-devops/proto/admin"
)

// LoginSet 注入Login
var LoginSet = wire.NewSet(wire.Struct(new(Login), "*"))

// Login 登录管理
type Login struct {
	LoginBll admin.ILogin
}

// GetCaptcha 获取验证码信息
func (a *Login) GetCaptcha(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.LoginBll.GetCaptcha(ctx, config.C.Captcha.Length)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	logger.Printf(ctx, item.CaptchaID)
	ginplus.ResSuccess(c, item)
}

// ResCaptcha 响应图形验证码
func (a *Login) ResCaptcha(c *gin.Context) {
	ctx := c.Request.Context()
	captchaID := c.Query("id")
	if captchaID == "" {
		ginplus.ResError(c, errors.New400Response("请提供验证码ID"))
		return
	}

	if c.Query("reload") != "" {
		if !captcha.Reload(captchaID) {
			ginplus.ResError(c, errors.New400Response("未找到验证码ID"))
			return
		}
	}

	cfg := config.C.Captcha
	err := a.LoginBll.ResCaptcha(ctx, c.Writer, captchaID, cfg.Width, cfg.Height)
	if err != nil {
		ginplus.ResError(c, err)
	}
}

// Login 用户登录
func (a *Login) Login(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.LoginParam
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}
	logger.Printf(ctx, item.UserName)
	/*if !captcha.VerifyString(item.CaptchaID, item.CaptchaCode) {
		ginplus.ResError(c, errors.New400Response("无效的验证码"))
		return
	}*/
	param := &proto.LoginParam{
		UserName: item.UserName,
		Password: item.Password,
	}
	rsp, err := loginClient.Verify(ctx, param)
	fmt.Print(rsp)
	user := &proto.User{}
	ptypes.UnmarshalAny(rsp.Items, user)
	//user, err := a.LoginBll.Verify(ctx, item.UserName, item.Password)

	if err != nil {
		ginplus.ResError(c, err)
		return
	}

	userID := user.ID
	// 将用户ID放入上下文
	ginplus.SetUserID(c, userID)

	ctx = logger.NewUserIDContext(ctx, userID)
	tokenInfo, err := a.LoginBll.GenerateToken(ctx, userID)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}

	logger.StartSpan(ctx, logger.SetSpanTitle("用户登录"), logger.SetSpanFuncName("Login")).Infof("登入系统")
	ginplus.ResSuccess(c, tokenInfo)
}

// Logout 用户登出
func (a *Login) Logout(c *gin.Context) {
	ctx := c.Request.Context()
	// 检查用户是否处于登录状态，如果是则执行销毁
	userID := ginplus.GetUserID(c)
	if userID != "" {
		err := a.LoginBll.DestroyToken(ctx, ginplus.GetToken(c))
		if err != nil {
			logger.Errorf(ctx, err.Error())
		}
		logger.StartSpan(ctx, logger.SetSpanTitle("用户登出"), logger.SetSpanFuncName("Logout")).Infof("登出系统")
	}
	ginplus.ResOK(c)
}

// RefreshToken 刷新令牌
func (a *Login) RefreshToken(c *gin.Context) {
	ctx := c.Request.Context()
	tokenInfo, err := a.LoginBll.GenerateToken(ctx, ginplus.GetUserID(c))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, tokenInfo)
}

// GetUserInfo 获取当前用户信息
func (a *Login) GetUserInfo(c *gin.Context) {
	ctx := c.Request.Context()
	param := &proto.UserLoginInfo{
		UserID: ginplus.GetUserID(c),
	}
	rsp, err := loginClient.GetLoginInfo(ctx, param)
	userInfo := &proto.UserLoginInfo{}
	ptypes.UnmarshalAny(rsp.Items, userInfo)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, userInfo)
}

// QueryUserMenuTree 查询当前用户菜单树
func (a *Login) QueryUserMenuTree(c *gin.Context) {
	ctx := c.Request.Context()
	param := &proto.UserLoginInfo{
		UserID: ginplus.GetUserID(c),
	}
	rsp, err := loginClient.QueryUserMenuTree(ctx, param)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	menuTrees := &proto.MenuTrees{}
	//这里又错误，未正确映射menuTrees的值
	ptypes.UnmarshalAny(rsp.Items, menuTrees)

	ginplus.ResList(c, menuTrees)
}

// UpdatePassword 更新个人密码
func (a *Login) UpdatePassword(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.UpdatePasswordParam
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}
	param := &proto.UpdatePasswordParam{
		UserID:      ginplus.GetUserID(c),
		OldPassword: item.OldPassword,
		NewPassword: item.NewPassword,
	}

	_, err := loginClient.UpdatePassword(ctx, param)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}
