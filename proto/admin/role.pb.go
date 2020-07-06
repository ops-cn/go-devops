// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: admin/role.proto

package admin

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Sequence int32  `protobuf:"varint,3,opt,name=Sequence,proto3" json:"Sequence,omitempty"`
	Memo     string `protobuf:"bytes,4,opt,name=Memo,proto3" json:"Memo,omitempty"`
	Status   int32  `protobuf:"varint,5,opt,name=Status,proto3" json:"Status,omitempty"`
	Creator  string `protobuf:"bytes,6,opt,name=Creator,proto3" json:"Creator,omitempty"`
	// @inject_tag: valid:"created_at"
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,7,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	// @inject_tag: valid:"updated_at"
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,8,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_role_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_admin_role_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_admin_role_proto_rawDescGZIP(), []int{0}
}

func (x *Role) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Role) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Role) GetSequence() int32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *Role) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *Role) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Role) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *Role) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Role) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type RoleQueryParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         []string `protobuf:"bytes,1,rep,name=ID,proto3" json:"ID,omitempty"`
	Name       string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	QueryValue string   `protobuf:"bytes,3,opt,name=QueryValue,proto3" json:"QueryValue,omitempty"`
	UserID     string   `protobuf:"bytes,4,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Status     int32    `protobuf:"varint,5,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *RoleQueryParam) Reset() {
	*x = RoleQueryParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_role_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleQueryParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleQueryParam) ProtoMessage() {}

func (x *RoleQueryParam) ProtoReflect() protoreflect.Message {
	mi := &file_admin_role_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleQueryParam.ProtoReflect.Descriptor instead.
func (*RoleQueryParam) Descriptor() ([]byte, []int) {
	return file_admin_role_proto_rawDescGZIP(), []int{1}
}

func (x *RoleQueryParam) GetID() []string {
	if x != nil {
		return x.ID
	}
	return nil
}

func (x *RoleQueryParam) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RoleQueryParam) GetQueryValue() string {
	if x != nil {
		return x.QueryValue
	}
	return ""
}

func (x *RoleQueryParam) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *RoleQueryParam) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_admin_role_proto protoreflect.FileDescriptor

var file_admin_role_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x02, 0x0a, 0x04, 0x52,
	0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x53, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4d, 0x65, 0x6d, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4d, 0x65, 0x6d, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x84, 0x01,
	0x0a, 0x0e, 0x52, 0x6f, 0x6c, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x73, 0x2d, 0x63, 0x6e, 0x2f, 0x67, 0x6f, 0x2d, 0x64, 0x65, 0x76,
	0x6f, 0x70, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_role_proto_rawDescOnce sync.Once
	file_admin_role_proto_rawDescData = file_admin_role_proto_rawDesc
)

func file_admin_role_proto_rawDescGZIP() []byte {
	file_admin_role_proto_rawDescOnce.Do(func() {
		file_admin_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_role_proto_rawDescData)
	})
	return file_admin_role_proto_rawDescData
}

var file_admin_role_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_admin_role_proto_goTypes = []interface{}{
	(*Role)(nil),                // 0: admin.Role
	(*RoleQueryParam)(nil),      // 1: admin.RoleQueryParam
	(*timestamp.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_admin_role_proto_depIdxs = []int32{
	2, // 0: admin.Role.CreatedAt:type_name -> google.protobuf.Timestamp
	2, // 1: admin.Role.UpdatedAt:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_admin_role_proto_init() }
func file_admin_role_proto_init() {
	if File_admin_role_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_admin_role_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_admin_role_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleQueryParam); i {
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
			RawDescriptor: file_admin_role_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_admin_role_proto_goTypes,
		DependencyIndexes: file_admin_role_proto_depIdxs,
		MessageInfos:      file_admin_role_proto_msgTypes,
	}.Build()
	File_admin_role_proto = out.File
	file_admin_role_proto_rawDesc = nil
	file_admin_role_proto_goTypes = nil
	file_admin_role_proto_depIdxs = nil
}
