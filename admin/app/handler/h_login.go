package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	proto "github.com/ops-cn/go-devops/proto/admin"
	"github.com/ops-cn/go-devops/proto/unified"
	"sort"

	"github.com/google/wire"
	"github.com/ops-cn/go-devops/admin/app/bll"
	"github.com/ops-cn/go-devops/admin/app/model"
	"github.com/ops-cn/go-devops/common/auth"
	"github.com/ops-cn/go-devops/common/errors"
	"github.com/ops-cn/go-devops/common/schema"
	"github.com/ops-cn/go-devops/common/thirdparty/captcha"
	"github.com/ops-cn/go-devops/common/util"
)

//type LoginService struct{}

var LoginSet = wire.NewSet(wire.Struct(new(LoginService), "*"), wire.Bind(new(bll.ILogin), new(*LoginService)))

// Login 登录管理
type LoginService struct {
	Auth            auth.Auther
	UserModel       model.IUser
	UserRoleModel   model.IUserRole
	RoleModel       model.IRole
	RoleMenuModel   model.IRoleMenu
	MenuModel       model.IMenu
	MenuActionModel model.IMenuAction
}

func (loginService *LoginService) GetCaptcha(ctx context.Context, length int) (*schema.LoginCaptcha, error) {
	captchaID := captcha.NewLen(length)
	item := &schema.LoginCaptcha{
		CaptchaID: captchaID,
	}
	return item, nil
}

// Verify 登录验证
func (loginService *LoginService) Verify(ctx context.Context, req *proto.LoginParam, res *unified.Response) error {
	// 检查是否是超级用户
	root := schema.GetRootUser()
	user := &proto.User{}
	if req.UserName == root.UserName && root.Password == req.Password {
		util.StructMapToStruct(&root, &user)
		res.Items, _ = ptypes.MarshalAny(user)
		return nil
	}

	result, err := loginService.UserModel.Query(ctx, schema.UserQueryParam{
		UserName: req.UserName,
	})
	if err != nil {
		return err
	} else if len(result.Data) == 0 {
		return errors.ErrInvalidUserName
	}

	item := result.Data[0]
	if item.Password != util.SHA1HashString(req.Password) {
		return errors.ErrInvalidPassword
	} else if item.Status != 1 {
		return errors.ErrUserDisable
	}
	util.StructMapToStruct(&item, &user)
	res.Items, _ = ptypes.MarshalAny(user)
	return nil
}

