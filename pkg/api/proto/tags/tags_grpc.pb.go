// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.0--rc2
// source: api/tags.proto

package tags

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
	TagsService_Create_FullMethodName = "/tags.TagsService/Create"
	TagsService_Get_FullMethodName    = "/tags.TagsService/Get"
	TagsService_GetAll_FullMethodName = "/tags.TagsService/GetAll"
)

// TagsServiceClient is the client API for TagsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TagsServiceClient interface {
	Create(ctx context.Context, in *TagsCreateRequest, opts ...grpc.CallOption) (*TagsCreateResponse, error)
	Get(ctx context.Context, in *TagsGetRequest, opts ...grpc.CallOption) (*TagsGetResponse, error)
	GetAll(ctx context.Context, in *TagsGetAllRequest, opts ...grpc.CallOption) (*TagsGetAllResponse, error)
}

type tagsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTagsServiceClient(cc grpc.ClientConnInterface) TagsServiceClient {
	return &tagsServiceClient{cc}
}

func (c *tagsServiceClient) Create(ctx context.Context, in *TagsCreateRequest, opts ...grpc.CallOption) (*TagsCreateResponse, error) {
	out := new(TagsCreateResponse)
	err := c.cc.Invoke(ctx, TagsService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagsServiceClient) Get(ctx context.Context, in *TagsGetRequest, opts ...grpc.CallOption) (*TagsGetResponse, error) {
	out := new(TagsGetResponse)
	err := c.cc.Invoke(ctx, TagsService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagsServiceClient) GetAll(ctx context.Context, in *TagsGetAllRequest, opts ...grpc.CallOption) (*TagsGetAllResponse, error) {
	out := new(TagsGetAllResponse)
	err := c.cc.Invoke(ctx, TagsService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TagsServiceServer is the server API for TagsService service.
// All implementations must embed UnimplementedTagsServiceServer
// for forward compatibility
type TagsServiceServer interface {
	Create(context.Context, *TagsCreateRequest) (*TagsCreateResponse, error)
	Get(context.Context, *TagsGetRequest) (*TagsGetResponse, error)
	GetAll(context.Context, *TagsGetAllRequest) (*TagsGetAllResponse, error)
	mustEmbedUnimplementedTagsServiceServer()
}

// UnimplementedTagsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTagsServiceServer struct {
}

func (UnimplementedTagsServiceServer) Create(context.Context, *TagsCreateRequest) (*TagsCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTagsServiceServer) Get(context.Context, *TagsGetRequest) (*TagsGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTagsServiceServer) GetAll(context.Context, *TagsGetAllRequest) (*TagsGetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedTagsServiceServer) mustEmbedUnimplementedTagsServiceServer() {}

// UnsafeTagsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TagsServiceServer will
// result in compilation errors.
type UnsafeTagsServiceServer interface {
	mustEmbedUnimplementedTagsServiceServer()
}

func RegisterTagsServiceServer(s grpc.ServiceRegistrar, srv TagsServiceServer) {
	s.RegisterService(&TagsService_ServiceDesc, srv)
}

func _TagsService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TagsCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagsService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsServiceServer).Create(ctx, req.(*TagsCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagsService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TagsGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagsService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsServiceServer).Get(ctx, req.(*TagsGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagsService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TagsGetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagsService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsServiceServer).GetAll(ctx, req.(*TagsGetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TagsService_ServiceDesc is the grpc.ServiceDesc for TagsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TagsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tags.TagsService",
	HandlerType: (*TagsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TagsService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _TagsService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _TagsService_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/tags.proto",
}
