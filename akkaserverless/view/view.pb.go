// Copyright 2021 Lightbend Inc.

// gRPC interface for Akka Serverless Views.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: view.proto

package view

import (
	protocol "github.com/lightbend/akkaserverless-go-sdk/akkaserverless/protocol"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Input message type for the gRPC stream in.
type ViewStreamIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//	*ViewStreamIn_Receive
	Message isViewStreamIn_Message `protobuf_oneof:"message"`
}

func (x *ViewStreamIn) Reset() {
	*x = ViewStreamIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_view_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewStreamIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewStreamIn) ProtoMessage() {}

func (x *ViewStreamIn) ProtoReflect() protoreflect.Message {
	mi := &file_view_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewStreamIn.ProtoReflect.Descriptor instead.
func (*ViewStreamIn) Descriptor() ([]byte, []int) {
	return file_view_proto_rawDescGZIP(), []int{0}
}

func (m *ViewStreamIn) GetMessage() isViewStreamIn_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *ViewStreamIn) GetReceive() *ReceiveEvent {
	if x, ok := x.GetMessage().(*ViewStreamIn_Receive); ok {
		return x.Receive
	}
	return nil
}

type isViewStreamIn_Message interface {
	isViewStreamIn_Message()
}

type ViewStreamIn_Receive struct {
	Receive *ReceiveEvent `protobuf:"bytes,1,opt,name=receive,proto3,oneof"`
}

func (*ViewStreamIn_Receive) isViewStreamIn_Message() {}

type ReceiveEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName           string             `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	CommandName           string             `protobuf:"bytes,2,opt,name=command_name,json=commandName,proto3" json:"command_name,omitempty"`
	Payload               *anypb.Any         `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Metadata              *protocol.Metadata `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	InitialTable          string             `protobuf:"bytes,5,opt,name=initial_table,json=initialTable,proto3" json:"initial_table,omitempty"`
	Key                   *Key               `protobuf:"bytes,6,opt,name=key,proto3" json:"key,omitempty"`
	BySubjectLookupResult *Row               `protobuf:"bytes,7,opt,name=by_subject_lookup_result,json=bySubjectLookupResult,proto3" json:"by_subject_lookup_result,omitempty"`
}

func (x *ReceiveEvent) Reset() {
	*x = ReceiveEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_view_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveEvent) ProtoMessage() {}

func (x *ReceiveEvent) ProtoReflect() protoreflect.Message {
	mi := &file_view_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveEvent.ProtoReflect.Descriptor instead.
func (*ReceiveEvent) Descriptor() ([]byte, []int) {
	return file_view_proto_rawDescGZIP(), []int{1}
}

func (x *ReceiveEvent) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *ReceiveEvent) GetCommandName() string {
	if x != nil {
		return x.CommandName
	}
	return ""
}

func (x *ReceiveEvent) GetPayload() *anypb.Any {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *ReceiveEvent) GetMetadata() *protocol.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *ReceiveEvent) GetInitialTable() string {
	if x != nil {
		return x.InitialTable
	}
	return ""
}

func (x *ReceiveEvent) GetKey() *Key {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *ReceiveEvent) GetBySubjectLookupResult() *Row {
	if x != nil {
		return x.BySubjectLookupResult
	}
	return nil
}

// Output message type for the gRPC stream out.
type ViewStreamOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//	*ViewStreamOut_Upsert
	Message isViewStreamOut_Message `protobuf_oneof:"message"`
}

func (x *ViewStreamOut) Reset() {
	*x = ViewStreamOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_view_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewStreamOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewStreamOut) ProtoMessage() {}

func (x *ViewStreamOut) ProtoReflect() protoreflect.Message {
	mi := &file_view_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewStreamOut.ProtoReflect.Descriptor instead.
func (*ViewStreamOut) Descriptor() ([]byte, []int) {
	return file_view_proto_rawDescGZIP(), []int{2}
}

func (m *ViewStreamOut) GetMessage() isViewStreamOut_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *ViewStreamOut) GetUpsert() *Upsert {
	if x, ok := x.GetMessage().(*ViewStreamOut_Upsert); ok {
		return x.Upsert
	}
	return nil
}

type isViewStreamOut_Message interface {
	isViewStreamOut_Message()
}

type ViewStreamOut_Upsert struct {
	Upsert *Upsert `protobuf:"bytes,1,opt,name=upsert,proto3,oneof"`
}

func (*ViewStreamOut_Upsert) isViewStreamOut_Message() {}

// to ignore an event and do nothing, use Upsert, but no/empty/null value for row
type Upsert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Row *Row `protobuf:"bytes,1,opt,name=row,proto3" json:"row,omitempty"`
}

func (x *Upsert) Reset() {
	*x = Upsert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_view_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Upsert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Upsert) ProtoMessage() {}

func (x *Upsert) ProtoReflect() protoreflect.Message {
	mi := &file_view_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Upsert.ProtoReflect.Descriptor instead.
func (*Upsert) Descriptor() ([]byte, []int) {
	return file_view_proto_rawDescGZIP(), []int{3}
}

func (x *Upsert) GetRow() *Row {
	if x != nil {
		return x.Row
	}
	return nil
}

type Key struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parts []*KeyPart `protobuf:"bytes,1,rep,name=parts,proto3" json:"parts,omitempty"`
}

func (x *Key) Reset() {
	*x = Key{}
	if protoimpl.UnsafeEnabled {
		mi := &file_view_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Key) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Key) ProtoMessage() {}

func (x *Key) ProtoReflect() protoreflect.Message {
	mi := &file_view_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Key.ProtoReflect.Descriptor instead.
func (*Key) Descriptor() ([]byte, []int) {
	return file_view_proto_rawDescGZIP(), []int{4}
}

func (x *Key) GetParts() []*KeyPart {
	if x != nil {
		return x.Parts
	}
	return nil
}

type KeyPart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Part:
	//	*KeyPart_StringPart
	//	*KeyPart_BytesPart
	//	*KeyPart_IntegerPart
	//	*KeyPart_FloatPart
	Part isKeyPart_Part `protobuf_oneof:"part"`
}

func (x *KeyPart) Reset() {
	*x = KeyPart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_view_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyPart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyPart) ProtoMessage() {}

func (x *KeyPart) ProtoReflect() protoreflect.Message {
	mi := &file_view_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyPart.ProtoReflect.Descriptor instead.
func (*KeyPart) Descriptor() ([]byte, []int) {
	return file_view_proto_rawDescGZIP(), []int{5}
}

func (m *KeyPart) GetPart() isKeyPart_Part {
	if m != nil {
		return m.Part
	}
	return nil
}

func (x *KeyPart) GetStringPart() string {
	if x, ok := x.GetPart().(*KeyPart_StringPart); ok {
		return x.StringPart
	}
	return ""
}

func (x *KeyPart) GetBytesPart() []byte {
	if x, ok := x.GetPart().(*KeyPart_BytesPart); ok {
		return x.BytesPart
	}
	return nil
}

func (x *KeyPart) GetIntegerPart() int64 {
	if x, ok := x.GetPart().(*KeyPart_IntegerPart); ok {
		return x.IntegerPart
	}
	return 0
}

func (x *KeyPart) GetFloatPart() float64 {
	if x, ok := x.GetPart().(*KeyPart_FloatPart); ok {
		return x.FloatPart
	}
	return 0
}

type isKeyPart_Part interface {
	isKeyPart_Part()
}

type KeyPart_StringPart struct {
	StringPart string `protobuf:"bytes,1,opt,name=string_part,json=stringPart,proto3,oneof"`
}

type KeyPart_BytesPart struct {
	BytesPart []byte `protobuf:"bytes,2,opt,name=bytes_part,json=bytesPart,proto3,oneof"`
}

type KeyPart_IntegerPart struct {
	IntegerPart int64 `protobuf:"varint,3,opt,name=integer_part,json=integerPart,proto3,oneof"`
}

type KeyPart_FloatPart struct {
	FloatPart float64 `protobuf:"fixed64,4,opt,name=float_part,json=floatPart,proto3,oneof"`
}

func (*KeyPart_StringPart) isKeyPart_Part() {}

func (*KeyPart_BytesPart) isKeyPart_Part() {}

func (*KeyPart_IntegerPart) isKeyPart_Part() {}

func (*KeyPart_FloatPart) isKeyPart_Part() {}

type Row struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index string     `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Key   *Key       `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value *anypb.Any `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"` // May be unset if a lookup found no row.
}

