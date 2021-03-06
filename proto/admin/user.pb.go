// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: admin/user.proto

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// @inject_tag: valid:"user_name"
	UserName string `protobuf:"bytes,2,opt,name=UserName,proto3" json:"UserName,omitempty"`
	// @inject_tag: valid:"real_name"
	RealName string `protobuf:"bytes,3,opt,name=RealName,proto3" json:"RealName,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=Password,proto3" json:"Password,omitempty"`
	Phone    string `protobuf:"bytes,5,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Email    string `protobuf:"bytes,6,opt,name=Email,proto3" json:"Email,omitempty"`
	Status   int32  `protobuf:"varint,7,opt,name=Status,proto3" json:"Status,omitempty"`
	Creator  string `protobuf:"bytes,8,opt,name=Creator,proto3" json:"Creator,omitempty"`
	// @inject_tag: valid:"created_at"
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,9,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UserRoles []*UserRole          `protobuf:"bytes,10,rep,name=UserRoles,proto3" json:"UserRoles,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_admin_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_admin_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *User) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *User) GetRealName() string {
	if x != nil {
		return x.RealName
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *User) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *User) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *User) GetUserRoles() []*UserRole {
	if x != nil {
		return x.UserRoles
	}
	return nil
}

type UserRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// @inject_tag: valid:"user_id"
	UserID string `protobuf:"bytes,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	// @inject_tag: valid:"role_id"
	RoleID string `protobuf:"bytes,3,opt,name=RoleID,proto3" json:"RoleID,omitempty"`
}

func (x *UserRole) Reset() {
	*x = UserRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRole) ProtoMessage() {}

func (x *UserRole) ProtoReflect() protoreflect.Message {
	mi := &file_admin_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRole.ProtoReflect.Descriptor instead.
func (*UserRole) Descriptor() ([]byte, []int) {
	return file_admin_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserRole) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *UserRole) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *UserRole) GetRoleID() string {
	if x != nil {
		return x.RoleID
	}
	return ""
}

type UserQueryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserQueryParam   *UserQueryParam   `protobuf:"bytes,1,opt,name=UserQueryParam,proto3" json:"UserQueryParam,omitempty"`
	UserQueryOptions *UserQueryOptions `protobuf:"bytes,2,opt,name=UserQueryOptions,proto3" json:"UserQueryOptions,omitempty"`
}

func (x *UserQueryReq) Reset() {
	*x = UserQueryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserQueryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserQueryReq) ProtoMessage() {}

func (x *UserQueryReq) ProtoReflect() protoreflect.Message {
	mi := &file_admin_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserQueryReq.ProtoReflect.Descriptor instead.
func (*UserQueryReq) Descriptor() ([]byte, []int) {
	return file_admin_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserQueryReq) GetUserQueryParam() *UserQueryParam {
	if x != nil {
		return x.UserQueryParam
	}
	return nil
}

func (x *UserQueryReq) GetUserQueryOptions() *UserQueryOptions {
	if x != nil {
		return x.UserQueryOptions
	}
	return nil
}

type UserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentID        string            `protobuf:"bytes,1,opt,name=CurrentID,proto3" json:"CurrentID,omitempty"`
	User             *User             `protobuf:"bytes,2,opt,name=User,proto3" json:"User,omitempty"`
	UserQueryOptions *UserQueryOptions `protobuf:"bytes,3,opt,name=UserQueryOptions,proto3" json:"UserQueryOptions,omitempty"`
}

func (x *UserReq) Reset() {
	*x = UserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserReq) ProtoMessage() {}

func (x *UserReq) ProtoReflect() protoreflect.Message {
	mi := &file_admin_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserReq.ProtoReflect.Descriptor instead.
func (*UserReq) Descriptor() ([]byte, []int) {
	return file_admin_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserReq) GetCurrentID() string {
	if x != nil {
		return x.CurrentID
	}
	return ""
}

func (x *UserReq) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserReq) GetUserQueryOptions() *UserQueryOptions {
	if x != nil {
		return x.UserQueryOptions
	}
	return nil
}

type UserQueryParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaginationParam *unified.PaginationParam `protobuf:"bytes,1,opt,name=PaginationParam,proto3" json:"PaginationParam,omitempty"`
	UserName        string                   `protobuf:"bytes,2,opt,name=UserName,proto3" json:"UserName,omitempty"`
	QueryValue      string                   `protobuf:"bytes,3,opt,name=QueryValue,proto3" json:"QueryValue,omitempty"`
	Status          int32                    `protobuf:"varint,4,opt,name=Status,proto3" json:"Status,omitempty"`
	RoleIDs         []string                 `protobuf:"bytes,5,rep,name=RoleIDs,proto3" json:"RoleIDs,omitempty"`
}

func (x *UserQueryParam) Reset() {
	*x = UserQueryParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserQueryParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserQueryParam) ProtoMessage() {}

func (x *UserQueryParam) ProtoReflect() protoreflect.Message {
	mi := &file_admin_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserQueryParam.ProtoReflect.Descriptor instead.
func (*UserQueryParam) Descriptor() ([]byte, []int) {
	return file_admin_user_proto_rawDescGZIP(), []int{4}
}

func (x *UserQueryParam) GetPaginationParam() *unified.PaginationParam {
	if x != nil {
		return x.PaginationParam
	}
	return nil
}

func (x *UserQueryParam) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserQueryParam) GetQueryValue() string {
	if x != nil {
		return x.QueryValue
	}
	return ""
}

func (x *UserQueryParam) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UserQueryParam) GetRoleIDs() []string {
	if x != nil {
		return x.RoleIDs
	}
	return nil
}

type UserQueryOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderField []*unified.OrderField `protobuf:"bytes,1,rep,name=OrderField,proto3" json:"OrderField,omitempty"`
}

func (x *UserQueryOptions) Reset() {
	*x = UserQueryOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserQueryOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserQueryOptions) ProtoMessage() {}

func (x *UserQueryOptions) ProtoReflect() protoreflect.Message {
	mi := &file_admin_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserQueryOptions.ProtoReflect.Descriptor instead.
func (*UserQueryOptions) Descriptor() ([]byte, []int) {
	return file_admin_user_proto_rawDescGZIP(), []int{5}
}

func (x *UserQueryOptions) GetOrderField() []*unified.OrderField {
	if x != nil {
		return x.OrderField
	}
	return nil
}

type UserQueryResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data       []*User                   `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
	PageResult *unified.PaginationResult `protobuf:"bytes,2,opt,name=PageResult,proto3" json:"PageResult,omitempty"`
}

