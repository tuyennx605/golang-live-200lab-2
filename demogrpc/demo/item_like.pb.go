// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: demogrpc/item_like.proto

package demo

import (
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

type GetItemLikeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int32 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *GetItemLikeReq) Reset() {
	*x = GetItemLikeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demogrpc_item_like_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemLikeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemLikeReq) ProtoMessage() {}

func (x *GetItemLikeReq) ProtoReflect() protoreflect.Message {
	mi := &file_demogrpc_item_like_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemLikeReq.ProtoReflect.Descriptor instead.
func (*GetItemLikeReq) Descriptor() ([]byte, []int) {
	return file_demogrpc_item_like_proto_rawDescGZIP(), []int{0}
}

func (x *GetItemLikeReq) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type ItemLikesResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result map[int32]int32 `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *ItemLikesResp) Reset() {
	*x = ItemLikesResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demogrpc_item_like_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemLikesResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemLikesResp) ProtoMessage() {}

func (x *ItemLikesResp) ProtoReflect() protoreflect.Message {
	mi := &file_demogrpc_item_like_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemLikesResp.ProtoReflect.Descriptor instead.
func (*ItemLikesResp) Descriptor() ([]byte, []int) {
	return file_demogrpc_item_like_proto_rawDescGZIP(), []int{1}
}

func (x *ItemLikesResp) GetResult() map[int32]int32 {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_demogrpc_item_like_proto protoreflect.FileDescriptor

var file_demogrpc_item_like_proto_rawDesc = []byte{
	0x0a, 0x18, 0x64, 0x65, 0x6d, 0x6f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x5f,
	0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x64, 0x65, 0x6d, 0x6f,
	0x22, 0x22, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x6b, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x03, 0x69, 0x64, 0x73, 0x22, 0x83, 0x01, 0x0a, 0x0d, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x6b,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x37, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a,
	0x39, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x4c, 0x0a, 0x0f, 0x49, 0x74,
	0x65, 0x6d, 0x4c, 0x69, 0x6b, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x14, 0x2e,
	0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x6b, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c,
	0x69, 0x6b, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x42, 0x07, 0x5a, 0x05, 0x64, 0x65, 0x6d, 0x6f,
	0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_demogrpc_item_like_proto_rawDescOnce sync.Once
	file_demogrpc_item_like_proto_rawDescData = file_demogrpc_item_like_proto_rawDesc
)

func file_demogrpc_item_like_proto_rawDescGZIP() []byte {
	file_demogrpc_item_like_proto_rawDescOnce.Do(func() {
		file_demogrpc_item_like_proto_rawDescData = protoimpl.X.CompressGZIP(file_demogrpc_item_like_proto_rawDescData)
	})
	return file_demogrpc_item_like_proto_rawDescData
}

var file_demogrpc_item_like_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_demogrpc_item_like_proto_goTypes = []any{
	(*GetItemLikeReq)(nil), // 0: demo.GetItemLikeReq
	(*ItemLikesResp)(nil),  // 1: demo.ItemLikesResp
	nil,                    // 2: demo.ItemLikesResp.ResultEntry
}
var file_demogrpc_item_like_proto_depIdxs = []int32{
	2, // 0: demo.ItemLikesResp.result:type_name -> demo.ItemLikesResp.ResultEntry
	0, // 1: demo.ItemLikeService.GetItemLikes:input_type -> demo.GetItemLikeReq
	1, // 2: demo.ItemLikeService.GetItemLikes:output_type -> demo.ItemLikesResp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_demogrpc_item_like_proto_init() }
func file_demogrpc_item_like_proto_init() {
	if File_demogrpc_item_like_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_demogrpc_item_like_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetItemLikeReq); i {
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
		file_demogrpc_item_like_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ItemLikesResp); i {
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
			RawDescriptor: file_demogrpc_item_like_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_demogrpc_item_like_proto_goTypes,
		DependencyIndexes: file_demogrpc_item_like_proto_depIdxs,
		MessageInfos:      file_demogrpc_item_like_proto_msgTypes,
	}.Build()
	File_demogrpc_item_like_proto = out.File
	file_demogrpc_item_like_proto_rawDesc = nil
	file_demogrpc_item_like_proto_goTypes = nil
	file_demogrpc_item_like_proto_depIdxs = nil
}
