// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package vk_messages

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

// MessagesServiceClient is the client API for MessagesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessagesServiceClient interface {
	GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*GetMessagesResponse, error)
}

type messagesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessagesServiceClient(cc grpc.ClientConnInterface) MessagesServiceClient {
	return &messagesServiceClient{cc}
}

func (c *messagesServiceClient) GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*GetMessagesResponse, error) {
	out := new(GetMessagesResponse)
	err := c.cc.Invoke(ctx, "/MessagesService/GetMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessagesServiceServer is the server API for MessagesService service.
// All implementations must embed UnimplementedMessagesServiceServer
// for forward compatibility
type MessagesServiceServer interface {
	GetMessages(context.Context, *GetMessagesRequest) (*GetMessagesResponse, error)
	mustEmbedUnimplementedMessagesServiceServer()
}

// UnimplementedMessagesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMessagesServiceServer struct {
}

func (UnimplementedMessagesServiceServer) GetMessages(context.Context, *GetMessagesRequest) (*GetMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}
func (UnimplementedMessagesServiceServer) mustEmbedUnimplementedMessagesServiceServer() {}

// UnsafeMessagesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessagesServiceServer will
// result in compilation errors.
type UnsafeMessagesServiceServer interface {
	mustEmbedUnimplementedMessagesServiceServer()
}

func RegisterMessagesServiceServer(s grpc.ServiceRegistrar, srv MessagesServiceServer) {
	s.RegisterService(&MessagesService_ServiceDesc, srv)
}

func _MessagesService_GetMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagesServiceServer).GetMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessagesService/GetMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagesServiceServer).GetMessages(ctx, req.(*GetMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MessagesService_ServiceDesc is the grpc.ServiceDesc for MessagesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessagesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MessagesService",
	HandlerType: (*MessagesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMessages",
			Handler:    _MessagesService_GetMessages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "messages/api.proto",
}
