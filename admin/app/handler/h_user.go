package handler

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"github.com/ops-cn/go-devops/admin/app/model"
	"github.com/ops-cn/go-devops/common/errors"
	"github.com/ops-cn/go-devops/common/noworker"
	"github.com/ops-cn/go-devops/common/schema"
	"github.com/ops-cn/go-devops/common/util"
	proto "github.com/ops-cn/go-devops/proto/admin"
	"github.com/ops-cn/go-devops/proto/unified"
)

// UserSet 注入User
var UserSet = wire.NewSet(wire.Struct(new(User), "*"))

// User 用户管理
type User struct {
	Enforcer      *casbin.SyncedEnforcer
	TransModel    model.ITrans
	UserModel     model.IUser
	UserRoleModel model.IUserRole
	RoleModel     model.IRole
}

// Query 查询数据
func (userMgr *User) Query(ctx context.Context, req *proto.UserQueryReq, res *unified.Response) error {
	param := schema.UserQueryParam{}
	copier.Copy(param, req.UserQueryParam)

	opts := schema.UserQueryOptions{}
	copier.Copy(opts, req.UserQueryOptions)
	result, err := userMgr.UserModel.Query(ctx, param, opts)
	if err != nil {
		return err
	}
	pbResult := &proto.UserQueryResult{}
	util.StructCopy(pbResult, result)
	res.Items, _ = ptypes.MarshalAny(pbResult)
	return nil
}

// QueryShow 查询显示项数据
func (userMgr *User) QueryShow(ctx context.Context, params schema.UserQueryParam, opts ...schema.UserQueryOptions) (*schema.UserShowQueryResult, error) {
	result, err := userMgr.UserModel.Query(ctx, params, opts...)
	if err != nil {
		return nil, err
	} else if result == nil {
		return nil, nil
	}

	userRoleResult, err := userMgr.UserRoleModel.Query(ctx, schema.UserRoleQueryParam{
		UserIDs: result.Data.ToIDs(),
	})
	if err != nil {
		return nil, err
	}

	roleResult, err := userMgr.RoleModel.Query(ctx, schema.RoleQueryParam{
		IDs: userRoleResult.Data.ToRoleIDs(),
	})
	if err != nil {
		return nil, err
	}

	return result.ToShowResult(userRoleResult.Data.ToUserIDMap(), roleResult.Data.ToMap()), nil
}

// Get 查询指定数据
func (userMgr *User) GetUser(ctx context.Context, id string, opts ...schema.UserQueryOptions) (*schema.User, error) {
	item, err := userMgr.UserModel.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	userRoleResult, err := userMgr.UserRoleModel.Query(ctx, schema.UserRoleQueryParam{
		UserID: id,
	})
	if err != nil {
		return nil, err
	}
	item.UserRoles = userRoleResult.Data

	return item, nil
}
func (userMgr *User) Get(ctx context.Context, req *proto.UserReq, res *unified.Response) error {
	opts := schema.UserQueryOptions{}
	copier.Copy(opts, req.UserQueryOptions)
	user, err := userMgr.UserModel.Get(ctx, req.CurrentID, opts)
	if err != nil {
		return err
	} else if user == nil {
		return errors.ErrNotFound
	}

	userRoleResult, err := userMgr.UserRoleModel.Query(ctx, schema.UserRoleQueryParam{
		UserID: req.CurrentID,
	})
	if err != nil {
		return err
	}
	user.UserRoles = userRoleResult.Data
	pbUser := &proto.User{}
	copier.Copy(pbUser, user)
	res.Items, _ = ptypes.MarshalAny(pbUser)
	return nil
}

