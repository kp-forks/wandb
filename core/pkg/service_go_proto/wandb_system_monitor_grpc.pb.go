// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.23.4
// source: wandb/proto/wandb_system_monitor.proto

package service_go_proto

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
	SystemMonitorService_GetStats_FullMethodName    = "/wandb_internal.SystemMonitorService/GetStats"
	SystemMonitorService_GetMetadata_FullMethodName = "/wandb_internal.SystemMonitorService/GetMetadata"
	SystemMonitorService_TearDown_FullMethodName    = "/wandb_internal.SystemMonitorService/TearDown"
)

// SystemMonitorServiceClient is the client API for SystemMonitorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// SystemMonitorService gRPC service.
//
// This service is used to collect system metrics from the host machine.
type SystemMonitorServiceClient interface {
	// GetStats samples system metrics.
	GetStats(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (*GetStatsResponse, error)
	// GetMetadata returns static metadata about the system.
	GetMetadata(ctx context.Context, in *GetMetadataRequest, opts ...grpc.CallOption) (*GetMetadataResponse, error)
	// TearDown instructs the system monitor to shut down.
	TearDown(ctx context.Context, in *TearDownRequest, opts ...grpc.CallOption) (*TearDownResponse, error)
}

type systemMonitorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemMonitorServiceClient(cc grpc.ClientConnInterface) SystemMonitorServiceClient {
	return &systemMonitorServiceClient{cc}
}

func (c *systemMonitorServiceClient) GetStats(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (*GetStatsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetStatsResponse)
	err := c.cc.Invoke(ctx, SystemMonitorService_GetStats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemMonitorServiceClient) GetMetadata(ctx context.Context, in *GetMetadataRequest, opts ...grpc.CallOption) (*GetMetadataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMetadataResponse)
	err := c.cc.Invoke(ctx, SystemMonitorService_GetMetadata_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemMonitorServiceClient) TearDown(ctx context.Context, in *TearDownRequest, opts ...grpc.CallOption) (*TearDownResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TearDownResponse)
	err := c.cc.Invoke(ctx, SystemMonitorService_TearDown_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemMonitorServiceServer is the server API for SystemMonitorService service.
// All implementations must embed UnimplementedSystemMonitorServiceServer
// for forward compatibility.
//
// SystemMonitorService gRPC service.
//
// This service is used to collect system metrics from the host machine.
type SystemMonitorServiceServer interface {
	// GetStats samples system metrics.
	GetStats(context.Context, *GetStatsRequest) (*GetStatsResponse, error)
	// GetMetadata returns static metadata about the system.
	GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error)
	// TearDown instructs the system monitor to shut down.
	TearDown(context.Context, *TearDownRequest) (*TearDownResponse, error)
	mustEmbedUnimplementedSystemMonitorServiceServer()
}

// UnimplementedSystemMonitorServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSystemMonitorServiceServer struct{}

func (UnimplementedSystemMonitorServiceServer) GetStats(context.Context, *GetStatsRequest) (*GetStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStats not implemented")
}
func (UnimplementedSystemMonitorServiceServer) GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetadata not implemented")
}
func (UnimplementedSystemMonitorServiceServer) TearDown(context.Context, *TearDownRequest) (*TearDownResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TearDown not implemented")
}
func (UnimplementedSystemMonitorServiceServer) mustEmbedUnimplementedSystemMonitorServiceServer() {}
func (UnimplementedSystemMonitorServiceServer) testEmbeddedByValue()                              {}

// UnsafeSystemMonitorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemMonitorServiceServer will
// result in compilation errors.
type UnsafeSystemMonitorServiceServer interface {
	mustEmbedUnimplementedSystemMonitorServiceServer()
}

func RegisterSystemMonitorServiceServer(s grpc.ServiceRegistrar, srv SystemMonitorServiceServer) {
	// If the following call pancis, it indicates UnimplementedSystemMonitorServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SystemMonitorService_ServiceDesc, srv)
}

func _SystemMonitorService_GetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemMonitorServiceServer).GetStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemMonitorService_GetStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemMonitorServiceServer).GetStats(ctx, req.(*GetStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemMonitorService_GetMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemMonitorServiceServer).GetMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemMonitorService_GetMetadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemMonitorServiceServer).GetMetadata(ctx, req.(*GetMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemMonitorService_TearDown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TearDownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemMonitorServiceServer).TearDown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemMonitorService_TearDown_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemMonitorServiceServer).TearDown(ctx, req.(*TearDownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SystemMonitorService_ServiceDesc is the grpc.ServiceDesc for SystemMonitorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SystemMonitorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wandb_internal.SystemMonitorService",
	HandlerType: (*SystemMonitorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStats",
			Handler:    _SystemMonitorService_GetStats_Handler,
		},
		{
			MethodName: "GetMetadata",
			Handler:    _SystemMonitorService_GetMetadata_Handler,
		},
		{
			MethodName: "TearDown",
			Handler:    _SystemMonitorService_TearDown_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wandb/proto/wandb_system_monitor.proto",
}