func (x *Row) Reset() {
	*x = Row{}
	if protoimpl.UnsafeEnabled {
		mi := &file_view_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Row) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Row) ProtoMessage() {}

func (x *Row) ProtoReflect() protoreflect.Message {
	mi := &file_view_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Row.ProtoReflect.Descriptor instead.
func (*Row) Descriptor() ([]byte, []int) {
	return file_view_proto_rawDescGZIP(), []int{6}
}

func (x *Row) GetIndex() string {
	if x != nil {
		return x.Index
	}
	return ""
}

func (x *Row) GetKey() *Key {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *Row) GetValue() *anypb.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_view_proto protoreflect.FileDescriptor

var file_view_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x61, 0x6b,
	0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x1a, 0x28, 0x61, 0x6b, 0x6b,
	0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x62, 0x0a, 0x0c, 0x56, 0x69, 0x65, 0x77, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x6e,
	0x12, 0x47, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65,
	0x73, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x69, 0x65,
	0x77, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00,
	0x52, 0x07, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0xfc, 0x02, 0x0a, 0x0c, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x3e, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x0d, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x34, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x4b, 0x65,
	0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x5b, 0x0a, 0x18, 0x62, 0x79, 0x5f, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x5f, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x6b, 0x6b, 0x61, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x52, 0x6f, 0x77, 0x52, 0x15, 0x62, 0x79,
	0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x5b, 0x0a, 0x0d, 0x56, 0x69, 0x65, 0x77, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x4f, 0x75, 0x74, 0x12, 0x3f, 0x0a, 0x06, 0x75, 0x70, 0x73, 0x65, 0x72, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x69, 0x65, 0x77, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x48, 0x00, 0x52, 0x06, 0x75,
	0x70, 0x73, 0x65, 0x72, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x3e, 0x0a, 0x06, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x12, 0x34, 0x0a, 0x03, 0x72, 0x6f,
	0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x52, 0x6f, 0x77, 0x52, 0x03, 0x72, 0x6f, 0x77,
	0x22, 0x43, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x3c, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x72, 0x74, 0x52, 0x05,
	0x70, 0x61, 0x72, 0x74, 0x73, 0x22, 0x9b, 0x01, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x72,
	0x74, 0x12, 0x21, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x61, 0x72, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x50, 0x61, 0x72, 0x74, 0x12, 0x1f, 0x0a, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x70, 0x61,
	0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x09, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x50, 0x61, 0x72, 0x74, 0x12, 0x23, 0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72,
	0x5f, 0x70, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0b, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x50, 0x61, 0x72, 0x74, 0x12, 0x1f, 0x0a, 0x0a, 0x66, 0x6c,
	0x6f, 0x61, 0x74, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00,
	0x52, 0x09, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x50, 0x61, 0x72, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x70,
	0x61, 0x72, 0x74, 0x22, 0x7d, 0x0a, 0x03, 0x52, 0x6f, 0x77, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x34, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x4b, 0x65,
	0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x32, 0x70, 0x0a, 0x05, 0x56, 0x69, 0x65, 0x77, 0x73, 0x12, 0x67, 0x0a, 0x06, 0x48,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x2b, 0x2e, 0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x49, 0x6e, 0x1a, 0x2c, 0x2e, 0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c,
	0x65, 0x73, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x69,
	0x65, 0x77, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4f, 0x75, 0x74,
	0x28, 0x01, 0x30, 0x01, 0x42, 0x62, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6b, 0x6b, 0x61,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x62, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2d, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f,
	0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2f, 0x76,
	0x69, 0x65, 0x77, 0x3b, 0x76, 0x69, 0x65, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_view_proto_rawDescOnce sync.Once
	file_view_proto_rawDescData = file_view_proto_rawDesc
)

