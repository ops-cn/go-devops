package admin

import (
	"context"
	"net/http"

	"github.com/ops-cn/go-devops/common/schema"
)

// ILogin 登录业务逻辑接口
type ILogin interface {
	// 获取图形验证码信息
	GetCaptcha(ctx context.Context, length int) (*schema.LoginCaptcha, error)
	// 生成并响应图形验证码
	ResCaptcha(ctx context.Context, w http.ResponseWriter, captchaID string, width, height int) error
}
