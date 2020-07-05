// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: admin/menu.proto

package admin

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	unified "github.com/ops-cn/go-devops/proto/unified"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Menu struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Sequence   int32  `protobuf:"varint,3,opt,name=Sequence,proto3" json:"Sequence,omitempty"`
	Icon       string `protobuf:"bytes,4,opt,name=Icon,proto3" json:"Icon,omitempty"`
	Router     string `protobuf:"bytes,5,opt,name=Router,proto3" json:"Router,omitempty"`
	ParentPath string `protobuf:"bytes,6,opt,name=ParentPath,proto3" json:"ParentPath,omitempty"`
	ShowStatus int32  `protobuf:"varint,7,opt,name=ShowStatus,proto3" json:"ShowStatus,omitempty"`
	Status     int32  `protobuf:"varint,8,opt,name=Status,proto3" json:"Status,omitempty"`
	Memo       string `protobuf:"bytes,9,opt,name=Memo,proto3" json:"Memo,omitempty"`
	Creator    string `protobuf:"bytes,10,opt,name=Creator,proto3" json:"Creator,omitempty"`
	// @inject_tag: valid:"created_at"
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,11,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	// @inject_tag: valid:"updated_at"
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,12,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	Actions   []*MenuAction        `protobuf:"bytes,13,rep,name=Actions,proto3" json:"Actions,omitempty"`
}

func (x *Menu) Reset() {
	*x = Menu{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_menu_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Menu) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Menu) ProtoMessage() {}

func (x *Menu) ProtoReflect() protoreflect.Message {
	mi := &file_admin_menu_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Menu.ProtoReflect.Descriptor instead.
func (*Menu) Descriptor() ([]byte, []int) {
	return file_admin_menu_proto_rawDescGZIP(), []int{0}
}

func (x *Menu) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Menu) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Menu) GetSequence() int32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *Menu) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *Menu) GetRouter() string {
	if x != nil {
		return x.Router
	}
	return ""
}

func (x *Menu) GetParentPath() string {
	if x != nil {
		return x.ParentPath
	}
	return ""
}

func (x *Menu) GetShowStatus() int32 {
	if x != nil {
		return x.ShowStatus
	}
	return 0
}

func (x *Menu) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Menu) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *Menu) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *Menu) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Menu) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Menu) GetActions() []*MenuAction {
	if x != nil {
		return x.Actions
	}
	return nil
}

type MenuAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        string                `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	MenuID    string                `protobuf:"bytes,2,opt,name=MenuID,proto3" json:"MenuID,omitempty"`
	Code      string                `protobuf:"bytes,3,opt,name=Code,proto3" json:"Code,omitempty"`
	Name      string                `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	Resources []*MenuActionResource `protobuf:"bytes,5,rep,name=Resources,proto3" json:"Resources,omitempty"`
}

func (x *MenuAction) Reset() {
	*x = MenuAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_menu_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuAction) ProtoMessage() {}

func (x *MenuAction) ProtoReflect() protoreflect.Message {
	mi := &file_admin_menu_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuAction.ProtoReflect.Descriptor instead.
func (*MenuAction) Descriptor() ([]byte, []int) {
	return file_admin_menu_proto_rawDescGZIP(), []int{1}
}

func (x *MenuAction) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *MenuAction) GetMenuID() string {
	if x != nil {
		return x.MenuID
	}
	return ""
}

func (x *MenuAction) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *MenuAction) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MenuAction) GetResources() []*MenuActionResource {
	if x != nil {
		return x.Resources
	}
	return nil
}

type MenuActionResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ActionID string `protobuf:"bytes,2,opt,name=ActionID,proto3" json:"ActionID,omitempty"`
	Method   string `protobuf:"bytes,3,opt,name=Method,proto3" json:"Method,omitempty"`
	Path     string `protobuf:"bytes,4,opt,name=Path,proto3" json:"Path,omitempty"`
}

func (x *MenuActionResource) Reset() {
	*x = MenuActionResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_menu_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuActionResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuActionResource) ProtoMessage() {}

func (x *MenuActionResource) ProtoReflect() protoreflect.Message {
	mi := &file_admin_menu_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuActionResource.ProtoReflect.Descriptor instead.
func (*MenuActionResource) Descriptor() ([]byte, []int) {
	return file_admin_menu_proto_rawDescGZIP(), []int{2}
}

func (x *MenuActionResource) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *MenuActionResource) GetActionID() string {
	if x != nil {
		return x.ActionID
	}
	return ""
}

