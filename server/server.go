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

// server implements the ProductCatalogServiceServer interface.
// It handles the gRPC requests and delegates the actual processing to
// the corresponding functions in the product package.
type server struct {
	todolist.UnimplementedProductCatalogServiceServer
	GrpcSrv *grpc.Server
	db      *db.MongoDb
}

// New creates a new instance of the server with the provided database client.
// It sets up the gRPC server, registers the product catalog service,
// and initializes reflection for gRPC server debugging.
func New(db *db.MongoDb) *server {
	grpcServer := grpc.NewServer()
	srv := &server{
		GrpcSrv: grpcServer,
		db:      db}
	todolist.RegisterProductCatalogServiceServer(grpcServer, srv)
	reflection.Register(grpcServer)
	return srv
}

// CreateProduct creates a new product in the catalog.
// It delegates the actual creation logic to the product package's Create function.
func (s *server) CreateProduct(ctx context.Context, in *todolist.Product) (*todolist.Product, error) {
	newProduct, err := mapper.ProductProtobufToProductModel(in)
	if err != nil {
		return nil, err
	}
	createdProduct, err := todo.Create(ctx, s.db, newProduct)
	if err != nil {
		return nil, err
	}
	protoResponse, err := mapper.ProductModelToProductProtobuf(createdProduct)
	if err != nil {
		return nil, err
	}
	return protoResponse, nil
}

// GetProduct retrieves a product by its ID from the catalog.
// It delegates the actual retrieval logic to the product package's Get function.
func (s *server) GetProduct(ctx context.Context, in *todolist.GetProductRequest) (*todolist.Product, error) {
	product, err := todo.Get(ctx, s.db, in)
	if err != nil {
		return nil, errors.Wrapf(err, "getting product with uuid %s", in.Uuid)
	}
	protoResponse, err := mapper.ProductModelToProductProtobuf(product)
	if err != nil {
		return nil, err
	}
	return protoResponse, nil
}
