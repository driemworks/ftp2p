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

// NodeServiceClient is the client API for NodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeServiceClient interface {
	// Obtains a node's name
	GetNodeStatus(ctx context.Context, in *NodeInfoRequest, opts ...grpc.CallOption) (*NodeInfoResponse, error)
	// // read/write to known peers
	// rpc ListKnownPeers(ListKnownPeersRequest) returns (stream ListKnownPeersResponse) {}
	// // read/write blocks
	ListBlocks(ctx context.Context, in *ListBlocksRequest, opts ...grpc.CallOption) (NodeService_ListBlocksClient, error)
	// add pending transaction
	AddTransaction(ctx context.Context, in *AddPendingTransactionRequest, opts ...grpc.CallOption) (*AddPendingTransactionResponse, error)
	Subscribe(ctx context.Context, in *JoinChannelRequest, opts ...grpc.CallOption) (NodeService_SubscribeClient, error)
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error)
}

type nodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeServiceClient(cc grpc.ClientConnInterface) NodeServiceClient {
	return &nodeServiceClient{cc}
}

func (c *nodeServiceClient) GetNodeStatus(ctx context.Context, in *NodeInfoRequest, opts ...grpc.CallOption) (*NodeInfoResponse, error) {
	out := new(NodeInfoResponse)
	err := c.cc.Invoke(ctx, "/proto.NodeService/GetNodeStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) ListBlocks(ctx context.Context, in *ListBlocksRequest, opts ...grpc.CallOption) (NodeService_ListBlocksClient, error) {
	stream, err := c.cc.NewStream(ctx, &NodeService_ServiceDesc.Streams[0], "/proto.NodeService/ListBlocks", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeServiceListBlocksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeService_ListBlocksClient interface {
	Recv() (*BlockResponse, error)
	grpc.ClientStream
}

type nodeServiceListBlocksClient struct {
	grpc.ClientStream
}

func (x *nodeServiceListBlocksClient) Recv() (*BlockResponse, error) {
	m := new(BlockResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeServiceClient) AddTransaction(ctx context.Context, in *AddPendingTransactionRequest, opts ...grpc.CallOption) (*AddPendingTransactionResponse, error) {
	out := new(AddPendingTransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.NodeService/AddTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) Subscribe(ctx context.Context, in *JoinChannelRequest, opts ...grpc.CallOption) (NodeService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &NodeService_ServiceDesc.Streams[1], "/proto.NodeService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeService_SubscribeClient interface {
	Recv() (*ChannelData, error)
	grpc.ClientStream
}

type nodeServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *nodeServiceSubscribeClient) Recv() (*ChannelData, error) {
	m := new(ChannelData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeServiceClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error) {
	out := new(PublishResponse)
	err := c.cc.Invoke(ctx, "/proto.NodeService/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServiceServer is the server API for NodeService service.
// All implementations must embed UnimplementedNodeServiceServer
// for forward compatibility
type NodeServiceServer interface {
	// Obtains a node's name
	GetNodeStatus(context.Context, *NodeInfoRequest) (*NodeInfoResponse, error)
	// // read/write to known peers
	// rpc ListKnownPeers(ListKnownPeersRequest) returns (stream ListKnownPeersResponse) {}
	// // read/write blocks
	ListBlocks(*ListBlocksRequest, NodeService_ListBlocksServer) error
	// add pending transaction
	AddTransaction(context.Context, *AddPendingTransactionRequest) (*AddPendingTransactionResponse, error)
	Subscribe(*JoinChannelRequest, NodeService_SubscribeServer) error
	Publish(context.Context, *PublishRequest) (*PublishResponse, error)
	mustEmbedUnimplementedNodeServiceServer()
}

// UnimplementedNodeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNodeServiceServer struct {
}

func (UnimplementedNodeServiceServer) GetNodeStatus(context.Context, *NodeInfoRequest) (*NodeInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeStatus not implemented")
}
func (UnimplementedNodeServiceServer) ListBlocks(*ListBlocksRequest, NodeService_ListBlocksServer) error {
	return status.Errorf(codes.Unimplemented, "method ListBlocks not implemented")
}
func (UnimplementedNodeServiceServer) AddTransaction(context.Context, *AddPendingTransactionRequest) (*AddPendingTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTransaction not implemented")
}
func (UnimplementedNodeServiceServer) Subscribe(*JoinChannelRequest, NodeService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedNodeServiceServer) Publish(context.Context, *PublishRequest) (*PublishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedNodeServiceServer) mustEmbedUnimplementedNodeServiceServer() {}

// UnsafeNodeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeServiceServer will
// result in compilation errors.
type UnsafeNodeServiceServer interface {
	mustEmbedUnimplementedNodeServiceServer()
}

func RegisterNodeServiceServer(s grpc.ServiceRegistrar, srv NodeServiceServer) {
	s.RegisterService(&NodeService_ServiceDesc, srv)
}

func _NodeService_GetNodeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).GetNodeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeService/GetNodeStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).GetNodeStatus(ctx, req.(*NodeInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_ListBlocks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListBlocksRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServiceServer).ListBlocks(m, &nodeServiceListBlocksServer{stream})
}

type NodeService_ListBlocksServer interface {
	Send(*BlockResponse) error
	grpc.ServerStream
}

type nodeServiceListBlocksServer struct {
	grpc.ServerStream
}

func (x *nodeServiceListBlocksServer) Send(m *BlockResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeService_AddTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPendingTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).AddTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeService/AddTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).AddTransaction(ctx, req.(*AddPendingTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(JoinChannelRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServiceServer).Subscribe(m, &nodeServiceSubscribeServer{stream})
}

type NodeService_SubscribeServer interface {
	Send(*ChannelData) error
	grpc.ServerStream
}

type nodeServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *nodeServiceSubscribeServer) Send(m *ChannelData) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeService_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeService/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeService_ServiceDesc is the grpc.ServiceDesc for NodeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.NodeService",
	HandlerType: (*NodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNodeStatus",
			Handler:    _NodeService_GetNodeStatus_Handler,
		},
		{
			MethodName: "AddTransaction",
			Handler:    _NodeService_AddTransaction_Handler,
		},
		{
			MethodName: "Publish",
			Handler:    _NodeService_Publish_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListBlocks",
			Handler:       _NodeService_ListBlocks_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Subscribe",
			Handler:       _NodeService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/node.proto",
}