func (x *UserQueryResult) Reset() {
	*x = UserQueryResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserQueryResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserQueryResult) ProtoMessage() {}

func (x *UserQueryResult) ProtoReflect() protoreflect.Message {
	mi := &file_admin_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserQueryResult.ProtoReflect.Descriptor instead.
func (*UserQueryResult) Descriptor() ([]byte, []int) {
	return file_admin_user_proto_rawDescGZIP(), []int{6}
}

func (x *UserQueryResult) GetData() []*User {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *UserQueryResult) GetPageResult() *unified.PaginationResult {
	if x != nil {
		return x.PageResult
	}
	return nil
}

var File_admin_user_proto protoreflect.FileDescriptor

var file_admin_user_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x75, 0x6e, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x02, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x52, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x52, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2d,
	0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x09, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x4a, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x22, 0x92, 0x01, 0x0a, 0x0c, 0x55, 0x73,
	0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x12, 0x3d, 0x0a, 0x0e, 0x55, 0x73,
	0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x43, 0x0a, 0x10, 0x55, 0x73, 0x65,
	0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x10, 0x55, 0x73,
	0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x8d,
	0x01, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1f, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x10, 0x55, 0x73, 0x65,
	0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x10, 0x55, 0x73,
	0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xc2,
	0x01, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x12, 0x42, 0x0a, 0x0f, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x75, 0x6e, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x52, 0x0f, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x6f, 0x6c,
	0x65, 0x49, 0x44, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x52, 0x6f, 0x6c, 0x65,
	0x49, 0x44, 0x73, 0x22, 0x47, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x33, 0x0a, 0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x75, 0x6e,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x52, 0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x6d, 0x0a, 0x0f,
	0x55, 0x73, 0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x1f, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x39, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x2e, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x0a, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xd5, 0x02, 0x0a, 0x07,
	0x55, 0x73, 0x65, 0x72, 0x4d, 0x67, 0x72, 0x12, 0x31, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x13, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x0e, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x1a, 0x11, 0x2e, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x0b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x11, 0x2e,
	0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x2d, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x75,
	0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x2a, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0b, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x11, 0x2e, 0x75, 0x6e, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x30, 0x0a,
	0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x11, 0x2e, 0x75, 0x6e, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x32, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x0b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x11,
	0x2e, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6f, 0x70, 0x73, 0x2d, 0x63, 0x6e, 0x2f, 0x67, 0x6f, 0x2d, 0x64, 0x65, 0x76, 0x6f,
	0x70, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_user_proto_rawDescOnce sync.Once
	file_admin_user_proto_rawDescData = file_admin_user_proto_rawDesc
)

