// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protoServidorInformantes

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

// ConnToServidorClient is the client API for ConnToServidor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnToServidorClient interface {
	UpdateName(ctx context.Context, in *MensajeToServidor, opts ...grpc.CallOption) (*Respuesta, error)
	UpdateNumber(ctx context.Context, in *MensajeToServidor, opts ...grpc.CallOption) (*Respuesta, error)
	AddCity(ctx context.Context, in *MensajeToServidor, opts ...grpc.CallOption) (*Respuesta, error)
	DeleteCity(ctx context.Context, in *MensajeToServidor, opts ...grpc.CallOption) (*Respuesta, error)
}

type connToServidorClient struct {
	cc grpc.ClientConnInterface
}

func NewConnToServidorClient(cc grpc.ClientConnInterface) ConnToServidorClient {
	return &connToServidorClient{cc}
}

func (c *connToServidorClient) UpdateName(ctx context.Context, in *MensajeToServidor, opts ...grpc.CallOption) (*Respuesta, error) {
	out := new(Respuesta)
	err := c.cc.Invoke(ctx, "/grpc.ConnToServidor/UpdateName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connToServidorClient) UpdateNumber(ctx context.Context, in *MensajeToServidor, opts ...grpc.CallOption) (*Respuesta, error) {
	out := new(Respuesta)
	err := c.cc.Invoke(ctx, "/grpc.ConnToServidor/UpdateNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connToServidorClient) AddCity(ctx context.Context, in *MensajeToServidor, opts ...grpc.CallOption) (*Respuesta, error) {
	out := new(Respuesta)
	err := c.cc.Invoke(ctx, "/grpc.ConnToServidor/AddCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connToServidorClient) DeleteCity(ctx context.Context, in *MensajeToServidor, opts ...grpc.CallOption) (*Respuesta, error) {
	out := new(Respuesta)
	err := c.cc.Invoke(ctx, "/grpc.ConnToServidor/DeleteCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnToServidorServer is the server API for ConnToServidor service.
// All implementations must embed UnimplementedConnToServidorServer
// for forward compatibility
type ConnToServidorServer interface {
	UpdateName(context.Context, *MensajeToServidor) (*Respuesta, error)
	UpdateNumber(context.Context, *MensajeToServidor) (*Respuesta, error)
	AddCity(context.Context, *MensajeToServidor) (*Respuesta, error)
	DeleteCity(context.Context, *MensajeToServidor) (*Respuesta, error)
	mustEmbedUnimplementedConnToServidorServer()
}

// UnimplementedConnToServidorServer must be embedded to have forward compatible implementations.
type UnimplementedConnToServidorServer struct {
}

func (UnimplementedConnToServidorServer) UpdateName(context.Context, *MensajeToServidor) (*Respuesta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateName not implemented")
}
func (UnimplementedConnToServidorServer) UpdateNumber(context.Context, *MensajeToServidor) (*Respuesta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNumber not implemented")
}
func (UnimplementedConnToServidorServer) AddCity(context.Context, *MensajeToServidor) (*Respuesta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCity not implemented")
}
func (UnimplementedConnToServidorServer) DeleteCity(context.Context, *MensajeToServidor) (*Respuesta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCity not implemented")
}
func (UnimplementedConnToServidorServer) mustEmbedUnimplementedConnToServidorServer() {}

// UnsafeConnToServidorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnToServidorServer will
// result in compilation errors.
type UnsafeConnToServidorServer interface {
	mustEmbedUnimplementedConnToServidorServer()
}

func RegisterConnToServidorServer(s grpc.ServiceRegistrar, srv ConnToServidorServer) {
	s.RegisterService(&ConnToServidor_ServiceDesc, srv)
}

func _ConnToServidor_UpdateName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MensajeToServidor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnToServidorServer).UpdateName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ConnToServidor/UpdateName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnToServidorServer).UpdateName(ctx, req.(*MensajeToServidor))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnToServidor_UpdateNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MensajeToServidor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnToServidorServer).UpdateNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ConnToServidor/UpdateNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnToServidorServer).UpdateNumber(ctx, req.(*MensajeToServidor))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnToServidor_AddCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MensajeToServidor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnToServidorServer).AddCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ConnToServidor/AddCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnToServidorServer).AddCity(ctx, req.(*MensajeToServidor))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnToServidor_DeleteCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MensajeToServidor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnToServidorServer).DeleteCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ConnToServidor/DeleteCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnToServidorServer).DeleteCity(ctx, req.(*MensajeToServidor))
	}
	return interceptor(ctx, in, info, handler)
}

// ConnToServidor_ServiceDesc is the grpc.ServiceDesc for ConnToServidor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConnToServidor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.ConnToServidor",
	HandlerType: (*ConnToServidorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateName",
			Handler:    _ConnToServidor_UpdateName_Handler,
		},
		{
			MethodName: "UpdateNumber",
			Handler:    _ConnToServidor_UpdateNumber_Handler,
		},
		{
			MethodName: "AddCity",
			Handler:    _ConnToServidor_AddCity_Handler,
		},
		{
			MethodName: "DeleteCity",
			Handler:    _ConnToServidor_DeleteCity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "servidorInformante.proto",
}