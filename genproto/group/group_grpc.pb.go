// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: group.proto

package group

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
	GroupService_CreateGroup_FullMethodName            = "/group.GroupService/CreateGroup"
	GroupService_UpdateGroup_FullMethodName            = "/group.GroupService/UpdateGroup"
	GroupService_DeleteGroup_FullMethodName            = "/group.GroupService/DeleteGroup"
	GroupService_GetGroupById_FullMethodName           = "/group.GroupService/GetGroupById"
	GroupService_GetAllGroups_FullMethodName           = "/group.GroupService/GetAllGroups"
	GroupService_AddStudentToGroup_FullMethodName      = "/group.GroupService/AddStudentToGroup"
	GroupService_DeleteStudentFromGroup_FullMethodName = "/group.GroupService/DeleteStudentFromGroup"
	GroupService_AddTeacherToGroup_FullMethodName      = "/group.GroupService/AddTeacherToGroup"
	GroupService_DeleteTeacherFromGroup_FullMethodName = "/group.GroupService/DeleteTeacherFromGroup"
	GroupService_GetStudentGroups_FullMethodName       = "/group.GroupService/GetStudentGroups"
	GroupService_GetTeacherGroups_FullMethodName       = "/group.GroupService/GetTeacherGroups"
	GroupService_GetGroupStudents_FullMethodName       = "/group.GroupService/GetGroupStudents"
)

// GroupServiceClient is the client API for GroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupServiceClient interface {
	CreateGroup(ctx context.Context, in *CreateGroupReq, opts ...grpc.CallOption) (*CreateGroupResp, error)
	UpdateGroup(ctx context.Context, in *UpdateGroupReq, opts ...grpc.CallOption) (*UpdateGroupResp, error)
	DeleteGroup(ctx context.Context, in *GroupId, opts ...grpc.CallOption) (*DeleteResp, error)
	GetGroupById(ctx context.Context, in *GroupId, opts ...grpc.CallOption) (*Group, error)
	GetAllGroups(ctx context.Context, in *GetAllGroupsReq, opts ...grpc.CallOption) (*GetAllGroupsResp, error)
	AddStudentToGroup(ctx context.Context, in *AddStudentReq, opts ...grpc.CallOption) (*AddStudentResp, error)
	DeleteStudentFromGroup(ctx context.Context, in *DeleteStudentReq, opts ...grpc.CallOption) (*DeleteResp, error)
	AddTeacherToGroup(ctx context.Context, in *AddTeacherReq, opts ...grpc.CallOption) (*AddTeacherResp, error)
	DeleteTeacherFromGroup(ctx context.Context, in *DeleteTeacherReq, opts ...grpc.CallOption) (*DeleteResp, error)
	GetStudentGroups(ctx context.Context, in *StudentId, opts ...grpc.CallOption) (*StudentGroups, error)
	GetTeacherGroups(ctx context.Context, in *TeacherId, opts ...grpc.CallOption) (*TeacherGroups, error)
	GetGroupStudents(ctx context.Context, in *GroupId, opts ...grpc.CallOption) (*GroupStudents, error)
}

type groupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupServiceClient(cc grpc.ClientConnInterface) GroupServiceClient {
	return &groupServiceClient{cc}
}

