// Copyright 2021 Lightbend Inc.

// Extension for specifying which field in a message is to be considered an
// entity key, for the purposes of associating gRPC calls with entities and
// sharding.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: legacy_entity_key.proto

package akkaserverless

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_legacy_entity_key_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50002,
		Name:          "akkaserverless.legacy_entity_key",
		Tag:           "varint,50002,opt,name=legacy_entity_key",
		Filename:      "legacy_entity_key.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional bool legacy_entity_key = 50002;
	E_LegacyEntityKey = &file_legacy_entity_key_proto_extTypes[0]
)

var File_legacy_entity_key_proto protoreflect.FileDescriptor

var file_legacy_entity_key_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f,
	0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x6b, 0x6b, 0x61, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x4b, 0x0a, 0x11, 0x6c,
	0x65, 0x67, 0x61, 0x63, 0x79, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x6b, 0x65, 0x79,
	0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xd2, 0x86, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x4b, 0x65, 0x79, 0x42, 0x60, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e,
	0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x50, 0x01,
	0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x62, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x6c, 0x65, 0x73, 0x73, 0x2d, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x6b, 0x6b,
	0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x3b, 0x61, 0x6b, 0x6b, 0x61,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_legacy_entity_key_proto_goTypes = []interface{}{
	(*descriptorpb.FieldOptions)(nil), // 0: google.protobuf.FieldOptions
}
var file_legacy_entity_key_proto_depIdxs = []int32{
	0, // 0: akkaserverless.legacy_entity_key:extendee -> google.protobuf.FieldOptions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_legacy_entity_key_proto_init() }
func file_legacy_entity_key_proto_init() {
	if File_legacy_entity_key_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_legacy_entity_key_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_legacy_entity_key_proto_goTypes,
		DependencyIndexes: file_legacy_entity_key_proto_depIdxs,
		ExtensionInfos:    file_legacy_entity_key_proto_extTypes,
	}.Build()
	File_legacy_entity_key_proto = out.File
	file_legacy_entity_key_proto_rawDesc = nil
	file_legacy_entity_key_proto_goTypes = nil
	file_legacy_entity_key_proto_depIdxs = nil
}
