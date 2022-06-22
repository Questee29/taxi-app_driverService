// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: proto/order.proto

package protob

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

// OrderGrpcClient is the client API for OrderGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderGrpcClient interface {
	OrderTaxi(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResponse, error)
	FindDriver(ctx context.Context, in *FindDriverRequest, opts ...grpc.CallOption) (*FindDriverResponse, error)
}

type orderGrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderGrpcClient(cc grpc.ClientConnInterface) OrderGrpcClient {
	return &orderGrpcClient{cc}
}

func (c *orderGrpcClient) OrderTaxi(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResponse, error) {
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, "/protob.OrderGrpc/OrderTaxi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderGrpcClient) FindDriver(ctx context.Context, in *FindDriverRequest, opts ...grpc.CallOption) (*FindDriverResponse, error) {
	out := new(FindDriverResponse)
	err := c.cc.Invoke(ctx, "/protob.OrderGrpc/FindDriver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderGrpcServer is the server API for OrderGrpc service.
// All implementations must embed UnimplementedOrderGrpcServer
// for forward compatibility
type OrderGrpcServer interface {
	OrderTaxi(context.Context, *OrderRequest) (*OrderResponse, error)
	FindDriver(context.Context, *FindDriverRequest) (*FindDriverResponse, error)
	mustEmbedUnimplementedOrderGrpcServer()
}

// UnimplementedOrderGrpcServer must be embedded to have forward compatible implementations.
type UnimplementedOrderGrpcServer struct {
}

func (UnimplementedOrderGrpcServer) OrderTaxi(context.Context, *OrderRequest) (*OrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderTaxi not implemented")
}
func (UnimplementedOrderGrpcServer) FindDriver(context.Context, *FindDriverRequest) (*FindDriverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindDriver not implemented")
}
func (UnimplementedOrderGrpcServer) mustEmbedUnimplementedOrderGrpcServer() {}

// UnsafeOrderGrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderGrpcServer will
// result in compilation errors.
type UnsafeOrderGrpcServer interface {
	mustEmbedUnimplementedOrderGrpcServer()
}

func RegisterOrderGrpcServer(s grpc.ServiceRegistrar, srv OrderGrpcServer) {
	s.RegisterService(&OrderGrpc_ServiceDesc, srv)
}

func _OrderGrpc_OrderTaxi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderGrpcServer).OrderTaxi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protob.OrderGrpc/OrderTaxi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderGrpcServer).OrderTaxi(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderGrpc_FindDriver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindDriverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderGrpcServer).FindDriver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protob.OrderGrpc/FindDriver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderGrpcServer).FindDriver(ctx, req.(*FindDriverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderGrpc_ServiceDesc is the grpc.ServiceDesc for OrderGrpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderGrpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protob.OrderGrpc",
	HandlerType: (*OrderGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OrderTaxi",
			Handler:    _OrderGrpc_OrderTaxi_Handler,
		},
		{
			MethodName: "FindDriver",
			Handler:    _OrderGrpc_FindDriver_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/order.proto",
}