func (c *groupServiceClient) CreateGroup(ctx context.Context, in *CreateGroupReq, opts ...grpc.CallOption) (*CreateGroupResp, error) {
	out := new(CreateGroupResp)
	err := c.cc.Invoke(ctx, "/group.GroupService/CreateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) UpdateGroup(ctx context.Context, in *UpdateGroupReq, opts ...grpc.CallOption) (*UpdateGroupResp, error) {
	out := new(UpdateGroupResp)
	err := c.cc.Invoke(ctx, "/group.GroupService/UpdateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) DeleteGroup(ctx context.Context, in *GroupId, opts ...grpc.CallOption) (*DeleteResp, error) {
	out := new(DeleteResp)
	err := c.cc.Invoke(ctx, "/group.GroupService/DeleteGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) GetGroupById(ctx context.Context, in *GroupId, opts ...grpc.CallOption) (*Group, error) {
	out := new(Group)
	err := c.cc.Invoke(ctx, "/group.GroupService/GetGroupById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) GetAllGroups(ctx context.Context, in *GetAllGroupsReq, opts ...grpc.CallOption) (*GetAllGroupsResp, error) {
	out := new(GetAllGroupsResp)
	err := c.cc.Invoke(ctx, "/group.GroupService/GetAllGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) AddStudentToGroup(ctx context.Context, in *AddStudentReq, opts ...grpc.CallOption) (*AddStudentResp, error) {
	out := new(AddStudentResp)
	err := c.cc.Invoke(ctx, "/group.GroupService/AddStudentToGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) DeleteStudentFromGroup(ctx context.Context, in *DeleteStudentReq, opts ...grpc.CallOption) (*DeleteResp, error) {
	out := new(DeleteResp)
	err := c.cc.Invoke(ctx, "/group.GroupService/DeleteStudentFromGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) AddTeacherToGroup(ctx context.Context, in *AddTeacherReq, opts ...grpc.CallOption) (*AddTeacherResp, error) {
	out := new(AddTeacherResp)
	err := c.cc.Invoke(ctx, "/group.GroupService/AddTeacherToGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) DeleteTeacherFromGroup(ctx context.Context, in *DeleteTeacherReq, opts ...grpc.CallOption) (*DeleteResp, error) {
	out := new(DeleteResp)
	err := c.cc.Invoke(ctx, "/group.GroupService/DeleteTeacherFromGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) GetStudentGroups(ctx context.Context, in *StudentId, opts ...grpc.CallOption) (*StudentGroups, error) {
	out := new(StudentGroups)
	err := c.cc.Invoke(ctx, "/group.GroupService/GetStudentGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) GetTeacherGroups(ctx context.Context, in *TeacherId, opts ...grpc.CallOption) (*TeacherGroups, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TeacherGroups)
	err := c.cc.Invoke(ctx, GroupService_GetTeacherGroups_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) GetGroupStudents(ctx context.Context, in *GroupId, opts ...grpc.CallOption) (*GroupStudents, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GroupStudents)
	err := c.cc.Invoke(ctx, GroupService_GetGroupStudents_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupServiceServer is the server API for GroupService service.
// All implementations must embed UnimplementedGroupServiceServer
// for forward compatibility
type GroupServiceServer interface {
	CreateGroup(context.Context, *CreateGroupReq) (*CreateGroupResp, error)
	UpdateGroup(context.Context, *UpdateGroupReq) (*UpdateGroupResp, error)
	DeleteGroup(context.Context, *GroupId) (*DeleteResp, error)
	GetGroupById(context.Context, *GroupId) (*Group, error)
	GetAllGroups(context.Context, *GetAllGroupsReq) (*GetAllGroupsResp, error)
	AddStudentToGroup(context.Context, *AddStudentReq) (*AddStudentResp, error)
	DeleteStudentFromGroup(context.Context, *DeleteStudentReq) (*DeleteResp, error)
	AddTeacherToGroup(context.Context, *AddTeacherReq) (*AddTeacherResp, error)
	DeleteTeacherFromGroup(context.Context, *DeleteTeacherReq) (*DeleteResp, error)
	GetStudentGroups(context.Context, *StudentId) (*StudentGroups, error)
	GetTeacherGroups(context.Context, *TeacherId) (*TeacherGroups, error)
	GetGroupStudents(context.Context, *GroupId) (*GroupStudents, error)
	mustEmbedUnimplementedGroupServiceServer()
}

// UnimplementedGroupServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGroupServiceServer struct {
}

func (UnimplementedGroupServiceServer) CreateGroup(context.Context, *CreateGroupReq) (*CreateGroupResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedGroupServiceServer) UpdateGroup(context.Context, *UpdateGroupReq) (*UpdateGroupResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroup not implemented")
}
func (UnimplementedGroupServiceServer) DeleteGroup(context.Context, *GroupId) (*DeleteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroup not implemented")
}
func (UnimplementedGroupServiceServer) GetGroupById(context.Context, *GroupId) (*Group, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupById not implemented")
}
func (UnimplementedGroupServiceServer) GetAllGroups(context.Context, *GetAllGroupsReq) (*GetAllGroupsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllGroups not implemented")
}
func (UnimplementedGroupServiceServer) AddStudentToGroup(context.Context, *AddStudentReq) (*AddStudentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddStudentToGroup not implemented")
}
func (UnimplementedGroupServiceServer) DeleteStudentFromGroup(context.Context, *DeleteStudentReq) (*DeleteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStudentFromGroup not implemented")
}
func (UnimplementedGroupServiceServer) AddTeacherToGroup(context.Context, *AddTeacherReq) (*AddTeacherResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTeacherToGroup not implemented")
}
func (UnimplementedGroupServiceServer) DeleteTeacherFromGroup(context.Context, *DeleteTeacherReq) (*DeleteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTeacherFromGroup not implemented")
}
func (UnimplementedGroupServiceServer) GetStudentGroups(context.Context, *StudentId) (*StudentGroups, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudentGroups not implemented")
}
func (UnimplementedGroupServiceServer) GetTeacherGroups(context.Context, *TeacherId) (*TeacherGroups, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeacherGroups not implemented")
}
func (UnimplementedGroupServiceServer) GetGroupStudents(context.Context, *GroupId) (*GroupStudents, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupStudents not implemented")
}
func (UnimplementedGroupServiceServer) mustEmbedUnimplementedGroupServiceServer() {}

// UnsafeGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupServiceServer will
// result in compilation errors.
type UnsafeGroupServiceServer interface {
	mustEmbedUnimplementedGroupServiceServer()
}

func RegisterGroupServiceServer(s grpc.ServiceRegistrar, srv GroupServiceServer) {
	s.RegisterService(&GroupService_ServiceDesc, srv)
}

func _GroupService_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/group.GroupService/CreateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).CreateGroup(ctx, req.(*CreateGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_UpdateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).UpdateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/group.GroupService/UpdateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).UpdateGroup(ctx, req.(*UpdateGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_DeleteGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).DeleteGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/group.GroupService/DeleteGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).DeleteGroup(ctx, req.(*GroupId))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_GetGroupById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).GetGroupById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/group.GroupService/GetGroupById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).GetGroupById(ctx, req.(*GroupId))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_GetAllGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllGroupsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).GetAllGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/group.GroupService/GetAllGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).GetAllGroups(ctx, req.(*GetAllGroupsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_AddStudentToGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddStudentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).AddStudentToGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/group.GroupService/AddStudentToGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).AddStudentToGroup(ctx, req.(*AddStudentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_DeleteStudentFromGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStudentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).DeleteStudentFromGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/group.GroupService/DeleteStudentFromGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).DeleteStudentFromGroup(ctx, req.(*DeleteStudentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_AddTeacherToGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTeacherReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).AddTeacherToGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/group.GroupService/AddTeacherToGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).AddTeacherToGroup(ctx, req.(*AddTeacherReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_DeleteTeacherFromGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTeacherReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).DeleteTeacherFromGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/group.GroupService/DeleteTeacherFromGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).DeleteTeacherFromGroup(ctx, req.(*DeleteTeacherReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_GetStudentGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StudentId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).GetStudentGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/group.GroupService/GetStudentGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).GetStudentGroups(ctx, req.(*StudentId))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_GetTeacherGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeacherId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).GetTeacherGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/group.GroupService/GetTeacherGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).GetTeacherGroups(ctx, req.(*TeacherId))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_GetGroupStudents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).GetGroupStudents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_GetGroupStudents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).GetGroupStudents(ctx, req.(*GroupId))
	}
	return interceptor(ctx, in, info, handler)
}

