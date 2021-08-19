// Copyright 2021 Lightbend Inc.

// gRPC interface for common messages used in protocols between the proxy and
// the user.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: component.proto

package protocol

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
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

// Transport-specific metadata associated with a message.
//
// The semantics of the metadata are not defined in this protocol, but rather,
// depend on the transport on which a particular instance of the metadata maps
// to. What keys or values are allowed or disallowed, whether duplicate values
// for the same key are allowed and how they are handled, and whether key names
// are case sensitive or not, are all undefined in the context of the Akka
// Serverless protocol.
//
// If a metadata entry associated with a message can't be expressed in an
// underlying transport, for example, due to invalid characters in a key or
// value, the behavior of the proxy is undefined. This is because metadata is
// transport specific, so if the user function chooses to use metadata, it is
// choosing to be specific to a particular transport, which is beyond the scope
// of the Akka Serverless protocol, and it's therefore the user function's
// responsibility to adhere to the semantics of that transport. The proxy MAY
// decide to drop metadata entries if it knows they are invalid or unsupported.
// If a metadata entry is dropped, the proxy MAY inform the user function that
// the entry was dropped by sending an error message to the
// Discovery.ReportError gRPC call.
//
// The metadata MAY also contain CloudEvent metadata. If a message comes from an
// Akka Serverless event source, the Akka Serverless proxy MUST attach
// CloudEvent metadata to it if the event doesn't already have CloudEvent
// metadata attached to it. This metadata SHALL be encoded according to the
// binary mode of the CloudEvent HTTP protocol binding, which can be found here:
//
// https://github.com/cloudevents/spec/blob/master/http-protocol-binding.md
//
// The Akka Serverless proxy MAY synthesize appropriate values for Akka
// Serverless metadata if no equivalent metadata exists in the event source, for
// example, if there is no type, the Akka Serverless proxy MAY use the name of
// the gRPC message as the CloudEvent type, and if there is no source, the Akka
// Serverless proxy MAY use the name of the topic as the source.
//
// If an incoming message does have CloudEvent metadata attached to it, the Akka
// Serverless proxy MUST transcode that CloudEvent metadata to the HTTP protocol
// binding as described above.
//
// Messages sent from the user function to an event destination MAY include
// CloudEvent metadata. If they include any CloudEvent metadata, they MUST
// include all required CloudEvent attributes, including id, source, specversion
// and type. The behavior of the proxy is undefined if some of these attributes,
// but not others, are included - the proxy MAY ignore them all, or MAY generate
// values itself, but SHOULD NOT fail sending the message. If the destination
// for the message is an event destination, the Akka Serverless proxy MUST
// transcode the supplied Akka Serverless metadata to a binding appropriate for
// the underlying transport for that event destination, it MUST NOT pass the
// CloudEvent metadata as is unless the transport uses the same binding rules.
type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The metadata entries.
	Entries []*MetadataEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_component_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_component_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_component_proto_rawDescGZIP(), []int{0}
}

func (x *Metadata) GetEntries() []*MetadataEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

// A metadata entry.
type MetadataEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Key for the entry. Valid keys are determined by the underlying transport.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The value.
	//
	// Types that are assignable to Value:
	//	*MetadataEntry_StringValue
	//	*MetadataEntry_BytesValue
	Value isMetadataEntry_Value `protobuf_oneof:"value"`
}

func (x *MetadataEntry) Reset() {
	*x = MetadataEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_component_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataEntry) ProtoMessage() {}

