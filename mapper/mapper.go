package mapper

import (
	"github.com/Tee2912/To-Do-List-Go/api/proto/auth"
	"github.com/Tee2912/To-Do-List-Go/api/proto/todolist"
	"github.com/Tee2912/To-Do-List-Go/db/models"
	"google.golang.org/protobuf/types/known/structpb"
)

// For ease of unit testing.
var structpbNewValue = structpb.NewValue

// TodoProtobufToProductModel converts a Protobuf Todo message to a MongoDB Todo model.
func TodoProtobufToTodoModel(todo *todolist.ToDo) (*models.ToDo, error) {
	dbTodo := &models.ToDo{
		Name:        todo.Name,
		Description: todo.Description,
		Status:      todo.Status,
		Id:          todo.Id,
	}

	return dbTodo, nil
}

// TodoModelToTodoProtobuf converts a MongoDB Todo model to a Protobuf Todo message.
func TodoModelToTodoProtobuf(dbProduct *models.ToDo) (*todolist.ToDo, error) {
	todo := &todolist.ToDo{
		Id:          dbProduct.Id,
		Name:        dbProduct.Name,
		Description: dbProduct.Description,
		Status:      dbProduct.Status,
	}

	return todo, nil
}

// TodoModelListToListTodosResponse converts a list of MongoDB Todo models to a Protobuf ListToDosRes message.
func TodoModelListToListTodosResponse(dbProducts []*models.ToDo) (*todolist.ListToDosRes, error) {
	response := &todolist.ListToDosRes{}
	todos := []*todolist.ToDo{}
	for _, dbProduct := range dbProducts {
		product, err := TodoModelToTodoProtobuf(dbProduct)
		if err != nil {
			return nil, err
		}
		todos = append(todos, product)
	}
	response.Todos = todos
	return response, nil
}

// UserProtobufToUserModel converts a Protobuf user message to a MongoDB user model.
func UserProtobufToUserModel(user *auth.RegisterRequest) (*models.User, error) {
	dbUser := &models.User{
		Username: user.Username,
		Password: user.Password,
	}

	return dbUser, nil
}
