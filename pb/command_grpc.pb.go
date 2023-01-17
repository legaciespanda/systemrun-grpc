// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/command.proto

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

// CommandClient is the client API for Command service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommandClient interface {
	//runs executable file on the server machine
	RunExecutableFile(ctx context.Context, in *RunExecutableRequest, opts ...grpc.CallOption) (*RunExecutableReply, error)
	//shutsdown a the remote server machine
	ShutdownSytem(ctx context.Context, in *ShutdownSytemRequest, opts ...grpc.CallOption) (*ShutdownSytemReply, error)
	//terminate or kill a process on the server machine
	TerminateProcess(ctx context.Context, in *TerminateProcessRequest, opts ...grpc.CallOption) (*TerminateProcessReply, error)
	//execute remote command on the server machine
	ExecuteRemoteCommand(ctx context.Context, in *RemoteCommandRequest, opts ...grpc.CallOption) (*RemoteCommandReply, error)
	//open a url on default browser of the server machine
	LaunchUrl(ctx context.Context, in *LaunchUrlRequest, opts ...grpc.CallOption) (*LaunchUrlReply, error)
}

type commandClient struct {
	cc grpc.ClientConnInterface
}

func NewCommandClient(cc grpc.ClientConnInterface) CommandClient {
	return &commandClient{cc}
}

func (c *commandClient) RunExecutableFile(ctx context.Context, in *RunExecutableRequest, opts ...grpc.CallOption) (*RunExecutableReply, error) {
	out := new(RunExecutableReply)
	err := c.cc.Invoke(ctx, "/command.Command/RunExecutableFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandClient) ShutdownSytem(ctx context.Context, in *ShutdownSytemRequest, opts ...grpc.CallOption) (*ShutdownSytemReply, error) {
	out := new(ShutdownSytemReply)
	err := c.cc.Invoke(ctx, "/command.Command/ShutdownSytem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandClient) TerminateProcess(ctx context.Context, in *TerminateProcessRequest, opts ...grpc.CallOption) (*TerminateProcessReply, error) {
	out := new(TerminateProcessReply)
	err := c.cc.Invoke(ctx, "/command.Command/TerminateProcess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandClient) ExecuteRemoteCommand(ctx context.Context, in *RemoteCommandRequest, opts ...grpc.CallOption) (*RemoteCommandReply, error) {
	out := new(RemoteCommandReply)
	err := c.cc.Invoke(ctx, "/command.Command/ExecuteRemoteCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandClient) LaunchUrl(ctx context.Context, in *LaunchUrlRequest, opts ...grpc.CallOption) (*LaunchUrlReply, error) {
	out := new(LaunchUrlReply)
	err := c.cc.Invoke(ctx, "/command.Command/LaunchUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommandServer is the server API for Command service.
// All implementations must embed UnimplementedCommandServer
// for forward compatibility
type CommandServer interface {
	//runs executable file on the server machine
	RunExecutableFile(context.Context, *RunExecutableRequest) (*RunExecutableReply, error)
	//shutsdown a the remote server machine
	ShutdownSytem(context.Context, *ShutdownSytemRequest) (*ShutdownSytemReply, error)
	//terminate or kill a process on the server machine
	TerminateProcess(context.Context, *TerminateProcessRequest) (*TerminateProcessReply, error)
	//execute remote command on the server machine
	ExecuteRemoteCommand(context.Context, *RemoteCommandRequest) (*RemoteCommandReply, error)
	//open a url on default browser of the server machine
	LaunchUrl(context.Context, *LaunchUrlRequest) (*LaunchUrlReply, error)
	mustEmbedUnimplementedCommandServer()
}

// UnimplementedCommandServer must be embedded to have forward compatible implementations.
type UnimplementedCommandServer struct {
}

func (UnimplementedCommandServer) RunExecutableFile(context.Context, *RunExecutableRequest) (*RunExecutableReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunExecutableFile not implemented")
}
func (UnimplementedCommandServer) ShutdownSytem(context.Context, *ShutdownSytemRequest) (*ShutdownSytemReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShutdownSytem not implemented")
}
func (UnimplementedCommandServer) TerminateProcess(context.Context, *TerminateProcessRequest) (*TerminateProcessReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TerminateProcess not implemented")
}
func (UnimplementedCommandServer) ExecuteRemoteCommand(context.Context, *RemoteCommandRequest) (*RemoteCommandReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteRemoteCommand not implemented")
}
func (UnimplementedCommandServer) LaunchUrl(context.Context, *LaunchUrlRequest) (*LaunchUrlReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LaunchUrl not implemented")
}
func (UnimplementedCommandServer) mustEmbedUnimplementedCommandServer() {}

// UnsafeCommandServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommandServer will
// result in compilation errors.
type UnsafeCommandServer interface {
	mustEmbedUnimplementedCommandServer()
}

func RegisterCommandServer(s grpc.ServiceRegistrar, srv CommandServer) {
	s.RegisterService(&Command_ServiceDesc, srv)
}

func _Command_RunExecutableFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunExecutableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServer).RunExecutableFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/command.Command/RunExecutableFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServer).RunExecutableFile(ctx, req.(*RunExecutableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Command_ShutdownSytem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShutdownSytemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServer).ShutdownSytem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/command.Command/ShutdownSytem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServer).ShutdownSytem(ctx, req.(*ShutdownSytemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Command_TerminateProcess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TerminateProcessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServer).TerminateProcess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/command.Command/TerminateProcess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServer).TerminateProcess(ctx, req.(*TerminateProcessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Command_ExecuteRemoteCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoteCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServer).ExecuteRemoteCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/command.Command/ExecuteRemoteCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServer).ExecuteRemoteCommand(ctx, req.(*RemoteCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Command_LaunchUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LaunchUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServer).LaunchUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/command.Command/LaunchUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServer).LaunchUrl(ctx, req.(*LaunchUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Command_ServiceDesc is the grpc.ServiceDesc for Command service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Command_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "command.Command",
	HandlerType: (*CommandServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunExecutableFile",
			Handler:    _Command_RunExecutableFile_Handler,
		},
		{
			MethodName: "ShutdownSytem",
			Handler:    _Command_ShutdownSytem_Handler,
		},
		{
			MethodName: "TerminateProcess",
			Handler:    _Command_TerminateProcess_Handler,
		},
		{
			MethodName: "ExecuteRemoteCommand",
			Handler:    _Command_ExecuteRemoteCommand_Handler,
		},
		{
			MethodName: "LaunchUrl",
			Handler:    _Command_LaunchUrl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/command.proto",
}
