// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: proto/activity.proto

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

const (
	ActivityService_AddActivity_FullMethodName    = "/Server.ActivityService/AddActivity"
	ActivityService_GetActivity_FullMethodName    = "/Server.ActivityService/GetActivity"
	ActivityService_GetAllActivity_FullMethodName = "/Server.ActivityService/GetAllActivity"
	ActivityService_UpdateActivity_FullMethodName = "/Server.ActivityService/UpdateActivity"
	ActivityService_DeleteActivity_FullMethodName = "/Server.ActivityService/DeleteActivity"
)

// ActivityServiceClient is the client API for ActivityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActivityServiceClient interface {
	AddActivity(ctx context.Context, in *AddActivityReq, opts ...grpc.CallOption) (*AddActivityRes, error)
	GetActivity(ctx context.Context, in *GetActivityReq, opts ...grpc.CallOption) (*GetActivityResp, error)
	GetAllActivity(ctx context.Context, in *GetAllReq, opts ...grpc.CallOption) (*GetAllResp, error)
	UpdateActivity(ctx context.Context, in *UpdateActivityReq, opts ...grpc.CallOption) (*UpdateActivityResp, error)
	DeleteActivity(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteResp, error)
}

type activityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewActivityServiceClient(cc grpc.ClientConnInterface) ActivityServiceClient {
	return &activityServiceClient{cc}
}

func (c *activityServiceClient) AddActivity(ctx context.Context, in *AddActivityReq, opts ...grpc.CallOption) (*AddActivityRes, error) {
	out := new(AddActivityRes)
	err := c.cc.Invoke(ctx, ActivityService_AddActivity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityServiceClient) GetActivity(ctx context.Context, in *GetActivityReq, opts ...grpc.CallOption) (*GetActivityResp, error) {
	out := new(GetActivityResp)
	err := c.cc.Invoke(ctx, ActivityService_GetActivity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityServiceClient) GetAllActivity(ctx context.Context, in *GetAllReq, opts ...grpc.CallOption) (*GetAllResp, error) {
	out := new(GetAllResp)
	err := c.cc.Invoke(ctx, ActivityService_GetAllActivity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityServiceClient) UpdateActivity(ctx context.Context, in *UpdateActivityReq, opts ...grpc.CallOption) (*UpdateActivityResp, error) {
	out := new(UpdateActivityResp)
	err := c.cc.Invoke(ctx, ActivityService_UpdateActivity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityServiceClient) DeleteActivity(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteResp, error) {
	out := new(DeleteResp)
	err := c.cc.Invoke(ctx, ActivityService_DeleteActivity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActivityServiceServer is the server API for ActivityService service.
// All implementations must embed UnimplementedActivityServiceServer
// for forward compatibility
type ActivityServiceServer interface {
	AddActivity(context.Context, *AddActivityReq) (*AddActivityRes, error)
	GetActivity(context.Context, *GetActivityReq) (*GetActivityResp, error)
	GetAllActivity(context.Context, *GetAllReq) (*GetAllResp, error)
	UpdateActivity(context.Context, *UpdateActivityReq) (*UpdateActivityResp, error)
	DeleteActivity(context.Context, *DeleteReq) (*DeleteResp, error)
	mustEmbedUnimplementedActivityServiceServer()
}

// UnimplementedActivityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedActivityServiceServer struct {
}

func (UnimplementedActivityServiceServer) AddActivity(context.Context, *AddActivityReq) (*AddActivityRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddActivity not implemented")
}
func (UnimplementedActivityServiceServer) GetActivity(context.Context, *GetActivityReq) (*GetActivityResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActivity not implemented")
}
func (UnimplementedActivityServiceServer) GetAllActivity(context.Context, *GetAllReq) (*GetAllResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllActivity not implemented")
}
func (UnimplementedActivityServiceServer) UpdateActivity(context.Context, *UpdateActivityReq) (*UpdateActivityResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateActivity not implemented")
}
func (UnimplementedActivityServiceServer) DeleteActivity(context.Context, *DeleteReq) (*DeleteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteActivity not implemented")
}
func (UnimplementedActivityServiceServer) mustEmbedUnimplementedActivityServiceServer() {}

// UnsafeActivityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActivityServiceServer will
// result in compilation errors.
type UnsafeActivityServiceServer interface {
	mustEmbedUnimplementedActivityServiceServer()
}

func RegisterActivityServiceServer(s grpc.ServiceRegistrar, srv ActivityServiceServer) {
	s.RegisterService(&ActivityService_ServiceDesc, srv)
}

func _ActivityService_AddActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddActivityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServiceServer).AddActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActivityService_AddActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServiceServer).AddActivity(ctx, req.(*AddActivityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityService_GetActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActivityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServiceServer).GetActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActivityService_GetActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServiceServer).GetActivity(ctx, req.(*GetActivityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityService_GetAllActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServiceServer).GetAllActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActivityService_GetAllActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServiceServer).GetAllActivity(ctx, req.(*GetAllReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityService_UpdateActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateActivityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServiceServer).UpdateActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActivityService_UpdateActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServiceServer).UpdateActivity(ctx, req.(*UpdateActivityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityService_DeleteActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServiceServer).DeleteActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActivityService_DeleteActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServiceServer).DeleteActivity(ctx, req.(*DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ActivityService_ServiceDesc is the grpc.ServiceDesc for ActivityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ActivityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Server.ActivityService",
	HandlerType: (*ActivityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddActivity",
			Handler:    _ActivityService_AddActivity_Handler,
		},
		{
			MethodName: "GetActivity",
			Handler:    _ActivityService_GetActivity_Handler,
		},
		{
			MethodName: "GetAllActivity",
			Handler:    _ActivityService_GetAllActivity_Handler,
		},
		{
			MethodName: "UpdateActivity",
			Handler:    _ActivityService_UpdateActivity_Handler,
		},
		{
			MethodName: "DeleteActivity",
			Handler:    _ActivityService_DeleteActivity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/activity.proto",
}