// GroupService_ServiceDesc is the grpc.ServiceDesc for GroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "group.GroupService",
	HandlerType: (*GroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGroup",
			Handler:    _GroupService_CreateGroup_Handler,
		},
		{
			MethodName: "UpdateGroup",
			Handler:    _GroupService_UpdateGroup_Handler,
		},
		{
			MethodName: "DeleteGroup",
			Handler:    _GroupService_DeleteGroup_Handler,
		},
		{
			MethodName: "GetGroupById",
			Handler:    _GroupService_GetGroupById_Handler,
		},
		{
			MethodName: "GetAllGroups",
			Handler:    _GroupService_GetAllGroups_Handler,
		},
		{
			MethodName: "AddStudentToGroup",
			Handler:    _GroupService_AddStudentToGroup_Handler,
		},
		{
			MethodName: "DeleteStudentFromGroup",
			Handler:    _GroupService_DeleteStudentFromGroup_Handler,
		},
		{
			MethodName: "AddTeacherToGroup",
			Handler:    _GroupService_AddTeacherToGroup_Handler,
		},
		{
			MethodName: "DeleteTeacherFromGroup",
			Handler:    _GroupService_DeleteTeacherFromGroup_Handler,
		},
		{
			MethodName: "GetStudentGroups",
			Handler:    _GroupService_GetStudentGroups_Handler,
		},
		{
			MethodName: "GetTeacherGroups",
			Handler:    _GroupService_GetTeacherGroups_Handler,
		},
		{
			MethodName: "GetGroupStudents",
			Handler:    _GroupService_GetGroupStudents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "group.proto",
}
