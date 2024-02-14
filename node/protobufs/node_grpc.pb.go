// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: node.proto

package protobufs

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
	NodeService_GetFrames_FullMethodName      = "/quilibrium.node.node.pb.NodeService/GetFrames"
	NodeService_GetFrameInfo_FullMethodName   = "/quilibrium.node.node.pb.NodeService/GetFrameInfo"
	NodeService_GetPeerInfo_FullMethodName    = "/quilibrium.node.node.pb.NodeService/GetPeerInfo"
	NodeService_GetNodeInfo_FullMethodName    = "/quilibrium.node.node.pb.NodeService/GetNodeInfo"
	NodeService_GetNetworkInfo_FullMethodName = "/quilibrium.node.node.pb.NodeService/GetNetworkInfo"
	NodeService_GetTokenInfo_FullMethodName   = "/quilibrium.node.node.pb.NodeService/GetTokenInfo"
)

// NodeServiceClient is the client API for NodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeServiceClient interface {
	GetFrames(ctx context.Context, in *GetFramesRequest, opts ...grpc.CallOption) (*FramesResponse, error)
	GetFrameInfo(ctx context.Context, in *GetFrameInfoRequest, opts ...grpc.CallOption) (*FrameInfoResponse, error)
	GetPeerInfo(ctx context.Context, in *GetPeerInfoRequest, opts ...grpc.CallOption) (*PeerInfoResponse, error)
	GetNodeInfo(ctx context.Context, in *GetNodeInfoRequest, opts ...grpc.CallOption) (*NodeInfoResponse, error)
	GetNetworkInfo(ctx context.Context, in *GetNetworkInfoRequest, opts ...grpc.CallOption) (*NetworkInfoResponse, error)
	GetTokenInfo(ctx context.Context, in *GetTokenInfoRequest, opts ...grpc.CallOption) (*TokenInfoResponse, error)
}

type nodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeServiceClient(cc grpc.ClientConnInterface) NodeServiceClient {
	return &nodeServiceClient{cc}
}

