package todo

import (
	"context"
	"fmt"

	"github.com/Tee2912/To-Do-List-Go/api/proto/todolist"
	"github.com/Tee2912/To-Do-List-Go/db"
	"github.com/Tee2912/To-Do-List-Go/db/todo/models"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "todo"

// Cursor is an interface that defines the methods necessary for iterating
// over query results in a data layer.
// This interface is particularly useful for simplifying unit tests
// by allowing the implementation of mock cursors that can be used
// for testing data retrieval and manipulation operations.
type Cursor interface {
	Decode(interface{}) error
	Err() error
	Close(context.Context) error
	Next(context.Context) bool
}

type cursorWrapper struct {
	*mongo.Cursor
}

// For ease of unit testing.
var (
	uuidProvider         = uuid.NewString
	insertIntoCollection = func(ctx context.Context, collection *mongo.Collection, document interface{}) (*mongo.InsertOneResult, error) {
		return collection.InsertOne(ctx, document)
	}
	find = func(ctx context.Context, collection *mongo.Collection, filter interface{}) (Cursor, error) {
		cur, err := collection.Find(ctx, filter)
		return &cursorWrapper{cur}, err
	}
	findOne = func(ctx context.Context, collection *mongo.Collection, filter interface{}, p *models.ToDo) error {
		sr := collection.FindOne(ctx, filter)
		return sr.Decode(p)
	}
	updateOne = func(ctx context.Context, collection *mongo.Collection, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
		return collection.UpdateOne(ctx, filter, update)
	}
	deleteOne = func(ctx context.Context, collection *mongo.Collection, filter interface{}) (*mongo.DeleteResult, error) {
		return collection.DeleteOne(ctx, filter)
	}
)

// GetToDo by id
func Get(ctx context.Context, db *db.MongoDb, req *todolist.ReadToDoReq) (*models.ToDo, error) {
	coll := db.Client.Database(db.DatabaseName).Collection(collectionName)
	var todo models.ToDo
	err := findOne(ctx, coll, bson.M{"id": req.GetId()}, &todo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf(`todo with id "%s" does not exist`, req.GetId())
		}
		return nil, errors.Wrapf(err, `getting todo with id "%s"`, req.GetId())
	}
	return &todo, nil
}

// Create a new todo
func Create(ctx context.Context, db *db.MongoDb, newToDo *models.ToDo) (*models.ToDo, error) {
	coll := db.Client.Database(db.DatabaseName).Collection(collectionName)
	newToDo.Id = uuidProvider()

	_, err := insertIntoCollection(ctx, coll, newToDo)
	if err != nil {
		return nil, errors.Wrap(err, "inserting todo")
	}
	return newToDo, nil
}

// Update todo
func Update(ctx context.Context, db *db.MongoDb, todoToUpdate *models.ToDo) (*models.ToDo, error) {
	coll := db.Client.Database(db.DatabaseName).Collection(collectionName)
	_, err := updateOne(ctx, coll, bson.M{"id": todoToUpdate.Id}, bson.M{"$set": todoToUpdate})
	if err != nil {
		return nil, errors.Wrapf(err, `updating todo with id "%s"`, todoToUpdate.Id)
	}
	return todoToUpdate, nil
}

// Delete todo by id
func Delete(ctx context.Context, db *db.MongoDb, req *todolist.DeleteToDoReq) (*todolist.DeleteToDoRes, error) {
	coll := db.Client.Database(db.DatabaseName).Collection(collectionName)
	_, err := deleteOne(ctx, coll, bson.M{"id": req.Id})
	if err != nil {
		return nil, errors.Wrapf(err, `deleting todo with id "%s"`, req.Id)
	}
	return &todolist.DeleteToDoRes{Result: "success"}, nil
}

// List lists all products in the database.
func List(ctx context.Context, db *db.MongoDb, req *todolist.ListToDosReq) ([]*models.ToDo, error) {
	coll := db.Client.Database(db.DatabaseName).Collection(collectionName)
	cur, err := find(ctx, coll, bson.M{})
	if err != nil {
		return nil, errors.Wrap(err, "finding todos")
	}
	defer cur.Close(ctx)
	var todos []*models.ToDo
	for cur.Next(ctx) {
		var todo models.ToDo
		if err = cur.Decode(&todo); err != nil {
			return nil, errors.Wrap(err, "decoding todo")
		}
		todos = append(todos, &todo)
	}
	if err := cur.Err(); err != nil {
		return nil, errors.Wrap(err, "cursor error")
	}
	return todos, nil
}
