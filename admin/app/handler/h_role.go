package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/copier"
	"github.com/ops-cn/go-devops/common/noworker"
	"github.com/ops-cn/go-devops/common/util"
	proto "github.com/ops-cn/go-devops/proto/admin"
	"github.com/ops-cn/go-devops/proto/unified"

	"github.com/casbin/casbin/v2"
	"github.com/google/wire"
	"github.com/ops-cn/go-devops/admin/app/model"
	"github.com/ops-cn/go-devops/common/errors"
	"github.com/ops-cn/go-devops/common/schema"
)

// RoleSet 注入Role
var RoleSet = wire.NewSet(wire.Struct(new(Role), "*"))

// Role 角色管理
type Role struct {
	Enforcer      *casbin.SyncedEnforcer
	TransModel    model.ITrans
	RoleModel     model.IRole
	RoleMenuModel model.IRoleMenu
	UserModel     model.IUser
}

// Query 查询数据
func (roleMgr *Role) Query(ctx context.Context, req *proto.RoleQueryReq, res *unified.Response) error {
	param := schema.RoleQueryParam{}
	copier.Copy(param, req.RoleQueryParam)

	opts := schema.RoleQueryOptions{}
	copier.Copy(opts, req.RoleQueryOptions)
	result, err := roleMgr.RoleModel.Query(ctx, param, opts)
	if err != nil {
		return err
	}
	pbResult := &proto.RoleQueryResult{}
	util.StructCopy(pbResult, result)
	res.Items, err = ptypes.MarshalAny(pbResult)
	return err
}

// Get 查询指定数据
func (roleMgr *Role) GetRole(ctx context.Context, id string, opts ...schema.RoleQueryOptions) (*schema.Role, error) {
	item, err := roleMgr.RoleModel.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	roleMenus, err := roleMgr.QueryRoleMenus(ctx, id)
	if err != nil {
		return nil, err
	}
	item.RoleMenus = roleMenus

	return item, nil
}
func (roleMgr *Role) Get(ctx context.Context, req *proto.RoleReq, res *unified.Response) error {
	opts := schema.RoleQueryOptions{}
	copier.Copy(opts, req.RoleQueryOptions)
	role, err := roleMgr.RoleModel.Get(ctx, req.CurrentID, opts)
	if err != nil {
		return err
	} else if role == nil {
		return errors.ErrNotFound
	}

	roleMenus, err := roleMgr.QueryRoleMenus(ctx, req.CurrentID)
	if err != nil {
		return err
	}
	role.RoleMenus = roleMenus
	pbRole := &proto.Role{}
	copier.Copy(pbRole, role)
	res.Items, err = ptypes.MarshalAny(pbRole)
	return err
}

// QueryRoleMenus 查询角色菜单列表
func (roleMgr *Role) QueryRoleMenus(ctx context.Context, roleID string) (schema.RoleMenus, error) {
	result, err := roleMgr.RoleMenuModel.Query(ctx, schema.RoleMenuQueryParam{
		RoleID: roleID,
	})
	if err != nil {
		return nil, err
	}
	return result.Data, nil
}

// Create 创建数据
func (roleMgr *Role) Create(ctx context.Context, req *proto.Role, res *unified.Response) error {
	role := schema.Role{}
	copier.Copy(role, req)
	err := roleMgr.checkName(ctx, role)
	if err != nil {
		return err
	}

	role.ID = noworker.NewID()
	err = ExecTrans(ctx, roleMgr.TransModel, func(ctx context.Context) error {
		for _, rmItem := range role.RoleMenus {
			rmItem.ID = noworker.NewID()
			rmItem.RoleID = role.ID
			err := roleMgr.RoleMenuModel.Create(ctx, *rmItem)
			if err != nil {
				return err
			}
		}
		return roleMgr.RoleModel.Create(ctx, role)
	})
	if err != nil {
		return err
	}
	LoadCasbinPolicy(ctx, roleMgr.Enforcer)
	newId := schema.NewIDResult(role.ID)
	pbRole := &proto.Role{
		ID: newId.ID,
	}
	res.Items, err = ptypes.MarshalAny(pbRole)
	return err
}
func (roleMgr *Role) CreateId(ctx context.Context, item schema.Role) (*schema.IDResult, error) {
	err := roleMgr.checkName(ctx, item)
	if err != nil {
		return nil, err
	}

	item.ID = noworker.NewID()
	err = ExecTrans(ctx, roleMgr.TransModel, func(ctx context.Context) error {
		for _, rmItem := range item.RoleMenus {
			rmItem.ID = noworker.NewID()
			rmItem.RoleID = item.ID
			err := roleMgr.RoleMenuModel.Create(ctx, *rmItem)
			if err != nil {
				return err
			}
		}
		return roleMgr.RoleModel.Create(ctx, item)
	})
	if err != nil {
		return nil, err
	}
	LoadCasbinPolicy(ctx, roleMgr.Enforcer)
	return schema.NewIDResult(item.ID), nil
}

