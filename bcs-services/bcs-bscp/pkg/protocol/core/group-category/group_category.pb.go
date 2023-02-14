// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: group_category.proto

package pbgc

import (
	base "bscp.io/pkg/protocol/core/base"
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

// GroupCategory source resource reference: pkg/dal/table/group_category.go
type GroupCategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Spec       *GroupCategorySpec       `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Attachment *GroupCategoryAttachment `protobuf:"bytes,3,opt,name=attachment,proto3" json:"attachment,omitempty"`
	Revision   *base.CreatedRevision    `protobuf:"bytes,4,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (x *GroupCategory) Reset() {
	*x = GroupCategory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_category_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupCategory) ProtoMessage() {}

func (x *GroupCategory) ProtoReflect() protoreflect.Message {
	mi := &file_group_category_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupCategory.ProtoReflect.Descriptor instead.
func (*GroupCategory) Descriptor() ([]byte, []int) {
	return file_group_category_proto_rawDescGZIP(), []int{0}
}

func (x *GroupCategory) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GroupCategory) GetSpec() *GroupCategorySpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *GroupCategory) GetAttachment() *GroupCategoryAttachment {
	if x != nil {
		return x.Attachment
	}
	return nil
}

func (x *GroupCategory) GetRevision() *base.CreatedRevision {
	if x != nil {
		return x.Revision
	}
	return nil
}

// GroupCategorySpec source resource reference: pkg/dal/table/group_category.go
type GroupCategorySpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GroupCategorySpec) Reset() {
	*x = GroupCategorySpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_category_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupCategorySpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupCategorySpec) ProtoMessage() {}

func (x *GroupCategorySpec) ProtoReflect() protoreflect.Message {
	mi := &file_group_category_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupCategorySpec.ProtoReflect.Descriptor instead.
func (*GroupCategorySpec) Descriptor() ([]byte, []int) {
	return file_group_category_proto_rawDescGZIP(), []int{1}
}

func (x *GroupCategorySpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// GroupCategoryAttachment source resource reference: pkg/dal/table/group_category.go
type GroupCategoryAttachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizId uint32 `protobuf:"varint,1,opt,name=biz_id,json=bizId,proto3" json:"biz_id,omitempty"`
	AppId uint32 `protobuf:"varint,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
}

func (x *GroupCategoryAttachment) Reset() {
	*x = GroupCategoryAttachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_category_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupCategoryAttachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupCategoryAttachment) ProtoMessage() {}

func (x *GroupCategoryAttachment) ProtoReflect() protoreflect.Message {
	mi := &file_group_category_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupCategoryAttachment.ProtoReflect.Descriptor instead.
func (*GroupCategoryAttachment) Descriptor() ([]byte, []int) {
	return file_group_category_proto_rawDescGZIP(), []int{2}
}

func (x *GroupCategoryAttachment) GetBizId() uint32 {
	if x != nil {
		return x.BizId
	}
	return 0
}

func (x *GroupCategoryAttachment) GetAppId() uint32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

var File_group_category_proto protoreflect.FileDescriptor

var file_group_category_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x62, 0x67, 0x63, 0x1a, 0x29, 0x62, 0x73,
	0x63, 0x70, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x01, 0x0a, 0x0d, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x04, 0x73, 0x70, 0x65,
	0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x62, 0x67, 0x63, 0x2e, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x70, 0x65, 0x63,
	0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x3d, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x62, 0x67,
	0x63, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x62, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x27, 0x0a, 0x11, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x47, 0x0a, 0x17, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x15,
	0x0a, 0x06, 0x62, 0x69, 0x7a, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x62, 0x69, 0x7a, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x42, 0x2f, 0x5a, 0x2d,
	0x62, 0x73, 0x63, 0x70, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2d,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x3b, 0x70, 0x62, 0x67, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_group_category_proto_rawDescOnce sync.Once
	file_group_category_proto_rawDescData = file_group_category_proto_rawDesc
)

func file_group_category_proto_rawDescGZIP() []byte {
	file_group_category_proto_rawDescOnce.Do(func() {
		file_group_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_group_category_proto_rawDescData)
	})
	return file_group_category_proto_rawDescData
}

var file_group_category_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_group_category_proto_goTypes = []interface{}{
	(*GroupCategory)(nil),           // 0: pbgc.GroupCategory
	(*GroupCategorySpec)(nil),       // 1: pbgc.GroupCategorySpec
	(*GroupCategoryAttachment)(nil), // 2: pbgc.GroupCategoryAttachment
	(*base.CreatedRevision)(nil),    // 3: pbbase.CreatedRevision
}
var file_group_category_proto_depIdxs = []int32{
	1, // 0: pbgc.GroupCategory.spec:type_name -> pbgc.GroupCategorySpec
	2, // 1: pbgc.GroupCategory.attachment:type_name -> pbgc.GroupCategoryAttachment
	3, // 2: pbgc.GroupCategory.revision:type_name -> pbbase.CreatedRevision
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_group_category_proto_init() }
func file_group_category_proto_init() {
	if File_group_category_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_group_category_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupCategory); i {
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
		file_group_category_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupCategorySpec); i {
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
		file_group_category_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupCategoryAttachment); i {
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
			RawDescriptor: file_group_category_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_group_category_proto_goTypes,
		DependencyIndexes: file_group_category_proto_depIdxs,
		MessageInfos:      file_group_category_proto_msgTypes,
	}.Build()
	File_group_category_proto = out.File
	file_group_category_proto_rawDesc = nil
	file_group_category_proto_goTypes = nil
	file_group_category_proto_depIdxs = nil
}
