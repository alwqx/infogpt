// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: admin/v1/admin.proto

package v1

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

const (
	Admin_HealthCheck_FullMethodName = "/admin.v1.Admin/HealthCheck"
	Admin_AppInfo_FullMethodName     = "/admin.v1.Admin/AppInfo"
	Admin_OpenaiChat_FullMethodName  = "/admin.v1.Admin/OpenaiChat"
)

// AdminClient is the client API for Admin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminClient interface {
	// Sends a greeting
	HealthCheck(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthReply, error)
	// Sends appinfo
	AppInfo(ctx context.Context, in *AppInfoRequest, opts ...grpc.CallOption) (*AppInfoReply, error)
	// proxy chat to openai
	OpenaiChat(ctx context.Context, in *OpenaiChatReuqest, opts ...grpc.CallOption) (*OpenaiChatReply, error)
}

type adminClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminClient(cc grpc.ClientConnInterface) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) HealthCheck(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthReply, error) {
	out := new(HealthReply)
	err := c.cc.Invoke(ctx, Admin_HealthCheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) AppInfo(ctx context.Context, in *AppInfoRequest, opts ...grpc.CallOption) (*AppInfoReply, error) {
	out := new(AppInfoReply)
	err := c.cc.Invoke(ctx, Admin_AppInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) OpenaiChat(ctx context.Context, in *OpenaiChatReuqest, opts ...grpc.CallOption) (*OpenaiChatReply, error) {
	out := new(OpenaiChatReply)
	err := c.cc.Invoke(ctx, Admin_OpenaiChat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServer is the server API for Admin service.
// All implementations must embed UnimplementedAdminServer
// for forward compatibility
type AdminServer interface {
	// Sends a greeting
	HealthCheck(context.Context, *HealthRequest) (*HealthReply, error)
	// Sends appinfo
	AppInfo(context.Context, *AppInfoRequest) (*AppInfoReply, error)
	// proxy chat to openai
	OpenaiChat(context.Context, *OpenaiChatReuqest) (*OpenaiChatReply, error)
	mustEmbedUnimplementedAdminServer()
}

// UnimplementedAdminServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServer struct {
}

func (UnimplementedAdminServer) HealthCheck(context.Context, *HealthRequest) (*HealthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedAdminServer) AppInfo(context.Context, *AppInfoRequest) (*AppInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppInfo not implemented")
}
func (UnimplementedAdminServer) OpenaiChat(context.Context, *OpenaiChatReuqest) (*OpenaiChatReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenaiChat not implemented")
}
func (UnimplementedAdminServer) mustEmbedUnimplementedAdminServer() {}

// UnsafeAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServer will
// result in compilation errors.
type UnsafeAdminServer interface {
	mustEmbedUnimplementedAdminServer()
}

func RegisterAdminServer(s grpc.ServiceRegistrar, srv AdminServer) {
	s.RegisterService(&Admin_ServiceDesc, srv)
}

func _Admin_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_HealthCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).HealthCheck(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_AppInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).AppInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_AppInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).AppInfo(ctx, req.(*AppInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_OpenaiChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenaiChatReuqest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).OpenaiChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_OpenaiChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).OpenaiChat(ctx, req.(*OpenaiChatReuqest))
	}
	return interceptor(ctx, in, info, handler)
}

// Admin_ServiceDesc is the grpc.ServiceDesc for Admin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Admin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.v1.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _Admin_HealthCheck_Handler,
		},
		{
			MethodName: "AppInfo",
			Handler:    _Admin_AppInfo_Handler,
		},
		{
			MethodName: "OpenaiChat",
			Handler:    _Admin_OpenaiChat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/v1/admin.proto",
}
