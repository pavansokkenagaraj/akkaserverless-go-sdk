// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.2
// source: hotitems.proto

package shoppingcart

import (
	_ "github.com/pavansokkenagaraj/akkaserverless-go-sdk/akkaserverless"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Quantity  int32  `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hotitems_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_hotitems_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_hotitems_proto_rawDescGZIP(), []int{0}
}

func (x *Item) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Item) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

var File_hotitems_proto protoreflect.FileDescriptor

var file_hotitems_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x68, 0x6f, 0x74, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x14, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x63, 0x61, 0x72, 0x74, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x55, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x32, 0x58, 0x0a, 0x0f, 0x48, 0x6f, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0f, 0x49, 0x74,
	0x65, 0x6d, 0x41, 0x64, 0x64, 0x65, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x12, 0x1a, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x63, 0x61, 0x72, 0x74, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x42, 0x65, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x5a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65, 0x69, 0x6f, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x63, 0x72, 0x64,
	0x74, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x73,
	0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x74, 0x3b, 0x73, 0x68, 0x6f, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x72, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hotitems_proto_rawDescOnce sync.Once
	file_hotitems_proto_rawDescData = file_hotitems_proto_rawDesc
)

func file_hotitems_proto_rawDescGZIP() []byte {
	file_hotitems_proto_rawDescOnce.Do(func() {
		file_hotitems_proto_rawDescData = protoimpl.X.CompressGZIP(file_hotitems_proto_rawDescData)
	})
	return file_hotitems_proto_rawDescData
}

var file_hotitems_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_hotitems_proto_goTypes = []interface{}{
	(*Item)(nil),        // 0: example.shoppingcart.Item
	(*empty.Empty)(nil), // 1: google.protobuf.Empty
}
var file_hotitems_proto_depIdxs = []int32{
	0, // 0: example.shoppingcart.HotItemsService.ItemAddedToCart:input_type -> example.shoppingcart.Item
	1, // 1: example.shoppingcart.HotItemsService.ItemAddedToCart:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hotitems_proto_init() }
func file_hotitems_proto_init() {
	if File_hotitems_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hotitems_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
			RawDescriptor: file_hotitems_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hotitems_proto_goTypes,
		DependencyIndexes: file_hotitems_proto_depIdxs,
		MessageInfos:      file_hotitems_proto_msgTypes,
	}.Build()
	File_hotitems_proto = out.File
	file_hotitems_proto_rawDesc = nil
	file_hotitems_proto_goTypes = nil
	file_hotitems_proto_depIdxs = nil
}
