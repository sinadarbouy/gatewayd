// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: plugin/v1/plugin.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GatewayDPluginServiceClient is the client API for GatewayDPluginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayDPluginServiceClient interface {
	// GetPluginConfig returns the plugin config upon registration
	GetPluginConfig(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error)
	// OnConfigLoaded is called when the config is loaded from any config provider
	OnConfigLoaded(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error)
	OnNewLogger(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error)
	OnNewPool(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error)
	OnNewProxy(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error)
	OnNewServer(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error)
	OnSignal(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error)
	OnRun(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error)
	OnBooting(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error)
	OnBooted(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error)
	OnOpening(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error)
	OnOpened(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error)
	OnClosing(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error)
	OnClosed(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error)
	OnTraffic(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error)
	OnIngressTraffic(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error)
	OnEgressTraffic(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error)
	OnShutdown(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error)
	OnTick(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error)
	OnNewClient(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error)
}

type gatewayDPluginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayDPluginServiceClient(cc grpc.ClientConnInterface) GatewayDPluginServiceClient {
	return &gatewayDPluginServiceClient{cc}
}

func (c *gatewayDPluginServiceClient) GetPluginConfig(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error) {
	out := new(structpb.Struct)
	err := c.cc.Invoke(ctx, "/plugin.v1.GatewayDPluginService/GetPluginConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayDPluginServiceClient) OnConfigLoaded(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error) {
	out := new(structpb.Struct)
	err := c.cc.Invoke(ctx, "/plugin.v1.GatewayDPluginService/OnConfigLoaded", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayDPluginServiceClient) OnNewLogger(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error) {
	out := new(structpb.Struct)
	err := c.cc.Invoke(ctx, "/plugin.v1.GatewayDPluginService/OnNewLogger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayDPluginServiceClient) OnNewPool(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error) {
	out := new(structpb.Struct)
	err := c.cc.Invoke(ctx, "/plugin.v1.GatewayDPluginService/OnNewPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayDPluginServiceClient) OnNewProxy(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error) {
	out := new(structpb.Struct)
	err := c.cc.Invoke(ctx, "/plugin.v1.GatewayDPluginService/OnNewProxy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayDPluginServiceClient) OnNewServer(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error) {
	out := new(structpb.Struct)
	err := c.cc.Invoke(ctx, "/plugin.v1.GatewayDPluginService/OnNewServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayDPluginServiceClient) OnSignal(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error) {
	out := new(structpb.Struct)
	err := c.cc.Invoke(ctx, "/plugin.v1.GatewayDPluginService/OnSignal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayDPluginServiceClient) OnRun(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error) {
	out := new(structpb.Struct)
	err := c.cc.Invoke(ctx, "/plugin.v1.GatewayDPluginService/OnRun", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayDPluginServiceClient) OnBooting(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error) {
	out := new(structpb.Struct)
	err := c.cc.Invoke(ctx, "/plugin.v1.GatewayDPluginService/OnBooting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayDPluginServiceClient) OnBooted(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error) {
	out := new(structpb.Struct)
	err := c.cc.Invoke(ctx, "/plugin.v1.GatewayDPluginService/OnBooted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayDPluginServiceClient) OnOpening(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error) {
	out := new(structpb.Struct)
	err := c.cc.Invoke(ctx, "/plugin.v1.GatewayDPluginService/OnOpening", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayDPluginServiceClient) OnOpened(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error) {
	out := new(structpb.Struct)
	err := c.cc.Invoke(ctx, "/plugin.v1.GatewayDPluginService/OnOpened", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayDPluginServiceClient) OnClosing(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error) {
	out := new(structpb.Struct)
	err := c.cc.Invoke(ctx, "/plugin.v1.GatewayDPluginService/OnClosing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayDPluginServiceClient) OnClosed(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error) {
	out := new(structpb.Struct)
	err := c.cc.Invoke(ctx, "/plugin.v1.GatewayDPluginService/OnClosed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayDPluginServiceClient) OnTraffic(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error) {
	out := new(structpb.Struct)
	err := c.cc.Invoke(ctx, "/plugin.v1.GatewayDPluginService/OnTraffic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayDPluginServiceClient) OnIngressTraffic(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error) {
	out := new(structpb.Struct)
	err := c.cc.Invoke(ctx, "/plugin.v1.GatewayDPluginService/OnIngressTraffic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayDPluginServiceClient) OnEgressTraffic(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error) {
	out := new(structpb.Struct)
	err := c.cc.Invoke(ctx, "/plugin.v1.GatewayDPluginService/OnEgressTraffic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayDPluginServiceClient) OnShutdown(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error) {
	out := new(structpb.Struct)
	err := c.cc.Invoke(ctx, "/plugin.v1.GatewayDPluginService/OnShutdown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayDPluginServiceClient) OnTick(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error) {
	out := new(structpb.Struct)
	err := c.cc.Invoke(ctx, "/plugin.v1.GatewayDPluginService/OnTick", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayDPluginServiceClient) OnNewClient(ctx context.Context, in *structpb.Struct, opts ...grpc.CallOption) (*structpb.Struct, error) {
	out := new(structpb.Struct)
	err := c.cc.Invoke(ctx, "/plugin.v1.GatewayDPluginService/OnNewClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayDPluginServiceServer is the server API for GatewayDPluginService service.
// All implementations must embed UnimplementedGatewayDPluginServiceServer
// for forward compatibility
type GatewayDPluginServiceServer interface {
	// GetPluginConfig returns the plugin config upon registration
	GetPluginConfig(context.Context, *structpb.Struct) (*structpb.Struct, error)
	// OnConfigLoaded is called when the config is loaded from any config provider
	OnConfigLoaded(context.Context, *structpb.Struct) (*structpb.Struct, error)
	OnNewLogger(context.Context, *structpb.Struct) (*structpb.Struct, error)
	OnNewPool(context.Context, *structpb.Struct) (*structpb.Struct, error)
	OnNewProxy(context.Context, *structpb.Struct) (*structpb.Struct, error)
	OnNewServer(context.Context, *structpb.Struct) (*structpb.Struct, error)
	OnSignal(context.Context, *structpb.Struct) (*structpb.Struct, error)
	OnRun(context.Context, *structpb.Struct) (*structpb.Struct, error)
	OnBooting(context.Context, *structpb.Struct) (*structpb.Struct, error)
	OnBooted(context.Context, *structpb.Struct) (*structpb.Struct, error)
	OnOpening(context.Context, *structpb.Struct) (*structpb.Struct, error)
	OnOpened(context.Context, *structpb.Struct) (*structpb.Struct, error)
	OnClosing(context.Context, *structpb.Struct) (*structpb.Struct, error)
	OnClosed(context.Context, *structpb.Struct) (*structpb.Struct, error)
	OnTraffic(context.Context, *structpb.Struct) (*structpb.Struct, error)
	OnIngressTraffic(context.Context, *structpb.Struct) (*structpb.Struct, error)
	OnEgressTraffic(context.Context, *structpb.Struct) (*structpb.Struct, error)
	OnShutdown(context.Context, *structpb.Struct) (*structpb.Struct, error)
	OnTick(context.Context, *structpb.Struct) (*structpb.Struct, error)
	OnNewClient(context.Context, *structpb.Struct) (*structpb.Struct, error)
	mustEmbedUnimplementedGatewayDPluginServiceServer()
}

// UnimplementedGatewayDPluginServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayDPluginServiceServer struct {
}

func (UnimplementedGatewayDPluginServiceServer) GetPluginConfig(context.Context, *structpb.Struct) (*structpb.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPluginConfig not implemented")
}
func (UnimplementedGatewayDPluginServiceServer) OnConfigLoaded(context.Context, *structpb.Struct) (*structpb.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnConfigLoaded not implemented")
}
func (UnimplementedGatewayDPluginServiceServer) OnNewLogger(context.Context, *structpb.Struct) (*structpb.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnNewLogger not implemented")
}
func (UnimplementedGatewayDPluginServiceServer) OnNewPool(context.Context, *structpb.Struct) (*structpb.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnNewPool not implemented")
}
func (UnimplementedGatewayDPluginServiceServer) OnNewProxy(context.Context, *structpb.Struct) (*structpb.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnNewProxy not implemented")
}
func (UnimplementedGatewayDPluginServiceServer) OnNewServer(context.Context, *structpb.Struct) (*structpb.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnNewServer not implemented")
}
func (UnimplementedGatewayDPluginServiceServer) OnSignal(context.Context, *structpb.Struct) (*structpb.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnSignal not implemented")
}
func (UnimplementedGatewayDPluginServiceServer) OnRun(context.Context, *structpb.Struct) (*structpb.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnRun not implemented")
}
func (UnimplementedGatewayDPluginServiceServer) OnBooting(context.Context, *structpb.Struct) (*structpb.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnBooting not implemented")
}
func (UnimplementedGatewayDPluginServiceServer) OnBooted(context.Context, *structpb.Struct) (*structpb.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnBooted not implemented")
}
func (UnimplementedGatewayDPluginServiceServer) OnOpening(context.Context, *structpb.Struct) (*structpb.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnOpening not implemented")
}
func (UnimplementedGatewayDPluginServiceServer) OnOpened(context.Context, *structpb.Struct) (*structpb.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnOpened not implemented")
}
func (UnimplementedGatewayDPluginServiceServer) OnClosing(context.Context, *structpb.Struct) (*structpb.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnClosing not implemented")
}
func (UnimplementedGatewayDPluginServiceServer) OnClosed(context.Context, *structpb.Struct) (*structpb.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnClosed not implemented")
}
func (UnimplementedGatewayDPluginServiceServer) OnTraffic(context.Context, *structpb.Struct) (*structpb.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnTraffic not implemented")
}
func (UnimplementedGatewayDPluginServiceServer) OnIngressTraffic(context.Context, *structpb.Struct) (*structpb.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnIngressTraffic not implemented")
}
func (UnimplementedGatewayDPluginServiceServer) OnEgressTraffic(context.Context, *structpb.Struct) (*structpb.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnEgressTraffic not implemented")
}
func (UnimplementedGatewayDPluginServiceServer) OnShutdown(context.Context, *structpb.Struct) (*structpb.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnShutdown not implemented")
}
func (UnimplementedGatewayDPluginServiceServer) OnTick(context.Context, *structpb.Struct) (*structpb.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnTick not implemented")
}
func (UnimplementedGatewayDPluginServiceServer) OnNewClient(context.Context, *structpb.Struct) (*structpb.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnNewClient not implemented")
}
func (UnimplementedGatewayDPluginServiceServer) mustEmbedUnimplementedGatewayDPluginServiceServer() {}

// UnsafeGatewayDPluginServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayDPluginServiceServer will
// result in compilation errors.
type UnsafeGatewayDPluginServiceServer interface {
	mustEmbedUnimplementedGatewayDPluginServiceServer()
}

func RegisterGatewayDPluginServiceServer(s grpc.ServiceRegistrar, srv GatewayDPluginServiceServer) {
	s.RegisterService(&GatewayDPluginService_ServiceDesc, srv)
}

func _GatewayDPluginService_GetPluginConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(structpb.Struct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayDPluginServiceServer).GetPluginConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugin.v1.GatewayDPluginService/GetPluginConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayDPluginServiceServer).GetPluginConfig(ctx, req.(*structpb.Struct))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayDPluginService_OnConfigLoaded_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(structpb.Struct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayDPluginServiceServer).OnConfigLoaded(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugin.v1.GatewayDPluginService/OnConfigLoaded",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayDPluginServiceServer).OnConfigLoaded(ctx, req.(*structpb.Struct))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayDPluginService_OnNewLogger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(structpb.Struct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayDPluginServiceServer).OnNewLogger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugin.v1.GatewayDPluginService/OnNewLogger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayDPluginServiceServer).OnNewLogger(ctx, req.(*structpb.Struct))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayDPluginService_OnNewPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(structpb.Struct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayDPluginServiceServer).OnNewPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugin.v1.GatewayDPluginService/OnNewPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayDPluginServiceServer).OnNewPool(ctx, req.(*structpb.Struct))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayDPluginService_OnNewProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(structpb.Struct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayDPluginServiceServer).OnNewProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugin.v1.GatewayDPluginService/OnNewProxy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayDPluginServiceServer).OnNewProxy(ctx, req.(*structpb.Struct))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayDPluginService_OnNewServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(structpb.Struct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayDPluginServiceServer).OnNewServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugin.v1.GatewayDPluginService/OnNewServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayDPluginServiceServer).OnNewServer(ctx, req.(*structpb.Struct))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayDPluginService_OnSignal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(structpb.Struct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayDPluginServiceServer).OnSignal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugin.v1.GatewayDPluginService/OnSignal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayDPluginServiceServer).OnSignal(ctx, req.(*structpb.Struct))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayDPluginService_OnRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(structpb.Struct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayDPluginServiceServer).OnRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugin.v1.GatewayDPluginService/OnRun",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayDPluginServiceServer).OnRun(ctx, req.(*structpb.Struct))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayDPluginService_OnBooting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(structpb.Struct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayDPluginServiceServer).OnBooting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugin.v1.GatewayDPluginService/OnBooting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayDPluginServiceServer).OnBooting(ctx, req.(*structpb.Struct))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayDPluginService_OnBooted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(structpb.Struct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayDPluginServiceServer).OnBooted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugin.v1.GatewayDPluginService/OnBooted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayDPluginServiceServer).OnBooted(ctx, req.(*structpb.Struct))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayDPluginService_OnOpening_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(structpb.Struct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayDPluginServiceServer).OnOpening(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugin.v1.GatewayDPluginService/OnOpening",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayDPluginServiceServer).OnOpening(ctx, req.(*structpb.Struct))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayDPluginService_OnOpened_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(structpb.Struct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayDPluginServiceServer).OnOpened(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugin.v1.GatewayDPluginService/OnOpened",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayDPluginServiceServer).OnOpened(ctx, req.(*structpb.Struct))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayDPluginService_OnClosing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(structpb.Struct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayDPluginServiceServer).OnClosing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugin.v1.GatewayDPluginService/OnClosing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayDPluginServiceServer).OnClosing(ctx, req.(*structpb.Struct))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayDPluginService_OnClosed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(structpb.Struct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayDPluginServiceServer).OnClosed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugin.v1.GatewayDPluginService/OnClosed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayDPluginServiceServer).OnClosed(ctx, req.(*structpb.Struct))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayDPluginService_OnTraffic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(structpb.Struct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayDPluginServiceServer).OnTraffic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugin.v1.GatewayDPluginService/OnTraffic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayDPluginServiceServer).OnTraffic(ctx, req.(*structpb.Struct))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayDPluginService_OnIngressTraffic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(structpb.Struct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayDPluginServiceServer).OnIngressTraffic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugin.v1.GatewayDPluginService/OnIngressTraffic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayDPluginServiceServer).OnIngressTraffic(ctx, req.(*structpb.Struct))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayDPluginService_OnEgressTraffic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(structpb.Struct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayDPluginServiceServer).OnEgressTraffic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugin.v1.GatewayDPluginService/OnEgressTraffic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayDPluginServiceServer).OnEgressTraffic(ctx, req.(*structpb.Struct))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayDPluginService_OnShutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(structpb.Struct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayDPluginServiceServer).OnShutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugin.v1.GatewayDPluginService/OnShutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayDPluginServiceServer).OnShutdown(ctx, req.(*structpb.Struct))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayDPluginService_OnTick_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(structpb.Struct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayDPluginServiceServer).OnTick(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugin.v1.GatewayDPluginService/OnTick",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayDPluginServiceServer).OnTick(ctx, req.(*structpb.Struct))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayDPluginService_OnNewClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(structpb.Struct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayDPluginServiceServer).OnNewClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugin.v1.GatewayDPluginService/OnNewClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayDPluginServiceServer).OnNewClient(ctx, req.(*structpb.Struct))
	}
	return interceptor(ctx, in, info, handler)
}

// GatewayDPluginService_ServiceDesc is the grpc.ServiceDesc for GatewayDPluginService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GatewayDPluginService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "plugin.v1.GatewayDPluginService",
	HandlerType: (*GatewayDPluginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPluginConfig",
			Handler:    _GatewayDPluginService_GetPluginConfig_Handler,
		},
		{
			MethodName: "OnConfigLoaded",
			Handler:    _GatewayDPluginService_OnConfigLoaded_Handler,
		},
		{
			MethodName: "OnNewLogger",
			Handler:    _GatewayDPluginService_OnNewLogger_Handler,
		},
		{
			MethodName: "OnNewPool",
			Handler:    _GatewayDPluginService_OnNewPool_Handler,
		},
		{
			MethodName: "OnNewProxy",
			Handler:    _GatewayDPluginService_OnNewProxy_Handler,
		},
		{
			MethodName: "OnNewServer",
			Handler:    _GatewayDPluginService_OnNewServer_Handler,
		},
		{
			MethodName: "OnSignal",
			Handler:    _GatewayDPluginService_OnSignal_Handler,
		},
		{
			MethodName: "OnRun",
			Handler:    _GatewayDPluginService_OnRun_Handler,
		},
		{
			MethodName: "OnBooting",
			Handler:    _GatewayDPluginService_OnBooting_Handler,
		},
		{
			MethodName: "OnBooted",
			Handler:    _GatewayDPluginService_OnBooted_Handler,
		},
		{
			MethodName: "OnOpening",
			Handler:    _GatewayDPluginService_OnOpening_Handler,
		},
		{
			MethodName: "OnOpened",
			Handler:    _GatewayDPluginService_OnOpened_Handler,
		},
		{
			MethodName: "OnClosing",
			Handler:    _GatewayDPluginService_OnClosing_Handler,
		},
		{
			MethodName: "OnClosed",
			Handler:    _GatewayDPluginService_OnClosed_Handler,
		},
		{
			MethodName: "OnTraffic",
			Handler:    _GatewayDPluginService_OnTraffic_Handler,
		},
		{
			MethodName: "OnIngressTraffic",
			Handler:    _GatewayDPluginService_OnIngressTraffic_Handler,
		},
		{
			MethodName: "OnEgressTraffic",
			Handler:    _GatewayDPluginService_OnEgressTraffic_Handler,
		},
		{
			MethodName: "OnShutdown",
			Handler:    _GatewayDPluginService_OnShutdown_Handler,
		},
		{
			MethodName: "OnTick",
			Handler:    _GatewayDPluginService_OnTick_Handler,
		},
		{
			MethodName: "OnNewClient",
			Handler:    _GatewayDPluginService_OnNewClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plugin/v1/plugin.proto",
}