// Create 创建数据
func (userMgr *User) Create(ctx context.Context, req *proto.User, res *unified.Response) error {

	user := schema.User{}
	copier.Copy(user, req)
	err := userMgr.checkUserName(ctx, user)
	if err != nil {
		return err
	}

	user.Password = util.SHA1HashString(user.Password)
	user.ID = noworker.NewID()
	err = ExecTrans(ctx, userMgr.TransModel, func(ctx context.Context) error {
		for _, urItem := range user.UserRoles {
			urItem.ID = noworker.NewID()
			urItem.UserID = user.ID
			err := userMgr.UserRoleModel.Create(ctx, *urItem)
			if err != nil {
				return err
			}
		}

		return userMgr.UserModel.Create(ctx, user)
	})
	if err != nil {
		return err
	}

	LoadCasbinPolicy(ctx, userMgr.Enforcer)
	newId := schema.NewIDResult(user.ID)
	pbUser := &proto.User{
		ID: newId.ID,
	}
	res.Items, _ = ptypes.MarshalAny(pbUser)
	return nil
}
func (userMgr *User) CreateID(ctx context.Context, item schema.User) (*schema.IDResult, error) {
	err := userMgr.checkUserName(ctx, item)
	if err != nil {
		return nil, err
	}

	item.Password = util.SHA1HashString(item.Password)
	item.ID = noworker.NewID()
	err = ExecTrans(ctx, userMgr.TransModel, func(ctx context.Context) error {
		for _, urItem := range item.UserRoles {
			urItem.ID = noworker.NewID()
			urItem.UserID = item.ID
			err := userMgr.UserRoleModel.Create(ctx, *urItem)
			if err != nil {
				return err
			}
		}

		return userMgr.UserModel.Create(ctx, item)
	})
	if err != nil {
		return nil, err
	}

	LoadCasbinPolicy(ctx, userMgr.Enforcer)
	return schema.NewIDResult(item.ID), nil
}

func (userMgr *User) checkUserName(ctx context.Context, item schema.User) error {
	if item.UserName == schema.GetRootUser().UserName {
		return errors.New400Response("用户名不合法")
	}

	result, err := userMgr.UserModel.Query(ctx, schema.UserQueryParam{
		PaginationParam: schema.PaginationParam{OnlyCount: true},
		UserName:        item.UserName,
	})
	if err != nil {
		return err
	} else if result.PageResult.Total > 0 {
		return errors.New400Response("用户名已经存在")
	}
	return nil
}

// Update 更新数据
func (userMgr *User) Update(ctx context.Context, id string, item schema.User) error {
	oldItem, err := userMgr.GetUser(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	} else if oldItem.UserName != item.UserName {
		err := userMgr.checkUserName(ctx, item)
		if err != nil {
			return err
		}
	}

	if item.Password != "" {
		item.Password = util.SHA1HashString(item.Password)
	} else {
		item.Password = oldItem.Password
	}

	item.ID = oldItem.ID
	item.Creator = oldItem.Creator
	item.CreatedAt = oldItem.CreatedAt
	err = ExecTrans(ctx, userMgr.TransModel, func(ctx context.Context) error {
		addUserRoles, delUserRoles := userMgr.compareUserRoles(ctx, oldItem.UserRoles, item.UserRoles)
		for _, rmitem := range addUserRoles {
			rmitem.ID = noworker.NewID()
			rmitem.UserID = id
			err := userMgr.UserRoleModel.Create(ctx, *rmitem)
			if err != nil {
				return err
			}
		}

		for _, rmitem := range delUserRoles {
			err := userMgr.UserRoleModel.Delete(ctx, rmitem.ID)
			if err != nil {
				return err
			}
		}

		return userMgr.UserModel.Update(ctx, id, item)
	})
	if err != nil {
		return err
	}

	LoadCasbinPolicy(ctx, userMgr.Enforcer)
	return nil
}

func (userMgr *User) compareUserRoles(ctx context.Context, oldUserRoles, newUserRoles schema.UserRoles) (addList, delList schema.UserRoles) {
	mOldUserRoles := oldUserRoles.ToMap()
	mNewUserRoles := newUserRoles.ToMap()

	for k, item := range mNewUserRoles {
		if _, ok := mOldUserRoles[k]; ok {
			delete(mOldUserRoles, k)
			continue
		}
		addList = append(addList, item)
	}

	for _, item := range mOldUserRoles {
		delList = append(delList, item)
	}
	return
}

// Delete 删除数据
func (userMgr *User) Delete(ctx context.Context, id string) error {
	oldItem, err := userMgr.UserModel.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	err = ExecTrans(ctx, userMgr.TransModel, func(ctx context.Context) error {
		err := userMgr.UserRoleModel.DeleteByUserID(ctx, id)
		if err != nil {
			return err
		}

		return userMgr.UserModel.Delete(ctx, id)
	})
	if err != nil {
		return err
	}

	LoadCasbinPolicy(ctx, userMgr.Enforcer)
	return nil
}

// UpdateStatus 更新状态
func (userMgr *User) UpdateStatus(ctx context.Context, id string, status int) error {
	oldItem, err := userMgr.UserModel.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}
	oldItem.Status = status

	err = userMgr.UserModel.UpdateStatus(ctx, id, status)
	if err != nil {
		return err
	}

	LoadCasbinPolicy(ctx, userMgr.Enforcer)
	return nil
}