func (x *MetadataEntry) ProtoReflect() protoreflect.Message {
	mi := &file_component_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataEntry.ProtoReflect.Descriptor instead.
func (*MetadataEntry) Descriptor() ([]byte, []int) {
	return file_component_proto_rawDescGZIP(), []int{1}
}

func (x *MetadataEntry) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (m *MetadataEntry) GetValue() isMetadataEntry_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *MetadataEntry) GetStringValue() string {
	if x, ok := x.GetValue().(*MetadataEntry_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (x *MetadataEntry) GetBytesValue() []byte {
	if x, ok := x.GetValue().(*MetadataEntry_BytesValue); ok {
		return x.BytesValue
	}
	return nil
}

type isMetadataEntry_Value interface {
	isMetadataEntry_Value()
}

type MetadataEntry_StringValue struct {
	// A string value. Valid values are determined by the underlying transport.
	//
	// If the transport does not support string values, the behavior of the Akka
	// Serverless proxy is undefined from the point of view of this protocol. If
	// there is a convention in the protocol for encoding string values as UTF-8
	// bytes, then the Akka Serverless proxy MAY do that.
	StringValue string `protobuf:"bytes,2,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type MetadataEntry_BytesValue struct {
	// A bytes value. Valid values are determined by the underlying transport.
	//
	// If the transport does not support bytes values, the behavior of the Akka
	// Serverless proxy is undefined from the point of view of this protocol. If
	// there is a convention in the protocol for encoding bytes values as Base64
	// encoded strings, then the Akka Serverless proxy MAY do that.
	BytesValue []byte `protobuf:"bytes,3,opt,name=bytes_value,json=bytesValue,proto3,oneof"`
}

func (*MetadataEntry_StringValue) isMetadataEntry_Value() {}

func (*MetadataEntry_BytesValue) isMetadataEntry_Value() {}

// A reply to the sender.
type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The reply payload
	Payload *anypb.Any `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	// Metadata for the reply
	//
	// Not all transports support per message metadata, for example, gRPC doesn't.
	// The Akka Serverless proxy MAY ignore the metadata in this case, or it MAY
	// lift the metadata into another place, for example, in gRPC, a unary call
	// MAY have its reply metadata placed in the headers of the HTTP response, or
	// the first reply to a streamed call MAY have its metadata placed in the
	// headers of the HTTP response.
	//
	// If the metadata is ignored, the Akka Serverless proxy MAY notify the user
	// function by sending an error message to the Discovery.ReportError gRPC
	// call.
	Metadata *Metadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_component_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_component_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_component_proto_rawDescGZIP(), []int{2}
}

func (x *Reply) GetPayload() *anypb.Any {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Reply) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// Forwards handling of this request to another component.
type Forward struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the service to forward to.
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// The name of the command.
	CommandName string `protobuf:"bytes,2,opt,name=command_name,json=commandName,proto3" json:"command_name,omitempty"`
	// The payload.
	Payload *anypb.Any `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	// The metadata to include with the forward
	Metadata *Metadata `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *Forward) Reset() {
	*x = Forward{}
	if protoimpl.UnsafeEnabled {
		mi := &file_component_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Forward) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Forward) ProtoMessage() {}

func (x *Forward) ProtoReflect() protoreflect.Message {
	mi := &file_component_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Forward.ProtoReflect.Descriptor instead.
func (*Forward) Descriptor() ([]byte, []int) {
	return file_component_proto_rawDescGZIP(), []int{3}
}

func (x *Forward) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *Forward) GetCommandName() string {
	if x != nil {
		return x.CommandName
	}
	return ""
}

func (x *Forward) GetPayload() *anypb.Any {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Forward) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// An action for the client
// (see ActionResponse in action.proto)
type ClientAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Action:
	//	*ClientAction_Reply
	//	*ClientAction_Forward
	//	*ClientAction_Failure
	Action isClientAction_Action `protobuf_oneof:"action"`
}

func (x *ClientAction) Reset() {
	*x = ClientAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_component_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientAction) ProtoMessage() {}

func (x *ClientAction) ProtoReflect() protoreflect.Message {
	mi := &file_component_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientAction.ProtoReflect.Descriptor instead.
func (*ClientAction) Descriptor() ([]byte, []int) {
	return file_component_proto_rawDescGZIP(), []int{4}
}

func (m *ClientAction) GetAction() isClientAction_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *ClientAction) GetReply() *Reply {
	if x, ok := x.GetAction().(*ClientAction_Reply); ok {
		return x.Reply
	}
	return nil
}

func (x *ClientAction) GetForward() *Forward {
	if x, ok := x.GetAction().(*ClientAction_Forward); ok {
		return x.Forward
	}
	return nil
}

func (x *ClientAction) GetFailure() *Failure {
	if x, ok := x.GetAction().(*ClientAction_Failure); ok {
		return x.Failure
	}
	return nil
}

type isClientAction_Action interface {
	isClientAction_Action()
}

