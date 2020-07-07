package admin

import (
	"context"
	"github.com/ops-cn/go-devops/gin-api/app/service/admin"
	"net/http"

	"github.com/google/wire"
	"github.com/ops-cn/go-devops/common/auth"
	"github.com/ops-cn/go-devops/common/errors"
	"github.com/ops-cn/go-devops/common/schema"
	"github.com/ops-cn/go-devops/common/thirdparty/captcha"
)

var _ admin.ILogin = (*Login)(nil)

// LoginSet 注入Login
var LoginSet = wire.NewSet(wire.Struct(new(Login), "*"), wire.Bind(new(admin.ILogin), new(*Login)))

// Login 登录管理
type Login struct {
	Auth auth.Auther
}

// GetCaptcha 获取图形验证码信息
func (a *Login) GetCaptcha(ctx context.Context, length int) (*schema.LoginCaptcha, error) {
	captchaID := captcha.NewLen(length)
	item := &schema.LoginCaptcha{
		CaptchaID: captchaID,
	}
	return item, nil
}

// ResCaptcha 生成并响应图形验证码
func (a *Login) ResCaptcha(ctx context.Context, w http.ResponseWriter, captchaID string, width, height int) error {
	err := captcha.WriteImage(w, captchaID, width, height)
	if err != nil {
		if err == captcha.ErrNotFound {
			return errors.ErrNotFound
		}
		return errors.WithStack(err)
	}

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Type", "image/png")
	return nil
}

// GenerateToken 生成令牌
func (a *Login) GenerateToken(ctx context.Context, userID string) (*schema.LoginTokenInfo, error) {
	tokenInfo, err := a.Auth.GenerateToken(ctx, userID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	item := &schema.LoginTokenInfo{
		AccessToken: tokenInfo.GetAccessToken(),
		TokenType:   tokenInfo.GetTokenType(),
		ExpiresAt:   tokenInfo.GetExpiresAt(),
	}
	return item, nil
}

// DestroyToken 销毁令牌
func (a *Login) DestroyToken(ctx context.Context, tokenString string) error {
	err := a.Auth.DestroyToken(ctx, tokenString)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
