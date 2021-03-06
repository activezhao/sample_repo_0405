// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package sample_package_0405

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

// Zhaoh0405Client is the client API for Zhaoh0405 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Zhaoh0405Client interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type zhaoh0405Client struct {
	cc grpc.ClientConnInterface
}

func NewZhaoh0405Client(cc grpc.ClientConnInterface) Zhaoh0405Client {
	return &zhaoh0405Client{cc}
}

func (c *zhaoh0405Client) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/sample_module_0405.sample_package_0405.Zhaoh0405/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Zhaoh0405Server is the server API for Zhaoh0405 service.
// All implementations must embed UnimplementedZhaoh0405Server
// for forward compatibility
type Zhaoh0405Server interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedZhaoh0405Server()
}

// UnimplementedZhaoh0405Server must be embedded to have forward compatible implementations.
type UnimplementedZhaoh0405Server struct {
}

func (UnimplementedZhaoh0405Server) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedZhaoh0405Server) mustEmbedUnimplementedZhaoh0405Server() {}

// UnsafeZhaoh0405Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Zhaoh0405Server will
// result in compilation errors.
type UnsafeZhaoh0405Server interface {
	mustEmbedUnimplementedZhaoh0405Server()
}

func RegisterZhaoh0405Server(s grpc.ServiceRegistrar, srv Zhaoh0405Server) {
	s.RegisterService(&Zhaoh0405_ServiceDesc, srv)
}

func _Zhaoh0405_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Zhaoh0405Server).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sample_module_0405.sample_package_0405.Zhaoh0405/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Zhaoh0405Server).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Zhaoh0405_ServiceDesc is the grpc.ServiceDesc for Zhaoh0405 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Zhaoh0405_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sample_module_0405.sample_package_0405.Zhaoh0405",
	HandlerType: (*Zhaoh0405Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Zhaoh0405_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Zhaoh0405.proto",
}
