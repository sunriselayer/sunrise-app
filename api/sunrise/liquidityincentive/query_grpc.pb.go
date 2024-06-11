// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: sunrise/liquidityincentive/query.proto

package liquidityincentive

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
	Query_Params_FullMethodName = "/sunrise.liquidityincentive.Query/Params"
	Query_Epoch_FullMethodName  = "/sunrise.liquidityincentive.Query/Epoch"
	Query_Epochs_FullMethodName = "/sunrise.liquidityincentive.Query/Epochs"
	Query_Gauge_FullMethodName  = "/sunrise.liquidityincentive.Query/Gauge"
	Query_Gauges_FullMethodName = "/sunrise.liquidityincentive.Query/Gauges"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of Epoch items.
	Epoch(ctx context.Context, in *QueryEpochRequest, opts ...grpc.CallOption) (*QueryEpochResponse, error)
	Epochs(ctx context.Context, in *QueryEpochsRequest, opts ...grpc.CallOption) (*QueryEpochsResponse, error)
	// Queries a list of Gauge items.
	Gauge(ctx context.Context, in *QueryGaugeRequest, opts ...grpc.CallOption) (*QueryGaugeResponse, error)
	Gauges(ctx context.Context, in *QueryGaugesRequest, opts ...grpc.CallOption) (*QueryGaugesResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Epoch(ctx context.Context, in *QueryEpochRequest, opts ...grpc.CallOption) (*QueryEpochResponse, error) {
	out := new(QueryEpochResponse)
	err := c.cc.Invoke(ctx, Query_Epoch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Epochs(ctx context.Context, in *QueryEpochsRequest, opts ...grpc.CallOption) (*QueryEpochsResponse, error) {
	out := new(QueryEpochsResponse)
	err := c.cc.Invoke(ctx, Query_Epochs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Gauge(ctx context.Context, in *QueryGaugeRequest, opts ...grpc.CallOption) (*QueryGaugeResponse, error) {
	out := new(QueryGaugeResponse)
	err := c.cc.Invoke(ctx, Query_Gauge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Gauges(ctx context.Context, in *QueryGaugesRequest, opts ...grpc.CallOption) (*QueryGaugesResponse, error) {
	out := new(QueryGaugesResponse)
	err := c.cc.Invoke(ctx, Query_Gauges_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of Epoch items.
	Epoch(context.Context, *QueryEpochRequest) (*QueryEpochResponse, error)
	Epochs(context.Context, *QueryEpochsRequest) (*QueryEpochsResponse, error)
	// Queries a list of Gauge items.
	Gauge(context.Context, *QueryGaugeRequest) (*QueryGaugeResponse, error)
	Gauges(context.Context, *QueryGaugesRequest) (*QueryGaugesResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) Epoch(context.Context, *QueryEpochRequest) (*QueryEpochResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Epoch not implemented")
}
func (UnimplementedQueryServer) Epochs(context.Context, *QueryEpochsRequest) (*QueryEpochsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Epochs not implemented")
}
func (UnimplementedQueryServer) Gauge(context.Context, *QueryGaugeRequest) (*QueryGaugeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Gauge not implemented")
}
func (UnimplementedQueryServer) Gauges(context.Context, *QueryGaugesRequest) (*QueryGaugesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Gauges not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Epoch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEpochRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Epoch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Epoch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Epoch(ctx, req.(*QueryEpochRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Epochs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEpochsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Epochs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Epochs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Epochs(ctx, req.(*QueryEpochsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Gauge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGaugeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Gauge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Gauge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Gauge(ctx, req.(*QueryGaugeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Gauges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGaugesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Gauges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Gauges_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Gauges(ctx, req.(*QueryGaugesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sunrise.liquidityincentive.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Epoch",
			Handler:    _Query_Epoch_Handler,
		},
		{
			MethodName: "Epochs",
			Handler:    _Query_Epochs_Handler,
		},
		{
			MethodName: "Gauge",
			Handler:    _Query_Gauge_Handler,
		},
		{
			MethodName: "Gauges",
			Handler:    _Query_Gauges_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sunrise/liquidityincentive/query.proto",
}
