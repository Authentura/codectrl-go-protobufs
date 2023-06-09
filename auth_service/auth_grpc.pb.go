// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.6
// source: auth.proto

package auth_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuthenticationClient is the client API for Authentication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticationClient interface {
	// Takes a token and verifies that it is:
	//
	//	a) In valid format
	//	b) Present in the database/filesystem (down to implementation)
	//	c) Has the correct permissions for the given intent
	VerifyToken(ctx context.Context, in *VerifyTokenRequest, opts ...grpc.CallOption) (*VerifyTokenRequestResult, error)
	GenerateToken(ctx context.Context, in *GenerateTokenRequest, opts ...grpc.CallOption) (*GenerateTokenRequestResult, error)
	RevokeToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*RevokeTokenRequestResult, error)
	RefreshToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Token, error)
	GithubLogin(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*LoginUrl, error)
}

type authenticationClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationClient(cc grpc.ClientConnInterface) AuthenticationClient {
	return &authenticationClient{cc}
}

func (c *authenticationClient) VerifyToken(ctx context.Context, in *VerifyTokenRequest, opts ...grpc.CallOption) (*VerifyTokenRequestResult, error) {
	out := new(VerifyTokenRequestResult)
	err := c.cc.Invoke(ctx, "/codectrl.auth_service.Authentication/VerifyToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) GenerateToken(ctx context.Context, in *GenerateTokenRequest, opts ...grpc.CallOption) (*GenerateTokenRequestResult, error) {
	out := new(GenerateTokenRequestResult)
	err := c.cc.Invoke(ctx, "/codectrl.auth_service.Authentication/GenerateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) RevokeToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*RevokeTokenRequestResult, error) {
	out := new(RevokeTokenRequestResult)
	err := c.cc.Invoke(ctx, "/codectrl.auth_service.Authentication/RevokeToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) RefreshToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/codectrl.auth_service.Authentication/RefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) GithubLogin(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*LoginUrl, error) {
	out := new(LoginUrl)
	err := c.cc.Invoke(ctx, "/codectrl.auth_service.Authentication/GithubLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServer is the server API for Authentication service.
// All implementations must embed UnimplementedAuthenticationServer
// for forward compatibility
type AuthenticationServer interface {
	// Takes a token and verifies that it is:
	//
	//	a) In valid format
	//	b) Present in the database/filesystem (down to implementation)
	//	c) Has the correct permissions for the given intent
	VerifyToken(context.Context, *VerifyTokenRequest) (*VerifyTokenRequestResult, error)
	GenerateToken(context.Context, *GenerateTokenRequest) (*GenerateTokenRequestResult, error)
	RevokeToken(context.Context, *Token) (*RevokeTokenRequestResult, error)
	RefreshToken(context.Context, *Token) (*Token, error)
	GithubLogin(context.Context, *emptypb.Empty) (*LoginUrl, error)
	mustEmbedUnimplementedAuthenticationServer()
}

// UnimplementedAuthenticationServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenticationServer struct {
}

func (UnimplementedAuthenticationServer) VerifyToken(context.Context, *VerifyTokenRequest) (*VerifyTokenRequestResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyToken not implemented")
}
func (UnimplementedAuthenticationServer) GenerateToken(context.Context, *GenerateTokenRequest) (*GenerateTokenRequestResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateToken not implemented")
}
func (UnimplementedAuthenticationServer) RevokeToken(context.Context, *Token) (*RevokeTokenRequestResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeToken not implemented")
}
func (UnimplementedAuthenticationServer) RefreshToken(context.Context, *Token) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedAuthenticationServer) GithubLogin(context.Context, *emptypb.Empty) (*LoginUrl, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GithubLogin not implemented")
}
func (UnimplementedAuthenticationServer) mustEmbedUnimplementedAuthenticationServer() {}

// UnsafeAuthenticationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticationServer will
// result in compilation errors.
type UnsafeAuthenticationServer interface {
	mustEmbedUnimplementedAuthenticationServer()
}

func RegisterAuthenticationServer(s grpc.ServiceRegistrar, srv AuthenticationServer) {
	s.RegisterService(&Authentication_ServiceDesc, srv)
}

func _Authentication_VerifyToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).VerifyToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/codectrl.auth_service.Authentication/VerifyToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).VerifyToken(ctx, req.(*VerifyTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_GenerateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).GenerateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/codectrl.auth_service.Authentication/GenerateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).GenerateToken(ctx, req.(*GenerateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_RevokeToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).RevokeToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/codectrl.auth_service.Authentication/RevokeToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).RevokeToken(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/codectrl.auth_service.Authentication/RefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).RefreshToken(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_GithubLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).GithubLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/codectrl.auth_service.Authentication/GithubLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).GithubLogin(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Authentication_ServiceDesc is the grpc.ServiceDesc for Authentication service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Authentication_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "codectrl.auth_service.Authentication",
	HandlerType: (*AuthenticationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VerifyToken",
			Handler:    _Authentication_VerifyToken_Handler,
		},
		{
			MethodName: "GenerateToken",
			Handler:    _Authentication_GenerateToken_Handler,
		},
		{
			MethodName: "RevokeToken",
			Handler:    _Authentication_RevokeToken_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _Authentication_RefreshToken_Handler,
		},
		{
			MethodName: "GithubLogin",
			Handler:    _Authentication_GithubLogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