func (x *MenuActionResource) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *MenuActionResource) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type MenuQueryParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaginationParam  *unified.PaginationParam `protobuf:"bytes,1,opt,name=PaginationParam,proto3" json:"PaginationParam,omitempty"`
	IDs              []string                 `protobuf:"bytes,2,rep,name=IDs,proto3" json:"IDs,omitempty"`
	Name             string                   `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	PrefixParentPath string                   `protobuf:"bytes,4,opt,name=PrefixParentPath,proto3" json:"PrefixParentPath,omitempty"`
	QueryValue       string                   `protobuf:"bytes,5,opt,name=QueryValue,proto3" json:"QueryValue,omitempty"`
	ParentID         string                   `protobuf:"bytes,6,opt,name=ParentID,proto3" json:"ParentID,omitempty"`
	ShowStatus       int32                    `protobuf:"varint,7,opt,name=ShowStatus,proto3" json:"ShowStatus,omitempty"`
	Status           int32                    `protobuf:"varint,8,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *MenuQueryParam) Reset() {
	*x = MenuQueryParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_menu_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuQueryParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuQueryParam) ProtoMessage() {}

func (x *MenuQueryParam) ProtoReflect() protoreflect.Message {
	mi := &file_admin_menu_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuQueryParam.ProtoReflect.Descriptor instead.
func (*MenuQueryParam) Descriptor() ([]byte, []int) {
	return file_admin_menu_proto_rawDescGZIP(), []int{3}
}

func (x *MenuQueryParam) GetPaginationParam() *unified.PaginationParam {
	if x != nil {
		return x.PaginationParam
	}
	return nil
}

func (x *MenuQueryParam) GetIDs() []string {
	if x != nil {
		return x.IDs
	}
	return nil
}

func (x *MenuQueryParam) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MenuQueryParam) GetPrefixParentPath() string {
	if x != nil {
		return x.PrefixParentPath
	}
	return ""
}

func (x *MenuQueryParam) GetQueryValue() string {
	if x != nil {
		return x.QueryValue
	}
	return ""
}

func (x *MenuQueryParam) GetParentID() string {
	if x != nil {
		return x.ParentID
	}
	return ""
}

func (x *MenuQueryParam) GetShowStatus() int32 {
	if x != nil {
		return x.ShowStatus
	}
	return 0
}

