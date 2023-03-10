// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: proto/proto.proto

package HashTable

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

// HashTableClient is the client API for HashTable service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HashTableClient interface {
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type hashTableClient struct {
	cc grpc.ClientConnInterface
}

func NewHashTableClient(cc grpc.ClientConnInterface) HashTableClient {
	return &hashTableClient{cc}
}

func (c *hashTableClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, "/HashTable.HashTable/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hashTableClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/HashTable.HashTable/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HashTableServer is the server API for HashTable service.
// All implementations must embed UnimplementedHashTableServer
// for forward compatibility
type HashTableServer interface {
	Put(context.Context, *PutRequest) (*PutResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	mustEmbedUnimplementedHashTableServer()
}

// UnimplementedHashTableServer must be embedded to have forward compatible implementations.
type UnimplementedHashTableServer struct {
}

func (UnimplementedHashTableServer) Put(context.Context, *PutRequest) (*PutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedHashTableServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedHashTableServer) mustEmbedUnimplementedHashTableServer() {}

// UnsafeHashTableServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HashTableServer will
// result in compilation errors.
type UnsafeHashTableServer interface {
	mustEmbedUnimplementedHashTableServer()
}

func RegisterHashTableServer(s grpc.ServiceRegistrar, srv HashTableServer) {
	s.RegisterService(&HashTable_ServiceDesc, srv)
}

func _HashTable_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HashTableServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HashTable.HashTable/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HashTableServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HashTable_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HashTableServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HashTable.HashTable/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HashTableServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HashTable_ServiceDesc is the grpc.ServiceDesc for HashTable service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HashTable_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "HashTable.HashTable",
	HandlerType: (*HashTableServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _HashTable_Put_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _HashTable_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/proto.proto",
}
