// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.1
// source: restoration.proto

package pbg

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
	RestorationService_RestorationCollection_FullMethodName = "/RestorationService/RestorationCollection"
)

// RestorationServiceClient is the client API for RestorationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RestorationServiceClient interface {
	RestorationCollection(ctx context.Context, in *RestorationCollectionRequest, opts ...grpc.CallOption) (*RestorationCollectionResponse, error)
}

type restorationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRestorationServiceClient(cc grpc.ClientConnInterface) RestorationServiceClient {
	return &restorationServiceClient{cc}
}

func (c *restorationServiceClient) RestorationCollection(ctx context.Context, in *RestorationCollectionRequest, opts ...grpc.CallOption) (*RestorationCollectionResponse, error) {
	out := new(RestorationCollectionResponse)
	err := c.cc.Invoke(ctx, RestorationService_RestorationCollection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RestorationServiceServer is the server API for RestorationService service.
// All implementations must embed UnimplementedRestorationServiceServer
// for forward compatibility
type RestorationServiceServer interface {
	RestorationCollection(context.Context, *RestorationCollectionRequest) (*RestorationCollectionResponse, error)
	mustEmbedUnimplementedRestorationServiceServer()
}

// UnimplementedRestorationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRestorationServiceServer struct {
}

func (UnimplementedRestorationServiceServer) RestorationCollection(context.Context, *RestorationCollectionRequest) (*RestorationCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestorationCollection not implemented")
}
func (UnimplementedRestorationServiceServer) mustEmbedUnimplementedRestorationServiceServer() {}

// UnsafeRestorationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RestorationServiceServer will
// result in compilation errors.
type UnsafeRestorationServiceServer interface {
	mustEmbedUnimplementedRestorationServiceServer()
}

func RegisterRestorationServiceServer(s grpc.ServiceRegistrar, srv RestorationServiceServer) {
	s.RegisterService(&RestorationService_ServiceDesc, srv)
}

func _RestorationService_RestorationCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestorationCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestorationServiceServer).RestorationCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RestorationService_RestorationCollection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestorationServiceServer).RestorationCollection(ctx, req.(*RestorationCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RestorationService_ServiceDesc is the grpc.ServiceDesc for RestorationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RestorationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RestorationService",
	HandlerType: (*RestorationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RestorationCollection",
			Handler:    _RestorationService_RestorationCollection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "restoration.proto",
}
