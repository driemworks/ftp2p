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
	GetNodeStatus(ctx context.Context, in *NodeInfoRequest, opts ...grpc.CallOption) (*NodeInfoResponse, error)
	// read/write to known peers
	ListKnownPeers(ctx context.Context, in *ListKnownPeersRequest, opts ...grpc.CallOption) (PublicNode_ListKnownPeersClient, error)
	// TODO: Authentication
	JoinKnownPeers(ctx context.Context, in *JoinKnownPeersRequest, opts ...grpc.CallOption) (*JoinKnownPeersResponse, error)
	// read/write blocks
	ListBlocks(ctx context.Context, in *ListBlocksRequest, opts ...grpc.CallOption) (PublicNode_ListBlocksClient, error)
	// add pending transaction
	AddPendingPublishCIDTransaction(ctx context.Context, in *AddPendingPublishCIDTransactionRequest, opts ...grpc.CallOption) (*AddPendingPublishCIDTransactionResponse, error)
	// list pending transactions
	ListPendingTransactions(ctx context.Context, in *ListPendingTransactionsRequest, opts ...grpc.CallOption) (PublicNode_ListPendingTransactionsClient, error)
}

type publicNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewPublicNodeClient(cc grpc.ClientConnInterface) PublicNodeClient {
	return &publicNodeClient{cc}
}