func (x *MenuQueryParam) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type MenuQueryResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data       []*Menu                   `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
	PageResult *unified.PaginationResult `protobuf:"bytes,2,opt,name=PageResult,proto3" json:"PageResult,omitempty"`
}

func (x *MenuQueryResult) Reset() {
	*x = MenuQueryResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_menu_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuQueryResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuQueryResult) ProtoMessage() {}

func (x *MenuQueryResult) ProtoReflect() protoreflect.Message {
	mi := &file_admin_menu_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuQueryResult.ProtoReflect.Descriptor instead.
func (*MenuQueryResult) Descriptor() ([]byte, []int) {
	return file_admin_menu_proto_rawDescGZIP(), []int{4}
}

func (x *MenuQueryResult) GetData() []*Menu {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *MenuQueryResult) GetPageResult() *unified.PaginationResult {
	if x != nil {
		return x.PageResult
	}
	return nil
}

type MenuTree struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         string        `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name       string        `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Sequence   int32         `protobuf:"varint,3,opt,name=Sequence,proto3" json:"Sequence,omitempty"`
	Icon       string        `protobuf:"bytes,4,opt,name=Icon,proto3" json:"Icon,omitempty"`
	Router     string        `protobuf:"bytes,5,opt,name=Router,proto3" json:"Router,omitempty"`
	ParentPath string        `protobuf:"bytes,6,opt,name=ParentPath,proto3" json:"ParentPath,omitempty"`
	ShowStatus int32         `protobuf:"varint,7,opt,name=ShowStatus,proto3" json:"ShowStatus,omitempty"`
	Status     int32         `protobuf:"varint,8,opt,name=Status,proto3" json:"Status,omitempty"`
	ParentID   string        `protobuf:"bytes,9,opt,name=ParentID,proto3" json:"ParentID,omitempty"`
	Actions    []*MenuAction `protobuf:"bytes,10,rep,name=Actions,proto3" json:"Actions,omitempty"`
	Children   []*MenuTree   `protobuf:"bytes,11,rep,name=Children,proto3" json:"Children,omitempty"`
}

func (x *MenuTree) Reset() {
	*x = MenuTree{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_menu_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuTree) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuTree) ProtoMessage() {}

func (x *MenuTree) ProtoReflect() protoreflect.Message {
	mi := &file_admin_menu_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuTree.ProtoReflect.Descriptor instead.
func (*MenuTree) Descriptor() ([]byte, []int) {
	return file_admin_menu_proto_rawDescGZIP(), []int{5}
}

func (x *MenuTree) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *MenuTree) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MenuTree) GetSequence() int32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *MenuTree) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *MenuTree) GetRouter() string {
	if x != nil {
		return x.Router
	}
	return ""
}

func (x *MenuTree) GetParentPath() string {
	if x != nil {
		return x.ParentPath
	}
	return ""
}

func (x *MenuTree) GetShowStatus() int32 {
	if x != nil {
		return x.ShowStatus
	}
	return 0
}

func (x *MenuTree) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *MenuTree) GetParentID() string {
	if x != nil {
		return x.ParentID
	}
	return ""
}

func (x *MenuTree) GetActions() []*MenuAction {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *MenuTree) GetChildren() []*MenuTree {
	if x != nil {
		return x.Children
	}
	return nil
}

type MenuActions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MenuAction []*MenuAction `protobuf:"bytes,1,rep,name=MenuAction,proto3" json:"MenuAction,omitempty"`
}

func (x *MenuActions) Reset() {
	*x = MenuActions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_menu_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuActions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuActions) ProtoMessage() {}

func (x *MenuActions) ProtoReflect() protoreflect.Message {
	mi := &file_admin_menu_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuActions.ProtoReflect.Descriptor instead.
func (*MenuActions) Descriptor() ([]byte, []int) {
	return file_admin_menu_proto_rawDescGZIP(), []int{6}
}

func (x *MenuActions) GetMenuAction() []*MenuAction {
	if x != nil {
		return x.MenuAction
	}
	return nil
}

type MenuTrees struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MenuTree []*MenuTree `protobuf:"bytes,1,rep,name=MenuTree,proto3" json:"MenuTree,omitempty"`
}

func (x *MenuTrees) Reset() {
	*x = MenuTrees{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_menu_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuTrees) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuTrees) ProtoMessage() {}

func (x *MenuTrees) ProtoReflect() protoreflect.Message {
	mi := &file_admin_menu_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuTrees.ProtoReflect.Descriptor instead.
func (*MenuTrees) Descriptor() ([]byte, []int) {
	return file_admin_menu_proto_rawDescGZIP(), []int{7}
}

func (x *MenuTrees) GetMenuTree() []*MenuTree {
	if x != nil {
		return x.MenuTree
	}
	return nil
}

var File_admin_menu_proto protoreflect.FileDescriptor

var file_admin_menu_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x75, 0x6e, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x03, 0x0a, 0x04, 0x4d, 0x65,
	0x6e, 0x75, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x63, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x49, 0x63, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x12, 0x1e,
	0x0a, 0x0a, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1e,
	0x0a, 0x0a, 0x53, 0x68, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x53, 0x68, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x4d, 0x65, 0x6d, 0x6f, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4d, 0x65, 0x6d, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38,
	0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2b, 0x0a, 0x07, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x95, 0x01, 0x0a, 0x0a, 0x4d, 0x65, 0x6e, 0x75, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x65, 0x6e, 0x75, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4d, 0x65, 0x6e, 0x75, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x4d, 0x65, 0x6e, 0x75, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x09, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0x6c, 0x0a,
	0x12, 0x4d, 0x65, 0x6e, 0x75, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12,
	0x16, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61, 0x74, 0x68, 0x22, 0x9a, 0x02, 0x0a, 0x0e,
	0x4d, 0x65, 0x6e, 0x75, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x42,
	0x0a, 0x0f, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x52, 0x0f, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x49, 0x44, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x03, 0x49, 0x44, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x50, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x68, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x53, 0x68, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x6d, 0x0a, 0x0f, 0x4d, 0x65, 0x6e, 0x75,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x39, 0x0a, 0x0a,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0a, 0x50, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0xc4, 0x02, 0x0a, 0x08, 0x4d, 0x65, 0x6e, 0x75,
	0x54, 0x72, 0x65, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x53, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x63, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x49, 0x63, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72,
	0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x74, 0x68, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x74, 0x68,
	0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x68, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x53, 0x68, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x49, 0x44, 0x12, 0x2b, 0x0a, 0x07, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4d, 0x65,
	0x6e, 0x75, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x2b, 0x0a, 0x08, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x0b, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4d, 0x65, 0x6e, 0x75,
	0x54, 0x72, 0x65, 0x65, 0x52, 0x08, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22, 0x40,
	0x0a, 0x0b, 0x4d, 0x65, 0x6e, 0x75, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x31, 0x0a,
	0x0a, 0x4d, 0x65, 0x6e, 0x75, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x4d, 0x65, 0x6e, 0x75, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x38, 0x0a, 0x09, 0x4d, 0x65, 0x6e, 0x75, 0x54, 0x72, 0x65, 0x65, 0x73, 0x12, 0x2b, 0x0a,
	0x08, 0x4d, 0x65, 0x6e, 0x75, 0x54, 0x72, 0x65, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x54, 0x72, 0x65, 0x65,
	0x52, 0x08, 0x4d, 0x65, 0x6e, 0x75, 0x54, 0x72, 0x65, 0x65, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x73, 0x2d, 0x63, 0x6e, 0x2f,
	0x67, 0x6f, 0x2d, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_menu_proto_rawDescOnce sync.Once
	file_admin_menu_proto_rawDescData = file_admin_menu_proto_rawDesc
)

func file_admin_menu_proto_rawDescGZIP() []byte {
	file_admin_menu_proto_rawDescOnce.Do(func() {
		file_admin_menu_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_menu_proto_rawDescData)
	})
	return file_admin_menu_proto_rawDescData
}

var file_admin_menu_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_admin_menu_proto_goTypes = []interface{}{
	(*Menu)(nil),                     // 0: admin.Menu
	(*MenuAction)(nil),               // 1: admin.MenuAction
	(*MenuActionResource)(nil),       // 2: admin.MenuActionResource
	(*MenuQueryParam)(nil),           // 3: admin.MenuQueryParam
	(*MenuQueryResult)(nil),          // 4: admin.MenuQueryResult
	(*MenuTree)(nil),                 // 5: admin.MenuTree
	(*MenuActions)(nil),              // 6: admin.MenuActions
	(*MenuTrees)(nil),                // 7: admin.MenuTrees
	(*timestamp.Timestamp)(nil),      // 8: google.protobuf.Timestamp
	(*unified.PaginationParam)(nil),  // 9: unified.PaginationParam
	(*unified.PaginationResult)(nil), // 10: unified.PaginationResult
}
var file_admin_menu_proto_depIdxs = []int32{
	8,  // 0: admin.Menu.CreatedAt:type_name -> google.protobuf.Timestamp
	8,  // 1: admin.Menu.UpdatedAt:type_name -> google.protobuf.Timestamp
	1,  // 2: admin.Menu.Actions:type_name -> admin.MenuAction
	2,  // 3: admin.MenuAction.Resources:type_name -> admin.MenuActionResource
	9,  // 4: admin.MenuQueryParam.PaginationParam:type_name -> unified.PaginationParam
	0,  // 5: admin.MenuQueryResult.Data:type_name -> admin.Menu
	10, // 6: admin.MenuQueryResult.PageResult:type_name -> unified.PaginationResult
	1,  // 7: admin.MenuTree.Actions:type_name -> admin.MenuAction
	5,  // 8: admin.MenuTree.Children:type_name -> admin.MenuTree
	1,  // 9: admin.MenuActions.MenuAction:type_name -> admin.MenuAction
	5,  // 10: admin.MenuTrees.MenuTree:type_name -> admin.MenuTree
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_admin_menu_proto_init() }
func file_admin_menu_proto_init() {
	if File_admin_menu_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_admin_menu_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Menu); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_admin_menu_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuAction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_admin_menu_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuActionResource); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_admin_menu_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuQueryParam); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_admin_menu_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuQueryResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_admin_menu_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuTree); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_admin_menu_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuActions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_admin_menu_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuTrees); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_admin_menu_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_admin_menu_proto_goTypes,
		DependencyIndexes: file_admin_menu_proto_depIdxs,
		MessageInfos:      file_admin_menu_proto_msgTypes,
	}.Build()
	File_admin_menu_proto = out.File
	file_admin_menu_proto_rawDesc = nil
	file_admin_menu_proto_goTypes = nil
	file_admin_menu_proto_depIdxs = nil
}
