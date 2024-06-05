// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: sunrise/liquiditypool/query.proto

package liquiditypool

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
<<<<<<< HEAD
	Query_Params_FullMethodName             = "/sunrise.liquiditypool.Query/Params"
	Query_Pool_FullMethodName               = "/sunrise.liquiditypool.Query/Pool"
	Query_PoolAll_FullMethodName            = "/sunrise.liquiditypool.Query/PoolAll"
	Query_Position_FullMethodName           = "/sunrise.liquiditypool.Query/Position"
	Query_PositionAll_FullMethodName        = "/sunrise.liquiditypool.Query/PositionAll"
	Query_PositionsByPool_FullMethodName    = "/sunrise.liquiditypool.Query/PositionsByPool"
	Query_PositionsByAddress_FullMethodName = "/sunrise.liquiditypool.Query/PositionsByAddress"
	Query_FeesByPositionId_FullMethodName   = "/sunrise.liquiditypool.Query/FeesByPositionId"
=======
	Query_Params_FullMethodName           = "/sunrise.liquiditypool.Query/Params"
	Query_Pool_FullMethodName             = "/sunrise.liquiditypool.Query/Pool"
	Query_Pools_FullMethodName            = "/sunrise.liquiditypool.Query/Pools"
	Query_Position_FullMethodName         = "/sunrise.liquiditypool.Query/Position"
	Query_Positions_FullMethodName        = "/sunrise.liquiditypool.Query/Positions"
	Query_AddressPositions_FullMethodName = "/sunrise.liquiditypool.Query/AddressPositions"
	Query_PositionFees_FullMethodName     = "/sunrise.liquiditypool.Query/PositionFees"
>>>>>>> 3988f81665ee85f01cc5adb24fcb7beb3e5cb010
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of Pool items.
	Pool(ctx context.Context, in *QueryPoolRequest, opts ...grpc.CallOption) (*QueryPoolResponse, error)
	Pools(ctx context.Context, in *QueryPoolsRequest, opts ...grpc.CallOption) (*QueryPoolsResponse, error)
	// Queries a list of Position items.
<<<<<<< HEAD
	Position(ctx context.Context, in *QueryGetPositionRequest, opts ...grpc.CallOption) (*QueryGetPositionResponse, error)
	PositionAll(ctx context.Context, in *QueryAllPositionRequest, opts ...grpc.CallOption) (*QueryAllPositionResponse, error)
	PositionsByPool(ctx context.Context, in *QueryPositionsByPoolRequest, opts ...grpc.CallOption) (*QueryPositionsByPoolResponse, error)
	PositionsByAddress(ctx context.Context, in *QueryPositionsByAddressRequest, opts ...grpc.CallOption) (*QueryPositionsByAddressResponse, error)
	// Query fees by position id
	FeesByPositionId(ctx context.Context, in *QueryFeesByPositionIdRequest, opts ...grpc.CallOption) (*QueryFeesByPositionIdResponse, error)
=======
	Position(ctx context.Context, in *QueryPositionRequest, opts ...grpc.CallOption) (*QueryPositionResponse, error)
	Positions(ctx context.Context, in *QueryPositionsRequest, opts ...grpc.CallOption) (*QueryPositionsResponse, error)
	AddressPositions(ctx context.Context, in *QueryAddressPositionsRequest, opts ...grpc.CallOption) (*QueryAddressPositionsResponse, error)
	PositionFees(ctx context.Context, in *QueryPositionFeesRequest, opts ...grpc.CallOption) (*QueryPositionFeesResponse, error)
>>>>>>> 3988f81665ee85f01cc5adb24fcb7beb3e5cb010
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