type ClientAction_Reply struct {
	// Send a reply
	Reply *Reply `protobuf:"bytes,1,opt,name=reply,proto3,oneof"`
}

type ClientAction_Forward struct {
	// Forward to another component
	Forward *Forward `protobuf:"bytes,2,opt,name=forward,proto3,oneof"`
}

type ClientAction_Failure struct {
	// Send a failure to the client
	Failure *Failure `protobuf:"bytes,3,opt,name=failure,proto3,oneof"`
}

func (*ClientAction_Reply) isClientAction_Action() {}

func (*ClientAction_Forward) isClientAction_Action() {}

func (*ClientAction_Failure) isClientAction_Action() {}

// A side effect to be done after this command is handled.
type SideEffect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the service to perform the side effect on.
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// The name of the command.
	CommandName string `protobuf:"bytes,2,opt,name=command_name,json=commandName,proto3" json:"command_name,omitempty"`
	// The payload of the command.
	Payload *anypb.Any `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	// Whether this side effect should be performed synchronously, ie, before the
	// reply is eventually sent, or not.
	Synchronous bool `protobuf:"varint,4,opt,name=synchronous,proto3" json:"synchronous,omitempty"`
	// The metadata to include with the side effect
	Metadata *Metadata `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *SideEffect) Reset() {
	*x = SideEffect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_component_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SideEffect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SideEffect) ProtoMessage() {}

func (x *SideEffect) ProtoReflect() protoreflect.Message {
	mi := &file_component_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SideEffect.ProtoReflect.Descriptor instead.
func (*SideEffect) Descriptor() ([]byte, []int) {
	return file_component_proto_rawDescGZIP(), []int{5}
}

func (x *SideEffect) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *SideEffect) GetCommandName() string {
	if x != nil {
		return x.CommandName
	}
	return ""
}

func (x *SideEffect) GetPayload() *anypb.Any {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *SideEffect) GetSynchronous() bool {
	if x != nil {
		return x.Synchronous
	}
	return false
}

func (x *SideEffect) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type StreamCancelled struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the component
	ComponentId string `protobuf:"bytes,1,opt,name=component_id,json=componentId,proto3" json:"component_id,omitempty"`
	// The command id
	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StreamCancelled) Reset() {
	*x = StreamCancelled{}
	if protoimpl.UnsafeEnabled {
		mi := &file_component_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamCancelled) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamCancelled) ProtoMessage() {}

func (x *StreamCancelled) ProtoReflect() protoreflect.Message {
	mi := &file_component_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamCancelled.ProtoReflect.Descriptor instead.
func (*StreamCancelled) Descriptor() ([]byte, []int) {
	return file_component_proto_rawDescGZIP(), []int{6}
}

func (x *StreamCancelled) GetComponentId() string {
	if x != nil {
		return x.ComponentId
	}
	return ""
}

func (x *StreamCancelled) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// A failure reply. If this is returned, it will be translated into a gRPC
// unknown error with the corresponding description if supplied.
type Failure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the command being replied to. Must match the input command.
	CommandId int64 `protobuf:"varint,1,opt,name=command_id,json=commandId,proto3" json:"command_id,omitempty"`
	// A description of the error.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Whether this failure should trigger an entity restart.
	Restart bool `protobuf:"varint,3,opt,name=restart,proto3" json:"restart,omitempty"`
}

func (x *Failure) Reset() {
	*x = Failure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_component_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Failure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Failure) ProtoMessage() {}

func (x *Failure) ProtoReflect() protoreflect.Message {
	mi := &file_component_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Failure.ProtoReflect.Descriptor instead.
func (*Failure) Descriptor() ([]byte, []int) {
	return file_component_proto_rawDescGZIP(), []int{7}
}

func (x *Failure) GetCommandId() int64 {
	if x != nil {
		return x.CommandId
	}
	return 0
}

func (x *Failure) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Failure) GetRestart() bool {
	if x != nil {
		return x.Restart
	}
	return false
}

var File_component_proto protoreflect.FileDescriptor

