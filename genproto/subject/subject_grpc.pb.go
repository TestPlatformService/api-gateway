// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: subject.proto

package subject

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
	SubjectService_CreateSubject_FullMethodName  = "/subject.SubjectService/CreateSubject"
	SubjectService_GetSubject_FullMethodName     = "/subject.SubjectService/GetSubject"
	SubjectService_GetAllSubjects_FullMethodName = "/subject.SubjectService/GetAllSubjects"
	SubjectService_UpdateSubject_FullMethodName  = "/subject.SubjectService/UpdateSubject"
	SubjectService_DeleteSubject_FullMethodName  = "/subject.SubjectService/DeleteSubject"
)

// SubjectServiceClient is the client API for SubjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubjectServiceClient interface {
	CreateSubject(ctx context.Context, in *CreateSubjectRequest, opts ...grpc.CallOption) (*Void, error)
	GetSubject(ctx context.Context, in *GetSubjectRequest, opts ...grpc.CallOption) (*GetSubjectResponse, error)
	GetAllSubjects(ctx context.Context, in *GetAllSubjectsRequest, opts ...grpc.CallOption) (*GetAllSubjectsResponse, error)
	UpdateSubject(ctx context.Context, in *UpdateSubjectRequest, opts ...grpc.CallOption) (*Void, error)
	DeleteSubject(ctx context.Context, in *DeleteSubjectRequest, opts ...grpc.CallOption) (*Void, error)
}

type subjectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubjectServiceClient(cc grpc.ClientConnInterface) SubjectServiceClient {
	return &subjectServiceClient{cc}
}

func (c *subjectServiceClient) CreateSubject(ctx context.Context, in *CreateSubjectRequest, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, SubjectService_CreateSubject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subjectServiceClient) GetSubject(ctx context.Context, in *GetSubjectRequest, opts ...grpc.CallOption) (*GetSubjectResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSubjectResponse)
	err := c.cc.Invoke(ctx, SubjectService_GetSubject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subjectServiceClient) GetAllSubjects(ctx context.Context, in *GetAllSubjectsRequest, opts ...grpc.CallOption) (*GetAllSubjectsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllSubjectsResponse)
	err := c.cc.Invoke(ctx, SubjectService_GetAllSubjects_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subjectServiceClient) UpdateSubject(ctx context.Context, in *UpdateSubjectRequest, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, SubjectService_UpdateSubject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subjectServiceClient) DeleteSubject(ctx context.Context, in *DeleteSubjectRequest, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, SubjectService_DeleteSubject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubjectServiceServer is the server API for SubjectService service.
// All implementations must embed UnimplementedSubjectServiceServer
// for forward compatibility
type SubjectServiceServer interface {
	CreateSubject(context.Context, *CreateSubjectRequest) (*Void, error)
	GetSubject(context.Context, *GetSubjectRequest) (*GetSubjectResponse, error)
	GetAllSubjects(context.Context, *GetAllSubjectsRequest) (*GetAllSubjectsResponse, error)
	UpdateSubject(context.Context, *UpdateSubjectRequest) (*Void, error)
	DeleteSubject(context.Context, *DeleteSubjectRequest) (*Void, error)
	mustEmbedUnimplementedSubjectServiceServer()
}

// UnimplementedSubjectServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSubjectServiceServer struct {
}

func (UnimplementedSubjectServiceServer) CreateSubject(context.Context, *CreateSubjectRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubject not implemented")
}
func (UnimplementedSubjectServiceServer) GetSubject(context.Context, *GetSubjectRequest) (*GetSubjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubject not implemented")
}
func (UnimplementedSubjectServiceServer) GetAllSubjects(context.Context, *GetAllSubjectsRequest) (*GetAllSubjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllSubjects not implemented")
}
func (UnimplementedSubjectServiceServer) UpdateSubject(context.Context, *UpdateSubjectRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSubject not implemented")
}
func (UnimplementedSubjectServiceServer) DeleteSubject(context.Context, *DeleteSubjectRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSubject not implemented")
}
func (UnimplementedSubjectServiceServer) mustEmbedUnimplementedSubjectServiceServer() {}

// UnsafeSubjectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubjectServiceServer will
// result in compilation errors.
type UnsafeSubjectServiceServer interface {
	mustEmbedUnimplementedSubjectServiceServer()
}

func RegisterSubjectServiceServer(s grpc.ServiceRegistrar, srv SubjectServiceServer) {
	s.RegisterService(&SubjectService_ServiceDesc, srv)
}

func _SubjectService_CreateSubject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSubjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubjectServiceServer).CreateSubject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubjectService_CreateSubject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubjectServiceServer).CreateSubject(ctx, req.(*CreateSubjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubjectService_GetSubject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubjectServiceServer).GetSubject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubjectService_GetSubject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubjectServiceServer).GetSubject(ctx, req.(*GetSubjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubjectService_GetAllSubjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllSubjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubjectServiceServer).GetAllSubjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubjectService_GetAllSubjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubjectServiceServer).GetAllSubjects(ctx, req.(*GetAllSubjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubjectService_UpdateSubject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSubjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubjectServiceServer).UpdateSubject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubjectService_UpdateSubject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubjectServiceServer).UpdateSubject(ctx, req.(*UpdateSubjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubjectService_DeleteSubject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSubjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubjectServiceServer).DeleteSubject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubjectService_DeleteSubject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubjectServiceServer).DeleteSubject(ctx, req.(*DeleteSubjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SubjectService_ServiceDesc is the grpc.ServiceDesc for SubjectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubjectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "subject.SubjectService",
	HandlerType: (*SubjectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSubject",
			Handler:    _SubjectService_CreateSubject_Handler,
		},
		{
			MethodName: "GetSubject",
			Handler:    _SubjectService_GetSubject_Handler,
		},
		{
			MethodName: "GetAllSubjects",
			Handler:    _SubjectService_GetAllSubjects_Handler,
		},
		{
			MethodName: "UpdateSubject",
			Handler:    _SubjectService_UpdateSubject_Handler,
		},
		{
			MethodName: "DeleteSubject",
			Handler:    _SubjectService_DeleteSubject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "subject.proto",
}