func (c *publicNodeClient) GetNodeStatus(ctx context.Context, in *NodeInfoRequest, opts ...grpc.CallOption) (*NodeInfoResponse, error) {
	out := new(NodeInfoResponse)
	err := c.cc.Invoke(ctx, "/proto.PublicNode/GetNodeStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicNodeClient) ListKnownPeers(ctx context.Context, in *ListKnownPeersRequest, opts ...grpc.CallOption) (PublicNode_ListKnownPeersClient, error) {
	stream, err := c.cc.NewStream(ctx, &PublicNode_ServiceDesc.Streams[0], "/proto.PublicNode/ListKnownPeers", opts...)
	if err != nil {
		return nil, err
	}
	x := &publicNodeListKnownPeersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PublicNode_ListKnownPeersClient interface {
	Recv() (*ListKnownPeersResponse, error)
	grpc.ClientStream
}

type publicNodeListKnownPeersClient struct {
	grpc.ClientStream
}

func (x *publicNodeListKnownPeersClient) Recv() (*ListKnownPeersResponse, error) {
	m := new(ListKnownPeersResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *publicNodeClient) JoinKnownPeers(ctx context.Context, in *JoinKnownPeersRequest, opts ...grpc.CallOption) (*JoinKnownPeersResponse, error) {
	out := new(JoinKnownPeersResponse)
	err := c.cc.Invoke(ctx, "/proto.PublicNode/JoinKnownPeers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicNodeClient) ListBlocks(ctx context.Context, in *ListBlocksRequest, opts ...grpc.CallOption) (PublicNode_ListBlocksClient, error) {
	stream, err := c.cc.NewStream(ctx, &PublicNode_ServiceDesc.Streams[1], "/proto.PublicNode/ListBlocks", opts...)
	if err != nil {
		return nil, err
	}
	x := &publicNodeListBlocksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PublicNode_ListBlocksClient interface {
	Recv() (*BlockResponse, error)
	grpc.ClientStream
}

type publicNodeListBlocksClient struct {
	grpc.ClientStream
}

func (x *publicNodeListBlocksClient) Recv() (*BlockResponse, error) {
	m := new(BlockResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *publicNodeClient) AddPendingPublishCIDTransaction(ctx context.Context, in *AddPendingPublishCIDTransactionRequest, opts ...grpc.CallOption) (*AddPendingPublishCIDTransactionResponse, error) {
	out := new(AddPendingPublishCIDTransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.PublicNode/AddPendingPublishCIDTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicNodeClient) ListPendingTransactions(ctx context.Context, in *ListPendingTransactionsRequest, opts ...grpc.CallOption) (PublicNode_ListPendingTransactionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &PublicNode_ServiceDesc.Streams[2], "/proto.PublicNode/ListPendingTransactions", opts...)
	if err != nil {
		return nil, err
	}
	x := &publicNodeListPendingTransactionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PublicNode_ListPendingTransactionsClient interface {
	Recv() (*PendingTransactionResponse, error)
	grpc.ClientStream
}

type publicNodeListPendingTransactionsClient struct {
	grpc.ClientStream
}

func (x *publicNodeListPendingTransactionsClient) Recv() (*PendingTransactionResponse, error) {
	m := new(PendingTransactionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PublicNodeServer is the server API for PublicNode service.
// All implementations must embed UnimplementedPublicNodeServer
// for forward compatibility
type PublicNodeServer interface {
	// Obtains a node's name
	GetNodeStatus(context.Context, *NodeInfoRequest) (*NodeInfoResponse, error)
	// read/write to known peers
	ListKnownPeers(*ListKnownPeersRequest, PublicNode_ListKnownPeersServer) error
	// TODO: Authentication
	JoinKnownPeers(context.Context, *JoinKnownPeersRequest) (*JoinKnownPeersResponse, error)
	// read/write blocks
	ListBlocks(*ListBlocksRequest, PublicNode_ListBlocksServer) error
	// add pending transaction
	AddPendingPublishCIDTransaction(context.Context, *AddPendingPublishCIDTransactionRequest) (*AddPendingPublishCIDTransactionResponse, error)
	// list pending transactions
	ListPendingTransactions(*ListPendingTransactionsRequest, PublicNode_ListPendingTransactionsServer) error
	mustEmbedUnimplementedPublicNodeServer()
}

// UnimplementedPublicNodeServer must be embedded to have forward compatible implementations.
type UnimplementedPublicNodeServer struct {
}

func (UnimplementedPublicNodeServer) GetNodeStatus(context.Context, *NodeInfoRequest) (*NodeInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeStatus not implemented")
}
func (UnimplementedPublicNodeServer) ListKnownPeers(*ListKnownPeersRequest, PublicNode_ListKnownPeersServer) error {
	return status.Errorf(codes.Unimplemented, "method ListKnownPeers not implemented")
}
func (UnimplementedPublicNodeServer) JoinKnownPeers(context.Context, *JoinKnownPeersRequest) (*JoinKnownPeersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinKnownPeers not implemented")
}
func (UnimplementedPublicNodeServer) ListBlocks(*ListBlocksRequest, PublicNode_ListBlocksServer) error {
	return status.Errorf(codes.Unimplemented, "method ListBlocks not implemented")
}
func (UnimplementedPublicNodeServer) AddPendingPublishCIDTransaction(context.Context, *AddPendingPublishCIDTransactionRequest) (*AddPendingPublishCIDTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPendingPublishCIDTransaction not implemented")
}
func (UnimplementedPublicNodeServer) ListPendingTransactions(*ListPendingTransactionsRequest, PublicNode_ListPendingTransactionsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListPendingTransactions not implemented")
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

func _PublicNode_GetNodeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicNodeServer).GetNodeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PublicNode/GetNodeStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicNodeServer).GetNodeStatus(ctx, req.(*NodeInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicNode_ListKnownPeers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListKnownPeersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PublicNodeServer).ListKnownPeers(m, &publicNodeListKnownPeersServer{stream})
}

type PublicNode_ListKnownPeersServer interface {
	Send(*ListKnownPeersResponse) error
	grpc.ServerStream
}

type publicNodeListKnownPeersServer struct {
	grpc.ServerStream
}

func (x *publicNodeListKnownPeersServer) Send(m *ListKnownPeersResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _PublicNode_JoinKnownPeers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinKnownPeersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicNodeServer).JoinKnownPeers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PublicNode/JoinKnownPeers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicNodeServer).JoinKnownPeers(ctx, req.(*JoinKnownPeersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicNode_ListBlocks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListBlocksRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PublicNodeServer).ListBlocks(m, &publicNodeListBlocksServer{stream})
}

type PublicNode_ListBlocksServer interface {
	Send(*BlockResponse) error
	grpc.ServerStream
}

type publicNodeListBlocksServer struct {
	grpc.ServerStream
}

func (x *publicNodeListBlocksServer) Send(m *BlockResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _PublicNode_AddPendingPublishCIDTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPendingPublishCIDTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicNodeServer).AddPendingPublishCIDTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PublicNode/AddPendingPublishCIDTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicNodeServer).AddPendingPublishCIDTransaction(ctx, req.(*AddPendingPublishCIDTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicNode_ListPendingTransactions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListPendingTransactionsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PublicNodeServer).ListPendingTransactions(m, &publicNodeListPendingTransactionsServer{stream})
}

type PublicNode_ListPendingTransactionsServer interface {
	Send(*PendingTransactionResponse) error
	grpc.ServerStream
}

type publicNodeListPendingTransactionsServer struct {
	grpc.ServerStream
}

func (x *publicNodeListPendingTransactionsServer) Send(m *PendingTransactionResponse) error {
	return x.ServerStream.SendMsg(m)
}

// PublicNode_ServiceDesc is the grpc.ServiceDesc for PublicNode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PublicNode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PublicNode",
	HandlerType: (*PublicNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNodeStatus",
			Handler:    _PublicNode_GetNodeStatus_Handler,
		},
		{
			MethodName: "JoinKnownPeers",
			Handler:    _PublicNode_JoinKnownPeers_Handler,
		},
		{
			MethodName: "AddPendingPublishCIDTransaction",
			Handler:    _PublicNode_AddPendingPublishCIDTransaction_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListKnownPeers",
			Handler:       _PublicNode_ListKnownPeers_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListBlocks",
			Handler:       _PublicNode_ListBlocks_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListPendingTransactions",
			Handler:       _PublicNode_ListPendingTransactions_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/node.proto",
}