func (c *nodeServiceClient) GetFrames(ctx context.Context, in *GetFramesRequest, opts ...grpc.CallOption) (*FramesResponse, error) {
	out := new(FramesResponse)
	err := c.cc.Invoke(ctx, NodeService_GetFrames_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) GetFrameInfo(ctx context.Context, in *GetFrameInfoRequest, opts ...grpc.CallOption) (*FrameInfoResponse, error) {
	out := new(FrameInfoResponse)
	err := c.cc.Invoke(ctx, NodeService_GetFrameInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) GetPeerInfo(ctx context.Context, in *GetPeerInfoRequest, opts ...grpc.CallOption) (*PeerInfoResponse, error) {
	out := new(PeerInfoResponse)
	err := c.cc.Invoke(ctx, NodeService_GetPeerInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) GetNodeInfo(ctx context.Context, in *GetNodeInfoRequest, opts ...grpc.CallOption) (*NodeInfoResponse, error) {
	out := new(NodeInfoResponse)
	err := c.cc.Invoke(ctx, NodeService_GetNodeInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) GetNetworkInfo(ctx context.Context, in *GetNetworkInfoRequest, opts ...grpc.CallOption) (*NetworkInfoResponse, error) {
	out := new(NetworkInfoResponse)
	err := c.cc.Invoke(ctx, NodeService_GetNetworkInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) GetTokenInfo(ctx context.Context, in *GetTokenInfoRequest, opts ...grpc.CallOption) (*TokenInfoResponse, error) {
	out := new(TokenInfoResponse)
	err := c.cc.Invoke(ctx, NodeService_GetTokenInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServiceServer is the server API for NodeService service.
// All implementations must embed UnimplementedNodeServiceServer
// for forward compatibility
type NodeServiceServer interface {
	GetFrames(context.Context, *GetFramesRequest) (*FramesResponse, error)
	GetFrameInfo(context.Context, *GetFrameInfoRequest) (*FrameInfoResponse, error)
	GetPeerInfo(context.Context, *GetPeerInfoRequest) (*PeerInfoResponse, error)
	GetNodeInfo(context.Context, *GetNodeInfoRequest) (*NodeInfoResponse, error)
	GetNetworkInfo(context.Context, *GetNetworkInfoRequest) (*NetworkInfoResponse, error)
	GetTokenInfo(context.Context, *GetTokenInfoRequest) (*TokenInfoResponse, error)
	mustEmbedUnimplementedNodeServiceServer()
}

// UnimplementedNodeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNodeServiceServer struct {
}

func (UnimplementedNodeServiceServer) GetFrames(context.Context, *GetFramesRequest) (*FramesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFrames not implemented")
}
func (UnimplementedNodeServiceServer) GetFrameInfo(context.Context, *GetFrameInfoRequest) (*FrameInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFrameInfo not implemented")
}
func (UnimplementedNodeServiceServer) GetPeerInfo(context.Context, *GetPeerInfoRequest) (*PeerInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPeerInfo not implemented")
}
func (UnimplementedNodeServiceServer) GetNodeInfo(context.Context, *GetNodeInfoRequest) (*NodeInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeInfo not implemented")
}
func (UnimplementedNodeServiceServer) GetNetworkInfo(context.Context, *GetNetworkInfoRequest) (*NetworkInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNetworkInfo not implemented")
}
func (UnimplementedNodeServiceServer) GetTokenInfo(context.Context, *GetTokenInfoRequest) (*TokenInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenInfo not implemented")
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

func _NodeService_GetFrames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFramesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).GetFrames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_GetFrames_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).GetFrames(ctx, req.(*GetFramesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_GetFrameInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFrameInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).GetFrameInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_GetFrameInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).GetFrameInfo(ctx, req.(*GetFrameInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_GetPeerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPeerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).GetPeerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_GetPeerInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).GetPeerInfo(ctx, req.(*GetPeerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_GetNodeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).GetNodeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_GetNodeInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).GetNodeInfo(ctx, req.(*GetNodeInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_GetNetworkInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNetworkInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).GetNetworkInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_GetNetworkInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).GetNetworkInfo(ctx, req.(*GetNetworkInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_GetTokenInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).GetTokenInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_GetTokenInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).GetTokenInfo(ctx, req.(*GetTokenInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeService_ServiceDesc is the grpc.ServiceDesc for NodeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "quilibrium.node.node.pb.NodeService",
	HandlerType: (*NodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFrames",
			Handler:    _NodeService_GetFrames_Handler,
		},
		{
			MethodName: "GetFrameInfo",
			Handler:    _NodeService_GetFrameInfo_Handler,
		},
		{
			MethodName: "GetPeerInfo",
			Handler:    _NodeService_GetPeerInfo_Handler,
		},
		{
			MethodName: "GetNodeInfo",
			Handler:    _NodeService_GetNodeInfo_Handler,
		},
		{
			MethodName: "GetNetworkInfo",
			Handler:    _NodeService_GetNetworkInfo_Handler,
		},
		{
			MethodName: "GetTokenInfo",
			Handler:    _NodeService_GetTokenInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node.proto",
}

const (
	NodeStats_PutNodeInfo_FullMethodName = "/quilibrium.node.node.pb.NodeStats/PutNodeInfo"
	NodeStats_PutPeerInfo_FullMethodName = "/quilibrium.node.node.pb.NodeStats/PutPeerInfo"
)

// NodeStatsClient is the client API for NodeStats service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeStatsClient interface {
	PutNodeInfo(ctx context.Context, in *PutNodeInfoRequest, opts ...grpc.CallOption) (*PutResponse, error)
	PutPeerInfo(ctx context.Context, in *PutPeerInfoRequest, opts ...grpc.CallOption) (*PutResponse, error)
}

type nodeStatsClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeStatsClient(cc grpc.ClientConnInterface) NodeStatsClient {
	return &nodeStatsClient{cc}
}

func (c *nodeStatsClient) PutNodeInfo(ctx context.Context, in *PutNodeInfoRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, NodeStats_PutNodeInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeStatsClient) PutPeerInfo(ctx context.Context, in *PutPeerInfoRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, NodeStats_PutPeerInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeStatsServer is the server API for NodeStats service.
// All implementations must embed UnimplementedNodeStatsServer
// for forward compatibility
type NodeStatsServer interface {
	PutNodeInfo(context.Context, *PutNodeInfoRequest) (*PutResponse, error)
	PutPeerInfo(context.Context, *PutPeerInfoRequest) (*PutResponse, error)
	mustEmbedUnimplementedNodeStatsServer()
}

// UnimplementedNodeStatsServer must be embedded to have forward compatible implementations.
type UnimplementedNodeStatsServer struct {
}

func (UnimplementedNodeStatsServer) PutNodeInfo(context.Context, *PutNodeInfoRequest) (*PutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutNodeInfo not implemented")
}
func (UnimplementedNodeStatsServer) PutPeerInfo(context.Context, *PutPeerInfoRequest) (*PutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutPeerInfo not implemented")
}
func (UnimplementedNodeStatsServer) mustEmbedUnimplementedNodeStatsServer() {}

// UnsafeNodeStatsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeStatsServer will
// result in compilation errors.
type UnsafeNodeStatsServer interface {
	mustEmbedUnimplementedNodeStatsServer()
}

func RegisterNodeStatsServer(s grpc.ServiceRegistrar, srv NodeStatsServer) {
	s.RegisterService(&NodeStats_ServiceDesc, srv)
}

func _NodeStats_PutNodeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutNodeInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeStatsServer).PutNodeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeStats_PutNodeInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeStatsServer).PutNodeInfo(ctx, req.(*PutNodeInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeStats_PutPeerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutPeerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeStatsServer).PutPeerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeStats_PutPeerInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeStatsServer).PutPeerInfo(ctx, req.(*PutPeerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeStats_ServiceDesc is the grpc.ServiceDesc for NodeStats service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeStats_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "quilibrium.node.node.pb.NodeStats",
	HandlerType: (*NodeStatsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutNodeInfo",
			Handler:    _NodeStats_PutNodeInfo_Handler,
		},
		{
			MethodName: "PutPeerInfo",
			Handler:    _NodeStats_PutPeerInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node.proto",
}
