// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: pb/lenven.proto

package pb

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
	Lenven_CreateBooking_FullMethodName  = "/pb.Lenven/CreateBooking"
	Lenven_GetBookings_FullMethodName    = "/pb.Lenven/GetBookings"
	Lenven_GetBookingByID_FullMethodName = "/pb.Lenven/GetBookingByID"
	Lenven_CancelBooking_FullMethodName  = "/pb.Lenven/CancelBooking"
)

// LenvenClient is the client API for Lenven service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LenvenClient interface {
	CreateBooking(ctx context.Context, in *CreateBookingParams, opts ...grpc.CallOption) (*CreateBookingResp, error)
	GetBookings(ctx context.Context, in *GetBookingsReq, opts ...grpc.CallOption) (*GetBookingsResp, error)
	GetBookingByID(ctx context.Context, in *GetBookingByIDReq, opts ...grpc.CallOption) (*GetBookingByIDResp, error)
	CancelBooking(ctx context.Context, in *CancelBookingReq, opts ...grpc.CallOption) (*CancelBookingResp, error)
}

type lenvenClient struct {
	cc grpc.ClientConnInterface
}

func NewLenvenClient(cc grpc.ClientConnInterface) LenvenClient {
	return &lenvenClient{cc}
}

func (c *lenvenClient) CreateBooking(ctx context.Context, in *CreateBookingParams, opts ...grpc.CallOption) (*CreateBookingResp, error) {
	out := new(CreateBookingResp)
	err := c.cc.Invoke(ctx, Lenven_CreateBooking_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lenvenClient) GetBookings(ctx context.Context, in *GetBookingsReq, opts ...grpc.CallOption) (*GetBookingsResp, error) {
	out := new(GetBookingsResp)
	err := c.cc.Invoke(ctx, Lenven_GetBookings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lenvenClient) GetBookingByID(ctx context.Context, in *GetBookingByIDReq, opts ...grpc.CallOption) (*GetBookingByIDResp, error) {
	out := new(GetBookingByIDResp)
	err := c.cc.Invoke(ctx, Lenven_GetBookingByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lenvenClient) CancelBooking(ctx context.Context, in *CancelBookingReq, opts ...grpc.CallOption) (*CancelBookingResp, error) {
	out := new(CancelBookingResp)
	err := c.cc.Invoke(ctx, Lenven_CancelBooking_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LenvenServer is the server API for Lenven service.
// All implementations must embed UnimplementedLenvenServer
// for forward compatibility
type LenvenServer interface {
	CreateBooking(context.Context, *CreateBookingParams) (*CreateBookingResp, error)
	GetBookings(context.Context, *GetBookingsReq) (*GetBookingsResp, error)
	GetBookingByID(context.Context, *GetBookingByIDReq) (*GetBookingByIDResp, error)
	CancelBooking(context.Context, *CancelBookingReq) (*CancelBookingResp, error)
	mustEmbedUnimplementedLenvenServer()
}

// UnimplementedLenvenServer must be embedded to have forward compatible implementations.
type UnimplementedLenvenServer struct {
}

func (UnimplementedLenvenServer) CreateBooking(context.Context, *CreateBookingParams) (*CreateBookingResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBooking not implemented")
}
func (UnimplementedLenvenServer) GetBookings(context.Context, *GetBookingsReq) (*GetBookingsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookings not implemented")
}
func (UnimplementedLenvenServer) GetBookingByID(context.Context, *GetBookingByIDReq) (*GetBookingByIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookingByID not implemented")
}
func (UnimplementedLenvenServer) CancelBooking(context.Context, *CancelBookingReq) (*CancelBookingResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelBooking not implemented")
}
func (UnimplementedLenvenServer) mustEmbedUnimplementedLenvenServer() {}

// UnsafeLenvenServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LenvenServer will
// result in compilation errors.
type UnsafeLenvenServer interface {
	mustEmbedUnimplementedLenvenServer()
}

func RegisterLenvenServer(s grpc.ServiceRegistrar, srv LenvenServer) {
	s.RegisterService(&Lenven_ServiceDesc, srv)
}

func _Lenven_CreateBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookingParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LenvenServer).CreateBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Lenven_CreateBooking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LenvenServer).CreateBooking(ctx, req.(*CreateBookingParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lenven_GetBookings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookingsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LenvenServer).GetBookings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Lenven_GetBookings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LenvenServer).GetBookings(ctx, req.(*GetBookingsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lenven_GetBookingByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookingByIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LenvenServer).GetBookingByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Lenven_GetBookingByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LenvenServer).GetBookingByID(ctx, req.(*GetBookingByIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lenven_CancelBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelBookingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LenvenServer).CancelBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Lenven_CancelBooking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LenvenServer).CancelBooking(ctx, req.(*CancelBookingReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Lenven_ServiceDesc is the grpc.ServiceDesc for Lenven service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Lenven_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Lenven",
	HandlerType: (*LenvenServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBooking",
			Handler:    _Lenven_CreateBooking_Handler,
		},
		{
			MethodName: "GetBookings",
			Handler:    _Lenven_GetBookings_Handler,
		},
		{
			MethodName: "GetBookingByID",
			Handler:    _Lenven_GetBookingByID_Handler,
		},
		{
			MethodName: "CancelBooking",
			Handler:    _Lenven_CancelBooking_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/lenven.proto",
}