// GenerateToken 生成令牌
func (loginService *LoginService) GenerateToken(ctx context.Context, userID string) (*schema.LoginTokenInfo, error) {
	tokenInfo, err := loginService.Auth.GenerateToken(ctx, userID)
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
func (loginService *LoginService) DestroyToken(ctx context.Context, tokenString string) error {
	err := loginService.Auth.DestroyToken(ctx, tokenString)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (loginService *LoginService) CheckAndGetUser(ctx context.Context, req *proto.UserLoginInfo, res *unified.Response) error {
	item, err := loginService.UserModel.Get(ctx, req.UserID)
	if err != nil {
		return err
	} else if item == nil {
		return errors.ErrInvalidUser
	} else if item.Status != 1 {
		return errors.ErrUserDisable
	}
	user := &proto.User{}
	util.StructMapToStruct(&item, &user)
	res.Items, _ = ptypes.MarshalAny(user)
	return nil
}

// GetLoginInfo 获取当前用户登录信息
func (loginService *LoginService) GetLoginInfo(ctx context.Context, req *proto.UserLoginInfo, res *unified.Response) error {
	if isRoot := schema.CheckIsRootUser(ctx, req.UserID); isRoot {
		root := schema.GetRootUser()
		loginInfo := &proto.UserLoginInfo{
			UserName: root.UserName,
			RealName: root.RealName,
		}

		res.Items, _ = ptypes.MarshalAny(loginInfo)
		return nil
	}

	user, err := loginService.UserModel.Get(ctx, req.UserID)
	if err != nil {
		return err
	}

	info := &proto.UserLoginInfo{
		UserID:   user.ID,
		UserName: user.UserName,
		RealName: user.RealName,
	}

	userRoleResult, err := loginService.UserRoleModel.Query(ctx, schema.UserRoleQueryParam{
		UserID: req.UserID,
	})
	if err != nil {
		return err
	}

	if roleIDs := userRoleResult.Data.ToRoleIDs(); len(roleIDs) > 0 {
		roleResult, err := loginService.RoleModel.Query(ctx, schema.RoleQueryParam{
			IDs:    roleIDs,
			Status: 1,
		})
		if err != nil {
			return err
		}
		var roles []*proto.Role
		for _, v := range roleResult.Data {
			role := &proto.Role{}
			util.StructMapToStruct(&v, &role)
			roles = append(roles, role)
		}
		info.Roles = roles
	}
	res.Items, _ = ptypes.MarshalAny(info)

	return nil
}

// QueryUserMenuTree 查询当前用户的权限菜单树
func (loginService *LoginService) QueryUserMenuTree(ctx context.Context, req *proto.UserLoginInfo, res *unified.Response) error {
	isRoot := schema.CheckIsRootUser(ctx, req.UserID)
	// 如果是root用户，则查询所有显示的菜单树
	if isRoot {
		result, err := loginService.MenuModel.Query(ctx, schema.MenuQueryParam{
			Status: 1,
		}, schema.MenuQueryOptions{
			OrderFields: schema.NewOrderFields(schema.NewOrderField("sequence", schema.OrderByDESC)),
		})
		if err != nil {
			return err
		}

		menuActionResult, err := loginService.MenuActionModel.Query(ctx, schema.MenuActionQueryParam{})
		if err != nil {
			return err
		}
		result.Data.FillMenuAction(menuActionResult.Data.ToMenuIDMap()).ToTree()
		return nil
	}

	userRoleResult, err := loginService.UserRoleModel.Query(ctx, schema.UserRoleQueryParam{
		UserID: req.UserID,
	})
	if err != nil {
		return err
	} else if len(userRoleResult.Data) == 0 {
		return errors.ErrNoPerm
	}

	roleMenuResult, err := loginService.RoleMenuModel.Query(ctx, schema.RoleMenuQueryParam{
		RoleIDs: userRoleResult.Data.ToRoleIDs(),
	})
	if err != nil {
		return err
	} else if len(roleMenuResult.Data) == 0 {
		return errors.ErrNoPerm
	}

	menuResult, err := loginService.MenuModel.Query(ctx, schema.MenuQueryParam{
		IDs:    roleMenuResult.Data.ToMenuIDs(),
		Status: 1,
	})
	if err != nil {
		return err
	} else if len(menuResult.Data) == 0 {
		return errors.ErrNoPerm
	}

	mData := menuResult.Data.ToMap()
	var qIDs []string
	for _, pid := range menuResult.Data.SplitParentIDs() {
		if _, ok := mData[pid]; !ok {
			qIDs = append(qIDs, pid)
		}
	}

	if len(qIDs) > 0 {
		pmenuResult, err := loginService.MenuModel.Query(ctx, schema.MenuQueryParam{
			IDs: menuResult.Data.SplitParentIDs(),
		})
		if err != nil {
			return err
		}
		menuResult.Data = append(menuResult.Data, pmenuResult.Data...)
	}

	sort.Sort(menuResult.Data)
	menuActionResult, err := loginService.MenuActionModel.Query(ctx, schema.MenuActionQueryParam{
		IDs: roleMenuResult.Data.ToActionIDs(),
	})
	if err != nil {
		return err
	}
	menuResult.Data.FillMenuAction(menuActionResult.Data.ToMenuIDMap()).ToTree()
	return nil
}

// UpdatePassword 更新当前用户登录密码
func (loginService *LoginService) UpdatePassword(ctx context.Context, req *proto.UpdatePasswordParam, res *unified.Response) error {
	if schema.CheckIsRootUser(ctx, req.UserID) {
		return errors.New400Response("root用户不允许更新密码")
	}

	user, err := loginService.UserModel.Get(ctx, req.UserID)
	if err != nil {
		return err
	} else if util.SHA1HashString(req.OldPassword) != user.Password {
		return errors.New400Response("旧密码不正确")
	}

	req.NewPassword = util.SHA1HashString(req.NewPassword)
	err = loginService.UserModel.UpdatePassword(ctx, req.UserID, req.NewPassword)

	return err
}
