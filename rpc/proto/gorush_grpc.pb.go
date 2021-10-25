// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// GorushClient is the client API for Gorush service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GorushClient interface {
	Send(ctx context.Context, in *NotificationRequest, opts ...grpc.CallOption) (*NotificationReply, error)
}

type gorushClient struct {
	cc grpc.ClientConnInterface
}

func NewGorushClient(cc grpc.ClientConnInterface) GorushClient {
	return &gorushClient{cc}
}

func (c *gorushClient) Send(ctx context.Context, in *NotificationRequest, opts ...grpc.CallOption) (*NotificationReply, error) {
	out := new(NotificationReply)
	err := c.cc.Invoke(ctx, "/proto.Gorush/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GorushServer is the server API for Gorush service.
// All implementations should embed UnimplementedGorushServer
// for forward compatibility
type GorushServer interface {
	Send(context.Context, *NotificationRequest) (*NotificationReply, error)
}

// UnimplementedGorushServer should be embedded to have forward compatible implementations.
type UnimplementedGorushServer struct {
}

func (UnimplementedGorushServer) Send(context.Context, *NotificationRequest) (*NotificationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}

// UnsafeGorushServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GorushServer will
// result in compilation errors.
type UnsafeGorushServer interface {
	mustEmbedUnimplementedGorushServer()
}

func RegisterGorushServer(s grpc.ServiceRegistrar, srv GorushServer) {
	s.RegisterService(&Gorush_ServiceDesc, srv)
}

func _Gorush_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GorushServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Gorush/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GorushServer).Send(ctx, req.(*NotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gorush_ServiceDesc is the grpc.ServiceDesc for Gorush service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gorush_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Gorush",
	HandlerType: (*GorushServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Gorush_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gorush.proto",
}

// HealthClient is the client API for Health service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HealthClient interface {
	Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type healthClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthClient(cc grpc.ClientConnInterface) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/proto.Health/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthServer is the server API for Health service.
// All implementations should embed UnimplementedHealthServer
// for forward compatibility
type HealthServer interface {
	Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
}

// UnimplementedHealthServer should be embedded to have forward compatible implementations.
type UnimplementedHealthServer struct {
}

func (UnimplementedHealthServer) Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}

// UnsafeHealthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthServer will
// result in compilation errors.
type UnsafeHealthServer interface {
	mustEmbedUnimplementedHealthServer()
}

func RegisterHealthServer(s grpc.ServiceRegistrar, srv HealthServer) {
	s.RegisterService(&Health_ServiceDesc, srv)
}

func _Health_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Health/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Check(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Health_ServiceDesc is the grpc.ServiceDesc for Health service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Health_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Health_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gorush.proto",
}
