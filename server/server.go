package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Tee2912/To-Do-List-Go/api/proto/auth"
	"github.com/Tee2912/To-Do-List-Go/api/proto/todolist"
	"github.com/Tee2912/To-Do-List-Go/db"
	"github.com/Tee2912/To-Do-List-Go/db/todo"
	"github.com/Tee2912/To-Do-List-Go/db/user"
	"github.com/Tee2912/To-Do-List-Go/mapper"
	"github.com/Tee2912/To-Do-List-Go/utils"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	clientID     = ""
	clientSecret = ""
)

type server struct {
	todolist.UnimplementedTodoServiceServer
	auth.UnimplementedAuthServiceServer
	GrpcSrv     *grpc.Server
	oauthConfig *oauth2.Config
	db          *db.MongoDb
}

func New(db *db.MongoDb) *server {

	// Set up GitHub OAuth config
	oauthConfig := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  "http://localhost:4000/callback", // Update with your callback URL
		Scopes:       []string{"user:email"},
		Endpoint:     oauth2.Endpoint{AuthURL: "https://github.com/login/oauth/authorize", TokenURL: "https://github.com/login/oauth/access_token"},
	}

	grpcServer := grpc.NewServer()
	srv := &server{
		GrpcSrv:     grpcServer,
		oauthConfig: oauthConfig,
		db:          db}
	todolist.RegisterTodoServiceServer(grpcServer, srv)
	auth.RegisterAuthServiceServer(grpcServer, srv)
	reflection.Register(grpcServer)

	http.HandleFunc("/login", srv.handleLogin)
	http.HandleFunc("/callback", srv.handleOAuthCallback)
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

// Register new user
func (s *server) Register(ctx context.Context, in *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	password, err := utils.HashPassword(in.Password)
	in.Password = password
	newUser, err := mapper.UserProtobufToUserModel(in)
	if err != nil {
		return nil, err
	}

	resp, err := user.SignUpUser(ctx, s.db, newUser)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Login new user
func (s *server) Login(ctx context.Context, in *auth.LoginRequest) (*auth.LoginResponse, error) {
	user, err := user.Get(ctx, s.db, in)
	valid := utils.CheckPasswordHash(in.Password, user.Password)
	if !valid {
		log.Println("Wrong password")
		return &auth.LoginResponse{Message: "Failed to login"}, errors.New("Wrong password")
	}

	if err != nil {
		return nil, err
	}

	return &auth.LoginResponse{Message: "Login Successfully"}, nil
}

// handleLogin initiates the GitHub OAuth login process
func (s *server) handleLogin(w http.ResponseWriter, r *http.Request) {
	url := s.oauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// handleOAuthCallback handles the callback from GitHub after successful login
func (s *server) handleOAuthCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	// Exchange the code for an access token
	token, err := s.oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to exchange code for token", http.StatusInternalServerError)
		return
	}
	fmt.Print(token)

	fmt.Fprintf(w, "Login successful!")
}
