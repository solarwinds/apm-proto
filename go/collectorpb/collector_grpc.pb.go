// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: collector.proto

package collectorpb

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

// TraceCollectorClient is the client API for TraceCollector service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TraceCollectorClient interface {
	// post events (traces) to collector.
	PostEvents(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResult, error)
	// post metrics (internal heartbeats, request counters, summary, runtime or custom metrics) to collector
	PostMetrics(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResult, error)
	// post [__Init](https://github.com/librato/trace/blob/master/docs/specs/KV/init.md) message to collector. May be used by APM library to validate api_key.
	PostStatus(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResult, error)
	// get sampling and other settings for this connection.  Note the SettingsRequest requirement for HostID fields. May be used by APM library to validate api_key.
	GetSettings(ctx context.Context, in *SettingsRequest, opts ...grpc.CallOption) (*SettingsResult, error)
	// ping is used for keep-alive purpose. The APM library is expected to ping the collector if the connection has been idled for 20 seconds (by default). Take note that keep-alive should only be performed if the connection was previously healthy - last API call gave a response
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*MessageResult, error)
}

type traceCollectorClient struct {
	cc grpc.ClientConnInterface
}

func NewTraceCollectorClient(cc grpc.ClientConnInterface) TraceCollectorClient {
	return &traceCollectorClient{cc}
}

func (c *traceCollectorClient) PostEvents(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResult, error) {
	out := new(MessageResult)
	err := c.cc.Invoke(ctx, "/collector.TraceCollector/postEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceCollectorClient) PostMetrics(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResult, error) {
	out := new(MessageResult)
	err := c.cc.Invoke(ctx, "/collector.TraceCollector/postMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceCollectorClient) PostStatus(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResult, error) {
	out := new(MessageResult)
	err := c.cc.Invoke(ctx, "/collector.TraceCollector/postStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceCollectorClient) GetSettings(ctx context.Context, in *SettingsRequest, opts ...grpc.CallOption) (*SettingsResult, error) {
	out := new(SettingsResult)
	err := c.cc.Invoke(ctx, "/collector.TraceCollector/getSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceCollectorClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*MessageResult, error) {
	out := new(MessageResult)
	err := c.cc.Invoke(ctx, "/collector.TraceCollector/ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TraceCollectorServer is the server API for TraceCollector service.
// All implementations should embed UnimplementedTraceCollectorServer
// for forward compatibility
type TraceCollectorServer interface {
	// post events (traces) to collector.
	PostEvents(context.Context, *MessageRequest) (*MessageResult, error)
	// post metrics (internal heartbeats, request counters, summary, runtime or custom metrics) to collector
	PostMetrics(context.Context, *MessageRequest) (*MessageResult, error)
	// post [__Init](https://github.com/librato/trace/blob/master/docs/specs/KV/init.md) message to collector. May be used by APM library to validate api_key.
	PostStatus(context.Context, *MessageRequest) (*MessageResult, error)
	// get sampling and other settings for this connection.  Note the SettingsRequest requirement for HostID fields. May be used by APM library to validate api_key.
	GetSettings(context.Context, *SettingsRequest) (*SettingsResult, error)
	// ping is used for keep-alive purpose. The APM library is expected to ping the collector if the connection has been idled for 20 seconds (by default). Take note that keep-alive should only be performed if the connection was previously healthy - last API call gave a response
	Ping(context.Context, *PingRequest) (*MessageResult, error)
}

// UnimplementedTraceCollectorServer should be embedded to have forward compatible implementations.
type UnimplementedTraceCollectorServer struct {
}

func (UnimplementedTraceCollectorServer) PostEvents(context.Context, *MessageRequest) (*MessageResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostEvents not implemented")
}
func (UnimplementedTraceCollectorServer) PostMetrics(context.Context, *MessageRequest) (*MessageResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostMetrics not implemented")
}
func (UnimplementedTraceCollectorServer) PostStatus(context.Context, *MessageRequest) (*MessageResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostStatus not implemented")
}
func (UnimplementedTraceCollectorServer) GetSettings(context.Context, *SettingsRequest) (*SettingsResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSettings not implemented")
}
func (UnimplementedTraceCollectorServer) Ping(context.Context, *PingRequest) (*MessageResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

// UnsafeTraceCollectorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TraceCollectorServer will
// result in compilation errors.
type UnsafeTraceCollectorServer interface {
	mustEmbedUnimplementedTraceCollectorServer()
}

func RegisterTraceCollectorServer(s grpc.ServiceRegistrar, srv TraceCollectorServer) {
	s.RegisterService(&TraceCollector_ServiceDesc, srv)
}

func _TraceCollector_PostEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceCollectorServer).PostEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collector.TraceCollector/postEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceCollectorServer).PostEvents(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraceCollector_PostMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceCollectorServer).PostMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collector.TraceCollector/postMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceCollectorServer).PostMetrics(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraceCollector_PostStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceCollectorServer).PostStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collector.TraceCollector/postStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceCollectorServer).PostStatus(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraceCollector_GetSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceCollectorServer).GetSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collector.TraceCollector/getSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceCollectorServer).GetSettings(ctx, req.(*SettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraceCollector_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceCollectorServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collector.TraceCollector/ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceCollectorServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TraceCollector_ServiceDesc is the grpc.ServiceDesc for TraceCollector service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TraceCollector_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "collector.TraceCollector",
	HandlerType: (*TraceCollectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "postEvents",
			Handler:    _TraceCollector_PostEvents_Handler,
		},
		{
			MethodName: "postMetrics",
			Handler:    _TraceCollector_PostMetrics_Handler,
		},
		{
			MethodName: "postStatus",
			Handler:    _TraceCollector_PostStatus_Handler,
		},
		{
			MethodName: "getSettings",
			Handler:    _TraceCollector_GetSettings_Handler,
		},
		{
			MethodName: "ping",
			Handler:    _TraceCollector_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "collector.proto",
}
