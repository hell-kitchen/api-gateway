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
	TagsService_Create_FullMethodName     = "/tags.TagsService/Create"
	TagsService_CreateMany_FullMethodName = "/tags.TagsService/CreateMany"
	TagsService_Get_FullMethodName        = "/tags.TagsService/Get"
	TagsService_GetAll_FullMethodName     = "/tags.TagsService/GetAll"
	TagsService_Delete_FullMethodName     = "/tags.TagsService/Delete"
	TagsService_Update_FullMethodName     = "/tags.TagsService/Update"
)

// TagsServiceClient is the client API for TagsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TagsServiceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	CreateMany(ctx context.Context, in *CreateManyRequest, opts ...grpc.CallOption) (*CreateManyResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
}

type tagsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTagsServiceClient(cc grpc.ClientConnInterface) TagsServiceClient {
	return &tagsServiceClient{cc}
}

func (c *tagsServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, TagsService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagsServiceClient) CreateMany(ctx context.Context, in *CreateManyRequest, opts ...grpc.CallOption) (*CreateManyResponse, error) {
	out := new(CreateManyResponse)
	err := c.cc.Invoke(ctx, TagsService_CreateMany_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagsServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, TagsService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagsServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, TagsService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagsServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, TagsService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagsServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, TagsService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TagsServiceServer is the server API for TagsService service.
// All implementations must embed UnimplementedTagsServiceServer
// for forward compatibility
type TagsServiceServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	CreateMany(context.Context, *CreateManyRequest) (*CreateManyResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	mustEmbedUnimplementedTagsServiceServer()
}

// UnimplementedTagsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTagsServiceServer struct {
}

func (UnimplementedTagsServiceServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTagsServiceServer) CreateMany(context.Context, *CreateManyRequest) (*CreateManyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMany not implemented")
}
func (UnimplementedTagsServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTagsServiceServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedTagsServiceServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTagsServiceServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
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
	in := new(CreateRequest)
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
		return srv.(TagsServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagsService_CreateMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateManyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsServiceServer).CreateMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagsService_CreateMany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsServiceServer).CreateMany(ctx, req.(*CreateManyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagsService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
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
		return srv.(TagsServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagsService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
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
		return srv.(TagsServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagsService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagsService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagsService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagsService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsServiceServer).Update(ctx, req.(*UpdateRequest))
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
			MethodName: "CreateMany",
			Handler:    _TagsService_CreateMany_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _TagsService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _TagsService_GetAll_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TagsService_Delete_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TagsService_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/tags.proto",
}