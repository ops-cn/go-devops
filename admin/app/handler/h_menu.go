package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/ops-cn/go-devops/common/noworker"
	proto "github.com/ops-cn/go-devops/proto/admin"
	"github.com/ops-cn/go-devops/proto/unified"
	"os"

	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"github.com/ops-cn/go-devops/admin/app/model"
	"github.com/ops-cn/go-devops/common/errors"
	"github.com/ops-cn/go-devops/common/schema"
	"github.com/ops-cn/go-devops/common/util"
)

// MenuSet 注入Menu
var MenuSet = wire.NewSet(wire.Struct(new(Menu), "*"))

// Menu 菜单管理
type Menu struct {
	TransModel              model.ITrans
	MenuModel               model.IMenu
	MenuActionModel         model.IMenuAction
	MenuActionResourceModel model.IMenuActionResource
}

// InitData 初始化菜单数据
func (menuMgr *Menu) InitData(ctx context.Context, dataFile string) error {
	result, err := menuMgr.MenuModel.Query(ctx, schema.MenuQueryParam{
		PaginationParam: schema.PaginationParam{OnlyCount: true},
	})
	if err != nil {
		return err
	} else if result.PageResult.Total > 0 {
		// 如果存在则不进行初始化
		return nil
	}

	data, err := menuMgr.readData(dataFile)
	if err != nil {
		return err
	}

	return menuMgr.createMenus(ctx, "", data)
}

func (menuMgr *Menu) readData(name string) (schema.MenuTrees, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data schema.MenuTrees
	d := util.YAMLNewDecoder(file)
	d.SetStrict(true)
	err = d.Decode(&data)
	return data, err
}

