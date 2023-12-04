// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.2
// source: todolist.proto

package todolist

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
	TodoService_CreateToDo_FullMethodName = "/todolist.TodoService/CreateToDo"
	TodoService_ReadToDo_FullMethodName   = "/todolist.TodoService/ReadToDo"
	TodoService_UpdateToDo_FullMethodName = "/todolist.TodoService/UpdateToDo"
	TodoService_DeleteToDo_FullMethodName = "/todolist.TodoService/DeleteToDo"
	TodoService_ListToDos_FullMethodName  = "/todolist.TodoService/ListToDos"
)

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoServiceClient interface {
	CreateToDo(ctx context.Context, in *ToDo, opts ...grpc.CallOption) (*ToDo, error)
	ReadToDo(ctx context.Context, in *ReadToDoReq, opts ...grpc.CallOption) (*ToDo, error)
	UpdateToDo(ctx context.Context, in *ToDo, opts ...grpc.CallOption) (*ToDo, error)
	DeleteToDo(ctx context.Context, in *DeleteToDoReq, opts ...grpc.CallOption) (*DeleteToDoRes, error)
	ListToDos(ctx context.Context, in *ListToDosReq, opts ...grpc.CallOption) (*ListToDosRes, error)
}

type todoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoServiceClient(cc grpc.ClientConnInterface) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) CreateToDo(ctx context.Context, in *ToDo, opts ...grpc.CallOption) (*ToDo, error) {
	out := new(ToDo)
	err := c.cc.Invoke(ctx, TodoService_CreateToDo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) ReadToDo(ctx context.Context, in *ReadToDoReq, opts ...grpc.CallOption) (*ToDo, error) {
	out := new(ToDo)
	err := c.cc.Invoke(ctx, TodoService_ReadToDo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) UpdateToDo(ctx context.Context, in *ToDo, opts ...grpc.CallOption) (*ToDo, error) {
	out := new(ToDo)
	err := c.cc.Invoke(ctx, TodoService_UpdateToDo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) DeleteToDo(ctx context.Context, in *DeleteToDoReq, opts ...grpc.CallOption) (*DeleteToDoRes, error) {
	out := new(DeleteToDoRes)
	err := c.cc.Invoke(ctx, TodoService_DeleteToDo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) ListToDos(ctx context.Context, in *ListToDosReq, opts ...grpc.CallOption) (*ListToDosRes, error) {
	out := new(ListToDosRes)
	err := c.cc.Invoke(ctx, TodoService_ListToDos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServiceServer is the server API for TodoService service.
// All implementations must embed UnimplementedTodoServiceServer
// for forward compatibility
type TodoServiceServer interface {
	CreateToDo(context.Context, *ToDo) (*ToDo, error)
	ReadToDo(context.Context, *ReadToDoReq) (*ToDo, error)
	UpdateToDo(context.Context, *ToDo) (*ToDo, error)
	DeleteToDo(context.Context, *DeleteToDoReq) (*DeleteToDoRes, error)
	ListToDos(context.Context, *ListToDosReq) (*ListToDosRes, error)
	mustEmbedUnimplementedTodoServiceServer()
}

// UnimplementedTodoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTodoServiceServer struct {
}

func (UnimplementedTodoServiceServer) CreateToDo(context.Context, *ToDo) (*ToDo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateToDo not implemented")
}
func (UnimplementedTodoServiceServer) ReadToDo(context.Context, *ReadToDoReq) (*ToDo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadToDo not implemented")
}
func (UnimplementedTodoServiceServer) UpdateToDo(context.Context, *ToDo) (*ToDo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateToDo not implemented")
}
func (UnimplementedTodoServiceServer) DeleteToDo(context.Context, *DeleteToDoReq) (*DeleteToDoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteToDo not implemented")
}
func (UnimplementedTodoServiceServer) ListToDos(context.Context, *ListToDosReq) (*ListToDosRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListToDos not implemented")
}
func (UnimplementedTodoServiceServer) mustEmbedUnimplementedTodoServiceServer() {}

// UnsafeTodoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoServiceServer will
// result in compilation errors.
type UnsafeTodoServiceServer interface {
	mustEmbedUnimplementedTodoServiceServer()
}

func RegisterTodoServiceServer(s grpc.ServiceRegistrar, srv TodoServiceServer) {
	s.RegisterService(&TodoService_ServiceDesc, srv)
}

func _TodoService_CreateToDo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ToDo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).CreateToDo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoService_CreateToDo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).CreateToDo(ctx, req.(*ToDo))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_ReadToDo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadToDoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).ReadToDo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoService_ReadToDo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).ReadToDo(ctx, req.(*ReadToDoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_UpdateToDo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ToDo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).UpdateToDo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoService_UpdateToDo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).UpdateToDo(ctx, req.(*ToDo))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_DeleteToDo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteToDoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).DeleteToDo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoService_DeleteToDo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).DeleteToDo(ctx, req.(*DeleteToDoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_ListToDos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListToDosReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).ListToDos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoService_ListToDos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).ListToDos(ctx, req.(*ListToDosReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TodoService_ServiceDesc is the grpc.ServiceDesc for TodoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "todolist.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateToDo",
			Handler:    _TodoService_CreateToDo_Handler,
		},
		{
			MethodName: "ReadToDo",
			Handler:    _TodoService_ReadToDo_Handler,
		},
		{
			MethodName: "UpdateToDo",
			Handler:    _TodoService_UpdateToDo_Handler,
		},
		{
			MethodName: "DeleteToDo",
			Handler:    _TodoService_DeleteToDo_Handler,
		},
		{
			MethodName: "ListToDos",
			Handler:    _TodoService_ListToDos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todolist.proto",
}