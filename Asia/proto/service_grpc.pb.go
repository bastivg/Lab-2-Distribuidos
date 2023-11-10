// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: service.proto

package proto

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

// MyServiceClient is the client API for MyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MyServiceClient interface {
	SendContinentMsg(ctx context.Context, in *MensajeContinente, opts ...grpc.CallOption) (*RespuestaGenerica, error)
	SendOMSdepositar(ctx context.Context, in *AlmacenarEnDN, opts ...grpc.CallOption) (*RespuestaGenerica, error)
	SendOMSask(ctx context.Context, in *PedirDN, opts ...grpc.CallOption) (*NombreCompleto, error)
	SendONUMsg(ctx context.Context, in *ConsultaPoblacion, opts ...grpc.CallOption) (*ListaNombres, error)
}

type myServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMyServiceClient(cc grpc.ClientConnInterface) MyServiceClient {
	return &myServiceClient{cc}
}

func (c *myServiceClient) SendContinentMsg(ctx context.Context, in *MensajeContinente, opts ...grpc.CallOption) (*RespuestaGenerica, error) {
	out := new(RespuestaGenerica)
	err := c.cc.Invoke(ctx, "/main.MyService/SendContinentMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myServiceClient) SendOMSdepositar(ctx context.Context, in *AlmacenarEnDN, opts ...grpc.CallOption) (*RespuestaGenerica, error) {
	out := new(RespuestaGenerica)
	err := c.cc.Invoke(ctx, "/main.MyService/SendOMSdepositar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myServiceClient) SendOMSask(ctx context.Context, in *PedirDN, opts ...grpc.CallOption) (*NombreCompleto, error) {
	out := new(NombreCompleto)
	err := c.cc.Invoke(ctx, "/main.MyService/SendOMSask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myServiceClient) SendONUMsg(ctx context.Context, in *ConsultaPoblacion, opts ...grpc.CallOption) (*ListaNombres, error) {
	out := new(ListaNombres)
	err := c.cc.Invoke(ctx, "/main.MyService/SendONUMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MyServiceServer is the server API for MyService service.
// All implementations must embed UnimplementedMyServiceServer
// for forward compatibility
type MyServiceServer interface {
	SendContinentMsg(context.Context, *MensajeContinente) (*RespuestaGenerica, error)
	SendOMSdepositar(context.Context, *AlmacenarEnDN) (*RespuestaGenerica, error)
	SendOMSask(context.Context, *PedirDN) (*NombreCompleto, error)
	SendONUMsg(context.Context, *ConsultaPoblacion) (*ListaNombres, error)
	mustEmbedUnimplementedMyServiceServer()
}

// UnimplementedMyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMyServiceServer struct {
}

func (UnimplementedMyServiceServer) SendContinentMsg(context.Context, *MensajeContinente) (*RespuestaGenerica, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendContinentMsg not implemented")
}
func (UnimplementedMyServiceServer) SendOMSdepositar(context.Context, *AlmacenarEnDN) (*RespuestaGenerica, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendOMSdepositar not implemented")
}
func (UnimplementedMyServiceServer) SendOMSask(context.Context, *PedirDN) (*NombreCompleto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendOMSask not implemented")
}
func (UnimplementedMyServiceServer) SendONUMsg(context.Context, *ConsultaPoblacion) (*ListaNombres, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendONUMsg not implemented")
}
func (UnimplementedMyServiceServer) mustEmbedUnimplementedMyServiceServer() {}

// UnsafeMyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MyServiceServer will
// result in compilation errors.
type UnsafeMyServiceServer interface {
	mustEmbedUnimplementedMyServiceServer()
}

func RegisterMyServiceServer(s grpc.ServiceRegistrar, srv MyServiceServer) {
	s.RegisterService(&MyService_ServiceDesc, srv)
}

func _MyService_SendContinentMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MensajeContinente)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyServiceServer).SendContinentMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.MyService/SendContinentMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyServiceServer).SendContinentMsg(ctx, req.(*MensajeContinente))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyService_SendOMSdepositar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlmacenarEnDN)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyServiceServer).SendOMSdepositar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.MyService/SendOMSdepositar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyServiceServer).SendOMSdepositar(ctx, req.(*AlmacenarEnDN))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyService_SendOMSask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PedirDN)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyServiceServer).SendOMSask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.MyService/SendOMSask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyServiceServer).SendOMSask(ctx, req.(*PedirDN))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyService_SendONUMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsultaPoblacion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyServiceServer).SendONUMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.MyService/SendONUMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyServiceServer).SendONUMsg(ctx, req.(*ConsultaPoblacion))
	}
	return interceptor(ctx, in, info, handler)
}

// MyService_ServiceDesc is the grpc.ServiceDesc for MyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.MyService",
	HandlerType: (*MyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendContinentMsg",
			Handler:    _MyService_SendContinentMsg_Handler,
		},
		{
			MethodName: "SendOMSdepositar",
			Handler:    _MyService_SendOMSdepositar_Handler,
		},
		{
			MethodName: "SendOMSask",
			Handler:    _MyService_SendOMSask_Handler,
		},
		{
			MethodName: "SendONUMsg",
			Handler:    _MyService_SendONUMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