func (roleMgr *Role) checkName(ctx context.Context, item schema.Role) error {
	result, err := roleMgr.RoleModel.Query(ctx, schema.RoleQueryParam{
		PaginationParam: schema.PaginationParam{OnlyCount: true},
		Name:            item.Name,
	})
	if err != nil {
		return err
	} else if result.PageResult.Total > 0 {
		return errors.New400Response("角色名称已经存在")
	}
	return nil
}

// Update 更新数据
func (roleMgr *Role) Update(ctx context.Context, req *proto.RoleReq, res *unified.Response) error {
	role := schema.Role{}
	copier.Copy(role, req.Role)
	oldItem, err := roleMgr.GetRole(ctx, req.CurrentID)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	} else if oldItem.Name != role.Name {
		err := roleMgr.checkName(ctx, role)
		if err != nil {
			return err
		}
	}

	role.ID = oldItem.ID
	role.Creator = oldItem.Creator
	role.CreatedAt = oldItem.CreatedAt
	err = ExecTrans(ctx, roleMgr.TransModel, func(ctx context.Context) error {
		addRoleMenus, delRoleMenus := roleMgr.compareRoleMenus(ctx, oldItem.RoleMenus, role.RoleMenus)
		for _, rmItem := range addRoleMenus {
			rmItem.ID = noworker.NewID()
			rmItem.RoleID = req.CurrentID
			err := roleMgr.RoleMenuModel.Create(ctx, *rmItem)
			if err != nil {
				return err
			}
		}

		for _, rmItem := range delRoleMenus {
			err := roleMgr.RoleMenuModel.Delete(ctx, rmItem.ID)
			if err != nil {
				return err
			}
		}

		return roleMgr.RoleModel.Update(ctx, req.CurrentID, role)
	})
	if err != nil {
		return err
	}
	LoadCasbinPolicy(ctx, roleMgr.Enforcer)
	return nil
}

func (roleMgr *Role) compareRoleMenus(ctx context.Context, oldRoleMenus, newRoleMenus schema.RoleMenus) (addList, delList schema.RoleMenus) {
	mOldRoleMenus := oldRoleMenus.ToMap()
	mNewRoleMenus := newRoleMenus.ToMap()

	for k, item := range mNewRoleMenus {
		if _, ok := mOldRoleMenus[k]; ok {
			delete(mOldRoleMenus, k)
			continue
		}
		addList = append(addList, item)
	}

	for _, item := range mOldRoleMenus {
		delList = append(delList, item)
	}
	return
}

// Delete 删除数据
func (roleMgr *Role) Delete(ctx context.Context, req *proto.Role, res *unified.Response) error {
	oldItem, err := roleMgr.RoleModel.Get(ctx, req.ID)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	userResult, err := roleMgr.UserModel.Query(ctx, schema.UserQueryParam{
		PaginationParam: schema.PaginationParam{OnlyCount: true},
		RoleIDs:         []string{req.ID},
	})
	if err != nil {
		return err
	} else if userResult.PageResult.Total > 0 {
		return errors.New400Response("该角色已被赋予用户，不允许删除")
	}

	err = ExecTrans(ctx, roleMgr.TransModel, func(ctx context.Context) error {
		err := roleMgr.RoleMenuModel.DeleteByRoleID(ctx, req.ID)
		if err != nil {
			return err
		}

		return roleMgr.RoleModel.Delete(ctx, req.ID)
	})
	if err != nil {
		return err
	}

	LoadCasbinPolicy(ctx, roleMgr.Enforcer)
	return nil
}

// UpdateStatus 更新状态
func (roleMgr *Role) UpdateStatus(ctx context.Context, req *proto.Role, res *unified.Response) error {
	oldItem, err := roleMgr.RoleModel.Get(ctx, req.ID)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	err = roleMgr.RoleModel.UpdateStatus(ctx, req.ID, int(req.Status))
	if err != nil {
		return err
	}
	LoadCasbinPolicy(ctx, roleMgr.Enforcer)
	return nil
}
