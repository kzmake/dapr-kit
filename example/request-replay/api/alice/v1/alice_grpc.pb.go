// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// AliceServiceClient is the client API for AliceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AliceServiceClient interface {
	Handle(ctx context.Context, in *HandleRequest, opts ...grpc.CallOption) (*HandleResponse, error)
}

type aliceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAliceServiceClient(cc grpc.ClientConnInterface) AliceServiceClient {
	return &aliceServiceClient{cc}
}

func (c *aliceServiceClient) Handle(ctx context.Context, in *HandleRequest, opts ...grpc.CallOption) (*HandleResponse, error) {
	out := new(HandleResponse)
	err := c.cc.Invoke(ctx, "/api.alice.v1.AliceService/Handle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AliceServiceServer is the server API for AliceService service.
// All implementations must embed UnimplementedAliceServiceServer
// for forward compatibility
type AliceServiceServer interface {
	Handle(context.Context, *HandleRequest) (*HandleResponse, error)
	mustEmbedUnimplementedAliceServiceServer()
}

// UnimplementedAliceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAliceServiceServer struct {
}

func (UnimplementedAliceServiceServer) Handle(context.Context, *HandleRequest) (*HandleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handle not implemented")
}
func (UnimplementedAliceServiceServer) mustEmbedUnimplementedAliceServiceServer() {}

// UnsafeAliceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AliceServiceServer will
// result in compilation errors.
type UnsafeAliceServiceServer interface {
	mustEmbedUnimplementedAliceServiceServer()
}

func RegisterAliceServiceServer(s grpc.ServiceRegistrar, srv AliceServiceServer) {
	s.RegisterService(&_AliceService_serviceDesc, srv)
}

func _AliceService_Handle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AliceServiceServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.alice.v1.AliceService/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AliceServiceServer).Handle(ctx, req.(*HandleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AliceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.alice.v1.AliceService",
	HandlerType: (*AliceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _AliceService_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/request-replay/api/alice/v1/alice.proto",
}