func (c *queryClient) Pool(ctx context.Context, in *QueryPoolRequest, opts ...grpc.CallOption) (*QueryPoolResponse, error) {
	out := new(QueryPoolResponse)
	err := c.cc.Invoke(ctx, Query_Pool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Pools(ctx context.Context, in *QueryPoolsRequest, opts ...grpc.CallOption) (*QueryPoolsResponse, error) {
	out := new(QueryPoolsResponse)
	err := c.cc.Invoke(ctx, Query_Pools_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Position(ctx context.Context, in *QueryPositionRequest, opts ...grpc.CallOption) (*QueryPositionResponse, error) {
	out := new(QueryPositionResponse)
	err := c.cc.Invoke(ctx, Query_Position_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Positions(ctx context.Context, in *QueryPositionsRequest, opts ...grpc.CallOption) (*QueryPositionsResponse, error) {
	out := new(QueryPositionsResponse)
	err := c.cc.Invoke(ctx, Query_Positions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AddressPositions(ctx context.Context, in *QueryAddressPositionsRequest, opts ...grpc.CallOption) (*QueryAddressPositionsResponse, error) {
	out := new(QueryAddressPositionsResponse)
	err := c.cc.Invoke(ctx, Query_AddressPositions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PositionFees(ctx context.Context, in *QueryPositionFeesRequest, opts ...grpc.CallOption) (*QueryPositionFeesResponse, error) {
	out := new(QueryPositionFeesResponse)
	err := c.cc.Invoke(ctx, Query_PositionFees_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PositionsByPool(ctx context.Context, in *QueryPositionsByPoolRequest, opts ...grpc.CallOption) (*QueryPositionsByPoolResponse, error) {
	out := new(QueryPositionsByPoolResponse)
	err := c.cc.Invoke(ctx, Query_PositionsByPool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PositionsByAddress(ctx context.Context, in *QueryPositionsByAddressRequest, opts ...grpc.CallOption) (*QueryPositionsByAddressResponse, error) {
	out := new(QueryPositionsByAddressResponse)
	err := c.cc.Invoke(ctx, Query_PositionsByAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) FeesByPositionId(ctx context.Context, in *QueryFeesByPositionIdRequest, opts ...grpc.CallOption) (*QueryFeesByPositionIdResponse, error) {
	out := new(QueryFeesByPositionIdResponse)
	err := c.cc.Invoke(ctx, Query_FeesByPositionId_FullMethodName, in, out, opts...)
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
	// Queries a list of Pool items.
	Pool(context.Context, *QueryPoolRequest) (*QueryPoolResponse, error)
	Pools(context.Context, *QueryPoolsRequest) (*QueryPoolsResponse, error)
	// Queries a list of Position items.
<<<<<<< HEAD
	Position(context.Context, *QueryGetPositionRequest) (*QueryGetPositionResponse, error)
	PositionAll(context.Context, *QueryAllPositionRequest) (*QueryAllPositionResponse, error)
	PositionsByPool(context.Context, *QueryPositionsByPoolRequest) (*QueryPositionsByPoolResponse, error)
	PositionsByAddress(context.Context, *QueryPositionsByAddressRequest) (*QueryPositionsByAddressResponse, error)
	// Query fees by position id
	FeesByPositionId(context.Context, *QueryFeesByPositionIdRequest) (*QueryFeesByPositionIdResponse, error)
=======
	Position(context.Context, *QueryPositionRequest) (*QueryPositionResponse, error)
	Positions(context.Context, *QueryPositionsRequest) (*QueryPositionsResponse, error)
	AddressPositions(context.Context, *QueryAddressPositionsRequest) (*QueryAddressPositionsResponse, error)
	PositionFees(context.Context, *QueryPositionFeesRequest) (*QueryPositionFeesResponse, error)
>>>>>>> 3988f81665ee85f01cc5adb24fcb7beb3e5cb010
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) Pool(context.Context, *QueryPoolRequest) (*QueryPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pool not implemented")
}
func (UnimplementedQueryServer) Pools(context.Context, *QueryPoolsRequest) (*QueryPoolsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pools not implemented")
}
func (UnimplementedQueryServer) Position(context.Context, *QueryPositionRequest) (*QueryPositionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Position not implemented")
}
func (UnimplementedQueryServer) Positions(context.Context, *QueryPositionsRequest) (*QueryPositionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Positions not implemented")
}
func (UnimplementedQueryServer) AddressPositions(context.Context, *QueryAddressPositionsRequest) (*QueryAddressPositionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddressPositions not implemented")
}
func (UnimplementedQueryServer) PositionFees(context.Context, *QueryPositionFeesRequest) (*QueryPositionFeesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PositionFees not implemented")
}
func (UnimplementedQueryServer) PositionsByPool(context.Context, *QueryPositionsByPoolRequest) (*QueryPositionsByPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PositionsByPool not implemented")
}
func (UnimplementedQueryServer) PositionsByAddress(context.Context, *QueryPositionsByAddressRequest) (*QueryPositionsByAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PositionsByAddress not implemented")
}
func (UnimplementedQueryServer) FeesByPositionId(context.Context, *QueryFeesByPositionIdRequest) (*QueryFeesByPositionIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeesByPositionId not implemented")
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

func _Query_Pool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Pool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Pool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Pool(ctx, req.(*QueryPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Pools_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPoolsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Pools(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Pools_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Pools(ctx, req.(*QueryPoolsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Position_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Position(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Position_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Position(ctx, req.(*QueryPositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Positions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPositionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Positions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Positions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Positions(ctx, req.(*QueryPositionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AddressPositions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAddressPositionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AddressPositions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_AddressPositions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AddressPositions(ctx, req.(*QueryAddressPositionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PositionFees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPositionFeesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PositionFees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_PositionFees_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PositionFees(ctx, req.(*QueryPositionFeesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PositionsByPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPositionsByPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PositionsByPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_PositionsByPool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PositionsByPool(ctx, req.(*QueryPositionsByPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PositionsByAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPositionsByAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PositionsByAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_PositionsByAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PositionsByAddress(ctx, req.(*QueryPositionsByAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_FeesByPositionId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFeesByPositionIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).FeesByPositionId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_FeesByPositionId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).FeesByPositionId(ctx, req.(*QueryFeesByPositionIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sunrise.liquiditypool.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Pool",
			Handler:    _Query_Pool_Handler,
		},
		{
			MethodName: "Pools",
			Handler:    _Query_Pools_Handler,
		},
		{
			MethodName: "Position",
			Handler:    _Query_Position_Handler,
		},
		{
			MethodName: "Positions",
			Handler:    _Query_Positions_Handler,
		},
		{
			MethodName: "AddressPositions",
			Handler:    _Query_AddressPositions_Handler,
		},
		{
			MethodName: "PositionFees",
			Handler:    _Query_PositionFees_Handler,
		},
		{
			MethodName: "PositionsByPool",
			Handler:    _Query_PositionsByPool_Handler,
		},
		{
			MethodName: "PositionsByAddress",
			Handler:    _Query_PositionsByAddress_Handler,
		},
		{
			MethodName: "FeesByPositionId",
			Handler:    _Query_FeesByPositionId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sunrise/liquiditypool/query.proto",
}
