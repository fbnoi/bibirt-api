// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.0
// source: user/v1/auth.proto

package v1

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
	Auth_RegisterAsAnonymous_FullMethodName = "/api.user.v1.Auth/RegisterAsAnonymous"
	Auth_UserInfo_FullMethodName            = "/api.user.v1.Auth/UserInfo"
	Auth_WSToken_FullMethodName             = "/api.user.v1.Auth/WSToken"
	Auth_RefreshToken_FullMethodName        = "/api.user.v1.Auth/RefreshToken"
	Auth_ConnUUID_FullMethodName            = "/api.user.v1.Auth/ConnUUID"
)

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	RegisterAsAnonymous(ctx context.Context, in *RegisterAsAnonymousRequest, opts ...grpc.CallOption) (*RegisterAsAnonymousReply, error)
	UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoReply, error)
	WSToken(ctx context.Context, in *WSTokenRequest, opts ...grpc.CallOption) (*WSTokenReply, error)
	RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenReply, error)
	ConnUUID(ctx context.Context, in *ConnUUIDRequest, opts ...grpc.CallOption) (*ConnUUIDReply, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) RegisterAsAnonymous(ctx context.Context, in *RegisterAsAnonymousRequest, opts ...grpc.CallOption) (*RegisterAsAnonymousReply, error) {
	out := new(RegisterAsAnonymousReply)
	err := c.cc.Invoke(ctx, Auth_RegisterAsAnonymous_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoReply, error) {
	out := new(UserInfoReply)
	err := c.cc.Invoke(ctx, Auth_UserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) WSToken(ctx context.Context, in *WSTokenRequest, opts ...grpc.CallOption) (*WSTokenReply, error) {
	out := new(WSTokenReply)
	err := c.cc.Invoke(ctx, Auth_WSToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenReply, error) {
	out := new(RefreshTokenReply)
	err := c.cc.Invoke(ctx, Auth_RefreshToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ConnUUID(ctx context.Context, in *ConnUUIDRequest, opts ...grpc.CallOption) (*ConnUUIDReply, error) {
	out := new(ConnUUIDReply)
	err := c.cc.Invoke(ctx, Auth_ConnUUID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations must embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	RegisterAsAnonymous(context.Context, *RegisterAsAnonymousRequest) (*RegisterAsAnonymousReply, error)
	UserInfo(context.Context, *UserInfoRequest) (*UserInfoReply, error)
	WSToken(context.Context, *WSTokenRequest) (*WSTokenReply, error)
	RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenReply, error)
	ConnUUID(context.Context, *ConnUUIDRequest) (*ConnUUIDReply, error)
	mustEmbedUnimplementedAuthServer()
}

// UnimplementedAuthServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) RegisterAsAnonymous(context.Context, *RegisterAsAnonymousRequest) (*RegisterAsAnonymousReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterAsAnonymous not implemented")
}
func (UnimplementedAuthServer) UserInfo(context.Context, *UserInfoRequest) (*UserInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
func (UnimplementedAuthServer) WSToken(context.Context, *WSTokenRequest) (*WSTokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WSToken not implemented")
}
func (UnimplementedAuthServer) RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedAuthServer) ConnUUID(context.Context, *ConnUUIDRequest) (*ConnUUIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConnUUID not implemented")
}
func (UnimplementedAuthServer) mustEmbedUnimplementedAuthServer() {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_RegisterAsAnonymous_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterAsAnonymousRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RegisterAsAnonymous(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_RegisterAsAnonymous_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RegisterAsAnonymous(ctx, req.(*RegisterAsAnonymousRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_UserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).UserInfo(ctx, req.(*UserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_WSToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WSTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).WSToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_WSToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).WSToken(ctx, req.(*WSTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_RefreshToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RefreshToken(ctx, req.(*RefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ConnUUID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnUUIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ConnUUID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_ConnUUID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ConnUUID(ctx, req.(*ConnUUIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.user.v1.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterAsAnonymous",
			Handler:    _Auth_RegisterAsAnonymous_Handler,
		},
		{
			MethodName: "UserInfo",
			Handler:    _Auth_UserInfo_Handler,
		},
		{
			MethodName: "WSToken",
			Handler:    _Auth_WSToken_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _Auth_RefreshToken_Handler,
		},
		{
			MethodName: "ConnUUID",
			Handler:    _Auth_ConnUUID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/v1/auth.proto",
}