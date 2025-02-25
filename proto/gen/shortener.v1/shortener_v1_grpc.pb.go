// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: shortener_v1.proto

package shortener_v1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Shortener_CreateLink_FullMethodName = "/shortener_v1.Shortener/CreateLink"
	Shortener_FetchLink_FullMethodName  = "/shortener_v1.Shortener/FetchLink"
)

// ShortenerClient is the client API for Shortener service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShortenerClient interface {
	CreateLink(ctx context.Context, in *CreateLinkRequest, opts ...grpc.CallOption) (*CreateLinkResponse, error)
	FetchLink(ctx context.Context, in *FetchLinkRequest, opts ...grpc.CallOption) (*FetchLinkResponse, error)
}

type shortenerClient struct {
	cc grpc.ClientConnInterface
}

func NewShortenerClient(cc grpc.ClientConnInterface) ShortenerClient {
	return &shortenerClient{cc}
}

func (c *shortenerClient) CreateLink(ctx context.Context, in *CreateLinkRequest, opts ...grpc.CallOption) (*CreateLinkResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateLinkResponse)
	err := c.cc.Invoke(ctx, Shortener_CreateLink_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenerClient) FetchLink(ctx context.Context, in *FetchLinkRequest, opts ...grpc.CallOption) (*FetchLinkResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FetchLinkResponse)
	err := c.cc.Invoke(ctx, Shortener_FetchLink_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShortenerServer is the server API for Shortener service.
// All implementations must embed UnimplementedShortenerServer
// for forward compatibility.
type ShortenerServer interface {
	CreateLink(context.Context, *CreateLinkRequest) (*CreateLinkResponse, error)
	FetchLink(context.Context, *FetchLinkRequest) (*FetchLinkResponse, error)
	mustEmbedUnimplementedShortenerServer()
}

// UnimplementedShortenerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedShortenerServer struct{}

func (UnimplementedShortenerServer) CreateLink(context.Context, *CreateLinkRequest) (*CreateLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLink not implemented")
}
func (UnimplementedShortenerServer) FetchLink(context.Context, *FetchLinkRequest) (*FetchLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchLink not implemented")
}
func (UnimplementedShortenerServer) mustEmbedUnimplementedShortenerServer() {}
func (UnimplementedShortenerServer) testEmbeddedByValue()                   {}

// UnsafeShortenerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShortenerServer will
// result in compilation errors.
type UnsafeShortenerServer interface {
	mustEmbedUnimplementedShortenerServer()
}

func RegisterShortenerServer(s grpc.ServiceRegistrar, srv ShortenerServer) {
	// If the following call pancis, it indicates UnimplementedShortenerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Shortener_ServiceDesc, srv)
}

func _Shortener_CreateLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).CreateLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Shortener_CreateLink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).CreateLink(ctx, req.(*CreateLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shortener_FetchLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).FetchLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Shortener_FetchLink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).FetchLink(ctx, req.(*FetchLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Shortener_ServiceDesc is the grpc.ServiceDesc for Shortener service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Shortener_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shortener_v1.Shortener",
	HandlerType: (*ShortenerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLink",
			Handler:    _Shortener_CreateLink_Handler,
		},
		{
			MethodName: "FetchLink",
			Handler:    _Shortener_FetchLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shortener_v1.proto",
}
