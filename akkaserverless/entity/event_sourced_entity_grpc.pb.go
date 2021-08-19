// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package entity

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EventSourcedEntitiesClient is the client API for EventSourcedEntities service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventSourcedEntitiesClient interface {
	// The stream. One stream will be established per active entity. Once
	// established, the first message sent will be Init, which contains the entity
	// ID, and, if the entity has previously persisted a snapshot, it will contain
	// that snapshot. It will then send zero to many event messages, one for each
	// event previously persisted. The entity is expected to apply these to its
	// state in a deterministic fashion. Once all the events are sent, one to many
	// commands are sent, with new commands being sent as new requests for the
	// entity come in. The entity is expected to reply to each command with
	// exactly one reply message. The entity should reply in order, and any events
	// that the entity requests to be persisted the entity should handle itself,
	// applying them to its own state, as if they had arrived as events when the
	// event stream was being replayed on load.
	Handle(ctx context.Context, opts ...grpc.CallOption) (EventSourcedEntities_HandleClient, error)
}

type eventSourcedEntitiesClient struct {
	cc grpc.ClientConnInterface
}

func NewEventSourcedEntitiesClient(cc grpc.ClientConnInterface) EventSourcedEntitiesClient {
	return &eventSourcedEntitiesClient{cc}
}

func (c *eventSourcedEntitiesClient) Handle(ctx context.Context, opts ...grpc.CallOption) (EventSourcedEntities_HandleClient, error) {
	stream, err := c.cc.NewStream(ctx, &EventSourcedEntities_ServiceDesc.Streams[0], "/akkaserverless.component.eventsourcedentity.EventSourcedEntities/Handle", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventSourcedEntitiesHandleClient{stream}
	return x, nil
}

type EventSourcedEntities_HandleClient interface {
	Send(*EventSourcedStreamIn) error
	Recv() (*EventSourcedStreamOut, error)
	grpc.ClientStream
}

type eventSourcedEntitiesHandleClient struct {
	grpc.ClientStream
}

func (x *eventSourcedEntitiesHandleClient) Send(m *EventSourcedStreamIn) error {
	return x.ClientStream.SendMsg(m)
}

func (x *eventSourcedEntitiesHandleClient) Recv() (*EventSourcedStreamOut, error) {
	m := new(EventSourcedStreamOut)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventSourcedEntitiesServer is the server API for EventSourcedEntities service.
// All implementations must embed UnimplementedEventSourcedEntitiesServer
// for forward compatibility
type EventSourcedEntitiesServer interface {
	// The stream. One stream will be established per active entity. Once
	// established, the first message sent will be Init, which contains the entity
	// ID, and, if the entity has previously persisted a snapshot, it will contain
	// that snapshot. It will then send zero to many event messages, one for each
	// event previously persisted. The entity is expected to apply these to its
	// state in a deterministic fashion. Once all the events are sent, one to many
	// commands are sent, with new commands being sent as new requests for the
	// entity come in. The entity is expected to reply to each command with
	// exactly one reply message. The entity should reply in order, and any events
	// that the entity requests to be persisted the entity should handle itself,
	// applying them to its own state, as if they had arrived as events when the
	// event stream was being replayed on load.
	Handle(EventSourcedEntities_HandleServer) error
	mustEmbedUnimplementedEventSourcedEntitiesServer()
}

// UnimplementedEventSourcedEntitiesServer must be embedded to have forward compatible implementations.
type UnimplementedEventSourcedEntitiesServer struct {
}

func (UnimplementedEventSourcedEntitiesServer) Handle(EventSourcedEntities_HandleServer) error {
	return status.Errorf(codes.Unimplemented, "method Handle not implemented")
}
func (UnimplementedEventSourcedEntitiesServer) mustEmbedUnimplementedEventSourcedEntitiesServer() {}

// UnsafeEventSourcedEntitiesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventSourcedEntitiesServer will
// result in compilation errors.
type UnsafeEventSourcedEntitiesServer interface {
	mustEmbedUnimplementedEventSourcedEntitiesServer()
}

func RegisterEventSourcedEntitiesServer(s grpc.ServiceRegistrar, srv EventSourcedEntitiesServer) {
	s.RegisterService(&EventSourcedEntities_ServiceDesc, srv)
}

func _EventSourcedEntities_Handle_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EventSourcedEntitiesServer).Handle(&eventSourcedEntitiesHandleServer{stream})
}

type EventSourcedEntities_HandleServer interface {
	Send(*EventSourcedStreamOut) error
	Recv() (*EventSourcedStreamIn, error)
	grpc.ServerStream
}

type eventSourcedEntitiesHandleServer struct {
	grpc.ServerStream
}

func (x *eventSourcedEntitiesHandleServer) Send(m *EventSourcedStreamOut) error {
	return x.ServerStream.SendMsg(m)
}

func (x *eventSourcedEntitiesHandleServer) Recv() (*EventSourcedStreamIn, error) {
	m := new(EventSourcedStreamIn)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventSourcedEntities_ServiceDesc is the grpc.ServiceDesc for EventSourcedEntities service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventSourcedEntities_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "akkaserverless.component.eventsourcedentity.EventSourcedEntities",
	HandlerType: (*EventSourcedEntitiesServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Handle",
			Handler:       _EventSourcedEntities_Handle_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "event_sourced_entity.proto",
}