var file_component_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x18, 0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x41, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07,
	0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x72, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x0c, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x21, 0x0a, 0x0b, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x77, 0x0a, 0x05, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x2e, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x3e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x22, 0xbf, 0x01, 0x0a, 0x07, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x3e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x6b, 0x6b, 0x61, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0xcf, 0x01, 0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x48, 0x00, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x3d, 0x0a, 0x07, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65,
	0x73, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x46, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x48, 0x00, 0x52, 0x07, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x12,
	0x3d, 0x0a, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x46, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x48, 0x00, 0x52, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x42, 0x08,
	0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xe4, 0x01, 0x0a, 0x0a, 0x53, 0x69, 0x64,
	0x65, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x6f, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x6f, 0x75, 0x73, 0x12,
	0x3e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65,
	0x73, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x44, 0x0a, 0x0f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c,
	0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x64, 0x0a, 0x07, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x42, 0x6a, 0x0a, 0x1b, 0x63,
	0x6f, 0x6d, 0x2e, 0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x62, 0x65, 0x6e, 0x64,
	0x2f, 0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2d,
	0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x6c, 0x65, 0x73, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_component_proto_rawDescOnce sync.Once
	file_component_proto_rawDescData = file_component_proto_rawDesc
)

func file_component_proto_rawDescGZIP() []byte {
	file_component_proto_rawDescOnce.Do(func() {
		file_component_proto_rawDescData = protoimpl.X.CompressGZIP(file_component_proto_rawDescData)
	})
	return file_component_proto_rawDescData
}

var file_component_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_component_proto_goTypes = []interface{}{
	(*Metadata)(nil),        // 0: akkaserverless.component.Metadata
	(*MetadataEntry)(nil),   // 1: akkaserverless.component.MetadataEntry
	(*Reply)(nil),           // 2: akkaserverless.component.Reply
	(*Forward)(nil),         // 3: akkaserverless.component.Forward
	(*ClientAction)(nil),    // 4: akkaserverless.component.ClientAction
	(*SideEffect)(nil),      // 5: akkaserverless.component.SideEffect
	(*StreamCancelled)(nil), // 6: akkaserverless.component.StreamCancelled
	(*Failure)(nil),         // 7: akkaserverless.component.Failure
	(*anypb.Any)(nil),       // 8: google.protobuf.Any
}
var file_component_proto_depIdxs = []int32{
	1,  // 0: akkaserverless.component.Metadata.entries:type_name -> akkaserverless.component.MetadataEntry
	8,  // 1: akkaserverless.component.Reply.payload:type_name -> google.protobuf.Any
	0,  // 2: akkaserverless.component.Reply.metadata:type_name -> akkaserverless.component.Metadata
	8,  // 3: akkaserverless.component.Forward.payload:type_name -> google.protobuf.Any
	0,  // 4: akkaserverless.component.Forward.metadata:type_name -> akkaserverless.component.Metadata
	2,  // 5: akkaserverless.component.ClientAction.reply:type_name -> akkaserverless.component.Reply
	3,  // 6: akkaserverless.component.ClientAction.forward:type_name -> akkaserverless.component.Forward
	7,  // 7: akkaserverless.component.ClientAction.failure:type_name -> akkaserverless.component.Failure
	8,  // 8: akkaserverless.component.SideEffect.payload:type_name -> google.protobuf.Any
	0,  // 9: akkaserverless.component.SideEffect.metadata:type_name -> akkaserverless.component.Metadata
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_component_proto_init() }
func file_component_proto_init() {
	if File_component_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_component_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_component_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataEntry); i {
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
		file_component_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
		file_component_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Forward); i {
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
		file_component_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientAction); i {
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
		file_component_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SideEffect); i {
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
		file_component_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamCancelled); i {
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
		file_component_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Failure); i {
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
	file_component_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*MetadataEntry_StringValue)(nil),
		(*MetadataEntry_BytesValue)(nil),
	}
	file_component_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*ClientAction_Reply)(nil),
		(*ClientAction_Forward)(nil),
		(*ClientAction_Failure)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_component_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_component_proto_goTypes,
		DependencyIndexes: file_component_proto_depIdxs,
		MessageInfos:      file_component_proto_msgTypes,
	}.Build()
	File_component_proto = out.File
	file_component_proto_rawDesc = nil
	file_component_proto_goTypes = nil
	file_component_proto_depIdxs = nil
}
