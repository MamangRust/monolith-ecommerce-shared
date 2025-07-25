// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.2
// source: review_detail.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ReviewDetailService_FindAll_FullMethodName                        = "/pb.ReviewDetailService/FindAll"
	ReviewDetailService_FindById_FullMethodName                       = "/pb.ReviewDetailService/FindById"
	ReviewDetailService_FindByActive_FullMethodName                   = "/pb.ReviewDetailService/FindByActive"
	ReviewDetailService_FindByTrashed_FullMethodName                  = "/pb.ReviewDetailService/FindByTrashed"
	ReviewDetailService_Create_FullMethodName                         = "/pb.ReviewDetailService/Create"
	ReviewDetailService_Update_FullMethodName                         = "/pb.ReviewDetailService/Update"
	ReviewDetailService_TrashedReviewDetail_FullMethodName            = "/pb.ReviewDetailService/TrashedReviewDetail"
	ReviewDetailService_RestoreReviewDetail_FullMethodName            = "/pb.ReviewDetailService/RestoreReviewDetail"
	ReviewDetailService_DeleteReviewDetailPermanent_FullMethodName    = "/pb.ReviewDetailService/DeleteReviewDetailPermanent"
	ReviewDetailService_RestoreAllReviewDetail_FullMethodName         = "/pb.ReviewDetailService/RestoreAllReviewDetail"
	ReviewDetailService_DeleteAllReviewDetailPermanent_FullMethodName = "/pb.ReviewDetailService/DeleteAllReviewDetailPermanent"
)

// ReviewDetailServiceClient is the client API for ReviewDetailService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReviewDetailServiceClient interface {
	FindAll(ctx context.Context, in *FindAllReviewRequest, opts ...grpc.CallOption) (*ApiResponsePaginationReviewDetails, error)
	FindById(ctx context.Context, in *FindByIdReviewDetailRequest, opts ...grpc.CallOption) (*ApiResponseReviewDetail, error)
	FindByActive(ctx context.Context, in *FindAllReviewRequest, opts ...grpc.CallOption) (*ApiResponsePaginationReviewDetailsDeleteAt, error)
	FindByTrashed(ctx context.Context, in *FindAllReviewRequest, opts ...grpc.CallOption) (*ApiResponsePaginationReviewDetailsDeleteAt, error)
	Create(ctx context.Context, in *CreateReviewDetailRequest, opts ...grpc.CallOption) (*ApiResponseReviewDetail, error)
	Update(ctx context.Context, in *UpdateReviewDetailRequest, opts ...grpc.CallOption) (*ApiResponseReviewDetail, error)
	TrashedReviewDetail(ctx context.Context, in *FindByIdReviewDetailRequest, opts ...grpc.CallOption) (*ApiResponseReviewDetailDeleteAt, error)
	RestoreReviewDetail(ctx context.Context, in *FindByIdReviewDetailRequest, opts ...grpc.CallOption) (*ApiResponseReviewDetailDeleteAt, error)
	DeleteReviewDetailPermanent(ctx context.Context, in *FindByIdReviewDetailRequest, opts ...grpc.CallOption) (*ApiResponseReviewDelete, error)
	RestoreAllReviewDetail(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ApiResponseReviewAll, error)
	DeleteAllReviewDetailPermanent(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ApiResponseReviewAll, error)
}

type reviewDetailServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReviewDetailServiceClient(cc grpc.ClientConnInterface) ReviewDetailServiceClient {
	return &reviewDetailServiceClient{cc}
}