func file_view_proto_rawDescGZIP() []byte {
	file_view_proto_rawDescOnce.Do(func() {
		file_view_proto_rawDescData = protoimpl.X.CompressGZIP(file_view_proto_rawDescData)
	})
	return file_view_proto_rawDescData
}

var file_view_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_view_proto_goTypes = []interface{}{
	(*ViewStreamIn)(nil),      // 0: akkaserverless.component.view.ViewStreamIn
	(*ReceiveEvent)(nil),      // 1: akkaserverless.component.view.ReceiveEvent
	(*ViewStreamOut)(nil),     // 2: akkaserverless.component.view.ViewStreamOut
	(*Upsert)(nil),            // 3: akkaserverless.component.view.Upsert
	(*Key)(nil),               // 4: akkaserverless.component.view.Key
	(*KeyPart)(nil),           // 5: akkaserverless.component.view.KeyPart
	(*Row)(nil),               // 6: akkaserverless.component.view.Row
	(*anypb.Any)(nil),         // 7: google.protobuf.Any
	(*protocol.Metadata)(nil), // 8: akkaserverless.component.Metadata
}
var file_view_proto_depIdxs = []int32{
	1,  // 0: akkaserverless.component.view.ViewStreamIn.receive:type_name -> akkaserverless.component.view.ReceiveEvent
	7,  // 1: akkaserverless.component.view.ReceiveEvent.payload:type_name -> google.protobuf.Any
	8,  // 2: akkaserverless.component.view.ReceiveEvent.metadata:type_name -> akkaserverless.component.Metadata
	4,  // 3: akkaserverless.component.view.ReceiveEvent.key:type_name -> akkaserverless.component.view.Key
	6,  // 4: akkaserverless.component.view.ReceiveEvent.by_subject_lookup_result:type_name -> akkaserverless.component.view.Row
	3,  // 5: akkaserverless.component.view.ViewStreamOut.upsert:type_name -> akkaserverless.component.view.Upsert
	6,  // 6: akkaserverless.component.view.Upsert.row:type_name -> akkaserverless.component.view.Row
	5,  // 7: akkaserverless.component.view.Key.parts:type_name -> akkaserverless.component.view.KeyPart
	4,  // 8: akkaserverless.component.view.Row.key:type_name -> akkaserverless.component.view.Key
	7,  // 9: akkaserverless.component.view.Row.value:type_name -> google.protobuf.Any
	0,  // 10: akkaserverless.component.view.Views.Handle:input_type -> akkaserverless.component.view.ViewStreamIn
	2,  // 11: akkaserverless.component.view.Views.Handle:output_type -> akkaserverless.component.view.ViewStreamOut
	11, // [11:12] is the sub-list for method output_type
	10, // [10:11] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_view_proto_init() }
func file_view_proto_init() {
	if File_view_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_view_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewStreamIn); i {
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
		file_view_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveEvent); i {
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
		file_view_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewStreamOut); i {
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
		file_view_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Upsert); i {
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
		file_view_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Key); i {
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
		file_view_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyPart); i {
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
		file_view_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Row); i {
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
	file_view_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ViewStreamIn_Receive)(nil),
	}
	file_view_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ViewStreamOut_Upsert)(nil),
	}
	file_view_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*KeyPart_StringPart)(nil),
		(*KeyPart_BytesPart)(nil),
		(*KeyPart_IntegerPart)(nil),
		(*KeyPart_FloatPart)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_view_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_view_proto_goTypes,
		DependencyIndexes: file_view_proto_depIdxs,
		MessageInfos:      file_view_proto_msgTypes,
	}.Build()
	File_view_proto = out.File
	file_view_proto_rawDesc = nil
	file_view_proto_goTypes = nil
	file_view_proto_depIdxs = nil
}
