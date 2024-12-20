// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: order.proto

package order

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
	OrderService_OrderList_FullMethodName                   = "/OrderService/OrderList"
	OrderService_OrderDetailed_FullMethodName               = "/OrderService/OrderDetailed"
	OrderService_CreateOrder_FullMethodName                 = "/OrderService/CreateOrder"
	OrderService_GetOrderToken_FullMethodName               = "/OrderService/GetOrderToken"
	OrderService_OrderRenewalLease_FullMethodName           = "/OrderService/OrderRenewalLease"
	OrderService_QueryOrderList_FullMethodName              = "/OrderService/QueryOrderList"
	OrderService_UserSettlementApplication_FullMethodName   = "/OrderService/UserSettlementApplication"
	OrderService_SettlementApplicationPage_FullMethodName   = "/OrderService/SettlementApplicationPage"
	OrderService_UpdateSettlementApplication_FullMethodName = "/OrderService/UpdateSettlementApplication"
)

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	//订单列表
	OrderList(ctx context.Context, in *OrderListRequest, opts ...grpc.CallOption) (*OrderListResponse, error)
	//订单详细
	OrderDetailed(ctx context.Context, in *OrderDetailedRequest, opts ...grpc.CallOption) (*OrderDetailResponse, error)
	//创建订单
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
	//获取订单提交token
	GetOrderToken(ctx context.Context, in *OrderEmptyRequest, opts ...grpc.CallOption) (*GetOrderTokenResponse, error)
	//订单续费
	OrderRenewalLease(ctx context.Context, in *OrderRenewalLeaseRequest, opts ...grpc.CallOption) (*OrderDefaultResponse, error)
	//管理端_查询订单列表
	QueryOrderList(ctx context.Context, in *QueryOrderListRequest, opts ...grpc.CallOption) (*QueryOrderListResponse, error)
	//管理端-结算申请
	UserSettlementApplication(ctx context.Context, in *SettlementApplicationPB, opts ...grpc.CallOption) (*SettlementApplicationResp, error)
	//管理端-结算申请分页对象
	SettlementApplicationPage(ctx context.Context, in *SettlementApplicationPB, opts ...grpc.CallOption) (*SettlementApplicationPageResp, error)
	//管理端-更新申请单
	UpdateSettlementApplication(ctx context.Context, in *SettlementApplicationPB, opts ...grpc.CallOption) (*UpdateSettlementApplicationResp, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) OrderList(ctx context.Context, in *OrderListRequest, opts ...grpc.CallOption) (*OrderListResponse, error) {
	out := new(OrderListResponse)
	err := c.cc.Invoke(ctx, OrderService_OrderList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) OrderDetailed(ctx context.Context, in *OrderDetailedRequest, opts ...grpc.CallOption) (*OrderDetailResponse, error) {
	out := new(OrderDetailResponse)
	err := c.cc.Invoke(ctx, OrderService_OrderDetailed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	out := new(CreateOrderResponse)
	err := c.cc.Invoke(ctx, OrderService_CreateOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrderToken(ctx context.Context, in *OrderEmptyRequest, opts ...grpc.CallOption) (*GetOrderTokenResponse, error) {
	out := new(GetOrderTokenResponse)
	err := c.cc.Invoke(ctx, OrderService_GetOrderToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) OrderRenewalLease(ctx context.Context, in *OrderRenewalLeaseRequest, opts ...grpc.CallOption) (*OrderDefaultResponse, error) {
	out := new(OrderDefaultResponse)
	err := c.cc.Invoke(ctx, OrderService_OrderRenewalLease_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) QueryOrderList(ctx context.Context, in *QueryOrderListRequest, opts ...grpc.CallOption) (*QueryOrderListResponse, error) {
	out := new(QueryOrderListResponse)
	err := c.cc.Invoke(ctx, OrderService_QueryOrderList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) UserSettlementApplication(ctx context.Context, in *SettlementApplicationPB, opts ...grpc.CallOption) (*SettlementApplicationResp, error) {
	out := new(SettlementApplicationResp)
	err := c.cc.Invoke(ctx, OrderService_UserSettlementApplication_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) SettlementApplicationPage(ctx context.Context, in *SettlementApplicationPB, opts ...grpc.CallOption) (*SettlementApplicationPageResp, error) {
	out := new(SettlementApplicationPageResp)
	err := c.cc.Invoke(ctx, OrderService_SettlementApplicationPage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) UpdateSettlementApplication(ctx context.Context, in *SettlementApplicationPB, opts ...grpc.CallOption) (*UpdateSettlementApplicationResp, error) {
	out := new(UpdateSettlementApplicationResp)
	err := c.cc.Invoke(ctx, OrderService_UpdateSettlementApplication_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations should embed UnimplementedOrderServiceServer
// for forward compatibility
type OrderServiceServer interface {
	//订单列表
	OrderList(context.Context, *OrderListRequest) (*OrderListResponse, error)
	//订单详细
	OrderDetailed(context.Context, *OrderDetailedRequest) (*OrderDetailResponse, error)
	//创建订单
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error)
	//获取订单提交token
	GetOrderToken(context.Context, *OrderEmptyRequest) (*GetOrderTokenResponse, error)
	//订单续费
	OrderRenewalLease(context.Context, *OrderRenewalLeaseRequest) (*OrderDefaultResponse, error)
	//管理端_查询订单列表
	QueryOrderList(context.Context, *QueryOrderListRequest) (*QueryOrderListResponse, error)
	//管理端-结算申请
	UserSettlementApplication(context.Context, *SettlementApplicationPB) (*SettlementApplicationResp, error)
	//管理端-结算申请分页对象
	SettlementApplicationPage(context.Context, *SettlementApplicationPB) (*SettlementApplicationPageResp, error)
	//管理端-更新申请单
	UpdateSettlementApplication(context.Context, *SettlementApplicationPB) (*UpdateSettlementApplicationResp, error)
}

// UnimplementedOrderServiceServer should be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (UnimplementedOrderServiceServer) OrderList(context.Context, *OrderListRequest) (*OrderListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderList not implemented")
}
func (UnimplementedOrderServiceServer) OrderDetailed(context.Context, *OrderDetailedRequest) (*OrderDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderDetailed not implemented")
}
func (UnimplementedOrderServiceServer) CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderServiceServer) GetOrderToken(context.Context, *OrderEmptyRequest) (*GetOrderTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderToken not implemented")
}
func (UnimplementedOrderServiceServer) OrderRenewalLease(context.Context, *OrderRenewalLeaseRequest) (*OrderDefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderRenewalLease not implemented")
}
func (UnimplementedOrderServiceServer) QueryOrderList(context.Context, *QueryOrderListRequest) (*QueryOrderListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryOrderList not implemented")
}
func (UnimplementedOrderServiceServer) UserSettlementApplication(context.Context, *SettlementApplicationPB) (*SettlementApplicationResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSettlementApplication not implemented")
}
func (UnimplementedOrderServiceServer) SettlementApplicationPage(context.Context, *SettlementApplicationPB) (*SettlementApplicationPageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SettlementApplicationPage not implemented")
}
func (UnimplementedOrderServiceServer) UpdateSettlementApplication(context.Context, *SettlementApplicationPB) (*UpdateSettlementApplicationResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSettlementApplication not implemented")
}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_OrderList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).OrderList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_OrderList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).OrderList(ctx, req.(*OrderListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_OrderDetailed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderDetailedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).OrderDetailed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_OrderDetailed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).OrderDetailed(ctx, req.(*OrderDetailedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrderToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderEmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrderToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_GetOrderToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrderToken(ctx, req.(*OrderEmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_OrderRenewalLease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRenewalLeaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).OrderRenewalLease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_OrderRenewalLease_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).OrderRenewalLease(ctx, req.(*OrderRenewalLeaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_QueryOrderList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryOrderListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).QueryOrderList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_QueryOrderList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).QueryOrderList(ctx, req.(*QueryOrderListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_UserSettlementApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SettlementApplicationPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).UserSettlementApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_UserSettlementApplication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).UserSettlementApplication(ctx, req.(*SettlementApplicationPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_SettlementApplicationPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SettlementApplicationPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).SettlementApplicationPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_SettlementApplicationPage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).SettlementApplicationPage(ctx, req.(*SettlementApplicationPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_UpdateSettlementApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SettlementApplicationPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).UpdateSettlementApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_UpdateSettlementApplication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).UpdateSettlementApplication(ctx, req.(*SettlementApplicationPB))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OrderList",
			Handler:    _OrderService_OrderList_Handler,
		},
		{
			MethodName: "OrderDetailed",
			Handler:    _OrderService_OrderDetailed_Handler,
		},
		{
			MethodName: "CreateOrder",
			Handler:    _OrderService_CreateOrder_Handler,
		},
		{
			MethodName: "GetOrderToken",
			Handler:    _OrderService_GetOrderToken_Handler,
		},
		{
			MethodName: "OrderRenewalLease",
			Handler:    _OrderService_OrderRenewalLease_Handler,
		},
		{
			MethodName: "QueryOrderList",
			Handler:    _OrderService_QueryOrderList_Handler,
		},
		{
			MethodName: "UserSettlementApplication",
			Handler:    _OrderService_UserSettlementApplication_Handler,
		},
		{
			MethodName: "SettlementApplicationPage",
			Handler:    _OrderService_SettlementApplicationPage_Handler,
		},
		{
			MethodName: "UpdateSettlementApplication",
			Handler:    _OrderService_UpdateSettlementApplication_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}
