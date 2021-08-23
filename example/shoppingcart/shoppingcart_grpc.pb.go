// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package shoppingcart

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ShoppingCartServiceClient is the client API for ShoppingCartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShoppingCartServiceClient interface {
	AddItem(ctx context.Context, in *AddLineItem, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RemoveItem(ctx context.Context, in *RemoveLineItem, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetCart(ctx context.Context, in *GetShoppingCart, opts ...grpc.CallOption) (*Cart, error)
}

type shoppingCartServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShoppingCartServiceClient(cc grpc.ClientConnInterface) ShoppingCartServiceClient {
	return &shoppingCartServiceClient{cc}
}

func (c *shoppingCartServiceClient) AddItem(ctx context.Context, in *AddLineItem, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/com.example.shoppingcart.ShoppingCartService/AddItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shoppingCartServiceClient) RemoveItem(ctx context.Context, in *RemoveLineItem, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/com.example.shoppingcart.ShoppingCartService/RemoveItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shoppingCartServiceClient) GetCart(ctx context.Context, in *GetShoppingCart, opts ...grpc.CallOption) (*Cart, error) {
	out := new(Cart)
	err := c.cc.Invoke(ctx, "/com.example.shoppingcart.ShoppingCartService/GetCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShoppingCartServiceServer is the server API for ShoppingCartService service.
// All implementations must embed UnimplementedShoppingCartServiceServer
// for forward compatibility
type ShoppingCartServiceServer interface {
	AddItem(context.Context, *AddLineItem) (*emptypb.Empty, error)
	RemoveItem(context.Context, *RemoveLineItem) (*emptypb.Empty, error)
	GetCart(context.Context, *GetShoppingCart) (*Cart, error)
	mustEmbedUnimplementedShoppingCartServiceServer()
}

// UnimplementedShoppingCartServiceServer must be embedded to have forward compatible implementations.
type UnimplementedShoppingCartServiceServer struct {
}

func (UnimplementedShoppingCartServiceServer) AddItem(context.Context, *AddLineItem) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddItem not implemented")
}
func (UnimplementedShoppingCartServiceServer) RemoveItem(context.Context, *RemoveLineItem) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveItem not implemented")
}
func (UnimplementedShoppingCartServiceServer) GetCart(context.Context, *GetShoppingCart) (*Cart, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCart not implemented")
}
func (UnimplementedShoppingCartServiceServer) mustEmbedUnimplementedShoppingCartServiceServer() {}

// UnsafeShoppingCartServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShoppingCartServiceServer will
// result in compilation errors.
type UnsafeShoppingCartServiceServer interface {
	mustEmbedUnimplementedShoppingCartServiceServer()
}

func RegisterShoppingCartServiceServer(s grpc.ServiceRegistrar, srv ShoppingCartServiceServer) {
	s.RegisterService(&ShoppingCartService_ServiceDesc, srv)
}

func _ShoppingCartService_AddItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddLineItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingCartServiceServer).AddItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.example.shoppingcart.ShoppingCartService/AddItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingCartServiceServer).AddItem(ctx, req.(*AddLineItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShoppingCartService_RemoveItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveLineItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingCartServiceServer).RemoveItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.example.shoppingcart.ShoppingCartService/RemoveItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingCartServiceServer).RemoveItem(ctx, req.(*RemoveLineItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShoppingCartService_GetCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShoppingCart)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingCartServiceServer).GetCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.example.shoppingcart.ShoppingCartService/GetCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingCartServiceServer).GetCart(ctx, req.(*GetShoppingCart))
	}
	return interceptor(ctx, in, info, handler)
}

// ShoppingCartService_ServiceDesc is the grpc.ServiceDesc for ShoppingCartService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShoppingCartService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.example.shoppingcart.ShoppingCartService",
	HandlerType: (*ShoppingCartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddItem",
			Handler:    _ShoppingCartService_AddItem_Handler,
		},
		{
			MethodName: "RemoveItem",
			Handler:    _ShoppingCartService_RemoveItem_Handler,
		},
		{
			MethodName: "GetCart",
			Handler:    _ShoppingCartService_GetCart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shoppingcart.proto",
}
