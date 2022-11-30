// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: proto/demoroo_api.proto

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

// BranchSearchServiceClient is the client API for BranchSearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BranchSearchServiceClient interface {
	Search(ctx context.Context, in *BranchSearchRequest, opts ...grpc.CallOption) (*BranchSearchResponse, error)
}

type branchSearchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBranchSearchServiceClient(cc grpc.ClientConnInterface) BranchSearchServiceClient {
	return &branchSearchServiceClient{cc}
}

func (c *branchSearchServiceClient) Search(ctx context.Context, in *BranchSearchRequest, opts ...grpc.CallOption) (*BranchSearchResponse, error) {
	out := new(BranchSearchResponse)
	err := c.cc.Invoke(ctx, "/BranchSearchService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BranchSearchServiceServer is the server API for BranchSearchService service.
// All implementations must embed UnimplementedBranchSearchServiceServer
// for forward compatibility
type BranchSearchServiceServer interface {
	Search(context.Context, *BranchSearchRequest) (*BranchSearchResponse, error)
	mustEmbedUnimplementedBranchSearchServiceServer()
}

// UnimplementedBranchSearchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBranchSearchServiceServer struct {
}

func (UnimplementedBranchSearchServiceServer) Search(context.Context, *BranchSearchRequest) (*BranchSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedBranchSearchServiceServer) mustEmbedUnimplementedBranchSearchServiceServer() {}

// UnsafeBranchSearchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BranchSearchServiceServer will
// result in compilation errors.
type UnsafeBranchSearchServiceServer interface {
	mustEmbedUnimplementedBranchSearchServiceServer()
}

func RegisterBranchSearchServiceServer(s grpc.ServiceRegistrar, srv BranchSearchServiceServer) {
	s.RegisterService(&BranchSearchService_ServiceDesc, srv)
}

func _BranchSearchService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BranchSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchSearchServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BranchSearchService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchSearchServiceServer).Search(ctx, req.(*BranchSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BranchSearchService_ServiceDesc is the grpc.ServiceDesc for BranchSearchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BranchSearchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BranchSearchService",
	HandlerType: (*BranchSearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _BranchSearchService_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/demoroo_api.proto",
}
