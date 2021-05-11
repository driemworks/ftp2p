// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package __

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

// PublicNodeClient is the client API for PublicNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PublicNodeClient interface {
	// Obtains a node's name
	// rpc GetNodeStatus(NodeInfoRequest) returns (NodeInfoResponse) {}
	// // read/write to known peers
	// rpc ListKnownPeers(ListKnownPeersRequest) returns (stream ListKnownPeersResponse) {}
	// // read/write blocks
	// rpc ListBlocks(ListBlocksRequest) returns (stream BlockResponse) {}
	// add pending transaction
	AddTransaction(ctx context.Context, in *AddPendingPublishCIDTransactionRequest, opts ...grpc.CallOption) (*AddPendingPublishCIDTransactionResponse, error)
}

type publicNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewPublicNodeClient(cc grpc.ClientConnInterface) PublicNodeClient {
	return &publicNodeClient{cc}
}

func (c *publicNodeClient) AddTransaction(ctx context.Context, in *AddPendingPublishCIDTransactionRequest, opts ...grpc.CallOption) (*AddPendingPublishCIDTransactionResponse, error) {
	out := new(AddPendingPublishCIDTransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.PublicNode/AddTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PublicNodeServer is the server API for PublicNode service.
// All implementations must embed UnimplementedPublicNodeServer
// for forward compatibility
type PublicNodeServer interface {
	// Obtains a node's name
	// rpc GetNodeStatus(NodeInfoRequest) returns (NodeInfoResponse) {}
	// // read/write to known peers
	// rpc ListKnownPeers(ListKnownPeersRequest) returns (stream ListKnownPeersResponse) {}
	// // read/write blocks
	// rpc ListBlocks(ListBlocksRequest) returns (stream BlockResponse) {}
	// add pending transaction
	AddTransaction(context.Context, *AddPendingPublishCIDTransactionRequest) (*AddPendingPublishCIDTransactionResponse, error)
	mustEmbedUnimplementedPublicNodeServer()
}

// UnimplementedPublicNodeServer must be embedded to have forward compatible implementations.
type UnimplementedPublicNodeServer struct {
}

func (UnimplementedPublicNodeServer) AddTransaction(context.Context, *AddPendingPublishCIDTransactionRequest) (*AddPendingPublishCIDTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTransaction not implemented")
}
func (UnimplementedPublicNodeServer) mustEmbedUnimplementedPublicNodeServer() {}

// UnsafePublicNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PublicNodeServer will
// result in compilation errors.
type UnsafePublicNodeServer interface {
	mustEmbedUnimplementedPublicNodeServer()
}

func RegisterPublicNodeServer(s grpc.ServiceRegistrar, srv PublicNodeServer) {
	s.RegisterService(&PublicNode_ServiceDesc, srv)
}

func _PublicNode_AddTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPendingPublishCIDTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicNodeServer).AddTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PublicNode/AddTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicNodeServer).AddTransaction(ctx, req.(*AddPendingPublishCIDTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PublicNode_ServiceDesc is the grpc.ServiceDesc for PublicNode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PublicNode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PublicNode",
	HandlerType: (*PublicNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTransaction",
			Handler:    _PublicNode_AddTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/node.proto",
}
