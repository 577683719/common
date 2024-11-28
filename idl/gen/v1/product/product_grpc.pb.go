// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: product.proto

package product

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
	ProductService_QueryInstancesPrice_FullMethodName             = "/ProductService/QueryInstancesPrice"
	ProductService_QueryAllRunningInstances_FullMethodName        = "/ProductService/QueryAllRunningInstances"
	ProductService_SelectProductList_FullMethodName               = "/ProductService/SelectProductList"
	ProductService_SelectProductDetailed_FullMethodName           = "/ProductService/SelectProductDetailed"
	ProductService_QueryServerAreaList_FullMethodName             = "/ProductService/QueryServerAreaList"
	ProductService_QueryGpuTypeList_FullMethodName                = "/ProductService/QueryGpuTypeList"
	ProductService_QueryMemoryList_FullMethodName                 = "/ProductService/QueryMemoryList"
	ProductService_QueryVideoMemoryList_FullMethodName            = "/ProductService/QueryVideoMemoryList"
	ProductService_QueryCpuTypeList_FullMethodName                = "/ProductService/QueryCpuTypeList"
	ProductService_QueryGpuType_FullMethodName                    = "/ProductService/QueryGpuType"
	ProductService_QueryCpuType_FullMethodName                    = "/ProductService/QueryCpuType"
	ProductService_AddGpuTypeIntroduce_FullMethodName             = "/ProductService/AddGpuTypeIntroduce"
	ProductService_AddCpuTypeIntroduce_FullMethodName             = "/ProductService/AddCpuTypeIntroduce"
	ProductService_GetRecommendedPersonProductList_FullMethodName = "/ProductService/GetRecommendedPersonProductList"
	ProductService_AddProductPurpose_FullMethodName               = "/ProductService/AddProductPurpose"
	ProductService_UpdateProductPurpose_FullMethodName            = "/ProductService/UpdateProductPurpose"
	ProductService_GetProductPurpose_FullMethodName               = "/ProductService/GetProductPurpose"
	ProductService_GetProductDetailed_FullMethodName              = "/ProductService/GetProductDetailed"
	ProductService_UpdateProductDetailed_FullMethodName           = "/ProductService/UpdateProductDetailed"
)

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductServiceClient interface {
	//查询当前实例的价格
	QueryInstancesPrice(ctx context.Context, in *QueryInstancesPriceRequest, opts ...grpc.CallOption) (*QueryInstancesPriceResponse, error)
	QueryAllRunningInstances(ctx context.Context, in *QueryAllRunningInstancesPriceRequest, opts ...grpc.CallOption) (*QueryAllRunningInstancesPriceResponse, error)
	//产品列表
	SelectProductList(ctx context.Context, in *SelectProductListRequest, opts ...grpc.CallOption) (*SelectProductListResponse, error)
	//产品详细
	SelectProductDetailed(ctx context.Context, in *SelectProductDetailedRequest, opts ...grpc.CallOption) (*SelectProductDetailedResponse, error)
	//查询区域列表
	QueryServerAreaList(ctx context.Context, in *QueryServerAreaListRequest, opts ...grpc.CallOption) (*QueryServerAreaListResponse, error)
	//查询gpu类型列表
	QueryGpuTypeList(ctx context.Context, in *QueryGpuTypeListRequest, opts ...grpc.CallOption) (*QueryGpuTypeListResponse, error)
	//查询内存列表
	QueryMemoryList(ctx context.Context, in *QueryMemoryListRequest, opts ...grpc.CallOption) (*QueryMemoryListResponse, error)
	//查询显存列表
	QueryVideoMemoryList(ctx context.Context, in *QueryVideoMemoryListRequest, opts ...grpc.CallOption) (*QueryVideoMemoryListResponse, error)
	//查询cup类型列表
	QueryCpuTypeList(ctx context.Context, in *QueryCpuTypeListRequest, opts ...grpc.CallOption) (*QueryCpuTypeListResponse, error)
	//-------------------------------管理端-------------------------------
	//管理端_查询gpu类型列表
	QueryGpuType(ctx context.Context, in *QueryGpuTypeRequest, opts ...grpc.CallOption) (*QueryGpuTypeResponse, error)
	//管理端_查询cpu类型列表
	QueryCpuType(ctx context.Context, in *QueryCpuTypeRequest, opts ...grpc.CallOption) (*QueryCpuTypeResponse, error)
	//管理端_添加gpu类型介绍
	AddGpuTypeIntroduce(ctx context.Context, in *AddGpuTypeIntroduceRequest, opts ...grpc.CallOption) (*AddGpuTypeIntroduceResponse, error)
	//管理端_添加cpu类型介绍
	AddCpuTypeIntroduce(ctx context.Context, in *AddCpuTypeIntroduceRequest, opts ...grpc.CallOption) (*AddCpuTypeIntroduceResponse, error)
	//管理端_获取推荐人产品列表
	GetRecommendedPersonProductList(ctx context.Context, in *GetRecommendedPersonProductListRequest, opts ...grpc.CallOption) (*GetRecommendedPersonProductListResponse, error)
	//管理端—添加产品用途
	AddProductPurpose(ctx context.Context, in *ProductPurpose, opts ...grpc.CallOption) (*ProductDefaultResponse, error)
	//管理端-修改产品用途
	UpdateProductPurpose(ctx context.Context, in *ProductPurpose, opts ...grpc.CallOption) (*ProductDefaultResponse, error)
	//管理端-根据产品id获取产品用途
	GetProductPurpose(ctx context.Context, in *ProductPurpose, opts ...grpc.CallOption) (*GetProductPurposeResponse, error)
	//管理端-获取产品详细
	GetProductDetailed(ctx context.Context, in *GetProductDetailedReq, opts ...grpc.CallOption) (*GetProductDetailedResp, error)
	//管理端-修改产品信息
	UpdateProductDetailed(ctx context.Context, in *Product, opts ...grpc.CallOption) (*UpdateProductDetailedResp, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) QueryInstancesPrice(ctx context.Context, in *QueryInstancesPriceRequest, opts ...grpc.CallOption) (*QueryInstancesPriceResponse, error) {
	out := new(QueryInstancesPriceResponse)
	err := c.cc.Invoke(ctx, ProductService_QueryInstancesPrice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) QueryAllRunningInstances(ctx context.Context, in *QueryAllRunningInstancesPriceRequest, opts ...grpc.CallOption) (*QueryAllRunningInstancesPriceResponse, error) {
	out := new(QueryAllRunningInstancesPriceResponse)
	err := c.cc.Invoke(ctx, ProductService_QueryAllRunningInstances_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) SelectProductList(ctx context.Context, in *SelectProductListRequest, opts ...grpc.CallOption) (*SelectProductListResponse, error) {
	out := new(SelectProductListResponse)
	err := c.cc.Invoke(ctx, ProductService_SelectProductList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) SelectProductDetailed(ctx context.Context, in *SelectProductDetailedRequest, opts ...grpc.CallOption) (*SelectProductDetailedResponse, error) {
	out := new(SelectProductDetailedResponse)
	err := c.cc.Invoke(ctx, ProductService_SelectProductDetailed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) QueryServerAreaList(ctx context.Context, in *QueryServerAreaListRequest, opts ...grpc.CallOption) (*QueryServerAreaListResponse, error) {
	out := new(QueryServerAreaListResponse)
	err := c.cc.Invoke(ctx, ProductService_QueryServerAreaList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) QueryGpuTypeList(ctx context.Context, in *QueryGpuTypeListRequest, opts ...grpc.CallOption) (*QueryGpuTypeListResponse, error) {
	out := new(QueryGpuTypeListResponse)
	err := c.cc.Invoke(ctx, ProductService_QueryGpuTypeList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) QueryMemoryList(ctx context.Context, in *QueryMemoryListRequest, opts ...grpc.CallOption) (*QueryMemoryListResponse, error) {
	out := new(QueryMemoryListResponse)
	err := c.cc.Invoke(ctx, ProductService_QueryMemoryList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) QueryVideoMemoryList(ctx context.Context, in *QueryVideoMemoryListRequest, opts ...grpc.CallOption) (*QueryVideoMemoryListResponse, error) {
	out := new(QueryVideoMemoryListResponse)
	err := c.cc.Invoke(ctx, ProductService_QueryVideoMemoryList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) QueryCpuTypeList(ctx context.Context, in *QueryCpuTypeListRequest, opts ...grpc.CallOption) (*QueryCpuTypeListResponse, error) {
	out := new(QueryCpuTypeListResponse)
	err := c.cc.Invoke(ctx, ProductService_QueryCpuTypeList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) QueryGpuType(ctx context.Context, in *QueryGpuTypeRequest, opts ...grpc.CallOption) (*QueryGpuTypeResponse, error) {
	out := new(QueryGpuTypeResponse)
	err := c.cc.Invoke(ctx, ProductService_QueryGpuType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) QueryCpuType(ctx context.Context, in *QueryCpuTypeRequest, opts ...grpc.CallOption) (*QueryCpuTypeResponse, error) {
	out := new(QueryCpuTypeResponse)
	err := c.cc.Invoke(ctx, ProductService_QueryCpuType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) AddGpuTypeIntroduce(ctx context.Context, in *AddGpuTypeIntroduceRequest, opts ...grpc.CallOption) (*AddGpuTypeIntroduceResponse, error) {
	out := new(AddGpuTypeIntroduceResponse)
	err := c.cc.Invoke(ctx, ProductService_AddGpuTypeIntroduce_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) AddCpuTypeIntroduce(ctx context.Context, in *AddCpuTypeIntroduceRequest, opts ...grpc.CallOption) (*AddCpuTypeIntroduceResponse, error) {
	out := new(AddCpuTypeIntroduceResponse)
	err := c.cc.Invoke(ctx, ProductService_AddCpuTypeIntroduce_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetRecommendedPersonProductList(ctx context.Context, in *GetRecommendedPersonProductListRequest, opts ...grpc.CallOption) (*GetRecommendedPersonProductListResponse, error) {
	out := new(GetRecommendedPersonProductListResponse)
	err := c.cc.Invoke(ctx, ProductService_GetRecommendedPersonProductList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) AddProductPurpose(ctx context.Context, in *ProductPurpose, opts ...grpc.CallOption) (*ProductDefaultResponse, error) {
	out := new(ProductDefaultResponse)
	err := c.cc.Invoke(ctx, ProductService_AddProductPurpose_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) UpdateProductPurpose(ctx context.Context, in *ProductPurpose, opts ...grpc.CallOption) (*ProductDefaultResponse, error) {
	out := new(ProductDefaultResponse)
	err := c.cc.Invoke(ctx, ProductService_UpdateProductPurpose_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetProductPurpose(ctx context.Context, in *ProductPurpose, opts ...grpc.CallOption) (*GetProductPurposeResponse, error) {
	out := new(GetProductPurposeResponse)
	err := c.cc.Invoke(ctx, ProductService_GetProductPurpose_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetProductDetailed(ctx context.Context, in *GetProductDetailedReq, opts ...grpc.CallOption) (*GetProductDetailedResp, error) {
	out := new(GetProductDetailedResp)
	err := c.cc.Invoke(ctx, ProductService_GetProductDetailed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) UpdateProductDetailed(ctx context.Context, in *Product, opts ...grpc.CallOption) (*UpdateProductDetailedResp, error) {
	out := new(UpdateProductDetailedResp)
	err := c.cc.Invoke(ctx, ProductService_UpdateProductDetailed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
// All implementations should embed UnimplementedProductServiceServer
// for forward compatibility
type ProductServiceServer interface {
	//查询当前实例的价格
	QueryInstancesPrice(context.Context, *QueryInstancesPriceRequest) (*QueryInstancesPriceResponse, error)
	QueryAllRunningInstances(context.Context, *QueryAllRunningInstancesPriceRequest) (*QueryAllRunningInstancesPriceResponse, error)
	//产品列表
	SelectProductList(context.Context, *SelectProductListRequest) (*SelectProductListResponse, error)
	//产品详细
	SelectProductDetailed(context.Context, *SelectProductDetailedRequest) (*SelectProductDetailedResponse, error)
	//查询区域列表
	QueryServerAreaList(context.Context, *QueryServerAreaListRequest) (*QueryServerAreaListResponse, error)
	//查询gpu类型列表
	QueryGpuTypeList(context.Context, *QueryGpuTypeListRequest) (*QueryGpuTypeListResponse, error)
	//查询内存列表
	QueryMemoryList(context.Context, *QueryMemoryListRequest) (*QueryMemoryListResponse, error)
	//查询显存列表
	QueryVideoMemoryList(context.Context, *QueryVideoMemoryListRequest) (*QueryVideoMemoryListResponse, error)
	//查询cup类型列表
	QueryCpuTypeList(context.Context, *QueryCpuTypeListRequest) (*QueryCpuTypeListResponse, error)
	//-------------------------------管理端-------------------------------
	//管理端_查询gpu类型列表
	QueryGpuType(context.Context, *QueryGpuTypeRequest) (*QueryGpuTypeResponse, error)
	//管理端_查询cpu类型列表
	QueryCpuType(context.Context, *QueryCpuTypeRequest) (*QueryCpuTypeResponse, error)
	//管理端_添加gpu类型介绍
	AddGpuTypeIntroduce(context.Context, *AddGpuTypeIntroduceRequest) (*AddGpuTypeIntroduceResponse, error)
	//管理端_添加cpu类型介绍
	AddCpuTypeIntroduce(context.Context, *AddCpuTypeIntroduceRequest) (*AddCpuTypeIntroduceResponse, error)
	//管理端_获取推荐人产品列表
	GetRecommendedPersonProductList(context.Context, *GetRecommendedPersonProductListRequest) (*GetRecommendedPersonProductListResponse, error)
	//管理端—添加产品用途
	AddProductPurpose(context.Context, *ProductPurpose) (*ProductDefaultResponse, error)
	//管理端-修改产品用途
	UpdateProductPurpose(context.Context, *ProductPurpose) (*ProductDefaultResponse, error)
	//管理端-根据产品id获取产品用途
	GetProductPurpose(context.Context, *ProductPurpose) (*GetProductPurposeResponse, error)
	//管理端-获取产品详细
	GetProductDetailed(context.Context, *GetProductDetailedReq) (*GetProductDetailedResp, error)
	//管理端-修改产品信息
	UpdateProductDetailed(context.Context, *Product) (*UpdateProductDetailedResp, error)
}

// UnimplementedProductServiceServer should be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (UnimplementedProductServiceServer) QueryInstancesPrice(context.Context, *QueryInstancesPriceRequest) (*QueryInstancesPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryInstancesPrice not implemented")
}
func (UnimplementedProductServiceServer) QueryAllRunningInstances(context.Context, *QueryAllRunningInstancesPriceRequest) (*QueryAllRunningInstancesPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryAllRunningInstances not implemented")
}
func (UnimplementedProductServiceServer) SelectProductList(context.Context, *SelectProductListRequest) (*SelectProductListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectProductList not implemented")
}
func (UnimplementedProductServiceServer) SelectProductDetailed(context.Context, *SelectProductDetailedRequest) (*SelectProductDetailedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectProductDetailed not implemented")
}
func (UnimplementedProductServiceServer) QueryServerAreaList(context.Context, *QueryServerAreaListRequest) (*QueryServerAreaListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryServerAreaList not implemented")
}
func (UnimplementedProductServiceServer) QueryGpuTypeList(context.Context, *QueryGpuTypeListRequest) (*QueryGpuTypeListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryGpuTypeList not implemented")
}
func (UnimplementedProductServiceServer) QueryMemoryList(context.Context, *QueryMemoryListRequest) (*QueryMemoryListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryMemoryList not implemented")
}
func (UnimplementedProductServiceServer) QueryVideoMemoryList(context.Context, *QueryVideoMemoryListRequest) (*QueryVideoMemoryListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryVideoMemoryList not implemented")
}
func (UnimplementedProductServiceServer) QueryCpuTypeList(context.Context, *QueryCpuTypeListRequest) (*QueryCpuTypeListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryCpuTypeList not implemented")
}
func (UnimplementedProductServiceServer) QueryGpuType(context.Context, *QueryGpuTypeRequest) (*QueryGpuTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryGpuType not implemented")
}
func (UnimplementedProductServiceServer) QueryCpuType(context.Context, *QueryCpuTypeRequest) (*QueryCpuTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryCpuType not implemented")
}
func (UnimplementedProductServiceServer) AddGpuTypeIntroduce(context.Context, *AddGpuTypeIntroduceRequest) (*AddGpuTypeIntroduceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGpuTypeIntroduce not implemented")
}
func (UnimplementedProductServiceServer) AddCpuTypeIntroduce(context.Context, *AddCpuTypeIntroduceRequest) (*AddCpuTypeIntroduceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCpuTypeIntroduce not implemented")
}
func (UnimplementedProductServiceServer) GetRecommendedPersonProductList(context.Context, *GetRecommendedPersonProductListRequest) (*GetRecommendedPersonProductListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecommendedPersonProductList not implemented")
}
func (UnimplementedProductServiceServer) AddProductPurpose(context.Context, *ProductPurpose) (*ProductDefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProductPurpose not implemented")
}
func (UnimplementedProductServiceServer) UpdateProductPurpose(context.Context, *ProductPurpose) (*ProductDefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProductPurpose not implemented")
}
func (UnimplementedProductServiceServer) GetProductPurpose(context.Context, *ProductPurpose) (*GetProductPurposeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductPurpose not implemented")
}
func (UnimplementedProductServiceServer) GetProductDetailed(context.Context, *GetProductDetailedReq) (*GetProductDetailedResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductDetailed not implemented")
}
func (UnimplementedProductServiceServer) UpdateProductDetailed(context.Context, *Product) (*UpdateProductDetailedResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProductDetailed not implemented")
}

// UnsafeProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServiceServer will
// result in compilation errors.
type UnsafeProductServiceServer interface {
	mustEmbedUnimplementedProductServiceServer()
}

func RegisterProductServiceServer(s grpc.ServiceRegistrar, srv ProductServiceServer) {
	s.RegisterService(&ProductService_ServiceDesc, srv)
}

func _ProductService_QueryInstancesPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryInstancesPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).QueryInstancesPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_QueryInstancesPrice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).QueryInstancesPrice(ctx, req.(*QueryInstancesPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_QueryAllRunningInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllRunningInstancesPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).QueryAllRunningInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_QueryAllRunningInstances_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).QueryAllRunningInstances(ctx, req.(*QueryAllRunningInstancesPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_SelectProductList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectProductListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).SelectProductList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_SelectProductList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).SelectProductList(ctx, req.(*SelectProductListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_SelectProductDetailed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectProductDetailedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).SelectProductDetailed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_SelectProductDetailed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).SelectProductDetailed(ctx, req.(*SelectProductDetailedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_QueryServerAreaList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryServerAreaListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).QueryServerAreaList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_QueryServerAreaList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).QueryServerAreaList(ctx, req.(*QueryServerAreaListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_QueryGpuTypeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGpuTypeListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).QueryGpuTypeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_QueryGpuTypeList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).QueryGpuTypeList(ctx, req.(*QueryGpuTypeListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_QueryMemoryList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMemoryListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).QueryMemoryList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_QueryMemoryList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).QueryMemoryList(ctx, req.(*QueryMemoryListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_QueryVideoMemoryList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryVideoMemoryListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).QueryVideoMemoryList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_QueryVideoMemoryList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).QueryVideoMemoryList(ctx, req.(*QueryVideoMemoryListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_QueryCpuTypeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCpuTypeListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).QueryCpuTypeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_QueryCpuTypeList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).QueryCpuTypeList(ctx, req.(*QueryCpuTypeListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_QueryGpuType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGpuTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).QueryGpuType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_QueryGpuType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).QueryGpuType(ctx, req.(*QueryGpuTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_QueryCpuType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCpuTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).QueryCpuType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_QueryCpuType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).QueryCpuType(ctx, req.(*QueryCpuTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_AddGpuTypeIntroduce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGpuTypeIntroduceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).AddGpuTypeIntroduce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_AddGpuTypeIntroduce_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).AddGpuTypeIntroduce(ctx, req.(*AddGpuTypeIntroduceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_AddCpuTypeIntroduce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCpuTypeIntroduceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).AddCpuTypeIntroduce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_AddCpuTypeIntroduce_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).AddCpuTypeIntroduce(ctx, req.(*AddCpuTypeIntroduceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetRecommendedPersonProductList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecommendedPersonProductListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetRecommendedPersonProductList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_GetRecommendedPersonProductList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetRecommendedPersonProductList(ctx, req.(*GetRecommendedPersonProductListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_AddProductPurpose_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductPurpose)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).AddProductPurpose(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_AddProductPurpose_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).AddProductPurpose(ctx, req.(*ProductPurpose))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_UpdateProductPurpose_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductPurpose)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).UpdateProductPurpose(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_UpdateProductPurpose_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).UpdateProductPurpose(ctx, req.(*ProductPurpose))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetProductPurpose_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductPurpose)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProductPurpose(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_GetProductPurpose_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProductPurpose(ctx, req.(*ProductPurpose))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetProductDetailed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductDetailedReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProductDetailed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_GetProductDetailed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProductDetailed(ctx, req.(*GetProductDetailedReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_UpdateProductDetailed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).UpdateProductDetailed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_UpdateProductDetailed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).UpdateProductDetailed(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductService_ServiceDesc is the grpc.ServiceDesc for ProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryInstancesPrice",
			Handler:    _ProductService_QueryInstancesPrice_Handler,
		},
		{
			MethodName: "QueryAllRunningInstances",
			Handler:    _ProductService_QueryAllRunningInstances_Handler,
		},
		{
			MethodName: "SelectProductList",
			Handler:    _ProductService_SelectProductList_Handler,
		},
		{
			MethodName: "SelectProductDetailed",
			Handler:    _ProductService_SelectProductDetailed_Handler,
		},
		{
			MethodName: "QueryServerAreaList",
			Handler:    _ProductService_QueryServerAreaList_Handler,
		},
		{
			MethodName: "QueryGpuTypeList",
			Handler:    _ProductService_QueryGpuTypeList_Handler,
		},
		{
			MethodName: "QueryMemoryList",
			Handler:    _ProductService_QueryMemoryList_Handler,
		},
		{
			MethodName: "QueryVideoMemoryList",
			Handler:    _ProductService_QueryVideoMemoryList_Handler,
		},
		{
			MethodName: "QueryCpuTypeList",
			Handler:    _ProductService_QueryCpuTypeList_Handler,
		},
		{
			MethodName: "QueryGpuType",
			Handler:    _ProductService_QueryGpuType_Handler,
		},
		{
			MethodName: "QueryCpuType",
			Handler:    _ProductService_QueryCpuType_Handler,
		},
		{
			MethodName: "AddGpuTypeIntroduce",
			Handler:    _ProductService_AddGpuTypeIntroduce_Handler,
		},
		{
			MethodName: "AddCpuTypeIntroduce",
			Handler:    _ProductService_AddCpuTypeIntroduce_Handler,
		},
		{
			MethodName: "GetRecommendedPersonProductList",
			Handler:    _ProductService_GetRecommendedPersonProductList_Handler,
		},
		{
			MethodName: "AddProductPurpose",
			Handler:    _ProductService_AddProductPurpose_Handler,
		},
		{
			MethodName: "UpdateProductPurpose",
			Handler:    _ProductService_UpdateProductPurpose_Handler,
		},
		{
			MethodName: "GetProductPurpose",
			Handler:    _ProductService_GetProductPurpose_Handler,
		},
		{
			MethodName: "GetProductDetailed",
			Handler:    _ProductService_GetProductDetailed_Handler,
		},
		{
			MethodName: "UpdateProductDetailed",
			Handler:    _ProductService_UpdateProductDetailed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product.proto",
}