// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package mall

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

// WechatPayServiceClient is the client API for WechatPayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WechatPayServiceClient interface {
	Mp(ctx context.Context, in *MpRequest, opts ...grpc.CallOption) (*MpReply, error)
	Mini(ctx context.Context, in *MpRequest, opts ...grpc.CallOption) (*MpReply, error)
	App(ctx context.Context, in *AppRequest, opts ...grpc.CallOption) (*AppReply, error)
	Scan(ctx context.Context, in *ScanRequest, opts ...grpc.CallOption) (*ScanReply, error)
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryReply, error)
	QueryRefund(ctx context.Context, in *QueryRefundRequest, opts ...grpc.CallOption) (*QueryRefundReply, error)
	Refund(ctx context.Context, in *RefundRequest, opts ...grpc.CallOption) (*RefundReply, error)
	Notify(ctx context.Context, in *NotifyRequest, opts ...grpc.CallOption) (*NotifyReply, error)
	NotifyRefund(ctx context.Context, in *NotifyRefundRequest, opts ...grpc.CallOption) (*NotifyRefundReply, error)
}

type wechatPayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWechatPayServiceClient(cc grpc.ClientConnInterface) WechatPayServiceClient {
	return &wechatPayServiceClient{cc}
}

func (c *wechatPayServiceClient) Mp(ctx context.Context, in *MpRequest, opts ...grpc.CallOption) (*MpReply, error) {
	out := new(MpReply)
	err := c.cc.Invoke(ctx, "/mall.WechatPayService/Mp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wechatPayServiceClient) Mini(ctx context.Context, in *MpRequest, opts ...grpc.CallOption) (*MpReply, error) {
	out := new(MpReply)
	err := c.cc.Invoke(ctx, "/mall.WechatPayService/Mini", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wechatPayServiceClient) App(ctx context.Context, in *AppRequest, opts ...grpc.CallOption) (*AppReply, error) {
	out := new(AppReply)
	err := c.cc.Invoke(ctx, "/mall.WechatPayService/App", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wechatPayServiceClient) Scan(ctx context.Context, in *ScanRequest, opts ...grpc.CallOption) (*ScanReply, error) {
	out := new(ScanReply)
	err := c.cc.Invoke(ctx, "/mall.WechatPayService/Scan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wechatPayServiceClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryReply, error) {
	out := new(QueryReply)
	err := c.cc.Invoke(ctx, "/mall.WechatPayService/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wechatPayServiceClient) QueryRefund(ctx context.Context, in *QueryRefundRequest, opts ...grpc.CallOption) (*QueryRefundReply, error) {
	out := new(QueryRefundReply)
	err := c.cc.Invoke(ctx, "/mall.WechatPayService/QueryRefund", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wechatPayServiceClient) Refund(ctx context.Context, in *RefundRequest, opts ...grpc.CallOption) (*RefundReply, error) {
	out := new(RefundReply)
	err := c.cc.Invoke(ctx, "/mall.WechatPayService/Refund", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wechatPayServiceClient) Notify(ctx context.Context, in *NotifyRequest, opts ...grpc.CallOption) (*NotifyReply, error) {
	out := new(NotifyReply)
	err := c.cc.Invoke(ctx, "/mall.WechatPayService/Notify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wechatPayServiceClient) NotifyRefund(ctx context.Context, in *NotifyRefundRequest, opts ...grpc.CallOption) (*NotifyRefundReply, error) {
	out := new(NotifyRefundReply)
	err := c.cc.Invoke(ctx, "/mall.WechatPayService/NotifyRefund", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WechatPayServiceServer is the server API for WechatPayService service.
// All implementations must embed UnimplementedWechatPayServiceServer
// for forward compatibility
type WechatPayServiceServer interface {
	Mp(context.Context, *MpRequest) (*MpReply, error)
	Mini(context.Context, *MpRequest) (*MpReply, error)
	App(context.Context, *AppRequest) (*AppReply, error)
	Scan(context.Context, *ScanRequest) (*ScanReply, error)
	Query(context.Context, *QueryRequest) (*QueryReply, error)
	QueryRefund(context.Context, *QueryRefundRequest) (*QueryRefundReply, error)
	Refund(context.Context, *RefundRequest) (*RefundReply, error)
	Notify(context.Context, *NotifyRequest) (*NotifyReply, error)
	NotifyRefund(context.Context, *NotifyRefundRequest) (*NotifyRefundReply, error)
	mustEmbedUnimplementedWechatPayServiceServer()
}

// UnimplementedWechatPayServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWechatPayServiceServer struct {
}

func (UnimplementedWechatPayServiceServer) Mp(context.Context, *MpRequest) (*MpReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mp not implemented")
}
func (UnimplementedWechatPayServiceServer) Mini(context.Context, *MpRequest) (*MpReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mini not implemented")
}
func (UnimplementedWechatPayServiceServer) App(context.Context, *AppRequest) (*AppReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method App not implemented")
}
func (UnimplementedWechatPayServiceServer) Scan(context.Context, *ScanRequest) (*ScanReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Scan not implemented")
}
func (UnimplementedWechatPayServiceServer) Query(context.Context, *QueryRequest) (*QueryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedWechatPayServiceServer) QueryRefund(context.Context, *QueryRefundRequest) (*QueryRefundReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryRefund not implemented")
}
func (UnimplementedWechatPayServiceServer) Refund(context.Context, *RefundRequest) (*RefundReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refund not implemented")
}
func (UnimplementedWechatPayServiceServer) Notify(context.Context, *NotifyRequest) (*NotifyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Notify not implemented")
}
func (UnimplementedWechatPayServiceServer) NotifyRefund(context.Context, *NotifyRefundRequest) (*NotifyRefundReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyRefund not implemented")
}
func (UnimplementedWechatPayServiceServer) mustEmbedUnimplementedWechatPayServiceServer() {}

// UnsafeWechatPayServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WechatPayServiceServer will
// result in compilation errors.
type UnsafeWechatPayServiceServer interface {
	mustEmbedUnimplementedWechatPayServiceServer()
}

func RegisterWechatPayServiceServer(s grpc.ServiceRegistrar, srv WechatPayServiceServer) {
	s.RegisterService(&WechatPayService_ServiceDesc, srv)
}

func _WechatPayService_Mp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WechatPayServiceServer).Mp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mall.WechatPayService/Mp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WechatPayServiceServer).Mp(ctx, req.(*MpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WechatPayService_Mini_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WechatPayServiceServer).Mini(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mall.WechatPayService/Mini",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WechatPayServiceServer).Mini(ctx, req.(*MpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WechatPayService_App_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WechatPayServiceServer).App(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mall.WechatPayService/App",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WechatPayServiceServer).App(ctx, req.(*AppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WechatPayService_Scan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WechatPayServiceServer).Scan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mall.WechatPayService/Scan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WechatPayServiceServer).Scan(ctx, req.(*ScanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WechatPayService_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WechatPayServiceServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mall.WechatPayService/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WechatPayServiceServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WechatPayService_QueryRefund_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRefundRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WechatPayServiceServer).QueryRefund(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mall.WechatPayService/QueryRefund",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WechatPayServiceServer).QueryRefund(ctx, req.(*QueryRefundRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WechatPayService_Refund_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefundRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WechatPayServiceServer).Refund(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mall.WechatPayService/Refund",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WechatPayServiceServer).Refund(ctx, req.(*RefundRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WechatPayService_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WechatPayServiceServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mall.WechatPayService/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WechatPayServiceServer).Notify(ctx, req.(*NotifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WechatPayService_NotifyRefund_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyRefundRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WechatPayServiceServer).NotifyRefund(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mall.WechatPayService/NotifyRefund",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WechatPayServiceServer).NotifyRefund(ctx, req.(*NotifyRefundRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WechatPayService_ServiceDesc is the grpc.ServiceDesc for WechatPayService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WechatPayService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mall.WechatPayService",
	HandlerType: (*WechatPayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Mp",
			Handler:    _WechatPayService_Mp_Handler,
		},
		{
			MethodName: "Mini",
			Handler:    _WechatPayService_Mini_Handler,
		},
		{
			MethodName: "App",
			Handler:    _WechatPayService_App_Handler,
		},
		{
			MethodName: "Scan",
			Handler:    _WechatPayService_Scan_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _WechatPayService_Query_Handler,
		},
		{
			MethodName: "QueryRefund",
			Handler:    _WechatPayService_QueryRefund_Handler,
		},
		{
			MethodName: "Refund",
			Handler:    _WechatPayService_Refund_Handler,
		},
		{
			MethodName: "Notify",
			Handler:    _WechatPayService_Notify_Handler,
		},
		{
			MethodName: "NotifyRefund",
			Handler:    _WechatPayService_NotifyRefund_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/mall/wechatpay.proto",
}
