package server

import (
	"context"

	"github.com/Tee2912/To-Do-List-Go/api/proto/todolist"
	"github.com/Tee2912/To-Do-List-Go/db"
	"github.com/Tee2912/To-Do-List-Go/db/todo"
	"github.com/Tee2912/To-Do-List-Go/mapper"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	todolist.UnimplementedTodoServiceServer
	GrpcSrv *grpc.Server
	db      *db.MongoDb
}

func New(db *db.MongoDb) *server {
	grpcServer := grpc.NewServer()
	srv := &server{
		GrpcSrv: grpcServer,
		db:      db}
	todolist.RegisterTodoServiceServer(grpcServer, srv)
	reflection.Register(grpcServer)
	return srv
}

// CreateToDo creates a new todo
func (s *server) CreateToDo(ctx context.Context, in *todolist.ToDo) (*todolist.ToDo, error) {
	newTodo, err := mapper.TodoProtobufToTodoModel(in)
	if err != nil {
		return nil, err
	}

	createdTodo, err := todo.Create(ctx, s.db, newTodo)
	if err != nil {
		return nil, err
	}
	protoResponse, err := mapper.TodoModelToTodoProtobuf(createdTodo)
	if err != nil {
		return nil, err
	}
	return protoResponse, nil
}

// ReadToDo retrieves a todo by its ID
func (s *server) ReadToDo(ctx context.Context, in *todolist.ReadToDoReq) (*todolist.ToDo, error) {
	todo, err := todo.Get(ctx, s.db, in)
	if err != nil {
		return nil, errors.Wrapf(err, "getting todo with id %s", in.Id)
	}
	protoResponse, err := mapper.TodoModelToTodoProtobuf(todo)
	if err != nil {
		return nil, err
	}
	return protoResponse, nil
}

// UpdateToDo updates an existing todo
func (s *server) UpdateToDo(ctx context.Context, in *todolist.ToDo) (*todolist.ToDo, error) {
	todooUpdate, err := mapper.TodoProtobufToTodoModel(in)
	if err != nil {
		return nil, err
	}
	updatedTodo, err := todo.Update(ctx, s.db, todooUpdate)
	if err != nil {
		return nil, err
	}
	protoResponse, err := mapper.TodoModelToTodoProtobuf(updatedTodo)
	if err != nil {
		return nil, err
	}
	return protoResponse, nil
}

// DeleteToDo deletes a todo
func (s *server) DeleteToDo(ctx context.Context, in *todolist.DeleteToDoReq) (*todolist.DeleteToDoRes, error) {
	resp, err := todo.Delete(ctx, s.db, in)
	if err != nil {
		return nil, errors.Wrapf(err, "deleting todo with id %s", in.Id)
	}

	return resp, nil
}

// ListToDos lists all the todos
func (s *server) ListToDos(ctx context.Context, in *todolist.ListToDosReq) (*todolist.ListToDosRes, error) {
	todos, err := todo.List(ctx, s.db, in)
	if err != nil {
		return nil, errors.Wrap(err, "listing todos")
	}
	protoResponse, err := mapper.TodoModelListToListTodosResponse(todos)
	if err != nil {
		return nil, err
	}
	return protoResponse, nil
}
