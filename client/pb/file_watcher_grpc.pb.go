// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: proto/file_watcher.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	FileWatcher_CreateFile_FullMethodName    = "/proto.FileWatcher/CreateFile"
	FileWatcher_CreateDir_FullMethodName     = "/proto.FileWatcher/CreateDir"
	FileWatcher_WriteFile_FullMethodName     = "/proto.FileWatcher/WriteFile"
	FileWatcher_RemoveFileDir_FullMethodName = "/proto.FileWatcher/RemoveFileDir"
	FileWatcher_RenameFileDir_FullMethodName = "/proto.FileWatcher/RenameFileDir"
)

// FileWatcherClient is the client API for FileWatcher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileWatcherClient interface {
	CreateFile(ctx context.Context, in *CreateFileRequest, opts ...grpc.CallOption) (*CreateFileResponse, error)
	CreateDir(ctx context.Context, in *CreateDirRequest, opts ...grpc.CallOption) (*CreateDirResponse, error)
	WriteFile(ctx context.Context, in *WriteFileRequest, opts ...grpc.CallOption) (*WriteFileResponse, error)
	RemoveFileDir(ctx context.Context, in *RemoveFileDirRequest, opts ...grpc.CallOption) (*RemoveFileDirResponse, error)
	RenameFileDir(ctx context.Context, in *RenameFileDirRequest, opts ...grpc.CallOption) (*RenameFileDirResponse, error)
}

type fileWatcherClient struct {
	cc grpc.ClientConnInterface
}

func NewFileWatcherClient(cc grpc.ClientConnInterface) FileWatcherClient {
	return &fileWatcherClient{cc}
}

func (c *fileWatcherClient) CreateFile(ctx context.Context, in *CreateFileRequest, opts ...grpc.CallOption) (*CreateFileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateFileResponse)
	err := c.cc.Invoke(ctx, FileWatcher_CreateFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileWatcherClient) CreateDir(ctx context.Context, in *CreateDirRequest, opts ...grpc.CallOption) (*CreateDirResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateDirResponse)
	err := c.cc.Invoke(ctx, FileWatcher_CreateDir_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileWatcherClient) WriteFile(ctx context.Context, in *WriteFileRequest, opts ...grpc.CallOption) (*WriteFileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WriteFileResponse)
	err := c.cc.Invoke(ctx, FileWatcher_WriteFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileWatcherClient) RemoveFileDir(ctx context.Context, in *RemoveFileDirRequest, opts ...grpc.CallOption) (*RemoveFileDirResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveFileDirResponse)
	err := c.cc.Invoke(ctx, FileWatcher_RemoveFileDir_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileWatcherClient) RenameFileDir(ctx context.Context, in *RenameFileDirRequest, opts ...grpc.CallOption) (*RenameFileDirResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RenameFileDirResponse)
	err := c.cc.Invoke(ctx, FileWatcher_RenameFileDir_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileWatcherServer is the server API for FileWatcher service.
// All implementations must embed UnimplementedFileWatcherServer
// for forward compatibility
type FileWatcherServer interface {
	CreateFile(context.Context, *CreateFileRequest) (*CreateFileResponse, error)
	CreateDir(context.Context, *CreateDirRequest) (*CreateDirResponse, error)
	WriteFile(context.Context, *WriteFileRequest) (*WriteFileResponse, error)
	RemoveFileDir(context.Context, *RemoveFileDirRequest) (*RemoveFileDirResponse, error)
	RenameFileDir(context.Context, *RenameFileDirRequest) (*RenameFileDirResponse, error)
	mustEmbedUnimplementedFileWatcherServer()
}

// UnimplementedFileWatcherServer must be embedded to have forward compatible implementations.
type UnimplementedFileWatcherServer struct {
}

func (UnimplementedFileWatcherServer) CreateFile(context.Context, *CreateFileRequest) (*CreateFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFile not implemented")
}
func (UnimplementedFileWatcherServer) CreateDir(context.Context, *CreateDirRequest) (*CreateDirResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDir not implemented")
}
func (UnimplementedFileWatcherServer) WriteFile(context.Context, *WriteFileRequest) (*WriteFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteFile not implemented")
}
func (UnimplementedFileWatcherServer) RemoveFileDir(context.Context, *RemoveFileDirRequest) (*RemoveFileDirResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFileDir not implemented")
}
func (UnimplementedFileWatcherServer) RenameFileDir(context.Context, *RenameFileDirRequest) (*RenameFileDirResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenameFileDir not implemented")
}
func (UnimplementedFileWatcherServer) mustEmbedUnimplementedFileWatcherServer() {}

// UnsafeFileWatcherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileWatcherServer will
// result in compilation errors.
type UnsafeFileWatcherServer interface {
	mustEmbedUnimplementedFileWatcherServer()
}

func RegisterFileWatcherServer(s grpc.ServiceRegistrar, srv FileWatcherServer) {
	s.RegisterService(&FileWatcher_ServiceDesc, srv)
}

func _FileWatcher_CreateFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileWatcherServer).CreateFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileWatcher_CreateFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileWatcherServer).CreateFile(ctx, req.(*CreateFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileWatcher_CreateDir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDirRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileWatcherServer).CreateDir(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileWatcher_CreateDir_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileWatcherServer).CreateDir(ctx, req.(*CreateDirRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileWatcher_WriteFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileWatcherServer).WriteFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileWatcher_WriteFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileWatcherServer).WriteFile(ctx, req.(*WriteFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileWatcher_RemoveFileDir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFileDirRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileWatcherServer).RemoveFileDir(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileWatcher_RemoveFileDir_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileWatcherServer).RemoveFileDir(ctx, req.(*RemoveFileDirRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileWatcher_RenameFileDir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameFileDirRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileWatcherServer).RenameFileDir(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileWatcher_RenameFileDir_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileWatcherServer).RenameFileDir(ctx, req.(*RenameFileDirRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FileWatcher_ServiceDesc is the grpc.ServiceDesc for FileWatcher service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileWatcher_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.FileWatcher",
	HandlerType: (*FileWatcherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFile",
			Handler:    _FileWatcher_CreateFile_Handler,
		},
		{
			MethodName: "CreateDir",
			Handler:    _FileWatcher_CreateDir_Handler,
		},
		{
			MethodName: "WriteFile",
			Handler:    _FileWatcher_WriteFile_Handler,
		},
		{
			MethodName: "RemoveFileDir",
			Handler:    _FileWatcher_RemoveFileDir_Handler,
		},
		{
			MethodName: "RenameFileDir",
			Handler:    _FileWatcher_RenameFileDir_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/file_watcher.proto",
}
