package mapper

import (
	"github.com/Tee2912/To-Do-List-Go/api/proto/todolist"
	"github.com/Tee2912/To-Do-List-Go/db/todo/models"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/structpb"
)

// For ease of unit testing.
var structpbNewValue = structpb.NewValue

// TodoProtobufToProductModel converts a Protobuf Todo message to a MongoDB Todo model.
func ProductProtobufToProductModel(todo *todolist.ToDo) (*models.ToDo, error) {
	dbProduct := &models.ToDo{
		Id:          todo.Id,
		Name:        todo.Name,
		Description: todo.Description,
		Status:      todo.Status,
	}
	attributes := make(map[string]interface{})
	for k, p := range todo.Attributes {
		attributes[k] = p.AsInterface()
	}
	dbProduct.Attributes = attributes
	return dbProduct, nil
}

// ProductModelToProductProtobuf converts a MongoDB Product model to a Protobuf Product message.
func ProductModelToProductProtobuf(dbProduct *models.Product) (*todolist.Product, error) {
	product := &todolist.Product{
		Uuid:        dbProduct.Uuid,
		Name:        dbProduct.Name,
		Description: dbProduct.Description,
		Price:       dbProduct.Price,
	}
	var err error
	attributes := make(map[string]*structpb.Value)
	for k, p := range dbProduct.Attributes {
		attributes[k], err = structpbNewValue(p)
		if err != nil {
			return nil, errors.Wrapf(err, `parsing attribute "%s"`, k)
		}
	}
	product.Attributes = attributes
	return product, nil
}

// ProductModelListToListProductsResponse converts a list of MongoDB Product models to a Protobuf ListProductsResponse message.
func ProductModelListToListProductsResponse(dbProducts []*models.Product) (*todolist.ListProductsResponse, error) {
	response := &todolist.ListProductsResponse{}
	products := []*todolist.Product{}
	for _, dbProduct := range dbProducts {
		product, err := ProductModelToProductProtobuf(dbProduct)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	response.Products = products
	return response, nil
}