func file_admin_user_proto_rawDescGZIP() []byte {
	file_admin_user_proto_rawDescOnce.Do(func() {
		file_admin_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_user_proto_rawDescData)
	})
	return file_admin_user_proto_rawDescData
}

var file_admin_user_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_admin_user_proto_goTypes = []interface{}{
	(*User)(nil),                     // 0: admin.User
	(*UserRole)(nil),                 // 1: admin.UserRole
	(*UserQueryReq)(nil),             // 2: admin.UserQueryReq
	(*UserReq)(nil),                  // 3: admin.UserReq
	(*UserQueryParam)(nil),           // 4: admin.UserQueryParam
	(*UserQueryOptions)(nil),         // 5: admin.UserQueryOptions
	(*UserQueryResult)(nil),          // 6: admin.UserQueryResult
	(*timestamp.Timestamp)(nil),      // 7: google.protobuf.Timestamp
	(*unified.PaginationParam)(nil),  // 8: unified.PaginationParam
	(*unified.OrderField)(nil),       // 9: unified.OrderField
	(*unified.PaginationResult)(nil), // 10: unified.PaginationResult
	(*unified.Response)(nil),         // 11: unified.Response
}
var file_admin_user_proto_depIdxs = []int32{
	7,  // 0: admin.User.CreatedAt:type_name -> google.protobuf.Timestamp
	1,  // 1: admin.User.UserRoles:type_name -> admin.UserRole
	4,  // 2: admin.UserQueryReq.UserQueryParam:type_name -> admin.UserQueryParam
	5,  // 3: admin.UserQueryReq.UserQueryOptions:type_name -> admin.UserQueryOptions
	0,  // 4: admin.UserReq.User:type_name -> admin.User
	5,  // 5: admin.UserReq.UserQueryOptions:type_name -> admin.UserQueryOptions
	8,  // 6: admin.UserQueryParam.PaginationParam:type_name -> unified.PaginationParam
	9,  // 7: admin.UserQueryOptions.OrderField:type_name -> unified.OrderField
	0,  // 8: admin.UserQueryResult.Data:type_name -> admin.User
	10, // 9: admin.UserQueryResult.PageResult:type_name -> unified.PaginationResult
	2,  // 10: admin.UserMgr.Query:input_type -> admin.UserQueryReq
	3,  // 11: admin.UserMgr.Get:input_type -> admin.UserReq
	0,  // 12: admin.UserMgr.Create:input_type -> admin.User
	3,  // 13: admin.UserMgr.Update:input_type -> admin.UserReq
	0,  // 14: admin.UserMgr.Delete:input_type -> admin.User
	0,  // 15: admin.UserMgr.UpdateStatus:input_type -> admin.User
	0,  // 16: admin.UserMgr.UpdatePassword:input_type -> admin.User
	11, // 17: admin.UserMgr.Query:output_type -> unified.Response
	11, // 18: admin.UserMgr.Get:output_type -> unified.Response
	11, // 19: admin.UserMgr.Create:output_type -> unified.Response
	11, // 20: admin.UserMgr.Update:output_type -> unified.Response
	11, // 21: admin.UserMgr.Delete:output_type -> unified.Response
	11, // 22: admin.UserMgr.UpdateStatus:output_type -> unified.Response
	11, // 23: admin.UserMgr.UpdatePassword:output_type -> unified.Response
	17, // [17:24] is the sub-list for method output_type
	10, // [10:17] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_admin_user_proto_init() }
func file_admin_user_proto_init() {
	if File_admin_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_admin_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_admin_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRole); i {
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
		file_admin_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserQueryReq); i {
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
		file_admin_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserReq); i {
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
		file_admin_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserQueryParam); i {
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
		file_admin_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserQueryOptions); i {
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
		file_admin_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserQueryResult); i {
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
			RawDescriptor: file_admin_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_user_proto_goTypes,
		DependencyIndexes: file_admin_user_proto_depIdxs,
		MessageInfos:      file_admin_user_proto_msgTypes,
	}.Build()
	File_admin_user_proto = out.File
	file_admin_user_proto_rawDesc = nil
	file_admin_user_proto_goTypes = nil
	file_admin_user_proto_depIdxs = nil
}