func (menuMgr *Menu) createMenus(ctx context.Context, parentID string, list schema.MenuTrees) error {
	return ExecTrans(ctx, menuMgr.TransModel, func(ctx context.Context) error {
		for _, item := range list {
			sitem := schema.Menu{
				Name:       item.Name,
				Sequence:   item.Sequence,
				Icon:       item.Icon,
				Router:     item.Router,
				ParentID:   parentID,
				Status:     1,
				ShowStatus: 1,
				Actions:    item.Actions,
			}
			if v := item.ShowStatus; v > 0 {
				sitem.ShowStatus = v
			}

			nsitem, err := menuMgr.CreateId(ctx, sitem)
			if err != nil {
				return err
			}

			if item.Children != nil && len(*item.Children) > 0 {
				err := menuMgr.createMenus(ctx, nsitem.ID, *item.Children)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}

// Query 查询数据
func (menuMgr *Menu) Query(ctx context.Context, req *proto.MenuQueryReq, res *unified.Response) error {
	menuActionResult, err := menuMgr.MenuActionModel.Query(ctx, schema.MenuActionQueryParam{})
	if err != nil {
		return err
	}
	param := schema.MenuQueryParam{}
	util.StructCopy(param, req.MenuQueryParam)
	pp := schema.PaginationParam{}
	rpp := req.MenuQueryParam.PaginationParam
	util.StructCopy(pp, rpp)
	param.PaginationParam = pp
	opts := schema.MenuQueryOptions{}
	result, err := menuMgr.MenuModel.Query(ctx, param, opts)
	if err != nil {
		return err
	}
	result.Data.FillMenuAction(menuActionResult.Data.ToMenuIDMap())
	pbResult := &proto.MenuQueryResult{}
	util.StructCopy(pbResult, result)
	res.Items, _ = ptypes.MarshalAny(pbResult)
	return nil
}

// Get 查询指定数据
func (menuMgr *Menu) GetMenu(ctx context.Context, id string, opts ...schema.MenuQueryOptions) (*schema.Menu, error) {
	item, err := menuMgr.MenuModel.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	actions, err := menuMgr.QueryActions(ctx, id)
	if err != nil {
		return nil, err
	}
	item.Actions = actions

	return item, nil
}
func (menuMgr *Menu) Get(ctx context.Context, req *proto.MenuReq, res *unified.Response) error {
	opts := schema.MenuQueryOptions{}
	copier.Copy(opts, req.MenuQueryOptions)
	menu, err := menuMgr.MenuModel.Get(ctx, req.CurrentID, opts)
	if err != nil {
		return err
	} else if menu == nil {
		return errors.ErrNotFound
	}

	actions, err := menuMgr.QueryActions(ctx, req.CurrentID)
	if err != nil {
		return err
	}
	menu.Actions = actions
	pbMenu := &proto.Menu{}
	copier.Copy(pbMenu, menu)
	res.Items, _ = ptypes.MarshalAny(pbMenu)
	return nil
}

// QueryActions 查询动作数据
func (menuMgr *Menu) QueryActions(ctx context.Context, id string) (schema.MenuActions, error) {
	result, err := menuMgr.MenuActionModel.Query(ctx, schema.MenuActionQueryParam{
		MenuID: id,
	})
	if err != nil {
		return nil, err
	} else if len(result.Data) == 0 {
		return nil, nil
	}

	resourceResult, err := menuMgr.MenuActionResourceModel.Query(ctx, schema.MenuActionResourceQueryParam{
		MenuID: id,
	})
	if err != nil {
		return nil, err
	}

	result.Data.FillResources(resourceResult.Data.ToActionIDMap())

	return result.Data, nil
}

func (menuMgr *Menu) checkName(ctx context.Context, item schema.Menu) error {
	result, err := menuMgr.MenuModel.Query(ctx, schema.MenuQueryParam{
		PaginationParam: schema.PaginationParam{
			OnlyCount: true,
		},
		ParentID: &item.ParentID,
		Name:     item.Name,
	})
	if err != nil {
		return err
	} else if result.PageResult.Total > 0 {
		return errors.New400Response("菜单名称已经存在")
	}
	return nil
}

// Create 创建数据
func (menuMgr *Menu) CreateId(ctx context.Context, item schema.Menu) (*schema.IDResult, error) {
	if err := menuMgr.checkName(ctx, item); err != nil {
		return nil, err
	}

	parentPath, err := menuMgr.getParentPath(ctx, item.ParentID)
	if err != nil {
		return nil, err
	}
	item.ParentPath = parentPath
	item.ID = noworker.NewID()

	err = ExecTrans(ctx, menuMgr.TransModel, func(ctx context.Context) error {
		err := menuMgr.createActions(ctx, item.ID, item.Actions)
		if err != nil {
			return err
		}

		return menuMgr.MenuModel.Create(ctx, item)
	})
	if err != nil {
		return nil, err
	}

	return schema.NewIDResult(item.ID), nil
}
func (menuMgr *Menu) Create(ctx context.Context, req *proto.Menu, res *unified.Response) error {
	menu := schema.Menu{}
	copier.Copy(menu, req)
	if err := menuMgr.checkName(ctx, menu); err != nil {
		return err
	}

	parentPath, err := menuMgr.getParentPath(ctx, menu.ParentID)
	if err != nil {
		return err
	}
	menu.ParentPath = parentPath
	menu.ID = noworker.NewID()

	err = ExecTrans(ctx, menuMgr.TransModel, func(ctx context.Context) error {
		err := menuMgr.createActions(ctx, menu.ID, menu.Actions)
		if err != nil {
			return err
		}

		return menuMgr.MenuModel.Create(ctx, menu)
	})
	if err != nil {
		return err
	}
	newId := schema.NewIDResult(menu.ID)
	pbMenu := &proto.Menu{
		ID: newId.ID,
	}
	res.Items, _ = ptypes.MarshalAny(pbMenu)
	return nil
}

// 创建动作数据
func (menuMgr *Menu) createActions(ctx context.Context, menuID string, items schema.MenuActions) error {
	for _, item := range items {
		item.ID = noworker.NewID()
		item.MenuID = menuID
		err := menuMgr.MenuActionModel.Create(ctx, *item)
		if err != nil {
			return err
		}

		for _, ritem := range item.Resources {
			ritem.ID = noworker.NewID()
			ritem.ActionID = item.ID
			err := menuMgr.MenuActionResourceModel.Create(ctx, *ritem)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

// 获取父级路径
func (menuMgr *Menu) getParentPath(ctx context.Context, parentID string) (string, error) {
	if parentID == "" {
		return "", nil
	}

	pitem, err := menuMgr.MenuModel.Get(ctx, parentID)
	if err != nil {
		return "", err
	} else if pitem == nil {
		return "", errors.ErrInvalidParent
	}

	return menuMgr.joinParentPath(pitem.ParentPath, pitem.ID), nil
}

func (menuMgr *Menu) joinParentPath(parent, id string) string {
	if parent != "" {
		return parent + "/" + id
	}
	return id
}

// Update 更新数据
func (menuMgr *Menu) Update(ctx context.Context, req *proto.MenuReq, res *unified.Response) error {
	menu := schema.Menu{}
	copier.Copy(menu, req.Menu)

	if req.CurrentID == menu.ParentID {
		return errors.ErrInvalidParent
	}

	oldItem, err := menuMgr.GetMenu(ctx, req.CurrentID)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	} else if oldItem.Name != menu.Name {
		if err := menuMgr.checkName(ctx, menu); err != nil {
			return err
		}
	}

	menu.ID = oldItem.ID
	menu.Creator = oldItem.Creator
	menu.CreatedAt = oldItem.CreatedAt

	if oldItem.ParentID != menu.ParentID {
		parentPath, err := menuMgr.getParentPath(ctx, menu.ParentID)
		if err != nil {
			return err
		}
		menu.ParentPath = parentPath
	} else {
		menu.ParentPath = oldItem.ParentPath
	}

	return ExecTrans(ctx, menuMgr.TransModel, func(ctx context.Context) error {
		err := menuMgr.updateActions(ctx, req.CurrentID, oldItem.Actions, menu.Actions)
		if err != nil {
			return err
		}

		err = menuMgr.updateChildParentPath(ctx, *oldItem, menu)
		if err != nil {
			return err
		}

		return menuMgr.MenuModel.Update(ctx, req.CurrentID, menu)
	})
}

// 更新动作数据
func (menuMgr *Menu) updateActions(ctx context.Context, menuID string, oldItems, newItems schema.MenuActions) error {
	addActions, delActions, updateActions := menuMgr.compareActions(ctx, oldItems, newItems)

	err := menuMgr.createActions(ctx, menuID, addActions)
	if err != nil {
		return err
	}

	for _, item := range delActions {
		err := menuMgr.MenuActionModel.Delete(ctx, item.ID)
		if err != nil {
			return err
		}

		err = menuMgr.MenuActionResourceModel.DeleteByActionID(ctx, item.ID)
		if err != nil {
			return err
		}
	}

	mOldItems := oldItems.ToMap()
	for _, item := range updateActions {
		oitem := mOldItems[item.Code]
		// 只更新动作名称
		if item.Name != oitem.Name {
			oitem.Name = item.Name
			err := menuMgr.MenuActionModel.Update(ctx, item.ID, *oitem)
			if err != nil {
				return err
			}
		}

		// 计算需要更新的资源配置（只包括新增和删除的，更新的不关心）
		addResources, delResources := menuMgr.compareResources(ctx, oitem.Resources, item.Resources)
		for _, aritem := range addResources {
			aritem.ID = noworker.NewID()
			aritem.ActionID = oitem.ID
			err := menuMgr.MenuActionResourceModel.Create(ctx, *aritem)
			if err != nil {
				return err
			}
		}

		for _, ditem := range delResources {
			err := menuMgr.MenuActionResourceModel.Delete(ctx, ditem.ID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// 对比动作列表
func (menuMgr *Menu) compareActions(ctx context.Context, oldActions, newActions schema.MenuActions) (addList, delList, updateList schema.MenuActions) {
	mOldActions := oldActions.ToMap()
	mNewActions := newActions.ToMap()

	for k, item := range mNewActions {
		if _, ok := mOldActions[k]; ok {
			updateList = append(updateList, item)
			delete(mOldActions, k)
			continue
		}
		addList = append(addList, item)
	}

	for _, item := range mOldActions {
		delList = append(delList, item)
	}
	return
}

// 对比资源列表
func (menuMgr *Menu) compareResources(ctx context.Context, oldResources, newResources schema.MenuActionResources) (addList, delList schema.MenuActionResources) {
	mOldResources := oldResources.ToMap()
	mNewResources := newResources.ToMap()

	for k, item := range mNewResources {
		if _, ok := mOldResources[k]; ok {
			delete(mOldResources, k)
			continue
		}
		addList = append(addList, item)
	}

	for _, item := range mOldResources {
		delList = append(delList, item)
	}
	return
}

// 检查并更新下级节点的父级路径
func (menuMgr *Menu) updateChildParentPath(ctx context.Context, oldItem, newItem schema.Menu) error {
	if oldItem.ParentID == newItem.ParentID {
		return nil
	}

	opath := menuMgr.joinParentPath(oldItem.ParentPath, oldItem.ID)
	result, err := menuMgr.MenuModel.Query(NewNoTrans(ctx), schema.MenuQueryParam{
		PrefixParentPath: opath,
	})
	if err != nil {
		return err
	}

	npath := menuMgr.joinParentPath(newItem.ParentPath, newItem.ID)
	for _, menu := range result.Data {
		err = menuMgr.MenuModel.UpdateParentPath(ctx, menu.ID, npath+menu.ParentPath[len(opath):])
		if err != nil {
			return err
		}
	}
	return nil
}

// Delete 删除数据
func (menuMgr *Menu) Delete(ctx context.Context, req *proto.Menu, res *unified.Response) error {
	oldItem, err := menuMgr.MenuModel.Get(ctx, req.ID)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	result, err := menuMgr.MenuModel.Query(ctx, schema.MenuQueryParam{
		PaginationParam: schema.PaginationParam{OnlyCount: true},
		ParentID:        &req.ID,
	})
	if err != nil {
		return err
	} else if result.PageResult.Total > 0 {
		return errors.ErrNotAllowDeleteWithChild
	}

	return ExecTrans(ctx, menuMgr.TransModel, func(ctx context.Context) error {
		err = menuMgr.MenuActionResourceModel.DeleteByMenuID(ctx, req.ID)
		if err != nil {
			return err
		}

		err := menuMgr.MenuActionModel.DeleteByMenuID(ctx, req.ID)
		if err != nil {
			return err
		}

		return menuMgr.MenuModel.Delete(ctx, req.ID)
	})
}

// UpdateStatus 更新状态
func (menuMgr *Menu) UpdateStatus(ctx context.Context, req *proto.Menu, res *unified.Response) error {
	oldItem, err := menuMgr.MenuModel.Get(ctx, req.ID)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return menuMgr.MenuModel.UpdateStatus(ctx, req.ID, int(req.Status))
}
