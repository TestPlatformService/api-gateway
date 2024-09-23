// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: notification.proto

package notification

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Notifications_CreateNotification_FullMethodName           = "/notification.Notifications/CreateNotification"
	Notifications_GetAllNotifications_FullMethodName          = "/notification.Notifications/GetAllNotifications"
	Notifications_GetAndMarkNotificationAsRead_FullMethodName = "/notification.Notifications/GetAndMarkNotificationAsRead"
)

// NotificationsClient is the client API for Notifications service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationsClient interface {
	CreateNotification(ctx context.Context, in *CreateNotificationsReq, opts ...grpc.CallOption) (*CreateNotificationsRes, error)
	GetAllNotifications(ctx context.Context, in *GetNotificationsReq, opts ...grpc.CallOption) (*GetNotificationsResponse, error)
	GetAndMarkNotificationAsRead(ctx context.Context, in *GetAndMarkNotificationAsReadReq, opts ...grpc.CallOption) (*GetAndMarkNotificationAsReadRes, error)
}

type notificationsClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationsClient(cc grpc.ClientConnInterface) NotificationsClient {
	return &notificationsClient{cc}
}

func (c *notificationsClient) CreateNotification(ctx context.Context, in *CreateNotificationsReq, opts ...grpc.CallOption) (*CreateNotificationsRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateNotificationsRes)
	err := c.cc.Invoke(ctx, Notifications_CreateNotification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsClient) GetAllNotifications(ctx context.Context, in *GetNotificationsReq, opts ...grpc.CallOption) (*GetNotificationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNotificationsResponse)
	err := c.cc.Invoke(ctx, Notifications_GetAllNotifications_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsClient) GetAndMarkNotificationAsRead(ctx context.Context, in *GetAndMarkNotificationAsReadReq, opts ...grpc.CallOption) (*GetAndMarkNotificationAsReadRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAndMarkNotificationAsReadRes)
	err := c.cc.Invoke(ctx, Notifications_GetAndMarkNotificationAsRead_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationsServer is the server API for Notifications service.
// All implementations must embed UnimplementedNotificationsServer
// for forward compatibility
type NotificationsServer interface {
	CreateNotification(context.Context, *CreateNotificationsReq) (*CreateNotificationsRes, error)
	GetAllNotifications(context.Context, *GetNotificationsReq) (*GetNotificationsResponse, error)
	GetAndMarkNotificationAsRead(context.Context, *GetAndMarkNotificationAsReadReq) (*GetAndMarkNotificationAsReadRes, error)
	mustEmbedUnimplementedNotificationsServer()
}

// UnimplementedNotificationsServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationsServer struct {
}

func (UnimplementedNotificationsServer) CreateNotification(context.Context, *CreateNotificationsReq) (*CreateNotificationsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNotification not implemented")
}
func (UnimplementedNotificationsServer) GetAllNotifications(context.Context, *GetNotificationsReq) (*GetNotificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllNotifications not implemented")
}
func (UnimplementedNotificationsServer) GetAndMarkNotificationAsRead(context.Context, *GetAndMarkNotificationAsReadReq) (*GetAndMarkNotificationAsReadRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAndMarkNotificationAsRead not implemented")
}
func (UnimplementedNotificationsServer) mustEmbedUnimplementedNotificationsServer() {}

// UnsafeNotificationsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationsServer will
// result in compilation errors.
type UnsafeNotificationsServer interface {
	mustEmbedUnimplementedNotificationsServer()
}

func RegisterNotificationsServer(s grpc.ServiceRegistrar, srv NotificationsServer) {
	s.RegisterService(&Notifications_ServiceDesc, srv)
}

func _Notifications_CreateNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNotificationsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).CreateNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notifications_CreateNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).CreateNotification(ctx, req.(*CreateNotificationsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifications_GetAllNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotificationsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).GetAllNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notifications_GetAllNotifications_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).GetAllNotifications(ctx, req.(*GetNotificationsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifications_GetAndMarkNotificationAsRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAndMarkNotificationAsReadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).GetAndMarkNotificationAsRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notifications_GetAndMarkNotificationAsRead_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).GetAndMarkNotificationAsRead(ctx, req.(*GetAndMarkNotificationAsReadReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Notifications_ServiceDesc is the grpc.ServiceDesc for Notifications service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Notifications_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notification.Notifications",
	HandlerType: (*NotificationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNotification",
			Handler:    _Notifications_CreateNotification_Handler,
		},
		{
			MethodName: "GetAllNotifications",
			Handler:    _Notifications_GetAllNotifications_Handler,
		},
		{
			MethodName: "GetAndMarkNotificationAsRead",
			Handler:    _Notifications_GetAndMarkNotificationAsRead_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notification.proto",
}