func (c *reviewDetailServiceClient) FindAll(ctx context.Context, in *FindAllReviewRequest, opts ...grpc.CallOption) (*ApiResponsePaginationReviewDetails, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiResponsePaginationReviewDetails)
	err := c.cc.Invoke(ctx, ReviewDetailService_FindAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewDetailServiceClient) FindById(ctx context.Context, in *FindByIdReviewDetailRequest, opts ...grpc.CallOption) (*ApiResponseReviewDetail, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiResponseReviewDetail)
	err := c.cc.Invoke(ctx, ReviewDetailService_FindById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewDetailServiceClient) FindByActive(ctx context.Context, in *FindAllReviewRequest, opts ...grpc.CallOption) (*ApiResponsePaginationReviewDetailsDeleteAt, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiResponsePaginationReviewDetailsDeleteAt)
	err := c.cc.Invoke(ctx, ReviewDetailService_FindByActive_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewDetailServiceClient) FindByTrashed(ctx context.Context, in *FindAllReviewRequest, opts ...grpc.CallOption) (*ApiResponsePaginationReviewDetailsDeleteAt, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiResponsePaginationReviewDetailsDeleteAt)
	err := c.cc.Invoke(ctx, ReviewDetailService_FindByTrashed_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewDetailServiceClient) Create(ctx context.Context, in *CreateReviewDetailRequest, opts ...grpc.CallOption) (*ApiResponseReviewDetail, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiResponseReviewDetail)
	err := c.cc.Invoke(ctx, ReviewDetailService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewDetailServiceClient) Update(ctx context.Context, in *UpdateReviewDetailRequest, opts ...grpc.CallOption) (*ApiResponseReviewDetail, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiResponseReviewDetail)
	err := c.cc.Invoke(ctx, ReviewDetailService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewDetailServiceClient) TrashedReviewDetail(ctx context.Context, in *FindByIdReviewDetailRequest, opts ...grpc.CallOption) (*ApiResponseReviewDetailDeleteAt, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiResponseReviewDetailDeleteAt)
	err := c.cc.Invoke(ctx, ReviewDetailService_TrashedReviewDetail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewDetailServiceClient) RestoreReviewDetail(ctx context.Context, in *FindByIdReviewDetailRequest, opts ...grpc.CallOption) (*ApiResponseReviewDetailDeleteAt, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiResponseReviewDetailDeleteAt)
	err := c.cc.Invoke(ctx, ReviewDetailService_RestoreReviewDetail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewDetailServiceClient) DeleteReviewDetailPermanent(ctx context.Context, in *FindByIdReviewDetailRequest, opts ...grpc.CallOption) (*ApiResponseReviewDelete, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiResponseReviewDelete)
	err := c.cc.Invoke(ctx, ReviewDetailService_DeleteReviewDetailPermanent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewDetailServiceClient) RestoreAllReviewDetail(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ApiResponseReviewAll, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiResponseReviewAll)
	err := c.cc.Invoke(ctx, ReviewDetailService_RestoreAllReviewDetail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewDetailServiceClient) DeleteAllReviewDetailPermanent(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ApiResponseReviewAll, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiResponseReviewAll)
	err := c.cc.Invoke(ctx, ReviewDetailService_DeleteAllReviewDetailPermanent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReviewDetailServiceServer is the server API for ReviewDetailService service.
// All implementations must embed UnimplementedReviewDetailServiceServer
// for forward compatibility.
type ReviewDetailServiceServer interface {
	FindAll(context.Context, *FindAllReviewRequest) (*ApiResponsePaginationReviewDetails, error)
	FindById(context.Context, *FindByIdReviewDetailRequest) (*ApiResponseReviewDetail, error)
	FindByActive(context.Context, *FindAllReviewRequest) (*ApiResponsePaginationReviewDetailsDeleteAt, error)
	FindByTrashed(context.Context, *FindAllReviewRequest) (*ApiResponsePaginationReviewDetailsDeleteAt, error)
	Create(context.Context, *CreateReviewDetailRequest) (*ApiResponseReviewDetail, error)
	Update(context.Context, *UpdateReviewDetailRequest) (*ApiResponseReviewDetail, error)
	TrashedReviewDetail(context.Context, *FindByIdReviewDetailRequest) (*ApiResponseReviewDetailDeleteAt, error)
	RestoreReviewDetail(context.Context, *FindByIdReviewDetailRequest) (*ApiResponseReviewDetailDeleteAt, error)
	DeleteReviewDetailPermanent(context.Context, *FindByIdReviewDetailRequest) (*ApiResponseReviewDelete, error)
	RestoreAllReviewDetail(context.Context, *emptypb.Empty) (*ApiResponseReviewAll, error)
	DeleteAllReviewDetailPermanent(context.Context, *emptypb.Empty) (*ApiResponseReviewAll, error)
	mustEmbedUnimplementedReviewDetailServiceServer()
}

// UnimplementedReviewDetailServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedReviewDetailServiceServer struct{}

func (UnimplementedReviewDetailServiceServer) FindAll(context.Context, *FindAllReviewRequest) (*ApiResponsePaginationReviewDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAll not implemented")
}
func (UnimplementedReviewDetailServiceServer) FindById(context.Context, *FindByIdReviewDetailRequest) (*ApiResponseReviewDetail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindById not implemented")
}
func (UnimplementedReviewDetailServiceServer) FindByActive(context.Context, *FindAllReviewRequest) (*ApiResponsePaginationReviewDetailsDeleteAt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByActive not implemented")
}
func (UnimplementedReviewDetailServiceServer) FindByTrashed(context.Context, *FindAllReviewRequest) (*ApiResponsePaginationReviewDetailsDeleteAt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByTrashed not implemented")
}
func (UnimplementedReviewDetailServiceServer) Create(context.Context, *CreateReviewDetailRequest) (*ApiResponseReviewDetail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedReviewDetailServiceServer) Update(context.Context, *UpdateReviewDetailRequest) (*ApiResponseReviewDetail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedReviewDetailServiceServer) TrashedReviewDetail(context.Context, *FindByIdReviewDetailRequest) (*ApiResponseReviewDetailDeleteAt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrashedReviewDetail not implemented")
}
func (UnimplementedReviewDetailServiceServer) RestoreReviewDetail(context.Context, *FindByIdReviewDetailRequest) (*ApiResponseReviewDetailDeleteAt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestoreReviewDetail not implemented")
}
func (UnimplementedReviewDetailServiceServer) DeleteReviewDetailPermanent(context.Context, *FindByIdReviewDetailRequest) (*ApiResponseReviewDelete, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReviewDetailPermanent not implemented")
}
func (UnimplementedReviewDetailServiceServer) RestoreAllReviewDetail(context.Context, *emptypb.Empty) (*ApiResponseReviewAll, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestoreAllReviewDetail not implemented")
}
func (UnimplementedReviewDetailServiceServer) DeleteAllReviewDetailPermanent(context.Context, *emptypb.Empty) (*ApiResponseReviewAll, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAllReviewDetailPermanent not implemented")
}
func (UnimplementedReviewDetailServiceServer) mustEmbedUnimplementedReviewDetailServiceServer() {}
func (UnimplementedReviewDetailServiceServer) testEmbeddedByValue()                             {}

// UnsafeReviewDetailServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReviewDetailServiceServer will
// result in compilation errors.
type UnsafeReviewDetailServiceServer interface {
	mustEmbedUnimplementedReviewDetailServiceServer()
}

func RegisterReviewDetailServiceServer(s grpc.ServiceRegistrar, srv ReviewDetailServiceServer) {
	// If the following call pancis, it indicates UnimplementedReviewDetailServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ReviewDetailService_ServiceDesc, srv)
}

func _ReviewDetailService_FindAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewDetailServiceServer).FindAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReviewDetailService_FindAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewDetailServiceServer).FindAll(ctx, req.(*FindAllReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewDetailService_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIdReviewDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewDetailServiceServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReviewDetailService_FindById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewDetailServiceServer).FindById(ctx, req.(*FindByIdReviewDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewDetailService_FindByActive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewDetailServiceServer).FindByActive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReviewDetailService_FindByActive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewDetailServiceServer).FindByActive(ctx, req.(*FindAllReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewDetailService_FindByTrashed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewDetailServiceServer).FindByTrashed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReviewDetailService_FindByTrashed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewDetailServiceServer).FindByTrashed(ctx, req.(*FindAllReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewDetailService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReviewDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewDetailServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReviewDetailService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewDetailServiceServer).Create(ctx, req.(*CreateReviewDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewDetailService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReviewDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewDetailServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReviewDetailService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewDetailServiceServer).Update(ctx, req.(*UpdateReviewDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewDetailService_TrashedReviewDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIdReviewDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewDetailServiceServer).TrashedReviewDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReviewDetailService_TrashedReviewDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewDetailServiceServer).TrashedReviewDetail(ctx, req.(*FindByIdReviewDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewDetailService_RestoreReviewDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIdReviewDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewDetailServiceServer).RestoreReviewDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReviewDetailService_RestoreReviewDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewDetailServiceServer).RestoreReviewDetail(ctx, req.(*FindByIdReviewDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewDetailService_DeleteReviewDetailPermanent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIdReviewDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewDetailServiceServer).DeleteReviewDetailPermanent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReviewDetailService_DeleteReviewDetailPermanent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewDetailServiceServer).DeleteReviewDetailPermanent(ctx, req.(*FindByIdReviewDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewDetailService_RestoreAllReviewDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewDetailServiceServer).RestoreAllReviewDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReviewDetailService_RestoreAllReviewDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewDetailServiceServer).RestoreAllReviewDetail(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewDetailService_DeleteAllReviewDetailPermanent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewDetailServiceServer).DeleteAllReviewDetailPermanent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReviewDetailService_DeleteAllReviewDetailPermanent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewDetailServiceServer).DeleteAllReviewDetailPermanent(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ReviewDetailService_ServiceDesc is the grpc.ServiceDesc for ReviewDetailService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReviewDetailService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ReviewDetailService",
	HandlerType: (*ReviewDetailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindAll",
			Handler:    _ReviewDetailService_FindAll_Handler,
		},
		{
			MethodName: "FindById",
			Handler:    _ReviewDetailService_FindById_Handler,
		},
		{
			MethodName: "FindByActive",
			Handler:    _ReviewDetailService_FindByActive_Handler,
		},
		{
			MethodName: "FindByTrashed",
			Handler:    _ReviewDetailService_FindByTrashed_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ReviewDetailService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ReviewDetailService_Update_Handler,
		},
		{
			MethodName: "TrashedReviewDetail",
			Handler:    _ReviewDetailService_TrashedReviewDetail_Handler,
		},
		{
			MethodName: "RestoreReviewDetail",
			Handler:    _ReviewDetailService_RestoreReviewDetail_Handler,
		},
		{
			MethodName: "DeleteReviewDetailPermanent",
			Handler:    _ReviewDetailService_DeleteReviewDetailPermanent_Handler,
		},
		{
			MethodName: "RestoreAllReviewDetail",
			Handler:    _ReviewDetailService_RestoreAllReviewDetail_Handler,
		},
		{
			MethodName: "DeleteAllReviewDetailPermanent",
			Handler:    _ReviewDetailService_DeleteAllReviewDetailPermanent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "review_detail.proto",
}
